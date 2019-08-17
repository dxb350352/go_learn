package main
import (
	"fmt"
	"encoding/xml"
	"strings"
	"errors"
	"github.com/parnurzeal/gorequest"
)

func main() {
	fmt.Println(SendSM("test2", "13540786214"))
}

func SendSM(content, phones string) error {
	request := gorequest.New()
	superAgent := request.Post("http://smsapi.c123.cn/OpenPlatform/OpenApi")
	superAgent.Query("action=sendOnce")
	superAgent.Query("ac=1001@501213780001")
	superAgent.Query("authkey=4D17063A693C1C2DCAE8FD98B7605519")
	superAgent.Query("cgid=52")
	superAgent.Query("csid=")
	superAgent.Query("c=" + content)
	superAgent.Query("m=" + phones)
	//	body := `<xml name="sendOnce" result="1"><Item cid="501213780001" sid="1001" msgid="158830675912720700" total="1" price="0.1" remain="0.300" /></xml>`
	_, body, _ := superAgent.End()
	fmt.Println(body)
	inputReader := strings.NewReader(body)
	decoder := xml.NewDecoder(inputReader)
	t, _ := decoder.Token()
	token := t.(xml.StartElement)
	if token.Name.Local == "xml" {
		for _, v := range token.Attr {
			if v.Name.Local == "result" && v.Value == "1" {
				return nil
			}
		}
	}
	return errors.New("Send message error")
}