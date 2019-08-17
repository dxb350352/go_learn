package models

import (
	"context"
	"crypto/tls"
	"encoding/json"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/cayleygraph/cayley"
	"github.com/cayleygraph/cayley/graph"
	_ "github.com/cayleygraph/cayley/graph/bolt"
	"github.com/cayleygraph/cayley/quad"
	"github.com/cayleygraph/cayley/query"
	"github.com/cayleygraph/cayley/query/gizmo"
	"github.com/jackiedong168/gorequest"
	pool "github.com/sas/modules/pool/app/models"
	"os"
	"path/filepath"
	"strings"
	"sync"
	"time"
)

var (
	Predicate_Pointto  = "pointto"
	Predicate_Children = "children"
	SysConfigKey       = "ipgraph_last_update_time"
	Src                = "src"
	Dst                = "dst"
	PageSize           = 100
	MainGraph          *cayley.Handle
	MainGraphSure      *cayley.Handle
	Current            string
	CurrentSure        string
	CurrentGraph       *cayley.Handle
	CurrentGraphSure   *cayley.Handle
	Lock               = new(sync.Mutex)
	LockFlag           bool
	Cond               = sync.NewCond(Lock)
)

type Node struct {
	Ip   string `json:"ip"`
	Area string `json:"area"`
	Coor string `json:"coor"`
	Sure bool   `json:"sure"`
}

//每天一个文件
func getIpgraphPath(sure bool, t ...time.Time) string {
	tt := time.Now()
	if len(t) > 0 {
		tt = t[0]
	}
	if sure {
		return filepath.Join(IpgraphPath, tt.Format("20060102")+"ipgraph_sure.db")
	} else {
		return filepath.Join(IpgraphPath, tt.Format("20060102")+"ipgraph.db")
	}
}

func GetGraph(sure bool, t ...time.Time) *cayley.Handle {
	fp := IpgraphMain
	if sure {
		fp = IpgraphSureMain
	}
	if len(t) > 0 {
		fp = getIpgraphPath(sure, t...)
	}
	_, err := os.OpenFile(fp, os.O_RDWR|os.O_CREATE, os.ModePerm)
	if err != nil {
		Logger.Error(err)
		return nil
	}
	os.Chmod(fp, os.ModePerm)
	// Initialize the database
	graph.InitQuadStore("bolt", fp, nil)
	// Open and use the database
	store, err := cayley.NewGraph("bolt", fp, nil)
	//MainGraph, err := cayley.NewMemoryGraph()
	if err != nil {
		Logger.Error(err)
		return nil
	}
	return store
}

//
func GetMainGraph(sure bool) *cayley.Handle {
	var store *cayley.Handle
	if sure {
		if MainGraphSure == nil {
			MainGraphSure = GetGraph(sure)
		}
		store = MainGraphSure
	} else {
		if MainGraph == nil {
			MainGraph = GetGraph(sure)
		}
		store = MainGraph
	}
	return store
}
func GetCurrentGraph(sure bool) *cayley.Handle {
	var store *cayley.Handle
	t := time.Now().Format("20060102")
	if sure {
		if t != CurrentSure {
			if CurrentGraphSure != nil {
				CurrentGraphSure.Close()
			}
			store = GetGraph(sure, time.Now())
			CurrentGraphSure = store
			CurrentSure = t
		} else {
			store = CurrentGraphSure
		}
	}else{
		if t != Current {
			if CurrentGraph != nil {
				CurrentGraph.Close()
			}
			store = GetGraph(sure, time.Now())
			CurrentGraph = store
			Current = t
		} else {
			store = CurrentGraph
		}
	}

	if t := time.Now().Format("20060102"); t != Current || CurrentGraphSure == nil || CurrentGraph == nil {
		if sure {
			if CurrentGraphSure != nil {
				CurrentGraphSure.Close()
			}
			store = GetGraph(sure, time.Now())
			CurrentGraphSure = store
		} else {
			if CurrentGraph != nil {
				CurrentGraph.Close()
			}
			store = GetGraph(sure, time.Now())
			CurrentGraph = store
		}
		Current = t
	} else {
		if sure {
			store = CurrentGraphSure
		} else {
			store = CurrentGraph
		}
	}
	return store
}

func GetMCGraph(sure, current bool) *cayley.Handle {
	if current {
		return GetCurrentGraph(sure)
	} else {
		return GetMainGraph(sure)
	}
}

//查询到这个节点的所有路径
func Path(ip, area string, sure, current bool) []*simplejson.Json {
	var all []*simplejson.Json
	store := GetMCGraph(sure, current)
	//begain
	node := strings.Replace(GetNode(ip, area, "", sure).(string), `"`, `\"`, -1)
	m := map[string]bool{ip: true}

	for i := 0; ; i++ {
		var qu = fmt.Sprintf(`g.V("%s")`, node)
		for ii := 0; ii < i; ii++ {
			qu += fmt.Sprintf(`.In("%s")`, Predicate_Pointto)
		}
		qu += fmt.Sprintf(`.Tag("%s").In("%s").Tag("%s").TagArray()`, Dst, Predicate_Pointto, Src)
		result := getResult(store, qu)
		if len(result) == 0 {
			break
		}

		var usefull bool
		for ii := 0; ii < len(result); ii++ {
			if m[result[ii].Get(Src + "_ip").MustString()] {
				continue
			}
			usefull = true
			m[result[ii].Get(Src + "_ip").MustString()] = true
			all = append(all, result[ii])
		}
		if !usefull {
			break
		}
	}
	//组建所有路径
	return all
}

//向前向后
func PathAll(ip, area, coor string, sure, current bool) []*simplejson.Json {
	var all []*simplejson.Json
	store := GetMCGraph(sure, current)
	//begain
	node := strings.Replace(GetNode(ip, area, coor, sure).(string), `"`, `\"`, -1)
	m := map[string]bool{ip: true}
	//向前
	for i := 0; ; i++ {
		var qu = fmt.Sprintf(`g.V("%s")`, node)
		for ii := 0; ii < i; ii++ {
			qu += fmt.Sprintf(`.In("%s")`, Predicate_Pointto)
		}
		qu += fmt.Sprintf(`.Tag("%s").In("%s").Tag("%s").TagArray()`, Dst, Predicate_Pointto, Src)
		result := getResult(store, qu)
		if len(result) == 0 {
			Logger.Info("被攻击次数", 0)
			break
		}

		var usefull bool
		for ii := 0; ii < len(result); ii++ {
			if m[result[ii].Get(Src + "_ip").MustString()] {
				continue
			}
			usefull = true
			m[result[ii].Get(Src + "_ip").MustString()] = true
			all = append(all, result[ii])
		}
		if !usefull {
			break
		}
	}
	//向后
	for i := 0; ; i++ {
		var qu = fmt.Sprintf(`g.V("%s")`, node)
		for ii := 0; ii < i; ii++ {
			qu += fmt.Sprintf(`.Out("%s")`, Predicate_Pointto)
		}
		qu += fmt.Sprintf(`.Tag("%s").Out("%s").Tag("%s").TagArray()`, Src, Predicate_Pointto, Dst)
		result := getResult(store, qu)
		if len(result) == 0 {
			Logger.Info("攻击次数", 0)
			break
		}

		var usefull bool
		for ii := 0; ii < len(result); ii++ {
			if m[result[ii].Get(Dst + "_ip").MustString()] {
				continue
			}
			usefull = true
			m[result[ii].Get(Dst + "_ip").MustString()] = true
			all = append(all, result[ii])
		}
		if !usefull {
			break
		}
	}
	//组建所有路径
	return all
}

//查询
func Map(parent, has string, limit, startLevel, endLevel int, sure, current bool) []*simplejson.Json {
	store := GetMCGraph(sure, current)
	//begain
	var out, in string
	for ; startLevel < 4; startLevel++ {
		out += fmt.Sprintf(`.Out("%s")`, Predicate_Children)
	}
	for ; endLevel < 4; endLevel++ {
		in += fmt.Sprintf(`.In("%s")`, Predicate_Children)
	}
	var qu string
	if has != "" {
		qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")%s.Is("%s")`, parent, out, Src, Predicate_Pointto, Dst, in, has)
		qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")%s.Is("%s")).TagArray()`, parent, out, Dst, Predicate_Pointto, Src, in, has)
	} else {
		qu = fmt.Sprintf(`g.V("%s")%s.Tag("%s").Out("%s").Tag("%s")`, parent, out, Src, Predicate_Pointto, Dst)
		qu += fmt.Sprintf(`.Or(g.V("%s")%s.Tag("%s").In("%s").Tag("%s")).TagArray()`, parent, out, Dst, Predicate_Pointto, Src)
	}
	return getResult(store, qu)
}

//获取查询结果
func getResult(store *cayley.Handle, qu string) []*simplejson.Json {
	ses := gizmo.NewSession(store)
	c := make(chan query.Result, 2)
	ses.Execute(context.TODO(), qu, c, 10)
	var got []*simplejson.Json
	for res := range c {
		if res.Result() == nil {
			return got
		}
		byt, err := json.Marshal(res.Result())
		if err != nil {
			Logger.Error(err)
			return got
		}
		j, err := simplejson.NewJson(byt)
		if err != nil {
			Logger.Error(err)
			return got
		}
		var i int
		for {
			tj := j.GetIndex(i)
			i++
			if tj.Interface() == nil {
				break
			}
			gt := simplejson.New()
			buildJson(gt, tj, Src)
			buildJson(gt, tj, Dst)
			got = append(got, gt)
		}
	}
	return Unique(got)
}

//把节点中数据抽取出来
func buildJson(j, tj *simplejson.Json, prefix string) {
	ttjj, err := simplejson.NewJson([]byte(tj.Get(prefix).MustString()))
	if err != nil {
		Logger.Error(err)
		return
	}
	m := ttjj.MustMap()
	for k, v := range m {
		j.Set(prefix+"_"+k, v)
	}
}

//添加各种数据到图中
func IpgraphCron() {
	Ipgraph(300)
}

func Ipgraph(updateSecond int64) {
	Lock.Lock()
	defer Lock.Unlock()
	if LockFlag {
		Cond.Wait()
	}
	LockFlag = true
	defer func() {
		LockFlag = false
		Cond.Signal()
	}()
	//不确定的
	store := GetCurrentGraph(false)
	mainstore := GetMainGraph(false)
	//确定的
	store_sure := GetCurrentGraph(true)
	mainstore_sure := GetMainGraph(true)

	//startTime := "0"
	startTime := fmt.Sprint((time.Now().Unix() - updateSecond) * 1000)
	endTime := fmt.Sprint(time.Now().Unix() * 1000)

	for _, v := range AppType {
		arr := strings.Split(v, "|")
		if len(arr) < 4 {
			continue
		}
		for i := len(arr); i < 6; i++ {
			arr = append(arr, "")
		}
		if arr[4] == "" {
			GainData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], arr[5], store_sure, mainstore_sure)
		} else {
			GainData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], arr[5], store, mainstore)
		}
	}
}

//查询并添加到图中
func GainData(__app, __type, startTime, endTime, sip, dip, scoreField, keywords string, store, mainstore *cayley.Handle) {
	obj, err := pool.Pool.BorrowObject()
	if err != nil || obj == nil {
		Logger.Error(err)
		return
	}
	request := obj.(*gorequest.SuperAgent)
	defer pool.Pool.ReturnObject(obj)
	if strings.HasPrefix(SearchUrl, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	request.Get(SearchUrl)

	total := 1

	request.Param("app", __app)
	request.Param("_type", __type)
	request.Param("keywords", keywords)
	request.Param("size", fmt.Sprint(PageSize))
	request.Param("startTime", startTime)
	request.Param("endTime", endTime)
	for start := 0; start < total; {
		request.QueryData.Set("from", fmt.Sprint(start))
		start += PageSize
		_, body, errs := request.EndBytes()
		if len(errs) != 0 {
			Logger.Error(errs)
			continue
		}
		j, err := simplejson.NewJson(body)
		if err != nil {
			Logger.Error(err)
			continue
		}
		total = j.Get("total").MustInt()
		if total > 10000 {
			total = 10000
		}
		j = j.Get("hits")
		var quadarr []quad.Quad
		for i := 0; ; i++ {
			tj := j.GetIndex(i)
			if tj == nil || tj.Interface() == nil {
				break
			}
			vsip := tj.GetPath("_source", sip).MustString()
			vdip := tj.GetPath("_source", dip).MustString()
			src_country := tj.GetPath("_source", "src_country").MustString()
			src_province := strings.TrimSuffix(tj.GetPath("_source", "src_province").MustString(), "市")
			src_city := tj.GetPath("_source", "src_city").MustString()
			src_district := tj.GetPath("_source", "src_district").MustString()
			src_coor := strings.Join(tj.GetPath("_source", "src_geolocation", "coordinates").MustStringArray(), ",")
			if src_country != "中国" {
				src_district = src_city
				src_city = src_province
				src_province = src_country
				src_country = "世界"
			}
			src_pcd := fmt.Sprintf(`%s,%s,%s,%s`, src_country, src_province, src_city, src_district)

			dst_country := tj.GetPath("_source", "dst_country").MustString()
			dst_province := strings.TrimSuffix(tj.GetPath("_source", "dst_province").MustString(), "市")
			dst_city := tj.GetPath("_source", "dst_city").MustString()
			dst_district := tj.GetPath("_source", "dst_district").MustString()
			dst_coor := strings.Join(tj.GetPath("_source", "dst_geolocation", "coordinates").MustStringArray(), ",")
			if dst_country != "中国" {
				dst_district = dst_city
				dst_city = dst_province
				dst_province = dst_country
				dst_country = "世界"
			}
			dst_pcd := fmt.Sprintf(`%s,%s,%s,%s`, dst_country, dst_province, dst_city, dst_district)

			if vsip == "" || src_province == "" || src_city == "" || src_district == "" || vdip == "" || dst_province == "" || dst_city == "" || dst_district == "" {
				continue
			}
			_id := ""
			//_id := tj.Get("_id").MustString()
			score := tj.GetPath("_source", scoreField).MustFloat64()

			//省
			quadarr = AddQuad("", "", "", "", src_country, Predicate_Children, src_province, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", "", "", dst_country, Predicate_Children, dst_province, _id, quadarr, score, scoreField)
			//市
			quadarr = AddQuad("", "", "", "", src_province, Predicate_Children, src_city, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", "", "", dst_province, Predicate_Children, dst_city, _id, quadarr, score, scoreField)
			//县
			quadarr = AddQuad("", "", "", "", src_city, Predicate_Children, src_pcd, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", "", "", "", dst_city, Predicate_Children, dst_pcd, _id, quadarr, score, scoreField)
			//县ip
			quadarr = AddQuad("", vsip, "", src_coor, src_pcd, Predicate_Children, src_pcd, _id, quadarr, score, scoreField)
			quadarr = AddQuad("", vdip, "", dst_coor, dst_pcd, Predicate_Children, dst_pcd, _id, quadarr, score, scoreField)
			//真数据
			quadarr = AddQuad(vsip, vdip, src_coor, dst_coor, src_pcd, Predicate_Pointto, dst_pcd, _id, quadarr, score, scoreField)
		}
		if len(quadarr) > 0 {
			err = store.AddQuadSet(quadarr)
			if err != nil {
				Logger.Error(err)
			}
			err = mainstore.AddQuadSet(quadarr)
			if err != nil {
				Logger.Error(err)
			}
		}
	}
}

//生成一条路径
func AddQuad(sip, dip, scoor, dcoor, subject, predicate, object, label string, quadarr []quad.Quad, score float64, scoreField string) []quad.Quad {
	if subject != "" && object != "" && strings.Trim(subject, ",") != "" && strings.Trim(object, ",") != "" {
		subnode := GetNode(sip, subject, scoor, scoreField == "")
		objnode := GetNode(dip, object, dcoor, scoreField == "")
		if scoreField == "" {
			quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
		} else {
			if score-ScoreMax < 0 {
				quadarr = append(quadarr, quad.Make(subnode, predicate, objnode, label))
			}
		}
	}
	return quadarr
}

//生成节点
func GetNode(ip, area, coor string, sure bool) interface{} {
	if ip == "" {
		return area
	}
	node := Node{Ip: ip, Area: area, Sure: sure, Coor: coor}
	byt, _ := json.Marshal(node)
	return string(byt)
}

//合并30天数据
func Combine30DataFile() {
	Lock.Lock()
	defer Lock.Unlock()
	if LockFlag {
		Cond.Wait()
	}
	LockFlag = true
	defer func() {
		LockFlag = false
		Cond.Signal()
	}()
	//关闭graph
	if MainGraph != nil {
		MainGraph.Close()
		MainGraph = nil
	}
	if MainGraphSure != nil {
		MainGraphSure.Close()
		MainGraphSure = nil
	}
	//删除之前文件
	err := os.Remove(IpgraphMain)
	if err != nil {
		Logger.Error(err)
	}
	err = os.Remove(IpgraphSureMain)
	if err != nil {
		Logger.Error(err)
	}
	//30天文件
	MainGraph = GetGraph(false)
	MainGraphSure = GetGraph(true)
	var data []quad.Quad
	t := time.Now()
	for _, sure := range []bool{true, false} {
		store := MainGraph
		if sure {
			store = MainGraphSure
		}
		for i := 0; i < CombineDataFileNum; i++ {
			//30天文件
			tt := t.Add(time.Duration(-24*i) * time.Hour)
			if _, err := os.Stat(getIpgraphPath(sure, tt)); err != nil {
				Logger.Error(err)
				continue
			}
			var fstore *cayley.Handle
			var close bool
			if tt.Format("20060102") == time.Now().Format("20060102") {
				fstore = GetCurrentGraph(sure)
			} else {
				close = true
				fstore = GetGraph(sure, tt)
			}
			qu := fmt.Sprintf(`g.V().Tag("%s").Out().Tag("%s").TagArray()`, Src, Dst)
			c := make(chan query.Result, 2)
			fses := gizmo.NewSession(fstore)
			fses.Execute(context.TODO(), qu, c, 2)
			for res := range c {
				if res.Result() == nil {
					break
				}
				byt, err := json.Marshal(res.Result())
				if err != nil {
					Logger.Error(err)
					break
				}
				j, err := simplejson.NewJson(byt)
				if err != nil {
					Logger.Error(err)
					break
				}
				var i int
				for {
					tj := j.GetIndex(i)
					i++
					if tj.Interface() == nil {
						break
					}
					src := tj.Get(Src).MustString()
					dst := tj.Get(Dst).MustString()
					if strings.Contains(tj.Get(Src).MustString(), "{") {
						data = append(data, quad.Make(src, Predicate_Pointto, dst, ""))
					} else {
						data = append(data, quad.Make(src, Predicate_Children, dst, ""))
					}
					if len(data) >= PageSize {
						err = store.AddQuadSet(data)
						data = []quad.Quad{}
						if err != nil {
							Logger.Error(err)
						}
					}
				}
			}
			if len(data) > 0 {
				store.AddQuadSet(data)
				data = []quad.Quad{}
				if err != nil {
					Logger.Error(err)
				}
			}
			if close {
				fstore.Close()
			}
		}
	}
}

func CurrentData(parent, has string) []*simplejson.Json {
	startTime := fmt.Sprint((time.Now().Unix() - ShowDataSeconds) * 1000)
	endTime := fmt.Sprint(time.Now().Unix() * 1000)
	sure := map[string]bool{}
	for _, v := range AppType {
		arr := strings.Split(v, "|")
		if len(arr) < 4 {
			continue
		}
		for i := len(arr); i < 6; i++ {
			arr = append(arr, "")
		}
		if arr[4] == "" {
			GainCurrentData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], arr[5], sure)
			//} else {
			//	GainCurrentData(arr[0], arr[1], startTime, endTime, arr[2], arr[3], arr[4], arr[5], not_sure)
		}
	}
	sure_result := []*simplejson.Json{}
	for k, _ := range sure {
		if strings.Contains(k, parent) && strings.Contains(k, has) {
			src_dst := strings.SplitN(k, "}{", 2)
			if len(src_dst) != 2 {
				continue
			}
			j := simplejson.New()
			err := getResult2(j, src_dst[0]+"}", Src)
			if err != nil {
				Logger.Error(err)
				continue
			}
			err = getResult2(j, "{"+src_dst[1], Dst)
			if err != nil {
				Logger.Error(err)
				continue
			}
			sure_result = append(sure_result, j)
		}
	}
	return sure_result
}

func getResult2(j *simplejson.Json, jsonstr, prefix string) error {
	src, err := simplejson.NewJson([]byte(jsonstr))
	if err != nil {
		return err
	}
	m := src.MustMap()
	for k, v := range m {
		j.Set(prefix+"_"+k, v)
	}
	return nil
}

//获取实时数据
func GainCurrentData(__app, __type, startTime, endTime, sip, dip, scoreField, keywords string, result map[string]bool) {
	obj, err := pool.Pool.BorrowObject()
	if err != nil || obj == nil {
		Logger.Error(err)
		return
	}
	request := obj.(*gorequest.SuperAgent)
	defer pool.Pool.ReturnObject(obj)
	if strings.HasPrefix(SearchUrl, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	request.Get(SearchUrl)

	total := 1

	request.Param("app", __app)
	request.Param("_type", __type)
	request.Param("keywords", keywords)
	request.Param("size", fmt.Sprint(PageSize))
	request.Param("startTime", startTime)
	request.Param("endTime", endTime)
	for start := 0; start < total; {
		request.QueryData.Set("from", fmt.Sprint(start))
		start += PageSize
		_, body, errs := request.EndBytes()
		if len(errs) != 0 {
			Logger.Error(errs)
			continue
		}

		j, err := simplejson.NewJson(body)
		if err != nil {
			Logger.Error(err)
			continue
		}
		total = j.Get("total").MustInt()
		if total > 10000 {
			total = 10000
		}
		j = j.Get("hits")
		for i := 0; ; i++ {
			tj := j.GetIndex(i)
			if tj == nil || tj.Interface() == nil {
				break
			}
			vsip := tj.GetPath("_source", sip).MustString()
			vdip := tj.GetPath("_source", dip).MustString()
			src_country := tj.GetPath("_source", "src_country").MustString()
			src_province := strings.TrimSuffix(tj.GetPath("_source", "src_province").MustString(), "市")
			src_city := tj.GetPath("_source", "src_city").MustString()
			src_district := tj.GetPath("_source", "src_district").MustString()
			src_coor := strings.Join(tj.GetPath("_source", "src_geolocation", "coordinates").MustStringArray(), ",")
			if src_country != "中国" {
				src_district = src_city
				src_city = src_province
				src_province = src_country
				src_country = "世界"
			}
			src_pcd := fmt.Sprintf(`%s,%s,%s,%s`, src_country, src_province, src_city, src_district)

			dst_country := tj.GetPath("_source", "dst_country").MustString()
			dst_province := strings.TrimSuffix(tj.GetPath("_source", "dst_province").MustString(), "市")
			dst_city := tj.GetPath("_source", "dst_city").MustString()
			dst_district := tj.GetPath("_source", "dst_district").MustString()
			dst_coor := strings.Join(tj.GetPath("_source", "dst_geolocation", "coordinates").MustStringArray(), ",")
			if dst_country != "中国" {
				dst_district = dst_city
				dst_city = dst_province
				dst_province = dst_country
				dst_country = "世界"
			}
			dst_pcd := fmt.Sprintf(`%s,%s,%s,%s`, dst_country, dst_province, dst_city, dst_district)

			if vsip == "" || src_province == "" || src_city == "" || src_district == "" || vdip == "" || dst_province == "" || dst_city == "" || dst_district == "" {
				continue
			}
			//_id := tj.Get("_id").MustString()
			score := tj.GetPath("_source", scoreField).MustFloat64()
			//真数据
			src := GetNode(vsip, src_pcd, src_coor, scoreField == "").(string)
			dst := GetNode(vdip, dst_pcd, dst_coor, scoreField == "").(string)
			if scoreField != "" && score-ScoreMax > 0 {
				continue
			}
			result[fmt.Sprintf("%s%s", src, dst)] = true
		}
	}
}

func CombineJsonArray(json1, json2 []*simplejson.Json) []*simplejson.Json {
	maps := make(map[string]bool)
	for _, tmpJson := range json2 {
		str, err := tmpJson.MarshalJSON()
		if err != nil {
			Logger.Errorf("can't transform json to string,err:%v", err.Error())
			continue
		}
		maps[string(str)] = true
	}
	for _, tmpJson := range json1 {
		str, err := tmpJson.MarshalJSON()
		if err != nil {
			Logger.Errorf("can't transform json to string,err:%v", err.Error())
			continue
		}
		if _, ok := maps[string(str)]; !ok {
			json2 = append(json2, tmpJson)
		}
	}

	return json2
}

//去重
func Unique(arr [] *simplejson.Json) [] *simplejson.Json {
	m := map[string]bool{}
	for i := 0; i < len(arr); i++ {
		ll := len(m)
		m[fmt.Sprint(arr[i].Get("src_ip").MustString(), "_", arr[i].Get("dst_ip").MustString())] = true
		if len(m) == ll {
			arr = append(arr[:i], arr[i+1:]...)
			i--
		}
	}
	return arr
}
