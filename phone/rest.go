package main

import (
	"crypto/tls"
	"fmt"
	"github.com/jackiedong168/gorequest"
	"strings"
)

var JuheAffiliationUri = `http://op.juhe.cn/onebox/phone/query`
var JuheAffiliationAppkey = `553cd38b7cf891d85ceca7d975049b8b`

func main() {
	phone:="13022142605"
	phone="13725181718"
	phone="13131313131"
	GetPhoneLocationJuhe(phone)
}

func GetPhoneLocationJuhe(phone string) {
	RequestGo := gorequest.New()
	if strings.HasPrefix(JuheAffiliationUri, "https") {
		RequestGo.Transport.TLSClientConfig = &tls.Config{InsecureSkipVerify: true}
	}
	RequestGo.Get(JuheAffiliationUri)
	RequestGo.Query(fmt.Sprintf("tel=%s", phone))
	RequestGo.Query("dtype=json")
	RequestGo.Query(fmt.Sprintf("key=%s", JuheAffiliationAppkey))
	_, body, errs := RequestGo.End()
	fmt.Println(string(body), errs)

}
