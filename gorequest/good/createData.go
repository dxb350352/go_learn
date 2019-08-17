package main

import (
	"time"
	"github.com/parnurzeal/gorequest"
	"crypto/tls"
	"strings"
	"fmt"
	"github.com/bitly/go-simplejson"
	"github.com/sas/utils"
)

var url = "http://192.168.130.201:9005/search/bulk"
//var url = "https://192.168.130.201:9005/search/bulk?by=month"

func main() {
	saveEveCables(`{"_key":"420527197812010018", "name":"董爱华", "tel":"13511111111", "landline":"02888888888","address":"成都市科环北路", "city":"成都"}`)
	saveEntIdCards(`{"_key":"420527197812010018"}`)
	saveEntPersons(`{"name":"董爱华", "id":"420527197812010018", "tels":["13511111111"], "landlines":["02888888888"],"addresses":["成都市科环北路"], "cities":["成都"]}`)
	saveEntTels(`{"_key":"13511111111"}`)
	saveEntLandlines(`{"_key":"02888888888"}`)

	saveEveCables(`{"_key":"520527197901010018", "name":"喻波", "tel":"13611111111", "landline":"01000000000","address":"北京市东城区", "city":"北京"}`)
	saveEntIdCards(`{"_key":"520527197901010018"}`)
	saveEntPersons(`{"name":"喻波", "id":"520527197901010018", "tels":["13611111111"], "landlines":["01000000000"], "addresses":["北京市东城区"], "cities":["北京"]}`)
	saveEntTels(`{"_key":"13611111111"}`)
	saveEntLandlines(`{"_key":"01000000000"}`)

	saveEveCables(`{"_key":"620527196901010018", "name":"王志海", "tel":"13711111111", "landline":"01011111111","address":"北京市西城区", "city":"北京"}`)
	saveEntIdCards(`{"_key":"620527196901010018"}`)
	saveEntPersons(`{"name":"王志海", "id":"620527196901010018", "tels":["13711111111"], "landlines":["01011111111"], "addresses":["北京市西城区"], "cities":["北京"]}`)
	saveEntTels(`{"_key":"13711111111"}`)
	saveEntLandlines(`{"_key":"01011111111"}`)

	saveEveCables(`{"_key":"720527197003010018", "name":"薛晓文", "tel":"13811111111", "landline":"01022222222","address":"北京市海淀区", "city":"北京"}`)
	saveEntIdCards(`{"_key":"720527197003010018"}`)
	saveEntPersons(`{"name":"薛晓文", "id":"720527197003010018", "tels":["13811111111"], "landlines":["01022222222"], "addresses":["北京市海淀区"], "cities":["北京"]}`)
	saveEntTels(`{"_key":"13811111111"}`)
	saveEntLandlines(`{"_key":"01022222222"}`)

	saveEveStudies(`{"_key":"420527200905010018", "name":"董懿萱", "sex":"女", "address":"成都市科环北路", "tel":"13568984989", "school":"贵溪小学"}`)
	saveEntSchools(`{"name":"贵溪小学","_key":"school1"}`)
	saveEntPersons(`{"name":"董懿萱", "sex":"女", "id":"420527200905010018", "tels":["13568984989"], "addresses":["成都市科环北路"], "schools":["贵溪小学"]}`)
	saveEntTels(`{"_key":"13568984989"}`)
	saveEntPersons(`{"name":"王一", "id":"420527198903010018"}`)
	saveEntIdCards(`{"_key":"420527200905010018"}`)
	saveEntIdCards(`{"_key":"420527198903010018"}`)

	saveEveStudies(`{"_key":"520527199904010018", "name":"喻小小", "sex":"男", "address":"北京市东城区", "tel":"13611111111", "school":"北京中学"}`)
	saveEntSchools(`{"name":"北京中学","_key":"school2"}`)
	saveEntPersons(`{"name":"喻小小", "sex":"男", "id":"520527199904010018", "tels":["13611111111"], "addresses":["北京市东城区"], "schools":["北京中学"]}`)
	saveEntPersons(`{"name":"李一", "id":"520527198802010018"}`)
	saveEntIdCards(`{"_key":"520527199904010018"}`)
	saveEntIdCards(`{"_key":"520527198802010018"}`)

	saveEveKuaidis(`{"_key":"12345678","from":{"address":"成都", "name":"董爱华", "tel":"13568984989"},"to":{"address":"北京", "name":"喻波", "tel":"13568984988"},"receiver":{"name":"王小二", "tel":"13534567891", "address":"成都", "company":"顺丰"},"sender":{"name":"李小二", "tel":"13423456781", "address":"北京"},"__time":1493668984000,"updatedAt":1493678984000}`)
	saveEveKuaidis(`{"_key":"22345678","from":{"address":"北京", "name":"喻波", "tel":"13568984988"},"to":{"address":"北京", "name":"王志海", "tel":"13568984987"},"receiver":{"name":"王小三", "tel":"14534567891", "address":"北京", "company":"顺丰"},"sender":{"name":"李小三", "tel":"13523456781", "address":"北京"},"__time":1493678984000,"updatedAt":1493778984000}`)
	saveEveKuaidis(`({"_key":"32345678","from":{"address":"北京", "name":"王志海", "tel":"13568984987"},"to":{"address":"北京", "name":"薛晓文", "tel":"13568984986"},"receiver":{"name":"王小三", "tel":"14534567891", "address":"北京", "company":"顺丰"},"sender":{"name":"李小三", "tel":"13523456781", "address":"北京"},"__time":1493778984000,"updatedAt":1493878984000}`)

	saveEntTels(`{"_key":"13568984988"}`)
	saveEntTels(`{"_key":"13568984987"}`)
	saveEntTels(`{"_key":"13568984986"}`)
}

func saveEveCables(data string) {
	save("eve_cables", "cables", data)
}
func saveEntIdCards(data string) {
	save("ent_id_cards", "id_cards", data)
}
func saveEntPersons(data string) {
	save("ent_persons", "persons", data)
}
func saveEntTels(data string) {
	save("ent_tels", "tels", data)
}
func saveEntLandlines(data string) {
	save("ent_landlines", "landlines", data)
}
func saveEveStudies(data string) {
	save("eve_studies", "studies", data)
}
func saveEntSchools(data string) {
	save("ent_schools", "schools", data)
}
func saveEveKuaidis(data string) {
	save("eve_kuaidis", "kuaidis", data)
}

func save(_app, _type, data string) {
	js, err := simplejson.NewJson([]byte(data))
	if err != nil {
		fmt.Println(err)
		return
	}
	id_pre := utils.GetField2ColumnUpper(_app)
	js.Set("_coll", id_pre)
	_, bl := js.CheckGet("_key")
	id, idbl := js.CheckGet("id")
	if !bl && idbl {
		js.Set("_key", id.MustString())
	}
	js.Set("__time", time.Now().Unix() * 1000)
	jspost := simplejson.New()
	jspost.Set("__app", _app)
	jspost.Set("__type", _type)
	jspost.Set("data", []*simplejson.Json{js})

	request := gorequest.New().Timeout(time.Minute).Post(url)
	if strings.HasPrefix(url, "https") {
		request.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify : true}
	}
	bytes, _ := jspost.Encode()
	fmt.Println(string(bytes))
	_, body, errs := request.SendString(string(bytes)).EndBytes()
	fmt.Println(string(body), errs)
}