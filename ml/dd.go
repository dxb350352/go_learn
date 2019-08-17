package main

import (
	"time"
	"github.com/bitly/go-simplejson"
	"fmt"
	"github.com/sas/utils"
	//pool "github.com/sas/modules/pool/app/models"
	"crypto/tls"
	"github.com/jackiedong168/gorequest"
)

func main() {
	carete_sparkml_data("","",0)
}

func carete_sparkml_data(dwTypeId1, dwTypeId2 string, second int64) error {
	if dwTypeId1 == "" {
		dwTypeId1 = "28672"
	}
	if dwTypeId2 == "" {
		dwTypeId2 = "28680"
	}
	if second == 0 {
		second = time.Now().Unix() * 1000
	}
	jsonObj := simplejson.New()
	jsonObj.Set("__time", second)
	jsonObj.Set("act_code", "28680")
	jsonObj.Set("area", "四川省,成都市,崇州市")
	jsonObj.Set("bSign", "0000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000000")
	jsonObj.Set("bVerify", utils.EncryptPassword(fmt.Sprint(time.Now().UnixNano()), "salt"))
	jsonObj.Set("bindUserSid", "S-1-5-21-2991239887-2667229774-1802184495-24015")
	jsonObj.Set("cAtt_GUID", "")
	jsonObj.Set("city", "成都市")
	jsonObj.Set("clientGroupId", "67")
	jsonObj.Set("clientName", "shron-PC")
	jsonObj.Set("clientVer", "3.1.2.509(11029)")
	jsonObj.Set("cltGuid", "26CD1C05-BD2D6003-4C060428-81D5913A")
	jsonObj.Set("cltId", "0")
	jsonObj.Set("district", "崇州市")
	jsonObj.Set("dwKey_Count", "2")
	jsonObj.Set("dwServerId", "117441034")
	jsonObj.Set("dwSignMode", "0")
	jsonObj.Set("dwTypeId1", dwTypeId1)
	jsonObj.Set("dwTypeId2", dwTypeId2)
	jsonObj.Set("dwTypeId3", "0")
	jsonObj.Set("dwTypeId4", "0")
	jsonObj.Set("dwTargetId1", "-1")
	jsonObj.Set("dwTargetId2", "-1")
	jsonObj.Set("dwTargetType1", "-1")
	jsonObj.Set("dwTargetType2", "-1")
	jsonObj.Set("dwVerifyMode", "1")
	jsonObj.Set("hdSn", "3SY09JGD094748")
	jsonObj.Set("id", "10")
	jsonObj.Set("ipClt", "192.168.31.24")
	jsonObj.Set("ipStr", "192.168.31.81")
	jsonObj.Set("logLevel", "3")
	jsonObj.Set("logType", "0")
	jsonObj.Set("modifyTime", second)
	jsonObj.Set("oSSessionId", "0")
	jsonObj.Set("opt", "客户端_升级")
	jsonObj.Set("opt_detail", "升级客户端。版本号：%s ->%s")
	jsonObj.Set("osUserName", "SYSTEM")
	jsonObj.Set("osVer", "Windows 7 64")
	jsonObj.Set("province", "四川省")
	jsonObj.Set("timeClt", "1460809492")
	jsonObj.Set("type_code", "4096")
	jsonObj.Set("userId", "0")
	jsonObj.Set("userSid", "S-1-5-21-2991239887-2667229774-1802184495-24015")
	jsonObj.Set("wKey_0", "3.1.2.509(9771)")
	jsonObj.Set("wKey_1", "(9771)")
	jsonObj.Set("wKey_2", "")
	jsonObj.Set("wKey_3", "")
	jsonObj.Set("wKey_4", "")
	jsonObj.Set("wKey_5", "")
	jsonObj.Set("wKey_6", "")
	jsonObj.Set("wKey_7", "")
	jsonObj.Set("wKey_8", "")
	jsonObj.Set("wKey_9", "")
	//return pool.IndexingBranch("https://192.168.130.201:9001/v2/sparkml", "sas_v3", "t_log_test", []simplejson.Json{*jsonObj})
	goReq := gorequest.New()
	rr := goReq.Timeout(time.Minute).Post("https://192.168.130.201:9001/v2/sparkml")
	rr.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	send := simplejson.New()
	send.Set("data", []simplejson.Json{*jsonObj})
	send.Set("__app", "sas_v3")
	send.Set("__type", "t_log_test")
	bytes, _ := send.Encode()
	_, body, errs := rr.SendString(string(bytes)).End()
	fmt.Println(body,errs)
	return nil
}

