package utils

import (
	"fmt"
	"github.com/tealeg/xlsx"
	"strings"
	"regexp"
	"github.com/miiton/kanaconv"
	"github.com/semicircle/gozhszht"
	"sort"
	"github.com/revel/revel"
)

var AreaMap map[rune]interface{}
//简称
var AreaMapShort map[rune]interface{}
var LevelKey rune = rune('_')
var ShengKey rune = rune('省')
var ShiKey rune = rune('市')
var LuKey rune = rune('路')
var DaoKey rune = rune('道')
var JieKey rune = rune('街')

type area struct {
	province string `json:"province"`
	city     string `json:"city"`
	district string `json:"district"`
	address  string `json:"adress"`
}

var special = []string{
	",东区",
	",南区",
	",西区",
	",北区",
	",中区",
	",中西区",
	",城区",
	",新城区",
	",郊区",
}

var cn2en = map[string]string{
	"〇":  "0",
	"零":  "0",
	"一":  "1",
	"二":  "2",
	"三":  "3",
	"四":  "4",
	"五":  "5",
	"六":  "6",
	"七":  "7",
	"八":  "8",
	"九":  "9",
	"栋":  "幢",
	"棟":  "幢",
	"撞":  "幢",
	"#":  "幢",
	"号楼": "幢",
	"组":  "单元",
	"层":  "楼",
	"室":  "号",
}
//栋、单元、门号共有数字
var dNumExp = "[0-9a-zA-Z〇零一二三四五六七八九十]"
//提取小区门号正则
var doorNumExp = regexp.MustCompile(`(` + dNumExp + `{1,3}(([栋棟撞幢\-#])|(号楼)))?(` + dNumExp + `{1,3}((单元)|[组\-]))?(` + dNumExp + `{1,4}[楼层\-])?` + dNumExp + `{1,5}[号室]?`)
//组成地址的字符
var addressExp = regexp.MustCompile("[\u4e00-\u9fa5A-Za-z0-9_\\-()# ]+")
//过长的有用字符
var toolongExp = regexp.MustCompile(`(_{2,})|( {2,})|(\({2,})|(\){2,})|(-{2,})`)

//初始化地址解析分词器
func InitAreaMap() {
	AreaMap = make(map[rune]interface{})
	AreaMapShort = make(map[rune]interface{})
	xlFile, err := xlsx.OpenFile(revel.BasePath + "/conf/area.xlsx")
	if err != nil {
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			if len(row.Cells) > 7 {
				Put(AreaMap, []rune(strings.TrimSpace(row.Cells[1].Value)), strings.TrimSpace(row.Cells[7].Value))
				Put(AreaMapShort, []rune(strings.TrimSpace(row.Cells[3].Value)), strings.TrimSpace(row.Cells[7].Value))
			}
		}
	}
}

//把词存入树中
func Put(areaMap map[rune]interface{}, key []rune, level string) {
	if len(key) == 0 || len(level) == 0 {
		return
	}
	var curr map[rune]interface{}
	curr = areaMap
	for i := 0; i < len(key); i++ {
		r := key[i]
		if _, ok := curr[r]; !ok {
			m := make(map[rune]interface{})
			curr[r] = m
			curr = m
			continue
		}
		if _, ok := curr[r].(map[rune]interface{}); !ok {
			curr[r] = make(map[rune]interface{})
		}
		curr = curr[r].(map[rune]interface{})
	}
	lkvarr, ok := curr[LevelKey].([]string)
	if ok {
		curr[LevelKey] = append(lkvarr, level)
	} else {
		curr[LevelKey] = []string{level}
	}
}
func AddValueAble(areaMap map[rune]interface{}, valueable map[string]float64, tarr []rune, start, i int, score float64, mm map[string]string) map[string]float64 {
	//后面有道或路字的丢掉
	if i != -1 && len(tarr) > i && (tarr[i] == LuKey || tarr[i] == DaoKey || tarr[i] == JieKey) {
		return valueable
	}
	//后面有道或路字的丢掉
	if i != -1 && len(tarr) > i+1 && (tarr[i+1] == LuKey || tarr[i+1] == DaoKey || tarr[i+1] == JieKey) {
		return valueable
	}
	//
	var kw string
	if i == -1 && start != -1 {
		kw = string(tarr[start:])
	} else if start != -1 && i != -1 {
		kw = string(tarr[start:i])
	}
	if level, ok := areaMap[LevelKey]; ok {
		values := level.([]string)
		for _, v := range values {
			mm[kw] = v
			valueable[v] = valueable[v] + score
		}
		return valueable
	}
	//处理没有省市
	//if tmp, ok := areaMap[ShengKey]; ok {
	//	areaMap = tmp.(map[rune]interface{})
	//}
	//if tmp, ok := areaMap[ShiKey]; ok {
	//	areaMap = tmp.(map[rune]interface{})
	//}

	//if level, ok := areaMap[LevelKey]; ok {
	//	values := level.([]string)
	//	for _, v := range values {
	//		//通过加省市出来的权值低一些
	//		valueable[v] = valueable[v] + 0.8
	//	}
	//}
	return valueable
}

//获取所有的省市县
func GetValueable(address string, score float64, valueable map[string]float64, mm map[string]string) (map[string]float64, map[string]string) {
	AreaMapVar := AreaMap
	if score != 1 {
		AreaMapVar = AreaMapShort
	}
	areaMap := AreaMapVar
	rarr := []rune(address)
	var start int = -1
	//所有的省市县
	for i := 0; i < len(rarr); i++ {
		v := rarr[i]
		if tmp, ok := areaMap[v]; ok && v != LevelKey {
			if start == -1 {
				start = i
			}
			areaMap = tmp.(map[rune]interface{})
			continue
		}

		valueable = AddValueAble(areaMap, valueable, rarr, start, i, score, mm)
		if start != -1 {
			i--
			start = -1
		}
		areaMap = AreaMapVar

	}
	valueable = AddValueAble(areaMap, valueable, rarr, start, -1, score, mm)
	//去掉特殊区
	for k, _ := range valueable {
		for _, vv := range special {
			if strings.HasSuffix(k, vv) {
				delete(valueable, k)
			}
		}
	}
	return valueable, mm
}

//处理成最终统一的地址信息
func HandleArea(address string, valueable map[string]float64, mm map[string]string) string {
	if len(valueable) == 0 {
		return fmt.Sprintf("%s,%s,%s,%s", "", "", "", address)
	}
	ar := getArea(valueable)
	getAddress(&ar, address, mm)

	return fmt.Sprintf("%s,%s,%s,%s", ar.province, ar.city, ar.district, ar.address)
}

//获取详细地址
func getAddress(ar *area, address string, mm map[string]string) {
	split := ar.district
	if split == "" {
		split = ar.city
		if split == "" {
			split = ar.province
		}
	}
	var spls []string
	for k, v := range mm {
		if strings.HasSuffix(v, ","+split) {
			spls = append(spls, k)
		}
	}
	if len(spls) > 0 {
		sort.Slice(spls, func(i, j int) bool {
			return len([]rune(spls[i])) > len([]rune(spls[j]))
		})
		split = spls[0]
	}
	//省略掉省市时
	//if split != "" && strings.Index(address, split) == -1 {
	//	split = strings.TrimRight(split, "省")
	//	split = strings.TrimRight(split, "市")
	//}
	addarr := strings.Split(address, split)
	if len(addarr) > 0 {
		//取最长的---也许可以把省市区县去掉了再比较长短
		sort.Slice(addarr, func(i, j int) bool {
			a1 := removePCD(ar, addarr[i])
			a2 := removePCD(ar, addarr[j])
			return len([]rune(a1)) > len([]rune(a2))
		})
		ar.address = getDoorFormat(strings.TrimSpace(addarr[0]))
	}
	//去掉地址前后带有省市区县
	ar.address = TrimLoopSuffixPrefix(ar.address, "中国")
	ar.address = trimPCD(ar.address, ar.city, "市")
	ar.address = trimPCD(ar.address, ar.province, "省")
	ar.address = trimPCD(ar.address, ar.district, "区")
}

func trimPCD(address, pcd, word string) string {
	if pcd == "" {
		return address
	}
	address = strings.TrimSpace(address)
	address = TrimLoopSuffixPrefix(address, pcd)
	if strings.HasSuffix(pcd, word) {
		address = TrimLoopSuffixPrefix(address, strings.TrimRight(pcd, word))
	}
	return strings.TrimSpace(address)
}

//去掉省市县
func removePCD(ar *area, address string) string {
	if ar.city != "" {
		address = strings.Replace(address, ar.city, "", -1)
		if strings.HasSuffix(ar.city, "市") {
			address = strings.Replace(address, strings.TrimRight(ar.city, "市"), "", -1)
		}
	}
	if ar.province != "" {
		address = strings.Replace(address, ar.province, "", -1)
		if strings.HasSuffix(ar.province, "省") {
			address = strings.Replace(address, strings.TrimRight(ar.province, "省"), "", -1)
		}
	}
	if ar.district != "" {
		address = strings.Replace(address, ar.district, "", -1)
	}
	return address
}

//找到最匹配的省市县区
func getArea(valueable map[string]float64) area {
	var ar area
	var prefix string
	province := getAreaPart(valueable, prefix, 1)
	if province != "" {
		prefix = province
	}
	city := getAreaPart(valueable, prefix, 2)
	if city != "" {
		prefix = city
	}
	district := getAreaPart(valueable, prefix, 3)

	if province != "" {
		ar.province = strings.Split(province, ",")[1]
	}
	if city != "" {
		carr := strings.Split(city, ",")
		ar.city = carr[2]
		if ar.province == "" {
			ar.province = carr[1]
		}
	}
	if district != "" {
		darr := strings.Split(district, ",")
		ar.district = darr[3]
		if ar.city == "" {
			ar.city = darr[2]
		}
		if ar.province == "" {
			ar.province = darr[1]
		}
	}

	return ar
}

//通过父区域找下一级区域
func getAreaPart(valueable map[string]float64, parent string, level int) string {
	var arstr string
	var arscore float64
	mm := map[float64][]string{}
	for k, v := range valueable {
		//取权值最高的
		if strings.Count(k, ",") != level || v < arscore {
			continue
		}

		if (parent != "" && strings.HasPrefix(k, parent+",") || parent == "" ) {
			mm[v] = append(mm[v], k)
			arstr = k
			arscore = v
		}
	}
	arr := mm[arscore]
	if len(arr) > 1 {
		var dalu []string
		var gat []string
		for _, v := range arr {
			if strings.Contains(v, "香港特别行政区") || strings.Contains(v, "澳门特别行政区") || strings.Contains(v, "台湾") {
				gat = append(gat, v)
			} else {
				dalu = append(dalu, v)
			}
		}
		sort.Strings(dalu)
		sort.Strings(gat)
		dalu = append(dalu, gat...)
		arstr = dalu[0]
	}
	return arstr
}

//有默认的前缀
func GetAddressValidatedPrefix(address, prefix string) string {
	//地址小于等于3个汉字，四个英文字母就忽略吧,只有省就去掉了
	if "社保登记" == address || len([]rune(address)) <= 3 || len(address) <= 4 {
		return ",,,"
	}
	if prefix == "" {
		return GetAddressValidated(address)
	}
	//有默认的前缀
	newaddress := GetAddressValidated(address)
	rnew := []rune(newaddress)
	//地址先解析，然后看原始数据前两个字是否和解析出来的省一致，如果不一致，就增加指定的字符串，然后再重新解析
	if !strings.HasPrefix(address, string(rnew[:2])) {
		newaddress = GetAddressValidated(prefix + address)
	}
	return newaddress
}

//经过统计，取出最有可能的
func GetAddressValidated(address string) string {
	//去掉特殊字符
	address = removeSpecialChar(address)
	m := map[string]float64{}
	mm := map[string]string{}
	m, mm = GetValueable(address, 1, m, mm)
	m, mm = GetValueable(address, 0.8, m, mm)
	valueable := map[string]float64{}
	for k, v := range m {
		valueable[k] = v
		karr := strings.Split(k, ",")
		//处理直辖市
		if len(karr) == 3 && karr[1]+"市" == karr[2] {
			kk := strings.Join(karr[:2], ",")
			valueable[kk] = valueable[kk] + v
		}
	}
	m = map[string]float64{}
	for k, v := range valueable {
		m[k] = v
	}
	for k, v := range valueable {
		for key, value := range m {
			if strings.HasPrefix(k, key+",") {
				m[key] = value + v
			}
		}
	}
	return HandleArea(address, m, mm)
}

//提取小区门号并转换统一格式
func getDoorFormat(src string) string {
	var key string
	arrarr := doorNumExp.FindAllStringSubmatch(src, -1)
	for _, v := range arrarr {
		//栋、单元、楼必须出现一个
		if v[1] == "" && v[5] == "" && v[8] == "" {
			continue
		}
		key = v[0]
		break
	}
	if key != "" {
		src = strings.Replace(src, key, formatDoor(key), 1)
	}
	return src
}

//转换统一小区门号格式
func formatDoor(src string) string {
	src = strings.ToUpper(src)
	//汉字转阿拉伯
	for k, v := range cn2en {
		src = strings.Replace(src, k, v, -1)
	}
	if strings.Index(src, "幢") == -1 {
		src = strings.Replace(src, "-", "幢", 1)
	}
	if strings.Index(src, "单元") == -1 {
		src = strings.Replace(src, "-", "单元", 1)
	}
	if strings.Index(src, "楼") == -1 {
		src = strings.Replace(src, "-", "楼", 1)
	}

	r := []rune("<" + src + ">")
	for i := 0; i < len(r); i++ {
		if r[i] == rune('十') {
			if IsNumber(string(r[i-1])) {
				if IsNumber(string(r[i+1])) {
					r = append(r[:i], r[i+1:]...)
				} else {
					r[i] = rune('0')
				}
			} else {
				if IsNumber(string(r[i+1])) {
					r[i] = rune('1')
				} else {
					tr := "10" + string(r[i+1:])
					r = append(r[:i], []rune(tr)...)
				}
			}
		}
	}
	return string(r[1:len(r)-1])
}

//去掉一些特殊字符
func removeSpecialChar(str string) string {
	//转半角
	str = kanaconv.SmartConv(str)
	//繁体转简体
	str = gozhszht.ToSimple(str)
	//繁体转简体BUG
	str = strings.Replace(str, "唿", "呼", -1)
	//抽取有用数据
	arr := addressExp.FindAllString(str, -1)
	result := strings.Join(arr, "")
	//去掉"其它", "其他"
	result = strings.Replace(result, "其它", "", -1)
	result = strings.Replace(result, "其他", "", -1)
	//删除过长但有点用的数据:--------_______((((((      )))))
	result = toolongExp.ReplaceAllStringFunc(result, func(hit string) string {
		if len(hit) > 0 {
			return hit[:1]
		}
		return ""
	})

	result = strings.Replace(result, "()", "", -1)
	return result
}
