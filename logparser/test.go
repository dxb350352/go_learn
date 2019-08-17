package main

import (
	"fmt"
	"github.com/dataence/sequence"
)

func main() {
	m:=map[string]map[string]string{}
	fmt.Println(m["ss"],1)
	fmt.Println(len(m["ss"]["323"]),2)

}

func gogogtest()  {
	log := `2016-08-21 00:01:43 20.0.18.20 POST /Microsoft-Server-ActiveSync/default.eas Cmd=Ping&User=cdrcb.com%5Cdengmu&DeviceId=androidc537050526&DeviceType=KSAndroid&Log=V141_LdapC1_RpcC28_RpcL31_Hb1680_S3_Error:PingCollisionDetected_Mbx:MAIL1.cdrcb.com_Throttle0_Budget:(A)Conn%3a1%2cHangingConn%3a0%2cAD%3a%24null%2f%24null%2f0%25%2cCAS%3a%24null%2f%24null%2f7%25%2cAB%3a%24null%2f%24null%2f0%25%2cRPC%3a%24null%2f%24null%2f4%25%2cFC%3a1000%2f0%2cPolicy%3aDefaultThrottlingPolicy%5F2b9a3e7f-eda1-4a1b-bf09-ade32bb5dc65%2cNorm_ 443 cdrcb.com\dengmu 117.136.63.27 KSAndroid/5.1.1-EAS-1.3 200 0 0 218`
	scanner := sequence.NewScanner()
	seq, err := scanner.Scan(log)
	if err!=nil{
		fmt.Println(err)
		return
	}
	for k,v:=range seq{
		fmt.Println(k,v)
	}
}
