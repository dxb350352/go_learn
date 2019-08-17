package main

import (
	"net/http"
	"strings"
	"io/ioutil"
	"fmt"
	"net/http/cookiejar"
)

func main() {
	gCurCookieJar, _ = cookiejar.New(nil)
	//Soap11("https://lisea.cn/test")
	Soap12("http://localhost:8080/axis2/services/es_service?wsdl")
	Soap12("http://localhost:8080/axis2/services/es_service?wsdl")
	Soap12("http://localhost:8080/axis2/services/es_service?wsdl")
	Soap12("http://localhost:8080/axis2/services/es_service?wsdl")
}

func Soap11(url string) {
	reqBody := `<soap:Envelope xmlns:xsi="http://www.w3.org/2001/XMLSchema-instance"
  xmlns:xsd="http://www.w3.org/2001/XMLSchema"
  xmlns:soap="http://schemas.xmlsoap.org/soap/envelope/">
  <soap:Body>
    <GetSpeech xmlns="http://xmlme.com/WebServices">
      <Request>string</Request>
    </GetSpeech>
  </soap:Body>
</soap:Envelope>`

	res, err := http.Post(url, "text/xml; charset=UTF-8", strings.NewReader(reqBody))
	if nil != err {
		fmt.Println("http post err:", err)
		return
	}
	defer res.Body.Close()

	// return status
	if http.StatusOK != res.StatusCode {
		fmt.Println("WebService soap1.1 request fail, status: %s\n", res.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if nil != err {
		fmt.Println("ioutil ReadAll err:", err)
		return
	}

	fmt.Println("webService soap1.1 response: ", string(data))
}

var gCurCookieJar *cookiejar.Jar
// soap1.2例子
func Soap12(url string) {
	reqBody := `<env:Envelope
					xmlns:env="http://www.w3.org/2003/05/soap-envelope">
					<env:Header/>
					<env:Body>
						<axis:monitor_call
							xmlns:axis="http://service">
							<token>d76d09b6bd1da892e4489a3513b4b5a7</token>
							<StartRow>0</StartRow>
							<Return_rows>100</Return_rows>
							<Call_Duration>30</Call_Duration>
							<Start_time>2017-08-23 00.00.00</Start_time>
							<End_time>2018-08-23 23.59.00</End_time>
						</axis:monitor_call>
					</env:Body>
				</env:Envelope>`
	client := http.Client{
		Jar: gCurCookieJar,
	}
	res, err := client.Post(url, "application/soap+xml; charset=utf-8", strings.NewReader(reqBody))
	if nil != err {
		fmt.Println("http post err:", err)
		return
	}
	defer res.Body.Close()

	// return status
	if http.StatusOK != res.StatusCode {
		fmt.Println("WebService soap1.2 request fail, status: %s\n", res.StatusCode)
		return
	}

	data, err := ioutil.ReadAll(res.Body)
	if nil != err {
		fmt.Println("ioutil ReadAll err:", err)
		return
	}

	fmt.Println("webService soap1.2 response: ", string(data))
}
