package main

import (
	"fmt"
	"time"
	"strconv"
	"math/rand"
	"os"
	"bufio"
	"io"
)

var zone_code []string = make([]string, 0)

func InitZoneCode() {

	f, err := os.OpenFile("E:/GOPATH/src/github.com/sas/tuomin/conf/zone_code.txt", os.O_RDONLY, os.ModePerm)
	if nil != err {
		fmt.Println("openfile err:", err)
		return
	}
	defer f.Close()
	buff := bufio.NewReader(f) //读入缓存
	for {
		line, err := buff.ReadString('\n') //以'\n'为结束符读入一行
		if err != nil || io.EOF == err {
			break
		}
		zone_code = append(zone_code, line[:6])
	}
	//fmt.Println("size:", len(zone_code))
	rand.Seed(time.Now().Unix())
	//	fmt.Println(time.Unix(time.Now().Unix()-rand.Int()%(90*365*24*3600), 0).Format("20060102"))
}

type IDCard struct{}

func (d *IDCard) MakeValue(data interface{}) interface{} {
	size_zone := len(zone_code)
	code := zone_code[rand.Int() % size_zone] + time.Unix(time.Now().Unix() - int64(rand.Int() % (90 * 365 * 24 * 3600)), 0).Format("20060102") + fmt.Sprintf("%03d", rand.Int() % 1000)
	weight := []int{7, 9, 10, 5, 8, 4, 2, 1, 6, 3, 7, 9, 10, 5, 8, 4, 2}
	sum := 0
	for iLoop := 0; iLoop < 17; iLoop++ {
		v, _ := strconv.Atoi(code[iLoop:iLoop + 1])
		//fmt.Println(iLoop, ",", v, ",", v * weight[iLoop])
		sum += v * weight[iLoop]
	}
	//fmt.Println("sum:", sum)
	modValue := sum % 11
	modValue = (12 - modValue) % 11
	if modValue < 10 {
		code += strconv.FormatInt(int64(modValue), 10)
	} else {
		code += "X"
	}
	return code
}
func main() {
	InitZoneCode()
	var mv IDCard
	fmt.Println(mv.MakeValue(nil))
}
