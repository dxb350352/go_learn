package main
import (
	"sync"
	"github.com/bitly/go-simplejson"
	"time"
	"fmt"
)

type Access struct {
	Time  int64 //访问时间
	Count int64 //次数
}
var mutex sync.Mutex
var accessRecode []simplejson.Json = make([]simplejson.Json, 0)
var IpCount map[string]*Access = make(map[string]*Access, 100) //记录一段时间某IP访问次数
var IpLimit map[string]int64 = make(map[string]int64, 100)   //记录哪些IP受限
var AcCount int64 = 50             //一秒内不得超过的访问次数
var AcLimit int64 = 10           //超过后，限制访问时间

func main() {
	for i := 0; i < 100; i++ {
		go gog()
	}
	time.Sleep(time.Second * 11)
}

func gog() {
	fmt.Println("------------------------------------start")
	defer fmt.Println("------------------------------------end")
	ip := "192.168.123.123"
	now := time.Now().Unix()
	limit := IpLimit[ip]
	fmt.Println(IpLimit, IpCount)
	//限制访问
	if limit > now {
		IpLimit[ip] = now + AcLimit
		fmt.Println("跑太快，闪了一下腰")
		return
	}

	access, ok := IpCount[ip]
	if ok && now == access.Time {
		access.Count += 1
		IpCount[ip] = access
		//超过访问次数
		if access.Count >= AcCount {
			IpLimit[ip] = now + AcLimit
		}
	} else {
		IpCount[ip] = &Access{Time: now, Count: 1}
		//删除过期数据
		for k, v := range IpCount {
			if v.Time != now {
				delete(IpCount, k)
			}
		}
		for k, v := range IpLimit {
			if v < now {
				delete(IpLimit, k)
			}
		}
	}

	//记录用户访问应用
	js := simplejson.New()
	js.Set("__time", now * 1000)
	js.Set("ip", ip)
	fmt.Println("accessRecode:len=", len(accessRecode), IpLimit[ip], IpCount[ip])
	mutex.Lock()
	accessRecode = append(accessRecode, *js)
	if len(accessRecode) >= 10 {
		time.Sleep(time.Second)
		fmt.Println("search.indexing...")
		accessRecode = make([]simplejson.Json, 0)
	}
	mutex.Unlock()
	fmt.Println("accessRecode:len=", len(accessRecode),"end")
}