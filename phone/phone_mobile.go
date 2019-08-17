package main

import (
	"github.com/tealeg/xlsx"
	"path/filepath"
	"strings"
	"os"
	"bufio"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/kr/pretty"
)

type PhoneData struct {
	PhoneCode string `json:"phone_code"`
	Province  string `json:"province"`
	City      string `json:"city"`
	AreaCode  string `json:"area_code"`
	PhoneLen  int    `json:"phone_len"`
	Phone     string `json:"phone"`
}

type MobileData struct {
	PhoneData
	PhonePrefix string `json:"phone_prefix"`
	Company     string `json:"company"`
}

var PhoneDataMap map[string]PhoneData
var MobileDataMap map[string]MobileData
//初始化电话解析
func InitPhoneMobile() {
	//固话
	PhoneDataMap = map[string]PhoneData{}
	xlFile, err := xlsx.OpenFile(filepath.Join("E:/GOPATH/src/testgo/phone", "phone.xlsx"))
	if err != nil {
		fmt.Println(err)
		return
	}
	for _, sheet := range xlFile.Sheets {
		for _, row := range sheet.Rows {
			phone, _ := row.Cells[0].String()
			area, _ := row.Cells[1].String()
			pc, _ := row.Cells[2].String()
			phone = strings.TrimSpace(phone)
			area = strings.TrimSpace(area)
			pc = strings.TrimSpace(pc)
			phonelen := 7
			if len(row.Cells) > 3 {
				plen, _ := row.Cells[3].String()
				if strings.TrimSpace(plen) == "8" {
					phonelen = 8
				}
			}
			if phone != "" && area != "" && strings.Count(pc, ",") == 1 {
				pcarr := strings.Split(pc, ",")
				PhoneDataMap[phone] = PhoneData{
					PhoneCode: phone,
					AreaCode:  area,
					Province:  pcarr[0],
					City:      pcarr[1],
					PhoneLen:  phonelen}
			}
		}
	}

	//手机
	MobileDataMap = map[string]MobileData{}
	f, err := os.Open(filepath.Join("E:/GOPATH/src/testgo/phone", "Mobile.txt"))
	if err != nil {
		fmt.Println(err)
		return
	}
	defer f.Close()
	sc := bufio.NewScanner(f)
	for sc.Scan() {
		line := sc.Text()
		line = strings.Replace(line, "\"", "", -1)
		arr := strings.Split(line, ",")
		if len(arr) >= 7 {
			MobileDataMap[arr[1]] = MobileData{
				PhonePrefix: arr[1],
				PhoneData: PhoneData{
					PhoneCode: arr[5],
					AreaCode:  PhoneDataMap[arr[5]].AreaCode,
					Province:  arr[2],
					City:      arr[3]},
				Company: arr[4],
			}
		}
	}
}

func main() {
	InitPhoneMobile()
	fmt.Println(MobileDataMap["1810614"])
	fmt.Println(MobileDataMap["1803278"])
	fmt.Println(MobileDataMap["1406939"])
	fmt.Println(MobileDataMap["1678485"])
	fmt.Println(MobileDataMap["1705929"])
	fmt.Println(MobileDataMap["1416789"])
	fmt.Println(MobileDataMap["1410329"])
	fmt.Println(MobileDataMap["1661309"])
	fmt.Println(MobileDataMap["1909573"])
	fmt.Println(MobileDataMap["1419234"])
	fmt.Println(MobileDataMap["1715454"])
	fmt.Println(MobileDataMap["1309645"])
	fmt.Println(MobileDataMap["1404907"])
	fmt.Println(MobileDataMap["1843589"])
	fmt.Println(MobileDataMap["1740509"])
	fmt.Println(MobileDataMap["1650283"])
	fmt.Println(MobileDataMap["1401378"])
	fmt.Println(MobileDataMap["1570553"])
	j := simplejson.New()
	j.Set("callertogatewaye164", "89797082317377751")
	handlePhone(j)
	pretty.Println(j)
}

var specialPhonePrefix map[string]bool = map[string]bool{"00852": true, "00853": true, "00886": true}
var phoneFiles []string = []string{"callere164", "calleee164", "calleraccesse164", "calleeaccesse164", "callertogatewaye164", "calleetogatewaye164"}

func handlePhone(j *simplejson.Json) {
	for _, key := range phoneFiles {
		phone := j.Get(key).MustString()
		phone = strings.Replace(phone, "#", "", -1)
		if len(phone) >= 11 {
			var mData PhoneData
			var tcode string
			//港澳台
			if strings.HasPrefix(phone, "00") {
				tcode = phone[:5]
			} else {
				tcode = "00" + phone[:3]
			}
			if specialPhonePrefix[tcode] {
				mData := PhoneDataMap[tcode]
				if mData.PhoneCode == "" {
					sp := phone[6:6]
					if sp == "0" {
						sp = phone[6:8]
					} else {
						sp = "0" + phone[6:7]
					}
					mData = PhoneDataMap[sp]
					if mData.PhoneCode == "" {
						mData = PhoneDataMap[sp[:2]]
					}
				}
			}
			mDataMap := map[string]PhoneData{}
			if mData.PhoneCode == "" {
				phonet := phone[len(phone)-11:]
				if MobileDataMap[phonet[:7]].PhonePrefix == "" {
					//7位固话
					phonet = phonet[0:4]
					search := phonet[1:4]
					search = "0" + search
					mData = PhoneDataMap[search]
					//7位号码，4位区号
					if mData.PhoneCode != "" && mData.PhoneLen == 7 {
						mData.Phone = search + phonet[4:]
						mDataMap[search] = mData
					}
					//4位区号已正确，但第一个不是0，需要解析3位区号的情况
					if phonet[:1] != "0" && mData.PhoneCode != "" && mData.PhoneLen == 7 || mData.PhoneCode == "" {
						search = phonet[2:4]
						search = "0" + search
						mData = PhoneDataMap[search]
						//7位号码，3位区号
						if mData.PhoneCode != "" && mData.PhoneLen == 7 {
							mData.Phone = search + phonet[4:]
							mDataMap[search] = mData
						}
					}
					//8位固话
					phonet = phonet[0:3]
					search = phonet
					search = "0" + search
					mData = PhoneDataMap[search]
					//8位号码，4位区号
					if mData.PhoneCode != "" && mData.PhoneLen == 8 {
						mData.Phone = search + phonet[3:]
						mDataMap[search] = mData
					}
					//4位区号已正确，但第一个不是0，需要解析3位区号的情况
					if len(phone) >= 12 && phone[len(phone)-12:len(phone)-11] != "0" && mData.PhoneCode != "" && mData.PhoneLen == 8 || mData.PhoneCode == "" {
						search = phonet[1:3]
						search = "0" + search
						mData = PhoneDataMap[search]
						//8位号码，3位区号
						if mData.PhoneCode != "" && mData.PhoneLen == 8 {
							mData.Phone = search + phonet[3:]
							mDataMap[search] = mData
						}
					}
				} else {
					//手机
					mData = MobileDataMap[phonet[:7]].PhoneData
					mData.Phone = phonet
				}
			} else {
				mData.Phone = phone
			}
			if len(mDataMap) == 1 {
				for _, v := range mDataMap {
					j.Set(key+"_province", v.Province)
					j.Set(key+"_city", v.City)
					j.Set(key+"_areacode", v.AreaCode)
					j.Set(key, v.Phone)
				}
			} else if len(mDataMap) > 1 {
				var arrProvince []string
				var arrCity []string
				var arrAreaCode []string
				var arrPhone []string
				for _, v := range mDataMap {
					arrProvince = append(arrProvince, v.Province)
					arrCity = append(arrCity, v.City)
					arrAreaCode = append(arrAreaCode, v.AreaCode)
					arrPhone = append(arrPhone, v.Phone)
				}
				j.Set(key+"_province", arrProvince)
				j.Set(key+"_city", arrCity)
				j.Set(key+"_areacode", arrAreaCode)
				j.Set(key, arrPhone)
			} else if mData.PhoneCode != "" {
				j.Set(key+"_province", mData.Province)
				j.Set(key+"_city", mData.City)
				j.Set(key+"_areacode", mData.AreaCode)
				j.Set(key, mData.Phone)
			} else {
				//if key == "calleetogatewaye164" {
				//	lock.L.Lock()
				//	record[j.Get(key).MustString()] = true
				//	lock.L.Unlock()
				//}
			}
		}
	}
}
