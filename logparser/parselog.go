package main

import (
	"github.com/jackiedong168/sequence"
	"fmt"
	"time"
	"github.com/bitly/go-simplejson"
	"github.com/sas/utils"
	"github.com/surge/xparse/xtime"
	"strings"
)

func main() {
	start := time.Now().UnixNano() / 1e6
	err := sequence.ReadConfig("E:/GOPATH/src/testgo/logparser/sequence.toml");
	printErr(err)
	readTime := time.Now().UnixNano() / 1e6
	fmt.Println(readTime - start)
	parse()
	parse1 := time.Now().UnixNano() / 1e6
	fmt.Println(parse1 - readTime)
	parse2file()
	parse2 := time.Now().UnixNano() / 1e6
	fmt.Println(parse2 - parse1)
}

func printErr(err error) {
	if err != nil {
		fmt.Println(err)
	}
}

func parse() {
	log := `2016-08-21 00:01:43 20.0.18.20 POST /Microsoft-Server-ActiveSync/default.eas Cmd=Ping&User=cdrcb.com%5Cdengmu&DeviceId=androidc537050526&DeviceType=KSAndroid&Log=V141_LdapC1_RpcC28_RpcL31_Hb1680_S3_Error:PingCollisionDetected_Mbx:MAIL1.cdrcb.com_Throttle0_Budget:(A)Conn%3a1%2cHangingConn%3a0%2cAD%3a%24null%2f%24null%2f0%25%2cCAS%3a%24null%2f%24null%2f7%25%2cAB%3a%24null%2f%24null%2f0%25%2cRPC%3a%24null%2f%24null%2f4%25%2cFC%3a1000%2f0%2cPolicy%3aDefaultThrottlingPolicy%5F2b9a3e7f-eda1-4a1b-bf09-ade32bb5dc65%2cNorm_ 443 cdrcb.com\dengmu 117.136.63.27 KSAndroid/5.1.1-EAS-1.3 200 0 0 218`
	scanner := sequence.NewScanner()
	seq, err := scanner.Scan(log)
	printErr(err)
	analyzer := sequence.NewAnalyzer()
	err = analyzer.Add(seq)
	printErr(err)
	aseq, err := analyzer.Analyze(seq)
	fmt.Println(aseq, "...............")
	parser := sequence.NewParser()
	err = parser.Add(aseq)
	printErr(err)
	seq2, err := parser.Parse(seq)
	printErr(err)
	json := Fields(seq2, "app")
	fmt.Println(json)
	for k,v:=range seq2{
		fmt.Println(k,v)
	}
	//err = utils.IndexingBranch("http://192.168.130.240:9005/search/bulk", "syslog", "syslog", []*simplejson.Json{json})
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func parse2file() {
	log := `2016-08-21 00:01:43 20.0.18.20 POST /Microsoft-Server-ActiveSync/default.eas Cmd=Ping&User=cdrcb.com%5Cdengmu&DeviceId=androidc537050526&DeviceType=KSAndroid&Log=V141_LdapC1_RpcC28_RpcL31_Hb1680_S3_Error:PingCollisionDetected_Mbx:MAIL1.cdrcb.com_Throttle0_Budget:(A)Conn%3a1%2cHangingConn%3a0%2cAD%3a%24null%2f%24null%2f0%25%2cCAS%3a%24null%2f%24null%2f7%25%2cAB%3a%24null%2f%24null%2f0%25%2cRPC%3a%24null%2f%24null%2f4%25%2cFC%3a1000%2f0%2cPolicy%3aDefaultThrottlingPolicy%5F2b9a3e7f-eda1-4a1b-bf09-ade32bb5dc65%2cNorm_ 443 cdrcb.com\dengmu 117.136.63.27 KSAndroid/5.1.1-EAS-1.3 200 0 0 218`
	scanner := sequence.NewScanner()
	seq, err := scanner.Scan(log)
	printErr(err)
	analyzer := sequence.NewAnalyzer()
	err = analyzer.Add(seq)
	printErr(err)
	aseq, err := analyzer.Analyze(seq)
	fmt.Println(aseq, "...............")
	log=strings.Replace(aseq.String(),"androidc537050526","%deviceid%",-1)
	fmt.Println(log, "...............")

	seq4parser, err:=scanner.Scan(log)
	printErr(err)

	parser := sequence.NewParser()

	err = parser.Add(seq4parser)
	printErr(err)
	seq2, err := parser.Parse(seq)
	printErr(err)
	json := Fields(seq2, "app")
	fmt.Println(json)
	//err = utils.IndexingBranch("http://192.168.130.240:9005/search/bulk", "syslog", "syslog", []*simplejson.Json{json})
	//if err != nil {
	//	fmt.Println(err)
	//}
}

func Fields(seq sequence.Sequence, app string) *simplejson.Json {
	j := simplejson.New()
	j.Set("__app", app)
	j.Set("__type", "syslog")
	isTimeSet := false
	for _, token := range seq {
		if token.Field != sequence.FieldUnknown {
			field := token.Field.String()
			vtype := token.Type.String()
			if vtype == "integer" {
				j.Set(field, utils.ParseInt(token.Value))
			} else if vtype == "float" {
				j.Set(field, utils.ParseFloat64(token.Value))
			} else {
				j.Set(field, token.Value)
			}
			if !isTimeSet && (field == "msgtime" || field == "date" || token.Type == sequence.TokenTime) {
				t, err := xtime.Parse(token.Value)
				if err == nil {
					j.Set("__time", t.Unix() * 1000)
					isTimeSet = true
				}
			}
		}
	}
	if !isTimeSet {
		j.Set("__time", time.Now().Unix() * 1000)
	}
	return j
}