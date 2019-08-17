package main

import "fmt"
import "errors"
import "math/rand"
import "time"
import "strconv"
import "github.com/bitly/go-simplejson"
import "github.com/parnurzeal/gorequest"

type Police struct {
	CertSN         string
	IndexId        int
	Loc            []float64
	PoliceID       string
	TeminalId      string
	TermOnlineTime time.Time
	TerminalHwId   string
	__time         int64
	__type         string
	address        *simplejson.Json
	change_status  bool
	duration       int
	online         bool
	police_number  string
	police_type    string
	user_depart    string
	user_name      string
	user_org       string
	user_org_name  string
}

var policeArr []Police

/**
@brief 设定速度：一分钟最多在横纵坐标上移动1公里也就是0.54'
*/
func getNextCoordinate1() (float64, float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//生成纬度
	latitude := float64(r.Intn(540)) / 1000 / 60

	//生成经度
	longitude := float64(r.Intn(540)) / 1000 / 60

	return latitude, longitude
}

/**
@brief 设定速度：一分钟最多在横纵坐标上移动1公里也就是0.54'
*/
func getNextCoordinate(latitude float64, longitude float64) (float64, float64) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))

	//生成纬度
	latitude += ((float64(r.Intn(540 * 2)) / 1000) - 0.54) / 60

	//生成经度
	longitude += ((float64(r.Intn(540 * 2)) / 1000) - 0.54) / 60

	return latitude, longitude
}

func getAddress(latitude float64, longitude float64) (*simplejson.Json, error) {
	request := gorequest.New()
	request.Header["Connection"] = "close"
	reponse, body, errs := request.Get("http://api.map.baidu.com/geocoder/v2/?output=json").Query(fmt.Sprint("ak=", "03KNhoTOqG60CIOErEw08qkSAIqDplrg")).Query(fmt.Sprint("location=", latitude, ",", longitude)).End()
	if errs != nil {
		fmt.Println("xxxx", errs)
		return nil, errs[0]
	}
	reponse.Body.Close()

	j2, err := simplejson.NewJson([]byte(body))
	//fmt.Println("body = ", string(body))
	if err != nil {
		fmt.Println("yyyy", err)
		return nil, err
	}

	var address *simplejson.Json
	if j2.Get("status").MustInt() == 0 {
		address = j2.GetPath("result", "addressComponent")
	}

	return address, nil
}


func move(i int) error {
	latitude, longitude := getNextCoordinate(policeArr[i].Loc[0], policeArr[i].Loc[1])
	policeArr[i].Loc[0] = latitude;
	policeArr[i].Loc[1] = longitude;

	tm := policeArr[i].TermOnlineTime.Add(time.Second * 60)
	policeArr[i].TermOnlineTime = tm
	policeArr[i].__time = tm.Unix() * 1000

	/*address, err := getAddress(latitude, longitude)
	if address != nil && err == nil {
		policeArr[i].address = address
		fmt.Println("address = ", policeArr[i].address)
	}else{
		return err
	}*/

	data := simplejson.New()
	data.Set("CertSN", policeArr[i].CertSN)
	data.Set("IndexId", policeArr[i].IndexId)
	data.Set("Loc", []float64{policeArr[i].Loc[1], policeArr[i].Loc[0]})
	data.Set("PoliceID", policeArr[i].PoliceID)
	data.Set("TeminalId", policeArr[i].TeminalId)
	data.Set("TermOnlineTime", policeArr[i].TermOnlineTime)
	data.Set("TerminalHwId", policeArr[i].TerminalHwId)
	data.Set("__time", policeArr[i].__time)
	data.Set("__type", policeArr[i].__type)//
	//data.Set("address", 		policeArr[i].address)//
	data.Set("change_status", policeArr[i].change_status)//
	data.Set("duration", policeArr[i].duration) //
	data.Set("online", policeArr[i].online) //
	data.Set("police_number", policeArr[i].police_number)
	data.Set("police_type", policeArr[i].police_type)//
	data.Set("user_depart", policeArr[i].user_depart)//
	data.Set("user_name", policeArr[i].user_name)
	data.Set("user_org", policeArr[i].user_org)//
	data.Set("user_org_name", policeArr[i].user_org_name)//

	dataArr := make([]simplejson.Json, 0)
	dataArr = append(dataArr, *data)

	send := simplejson.New()
	send.Set("data", dataArr)
	send.Set("__app", "szga_monitorsvr")
	send.Set("__type", "t_termonline")

	fmt.Println(send)

	bytes, err := send.Encode()
	if err != nil {
		return err
	}
	request := gorequest.New()
	_, body, errs := request.Post("http://192.168.130.240:9006/search/indexing").Send(string(bytes)).End()
	if errs != nil {
		return errors.New(fmt.Sprintln(errs))
	}
	j2, err := simplejson.NewJson([]byte(body))
	if err != nil {
		return errors.New("not return valid json")
	}
	if j2.Get("status").MustInt() != 200 {
		return errors.New(body)
	}
	return nil
}

func initPolice(latitude float64, longitude float64, num int) {
	r := rand.New(rand.NewSource(time.Now().UnixNano()))
	tm := time.Now()
	var police Police

	for i := 0; i < num; i++ {
		police.CertSN = fmt.Sprintf("%06d", 030303 + i)
		police.IndexId = 13 + i
		police.Loc = []float64{latitude + float64(r.Intn(540 * 20)) / 1000 / 60, longitude + float64(r.Intn(540 * 20)) / 1000 / 60}
		police.PoliceID = strconv.Itoa(3 + i)
		police.TeminalId = fmt.Sprintf("%04d", 0003 + i)
		police.TermOnlineTime = tm
		police.TerminalHwId = strconv.Itoa(9875899495895 + i)
		police.__time = tm.Unix() * 1000;
		police.__type = "t_termonline"
		police.change_status = false
		police.duration = 60000
		police.online = true
		police.police_number = strconv.Itoa(3 + i)
		police.police_type = "officer"
		police.user_depart = "刑侦大队"
		police.user_name = fmt.Sprint("Sherlock.", i, ".Holmes")
		police.user_org = "8080"
		police.user_org_name = "刑事侦查"

		policeArr = append(policeArr, police)
	}
}

func main() {
	//	var argErr 	bool = false
	var latitude float64 = 34.567
	var longitude float64 = 104.567
	var policeNum int = 100
	var stepNum int = 100
	//	var err 	error
	//	for i := 0; i < 1; i++ {
	//		if (len(os.Args) != 5) {
	//			argErr = true
	//			break
	//		}
	//
	//		latitude, err = strconv.ParseFloat(os.Args[1], 64)
	//		if (err != nil) {
	//			argErr = true
	//			break
	//		}
	//
	//		longitude, err = strconv.ParseFloat(os.Args[2], 64)
	//		if (err != nil) {
	//			argErr = true
	//			break
	//		}
	//
	//		var num int64
	//		num, err = strconv.ParseInt(os.Args[3], 10, 32)
	//		if (err != nil) {
	//			argErr = true
	//			break
	//		}
	//		policeNum = int(num)
	//
	//		var st int64
	//		st, err = strconv.ParseInt(os.Args[4], 10, 32)
	//		if (err != nil) {
	//			argErr = true
	//			break
	//		}
	//		stepNum = int(st)
	//	}
	//
	//	if (argErr == true) {
	//		fmt.Println("Usage: cmd latitude longitude police_num step_num")
	//		return
	//	}

	initPolice(latitude, longitude, policeNum)

	for i := 0; i < stepNum; i++ {
		for j := 0; j < policeNum; j++ {
			move(j)
		}
	}

	//fmt.Println(getNextCoordinate())
}

