package main

import (
	"testgo/cayley/ipgraph/models"
	"github.com/cayleygraph/cayley/query/gizmo"
	"fmt"
	"github.com/cayleygraph/cayley/query"
	"context"
	"github.com/bitly/go-simplejson"
)

func init() {
	//models.Ipgraph(60 * 60 * 10)
	//c := cron.New()
	//c.Start()
	//c.AddFunc("0/10 * * * *", func() {
	//	models.IpgraphCron()
	//})

}

func main() {
	//mapTest()
	mapHasTest()
	//count(true, true)
}

func count(sure, cuurent bool) {
	fmt.Println(models.GetMCGraph(sure, cuurent).Size())
	ses := gizmo.NewSession(models.GetMCGraph(sure, cuurent))
	c := make(chan query.Result, 2)
	qu := `g.V().Tag("source").Out().Tag("dest").Count()`
	ses.Execute(context.TODO(), qu, c, 10)
	for res := range c {
		//spew.Dump(res.Result())
		fmt.Println(res.Result())
	}

}

func mapTest() {
	arr := models.Map("世界", "", 10000, 0, 0, false, true)
	for i, v := range arr {
		fmt.Println(i, v.Get("src_ip"), v.Get("dst_ip"))
	}
	fmt.Println("............................")
	arr = models.Map("中国", "", 10000, 0, 0, false, true)
	for i, v := range arr {
		fmt.Println(i, v.Get("src_ip"), v.Get("dst_ip"))
	}
	fmt.Println("............................")
	arr = models.Map("四川省", "", 10000, 1, 1, false, true)
	for i, v := range arr {
		fmt.Println(i, v.Get("src_ip"), v.Get("dst_ip"))
	}
	fmt.Println("............................")
	arr = models.Map("成都市", "", 10000, 2, 2, false, true)
	for i, v := range arr {
		fmt.Println(i, v.Get("src_ip"), v.Get("dst_ip"))
	}
	fmt.Println("............................")
}

func mapHasTest() {
	printlnArr(models.Map("新加坡", "中国", 10000, 1, 0, false, true))
	printlnArr(models.Map("Singapore", "四川省", 10000, 2, 1, false, true))
	printlnArr(models.Map("世界,新加坡,Singapore,Singapore", "成都市", 10000, 3, 2, false, true))
	printlnArr(models.Map("世界,新加坡,Singapore,Singapore", "中国,四川省,成都市,青羊区", 10000, 3, 3, false, true))

	printlnArr(models.Map("世界,新加坡,Singapore,Singapore", "成都市", 10000, 3, 2, false, true))
}

func printlnArr(arr []*simplejson.Json, iface ...interface{}) {
	for i, v := range arr {
		fmt.Println(i, v.Get("src_ip"), v.Get("dst_ip"))
	}
	fmt.Println(iface...)
}
