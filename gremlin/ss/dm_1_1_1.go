package main

import (
	"fmt"
	"github.com/jackiedong168/gorequest"
	"github.com/bitly/go-simplejson"
	"log"
	"github.com/go-gremlin/gremlin"
	//"github.com/jackiedong168/gremlin"
	"time"
)

func sendDataTestHandle_md_1_1_1() {
	data := `{"type":"Handle_md_1_1_1","data":{"Address":{"attr":"address","code":"address","value":"N,N"},"BirthDay":"1955-07-10","City":"Port Orchard","Country":"1","Email":{"code":"email","value":"Susurus@hotmail.com"},"FullName":{"code":"name","value":"N N"},"GPS":"47.5066,122.6054","Gender":"Male","Height":"178","ID":"225990","LoginKey":{"code":"pwd","value":"3dcae1653e1360d8e7de63d014b6c1ee"},"NickName":{"code":"user","value":"N"},"Password":{"code":"pwd","value":"$2a$12$MQILH2KXCBLzmp6ifsoKx.nauZ3Y8VdcfleIRTCJ7Wv5i5gWUgQmu"},"Phone":"","SecurityAnswer":"east","SecurityQuestion":"2","State":"50","UserName":{"code":"user","value":"Snuggertunks"},"WorkPhone":"","Zip":"98366","_归属地":["其他"],"lobster_info":{"info_rowid":"699c7055eb5afe3e9121266722ade246","info_table":"lb_info","info_type":"/关联查询/互联网数据/境外网站"}}}`
	request := gorequest.New()
	_, body, errs := request.Post("http://192.168.128.198:9000/gaxz/indexing").SendString(data).End()
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	fmt.Println(body)
}
func sendDataTestHandle_md_2_1_7() {
	data := `{"type":"Handle_md_2_1_7","data":[{
	    "phone": {
		"code": "mobile",
		"attr": "mobile",
		"value": "13227071299"
	    },
	    "QQ": "4586243325",
	    "location": [
		"陕西省"
	    ],
	    "lobster_info": {
		"info_type": "/关联查询/互联网数据/网络社交",
		"info_table": "lb_info",
		"info_rowid": "6160e71d154402bfbd55e7aecd43089d"
	    },
	    "name": {
		"code": "name",
		"attr": "name",
		"value": "余彬"
	    },
	    "Email": "amdin@163.com"
	}]}`
	request := gorequest.New()
	_, body, errs := request.Post("http://192.168.128.198:9000/gaxz/indexing").SendString(data).End()
	if len(errs) > 0 {
		log.Fatal(errs)
	}
	fmt.Println(body)
}
func main() {
	//sendDataTestHandle_md_2_1_7()
	//data := `{"type":"Handle_md_1_1_1","data":{"Address":{"attr":"address","code":"address","value":"N,N"},"BirthDay":"1955-07-10","City":"Port Orchard","Country":"1","Email":{"code":"email","value":"Susurus@hotmail.com"},"FullName":{"code":"name","value":"N N"},"GPS":"47.5066,122.6054","Gender":"Male","Height":"178","ID":"225990","LoginKey":{"code":"pwd","value":"3dcae1653e1360d8e7de63d014b6c1ee"},"NickName":{"code":"user","value":"N"},"Password":{"code":"pwd","value":"$2a$12$MQILH2KXCBLzmp6ifsoKx.nauZ3Y8VdcfleIRTCJ7Wv5i5gWUgQmu"},"Phone":"","SecurityAnswer":"east","SecurityQuestion":"2","State":"50","UserName":{"code":"user","value":"Snuggertunks"},"WorkPhone":"","Zip":"98366","_归属地":["其他"],"lobster_info":{"info_rowid":"699c7055eb5afe3e9121266722ade246","info_table":"lb_info","info_type":"/关联查询/互联网数据/境外网站"}}}`
	//data := `{
	//    "phone": {
	//	"code": "mobile",
	//	"attr": "mobile",
	//	"value": "13227071299"
	//    },
	//    "QQ": "4586243325",
	//    "location": [
	//	"陕西省"
	//    ],
	//    "lobster_info": {
	//	"info_type": "/关联查询/互联网数据/网络社交",
	//	"info_table": "lb_info",
	//	"info_rowid": "6160e71d154402bfbd55e7aecd43089d"
	//    },
	//    "name": {
	//	"code": "name",
	//	"attr": "name",
	//	"value": "余彬"
	//    },
	//    "Email": "amdin@163.com"
	//}`
	//j, err := simplejson.NewJson([]byte(data))
	//if err != nil {
	//	log.Fatal(err.Error(), 1)
	//}
	//err = Handle_md_2_1_7(j)
	//if err != nil {
	//	log.Fatal(err.Error(), 2)
	//}
}
func init() {
	fmt.Println("init..............")
	if err := gremlin.NewCluster("ws://192.168.131.114:8182/gremlin", "ws://192.168.131.115:8182/gremlin", "ws://192.168.131.116:8182/gremlin"); err != nil {
	//if err := gremlin.NewCluster("ws://192.168.1.100:8182/gremlin"); err != nil {
	//if err := gremlin.NewCluster("ws://192.168.132.175:8182/gremlin"); err != nil {
		println(err.Error(), "..")
		return
	}
	st := time.Now().UnixNano()
	by, err := gremlin.Query(`
	p = g.V().has("ent_idcards", "person_id", "person_id_114").limit(1).toList()
	if (p.size() > 0) {
	    p = p.get(0)
	} else {
	    p = graph.addVertex(label, "ent_idcards", "person_id", "person_id_114")
	}
	`).Exec()
	fmt.Println((time.Now().UnixNano() - st) / 1e6)
	if err != nil {
		log.Fatal(err.Error(), "...............init")
	}
	fmt.Println(string(by))
}
func Query(query string) *gremlin.Request {
	return gremlin.Query(query)
}

func QueryExec(query string) (*simplejson.Json, error) {
	data, err := Query(query).Exec()
	if err != nil {
		return nil, err
	}
	return simplejson.NewJson(data)
}

func Handle_md_2_1_7(j *simplejson.Json) error {
	//添加表数据
	Person_Phone := j.Get("phone").Get("value").MustString()
	gql := fmt.Sprint(`graph = JanusGraphFactory.open('conf/janusgraph-cassandra-es.properties')` + "\n")
	gql += fmt.Sprint(`g = graph.traversal()` + "\n")
	gql += fmt.Sprintf(`pphone=graph.addVertex(label, "ent_phones","person_phone","%s")` + "\n", Person_Phone)
	//gql := fmt.Sprintf(`pphone=graph.addVertex(label, "ent_phones","person_phone","%s")` + "\n", Person_Phone)
	Person_Name := j.Get("name").Get("value").MustString()
	gql += fmt.Sprintf(`pname=graph.addVertex(label, "ent_persons","person_name","%s")` + "\n", Person_Name)
	Person_Email := j.Get("Email").MustString()
	gql += fmt.Sprintf(`pemail=graph.addVertex(label, "ent_emails","person_email","%s")` + "\n", Person_Email)
	Person_QQ := j.Get("QQ").MustString()
	gql += fmt.Sprintf(`pqq=graph.addVertex(label, "ent_qqs","person_qq","%s")` + "\n", Person_QQ)
	Person_Location_Desc := j.Get("location").MustArray()
	gql += fmt.Sprintf(`peve=graph.addVertex(label, "evt_dm","_type","md_2_1_7","person_phone","%s","person_name","%s","person_email","%s","person_qq","%s","person_location_desc","%s")` + "\n", Person_Phone, Person_Name, Person_Email, Person_QQ, Person_Location_Desc)
	//添加关系
	gql += fmt.Sprint(`pqq.addEdge("rel_phones", pphone)`, "\n")
	gql += fmt.Sprint(`pqq.addEdge("rel_persons", pname)`, "\n")
	gql += fmt.Sprint(`pqq.addEdge("rel_emails", pemail)`, "\n")
	gql += fmt.Sprint(`pqq.addEdge("evt_dm", peve)`, "\n")
	//gql += fmt.Sprint(`id1 = g.V().has("person_qq", "4586243325").valueMap()`, "\n")
	gql += fmt.Sprint(`id1 = g.V().has("person_qq", "4586243325").valueMap()`, "\n")
	fmt.Println(gql)
	by, err := gremlin.Query(gql).Exec()
	fmt.Println(string(by))
	return err
}