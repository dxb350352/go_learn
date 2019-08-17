package main

import (
	"github.com/bitly/go-simplejson"
	"log"
	"net/http"
	"fmt"
	"strings"
)

var jstr = `{"004_602_2_49000000000_49009999999_49003760466_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49000000000_49009999999_49003760466_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49003760466,"update_time":1513143942000},"004_602_2_49000000000_49009999999_49006919872_21":{"__index":"nbo_webid_2013_06","__time":1370044800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49000000000_49009999999_49006919872_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49006919872,"update_time":1513144024000},"004_602_2_49010000000_49019999999_49011963571_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49010000000_49019999999_49011963571_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49011963571,"update_time":1513144156000},"004_602_2_49040000000_49049999999_49041342488_21":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49040000000_49049999999_49041342488_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49041342488,"update_time":1513144903000},"004_602_2_49060000000_49069999999_49060032051_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49060000000_49069999999_49060032051_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49060032051,"update_time":1513145347000},"004_602_2_49070000000_49079999999_49079711218_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49070000000_49079999999_49079711218_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49079711218,"update_time":1513145906000},"004_602_2_49080000000_49089999999_49080593840_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49080000000_49089999999_49080593840_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49080593840,"update_time":1513145931000},"004_602_2_49090000000_49099999999_49091535048_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49090000000_49099999999_49091535048_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49091535048,"update_time":1513146187000},"004_602_2_49090000000_49099999999_49091803417_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49090000000_49099999999_49091803417_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49091803417,"update_time":1513146192000},"004_602_2_49090000000_49099999999_49093501631_21":{"__index":"nbo_webid_2013_06","__time":1370044800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49090000000_49099999999_49093501631_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49093501631,"update_time":1513146227000},"004_602_2_49100000000_49107384607_49103370350_21":{"__index":"nbo_webid_2013_06","__time":1370044800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49100000000_49107384607_49103370350_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49103370350,"update_time":1513146464000},"004_602_2_49100000000_49107384607_49105132236_21":{"__index":"nbo_webid_2014_06","__time":1401580800000,"__type":"webid","event_type":"004_602","id":"004_602_2_49100000000_49107384607_49105132236_21","is_relative":false,"multi":"2","person2_email":"177256686@qq.com","person2_qq":"177256686","person_email":"223243377@qq.com","person_nick":"土-土生金","person_qq":"223243377","seq":49105132236,"update_time":1513146502000},"004_605_1_58460000000_58469999999_58469403284":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_58460000000_58469999999_58469403284","is_relative":false,"multi":"1","object_groupid":"4996507","person_email":"223243377@qq.com","person_qq":"223243377","seq":58469403284,"update_time":1513493735000},"004_605_1_58880000000_58889999999_58885419224":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_58880000000_58889999999_58885419224","is_relative":false,"multi":"1","object_groupid":"15284252","object_groupname":"天与地","person_email":"223243377@qq.com","person_qq":"223243377","seq":58885419224,"update_time":1513507569000},"004_605_1_59280000000_59289999999_59286628412":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59280000000_59289999999_59286628412","is_relative":false,"multi":"1","object_groupid":"2250092","person_email":"223243377@qq.com","person_qq":"223243377","seq":59286628412,"update_time":1513520103000},"004_605_1_59550000000_59559999999_59555050303":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59550000000_59559999999_59555050303","is_relative":false,"multi":"1","object_groupid":"3895336","object_groupname":"龙行天下1群","person_email":"223243377@qq.com","person_qq":"223243377","seq":59555050303,"update_time":1513527539000},"004_605_1_59680000000_59689999999_59687869178":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59680000000_59689999999_59687869178","is_relative":false,"multi":"1","object_groupid":"15284252","person_email":"223243377@qq.com","person_qq":"223243377","seq":59687869178,"update_time":1513531204000},"004_605_1_59700000000_59709999999_59705493022":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59700000000_59709999999_59705493022","is_relative":false,"multi":"1","object_groupid":"37669050","object_groupname":"株潭中学94毕业届","person_email":"223243377@qq.com","person_qq":"223243377","seq":59705493022,"update_time":1513531725000},"004_605_1_59710000000_59719999999_59710038699":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59710000000_59719999999_59710038699","is_relative":false,"multi":"1","object_groupid":"3895336","person_email":"223243377@qq.com","person_qq":"223243377","seq":59710038699,"update_time":1513531858000},"004_605_1_59820000000_59829999999_59827545055":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59820000000_59829999999_59827545055","is_relative":false,"multi":"1","object_groupid":"42368111","object_groupname":"閞杺一刻","person_email":"223243377@qq.com","person_qq":"223243377","seq":59827545055,"update_time":1513535885000},"004_605_1_59910000000_59919999999_59910562179":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59910000000_59919999999_59910562179","is_relative":false,"multi":"1","object_groupid":"2250092","object_groupname":"克","person_email":"223243377@qq.com","person_qq":"223243377","seq":59910562179,"update_time":1513539003000},"004_605_1_59920000000_59929999999_59920406129":{"__index":"nbo_webid_2010_06","__time":1275350400000,"__type":"webid","event_type":"004_605","id":"004_605_1_59920000000_59929999999_59920406129","is_relative":false,"multi":"1","object_groupid":"4996507","object_groupname":"四站连(兄弟情)","person_email":"223243377@qq.com","person_qq":"223243377","seq":59920406129,"update_time":1513539329000},"004_605_1_60320000000_60329999999_60320983926":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_60320000000_60329999999_60320983926","is_relative":false,"multi":"1","object_groupid":"42368111","person_email":"223243377@qq.com","person_qq":"223243377","seq":60320983926,"update_time":1513554828000},"004_605_1_60340000000_60349999999_60345308868":{"__index":"nbo_webid_1970_06","__time":13046400000,"__type":"webid","event_type":"004_605","id":"004_605_1_60340000000_60349999999_60345308868","is_relative":false,"multi":"1","object_groupid":"37669050","person_email":"223243377@qq.com","person_qq":"223243377","seq":60345308868,"update_time":1513556193000},"004_607_1_60150000000_60159999999_60158575279":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_607","id":"004_607_1_60150000000_60159999999_60158575279","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_phone":"13702420842","person_qq":"223243377","seq":60158575279,"update_time":1513311777000},"004_613_1_66360000000_66369999999_66367115453":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_613","id":"004_613_1_66360000000_66369999999_66367115453","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_location_desc":"广东省,佛山市,,","person_nick":"奎星楼","person_qq":"223243377","seq":66367115453,"update_time":1513323568000},"004_618_2_87070000000_87079999999_87077181365_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_87070000000_87079999999_87077181365_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"86372556@qq.com","person_nick":"波赛东","person_qq":"86372556"},"004_618_2_87120000000_87129999999_87122585173_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_87120000000_87129999999_87122585173_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"401789868@qq.com","person_nick":"蓉","person_qq":"401789868"},"004_618_2_87920000000_87929999999_87927632529_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_87920000000_87929999999_87927632529_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"497194288@qq.com","person_nick":"安凌哥","person_qq":"497194288"},"004_618_2_88060000000_88069999999_88069783551_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_88060000000_88069999999_88069783551_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"土土生金","person2_qq":"223243377","person_email":"1440620820@qq.com","person_nick":"超越想象","person_qq":"1440620820"},"004_618_2_88570000000_88579999999_88570817467_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_88570000000_88579999999_88570817467_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"76522135@qq.com","person_nick":"守望的距离","person_qq":"76522135"},"004_618_2_89140000000_89149999999_89140231044_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_89140000000_89149999999_89140231044_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"177256686@qq.com","person_nick":"【紫se茉莉】","person_qq":"177256686"},"004_618_2_89310000000_89319999999_89318213434_22":{"__index":"nbo_webid_2015_06","__time":1433116800000,"__type":"webid","event_type":"004_618","id":"004_618_2_89310000000_89319999999_89318213434_22","is_relative":false,"multi":"2","person2_email":"223243377@qq.com","person2_nick":"奎星楼","person2_qq":"223243377","person_email":"1668260400@qq.com","person_nick":"指点数码售后","person_qq":"1668260400"},"006_005_1_35370000000_35379999999_35375260923":{"__index":"nbo_webid_2011_06","__time":1306886400000,"__type":"webid","event_type":"006_005","id":"006_005_1_35370000000_35379999999_35375260923","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_nick":"金奎","person_qq":"223243377","person_weibo_id":"cjk10000","seq":35375260923,"update_time":1513459044000},"025_001_1_14015933489_14016416321_14016161703":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_14015933489_14016416321_14016161703","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":14016161703,"update_time":1532744534617,"userid":"025_9d122dad6621126288b881bb9581dd66"},"025_001_1_26008723034_26009207646_26009017742":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_26008723034_26009207646_26009017742","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":26009017742,"update_time":1532746906113,"userid":"025_d547c268e5959ac805c9a59601e4db5a"},"025_001_1_30005357583_30005844635_30005389476":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_30005357583_30005844635_30005389476","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":30005389476,"update_time":1532747789885,"userid":"025_c2a927b4e03fc4f68cb8bf67a0ac2cee"},"025_001_1_42008716410_42009200654_42008734013":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_42008716410_42009200654_42008734013","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":42008734013,"update_time":1532750772615,"userid":"025_6dc1965b522e46d3860dfcf5d6ea5351"},"025_001_1_5010173030_5010657459_5010192089":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_5010173030_5010657459_5010192089","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":5010192089,"update_time":1532751671216,"userid":"025_ba10698c7356e6bd829dd224245e2de8"},"025_001_1_52004498191_52004997989_52004811229":{"__index":"nbo_webid_2018_01","__time":1530374400000,"__type":"webid","event_type":"025_001","id":"025_001_1_52004498191_52004997989_52004811229","is_relative":false,"multi":"1","person_email":"223243377@qq.com","person_qq":"223243377","seq":52004811229,"update_time":1533286220358,"userid":"025_8611b4352505ba8e3c4b024811d25efc"}}
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
	m := j.MustMap()
	for key, _ := range m {
		fmt.Println(key)
		jj := j.Get(key)
		if jj.Interface() == nil {
			break
		}

		__index := jj.Get("__index").MustString()
		__type := jj.Get("__type").MustString()
		id := jj.Get("id").MustString()
		jj.Del("__index")
		jj.Del("__type")
		jj.Del("id")
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
