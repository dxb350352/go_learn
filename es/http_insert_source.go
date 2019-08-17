package main

import (
	"github.com/bitly/go-simplejson"
	"log"
	"net/http"
	"fmt"
	"strings"
)

var jstr = `
[{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_1009208996_1043255989_1041668472","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"林生客气so","person_phone":"13302731177","seq":1041668472,"update_time":1530834014210,"userid":"025_448bde2ac01fbd4183563d18eda3e3bb"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_7003860656_7043407160_7008970998","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"澄树凯","person_phone":"13302731177","seq":7008970998,"update_time":1530877231598,"userid":"025_e74caf9bdaa89af0d1185b7746bf5ca4"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_16000966692_16029823740_16013895135","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"中旅经理","person_phone":["13302731177","13902731177"],"seq":16013895135,"update_time":1531146287903,"userid":"025_c92b60827641a76859becec610c7bf6d"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_12019799761_12008316769_12020158668","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"120门客","person_phone":"13302731177","seq":12020158668,"update_time":1531130223029,"userid":"025_8678b0f63db5d6eac8e066579ee25f51"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_18001453431_18023655180_18001501657","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"澄海顺安国旅林树楷","person_phone":["075483660666","075483660777","13302731177","13902731177"],"seq":18001501657,"update_time":1531163753915,"userid":"025_84b065d339dfc92bf775c56713c4037d"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_18014534310_18031418332_18029630652","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"林树楷","person_phone":["075488784799","13302731177","13902731177"],"seq":18029630652,"update_time":1531170907924,"userid":"025_3a156f85329046b7a7441b112759fe78"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_20008728416_20000513626_20008754555","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"中旅经理","person_phone":["13302731177","13902731177"],"seq":20008754555,"update_time":1531483659892,"userid":"025_c472ba3b87dd1992a65b8f73c90649d7"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_20002909472_20016534744_20011859237","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"树楷新电话","person_phone":"13302731177","seq":20011859237,"update_time":1531478191299,"userid":"025_c9e9a67ddcb7fbefb72a202c60c7e96e"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_28000000000_28042773198_28000350674","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"中旅楷2","person_phone":["13302731177","13318054321"],"seq":28000350674,"update_time":1531540405398,"userid":"025_5a47281a77954e89522b9a374a21630a"},"fields":null,"sort":[1530374400000]},{"_index":"nbo_person_2018_01","_type":"person","_id":"025_002_2_24001453932_24005224621_24001729842","_source":{"__time":1530374400000,"event_type":"025_002","multi":"2","person_nick":"中旅楷2","person_phone":["13302731177","13318054321"],"seq":24001729842,"update_time":1531511781825,"userid":"025_69c6a72f875dc49958eab79843c502fb"},"fields":null,"sort":[1530374400000]}]
`

var FieldToTypeMap = map[string][]string{
	"attention": {"phone"},
	"person": {
		"person_phone",
	},
	"id": {
		"person_id",
		"person_bankcard_id",
		"phone_id_imei5",
		"phone_id_imsi5",
		"vehicle_id",
		"person_id_appuser",
		"phone_id_mac",
		"phone_id_iccid5",
	},
	"address": {
		"person_location_desc",
	},
	"hardware": {
		"device_imei",
		"device_mac",
	},
	"webid": {
		"person_weixin",
		"person_qq",
		"person_email",
		"person_weibo_id",
		"person_taobao_id",
		"person_alipay_id",
		"person_loginname",
		"person_uid",
	},
	"webgroup": {
		"object_groupid",
	},
	"contacts": {
		"userid",
	},
}

func main() {
	j, err := simplejson.NewJson([]byte(jstr))
	if err != nil {
		log.Fatal(err)
	}
	for i := 0; true; i++ {
		fmt.Println(i)
		jj := j.GetIndex(i)
		if jj.Interface() == nil {
			break
		}

		__index := jj.Get("_index").MustString()
		__type := jj.Get("_type").MustString()
		id := jj.Get("_id").MustString()
		jj = jj.Get("_source")
		jj.Set("id", id)
		urlstr := fmt.Sprintf("http://192.168.130.201:9200/%s/%s/%s", __index, __type, id)
		bys, err := jj.MarshalJSON()
		if err != nil {
			fmt.Println(err)
			continue
		}
		for _, field := range FieldToTypeMap[__type] {
			ros := getArr(jj.Get(field))
			if len(ros) == 0 {
				fmt.Println(id, field)
				continue
			}
			for _, routing := range ros {
				resp, err := http.Post(urlstr+"?routing="+routing, "application/json; charset=utf-8", strings.NewReader(string(bys)))
				if err != nil {
					fmt.Println(err)
					continue
				}
				fmt.Println(id, field, routing, resp.Status)
			}
		}
	}
}

func getArr(j *simplejson.Json) []string {
	s := j.MustString()
	if s != "" {
		return []string{s}
	}
	return j.MustStringArray()
}
