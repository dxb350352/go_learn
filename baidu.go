package main
import (
	"github.com/parnurzeal/gorequest"
	"fmt"
	"github.com/bitly/go-simplejson"
	"net/http"
	"io/ioutil"
)

func main() {
	httpGet()
	fmt.Println("..........................")
	request := gorequest.New()
	_, body, errs := request.Get("http://api.map.baidu.com/geocoder/v2/?output=json").Query(fmt.Sprint("ak=", "03KNhoTOqG60CIOErEw08qkSAIqDplrg")).Query(fmt.Sprint("location=", 39.983424, ",", 116.322987)).End()
	if errs != nil {
		fmt.Println(errs)
		return
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	if j2.Get("status").MustInt() == 0 {
		fmt.Println(j2.GetPath("result", "addressComponent"))
	}
}

func httpGet() {
	resp, err := http.Get("http://api.map.baidu.com/geocoder/v2/?output=json&ak=03KNhoTOqG60CIOErEw08qkSAIqDplrg&" + fmt.Sprint("location=", 22.67358, ",", 114.01886))
	if err != nil {
		fmt.Println(err)
	}

	defer resp.Body.Close()
	body, err := ioutil.ReadAll(resp.Body)
	if err != nil {
		fmt.Println(err)
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		fmt.Println(err)
		return
	}
	if j2.Get("status").MustInt() == 0 {
		b, _ := j2.GetPath("result", "addressComponent").Encode()
		fmt.Println(string(b))
	}
}
