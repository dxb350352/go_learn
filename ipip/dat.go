package main

import (
	"github.com/17mon/go-ipip/ipip"
	"fmt"
	"time"
	"math/rand"
	"os"
	"log"
)

type IPIPLocation struct {
	Country   string
	Province  string
	City      string
	District  string
	ChinaCode string
}

func main() {
	datpath := "E:/GOPATH/src/testgo/ipip/17monipdb.dat"
	log.SetFlags(log.LstdFlags | log.Lshortfile)

	var s = "123.121.18.79"

	s = "118.8.8.8"

	disDb := ipip.NewDistrictDB()
	if e := disDb.Load(datpath); e != nil {
		// load db file failed
	}

	cityDb := ipip.NewCityDB()
	if e := cityDb.Load(datpath); e != nil {
		// load db file failed
	}

	var loc IPIPLocation
	dis, err := disDb.Find(s)
	if err == nil {
		loc.Country = dis.Country
		loc.Province = dis.Province
		loc.City = dis.City
		loc.District = dis.District
		loc.ChinaCode = dis.Code
	} else if err == ipip.DistrictEmptyError {
		if city, err := cityDb.Find(s); err == nil {
			loc.Country = city.Country
			loc.Province = city.Province
			loc.City = city.City
			loc.ChinaCode = city.ChinaCode
		}
	} else {
		log.Fatal(err)
	}

	fmt.Println(loc)

	os.Exit(0)
	disTest()
	cityTest()
}

func disTest() {
	dis := ipip.NewDistrictDB()
	dis.Load("c:/WORK/tiantexin/framework/library/ip/quxian.datx")

	loc, err := dis.Find("123.121.18.79")
	if err == nil {
		fmt.Println(loc)
	}
}

func cityTest() {
	ipip.Load("C:/WORK/tiantexin/17mon/mydata4vipday2.datx")
	var now time.Time
	/*
	now = time.Now()
	var wg sync.WaitGroup
	wg.Add(1000000)
	for i := 0; i < 1000000; i++ {
		go func(){
			ipip.Find(randomIp())
			wg.Done()
		}()
	}
	wg.Wait()
	fmt.Println(time.Now().Sub(now))
	now = time.Now()
	for i := 0; i < 1000000; i++ {
		ipip.Find2(randomIp())
	}
	fmt.Println(time.Now().Sub(now))
*/
	now = time.Now()
	for i := 0; i < 1000000; i++ {
		//ipip.FindLocation(randomIp())
	}
	fmt.Println(time.Now().Sub(now))

	fmt.Println(ipip.FindLocation("1.0.140.156"))
	fmt.Println(ipip.Find2("1.0.206.85"))
	fmt.Println(ipip.Find2("8.8.8.8"))
	fmt.Println(ipip.Find2("118.28.8.8"))
}

func randomIp() string {
	return fmt.Sprintf("%d.%d.%d.0", rand.Intn(255), rand.Intn(255), rand.Intn(255))
}
