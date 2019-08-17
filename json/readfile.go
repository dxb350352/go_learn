package main
import (
	"io/ioutil"
	js "github.com/bitly/go-simplejson"
	"fmt"
)

func main() {
	dat, _ := ioutil.ReadFile("E:/GOPATH_GO/src/github.com/sas/sasslvpn_web/conf/config.json")
	VPNConfig, _ := js.NewJson(dat)
	fmt.Println(VPNConfig == nil)
	servers := VPNConfig.Get("services")
	length := len(servers.MustArray())
	fmt.Println(length)
	for i := 0; i < length; i++ {
		fmt.Println(servers.GetIndex(i).Get("AppUrl").MustString())
		fmt.Println(servers.GetIndex(i).Get("AppName").MustString())
		fmt.Println(servers.GetIndex(i).Get("AppIp").MustString())
	}

}
