package main

import (
	"regexp"
	"fmt"
	"github.com/bitly/go-simplejson"
	"strings"
	"github.com/sas/utils"
)

func main() {
	j, err := simplejson.NewJson([]byte(`{
  "hits": [
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3c9cdeb5-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3c9cdeb5-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "lqba终端",
          "terminal_hardware_id": "0abfeadc3bb53af8"
        },
        "terminal_hardware_id": "0abfeadc3bb53af8",
        "user": {},
        "user_id": "513721199909058425"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3c87eb5c-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3c87eb5c-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "lqba终端",
          "terminal_hardware_id": "0abdd40e5feabec8"
        },
        "terminal_hardware_id": "0abdd40e5feabec8",
        "user": {},
        "user_id": "513721199909010207"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3ce15789-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3ce15789-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "iphone 7S tuhao gold",
          "terminal_hardware_id": "0adec81edec4bcff"
        },
        "terminal_hardware_id": "0adec81edec4bcff",
        "user": {},
        "user_id": "513721199909080766"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3c73c279-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3c73c279-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "iphone 7S tuhao gold",
          "terminal_hardware_id": "0aba1d8aecc01ed2"
        },
        "terminal_hardware_id": "0aba1d8aecc01ed2",
        "user": {},
        "user_id": "513721199909081480"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3c91f9d9-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3c91f9d9-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "lqba终端",
          "terminal_hardware_id": "0abe496eeeca1ffe"
        },
        "terminal_hardware_id": "0abe496eeeca1ffe",
        "user": {},
        "user_id": "513721199909013669"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3cc62670-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3cc62670-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "外星人顶配",
          "terminal_hardware_id": "0adb6e6f1ef8847b"
        },
        "terminal_hardware_id": "0adb6e6f1ef8847b",
        "user": {},
        "user_id": "513721199909078425"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3cd5eb31-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3cd5eb31-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {
          "city": "天津市",
          "district": "河北区",
          "org_name": "公安部,天津市公安局,天津市公安局河北分局,天津市公安局河北分局刑事侦查支队",
          "parent": "120105000000",
          "parent_chain": "010000000000,120000000000,120105000000,",
          "province": "天津"
        },
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "宝马终端",
          "terminal_hardware_id": "0adba121efafbbde"
        },
        "terminal_hardware_id": "0adba121efafbbde",
        "user": {
          "police_id": "4d640b0ff4d64570a57ed16c550c5b56",
          "police_type": "025",
          "user_depart": "",
          "user_id": "513721199909081990",
          "user_name": "马英547",
          "user_org": "120105150000"
        },
        "user_id": "513721199909081990"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3ca7307d-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3ca7307d-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943818000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "宝马终端",
          "terminal_hardware_id": "0accdfddce471b8e"
        },
        "terminal_hardware_id": "0accdfddce471b8e",
        "user": {
          "police_id": "176a0ab02c7a4448b2541ed72a9dfb20",
          "police_type": "023",
          "user_depart": "",
          "user_id": "513721199909880424",
          "user_name": "赵本4276",
          "user_org": "15"
        },
        "user_id": "513721199909880424"
      },
      "fields": null,
      "sort": [
        1495943818000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3c2a874f-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3c2a874f-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943817000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "iphone 7S tuhao gold",
          "terminal_hardware_id": "0a5df2aadb94dacb"
        },
        "terminal_hardware_id": "0a5df2aadb94dacb",
        "user": {
          "police_id": "1153435c72df4a379591624571023494",
          "police_type": "007",
          "user_depart": "",
          "user_id": "513721199909991631",
          "user_name": "路易1798",
          "user_org": "50"
        },
        "user_id": "513721199909991631"
      },
      "fields": null,
      "sort": [
        1495943817000
      ]
    },
    {
      "_index": "sas_mdm_terminalsecurity_event_2017_05_28",
      "_type": "terminalsecurity_event",
      "_id": "3bffcd5b-4359-11e7-a1e5-0010f3369cc2",
      "_source": {
        "__app": "mdm",
        "__id": "3bffcd5b-4359-11e7-a1e5-0010f3369cc2",
        "__index": "sas_mdm_terminalsecurity_event_2017_05_28",
        "__time": 1495943817000,
        "__type": "terminalsecurity_event",
        "address": {},
        "body": "超过最大允许连接断开时间",
        "event_code": "event_lost",
        "event_desc": "超过最大允许连接断开时间",
        "event_name": "失联(长时间未登录设备)",
        "is_handled": "0",
        "org": {},
        "terminal": {
          "latitude": 0,
          "longitude": 0,
          "terminal_band": "lqba终端",
          "terminal_hardware_id": "09ecfcdcbde04ccd"
        },
        "terminal_hardware_id": "09ecfcdcbde04ccd",
        "user": {},
        "user_id": "513721199909016215"
      },
      "fields": null,
      "sort": [
        1495943817000
      ]
    }
  ],
  "took": 1,
  "total": 112
}`))
	params := "terminal_hardware_id={_source.terminal_hardware_id}&action=setLockScreen&user_id={_source.user_id}"
	content := "["
	if (err != nil) {
		return
	} else {
		_, bl := j.CheckGet("total")
		//查询
		if bl {
			j = j.Get("hits")
		}
		length := len(j.MustArray())
		for i := 0; i < length; i++ {
			if i > 0 {
				content += ","
			}
			//fmt.Println(j.GetIndex(i).Get("_source").Get("user_id"))
			content += `{"params":"` + GetUrlParameters(params, j.GetIndex(i)) + `"}`
			break
		}
	}
	content += "]"
}

func GetUrlParameters(exp string, js *simplejson.Json) string {
	var reg = regexp.MustCompile(`\{\s*[\w\\.]+\s*\}`)
	r := reg.FindAllStringSubmatch(exp, -1)
	for _, v := range r {
		arr := strings.Split(v[0][1:len(v[0]) - 1], ".")
		nstr := fmt.Sprint(getJsonValue(js, arr).Interface())
		exp = strings.Replace(exp, v[0], nstr, -1)
	}
	return exp
}

func getJsonValue(json *simplejson.Json, paths []string) *simplejson.Json {
	var result *simplejson.Json
	if len(paths) >= 1 {
		if match, err := regexp.Match("\\[\\d*\\]", []byte(paths[0])); match && err == nil {
			//数组
			if strings.HasPrefix(paths[0], "[") {
				result = json.GetIndex(utils.ParseInt(paths[0][1 : len(paths[0]) - 1]))
			} else {
				result = json.Get(paths[0][0:strings.Index(paths[0], "[")]).GetIndex(utils.ParseInt(paths[0][strings.Index(paths[0], "[") + 1 : len(paths[0]) - 1]))
			}
		} else {
			//字段
			result = json.Get(paths[0])
		}
	} else {
		return json
	}
	return getJsonValue(result, paths[1:])
}