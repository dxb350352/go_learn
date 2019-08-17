package models

import "github.com/davecgh/go-spew/spew"

type Loggerer struct {
}

func (l Loggerer) Error(i ... interface{}) {
	spew.Dump(i)
}
func (l Loggerer) Errorf(i ... interface{}) {
	spew.Dump(i)
}
func (l Loggerer) Info(i ... interface{}) {
	spew.Dump(i)
}

var Logger Loggerer
var AppType = []string{
	//"cdrcb|score_netflow|sip|dip|score|",
	"cdrcb|qy_ids|src_ip|dst_ip||",
	//"cdrcb|qy_ads_01|src_ip|dst_ip||",
	//"cdrcb|qy_ads_02|src_ip|dst_ip||",
	//"cdrcb|qy_waf_01|src_ip|dst_ip||",
	//"cdrcb|qy_waf_02|src_ip|dst_ip||",
	//"cdrcb|qy_waf_03|src_ip|dst_ip||",
}
var SearchUrl = "https://192.168.131.114:9005/search/search"
var ShowDataMinutes int64 = 10
var ScoreMax = 0.1
var CombineDataFileNum = 7
var ShowDataSeconds int64 = 30
var IpgraphPath = "D:/ipgraph"
var IpgraphMain = "D:/ipgraph/ipgraph.db"
var IpgraphSureMain = "D:/ipgraph/ipgraph_sure.db"
