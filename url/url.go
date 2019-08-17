package main

import (
	"net/url"
	"fmt"
	"github.com/sas/utils"
)

func main() {
	//urlstr := `https://localhost:8888/bmalerts/alerts?fdsaf=fds&aaa=23&keyword=da=dd and fdas=dd&startTime=now-1h&endTime=now`
	//urlstr := `/bmalerts/alerts?fdsaf=fds&aaa=23&keyword=da=dd and fdas=dd&startTime=now-1h&endTime=now`
	urlstr:="baidu.com"
	u, err := url.Parse(urlstr)
	if err != nil {
		fmt.Println(err.Error())
	}

	fmt.Println(u.Fragment, "-Fragment")
	fmt.Println(u.Host, "-Host")
	fmt.Println(u.Opaque, "-Opaque")
	fmt.Println(u.Path, "-Path")
	fmt.Println(u.RawPath, "-RawPath")
	fmt.Println(u.RawQuery, "-RawQuery")
	fmt.Println(u.Scheme, "-Scheme")
	fmt.Println(u.EscapedPath(), "-EscapedPath")
	fmt.Println(u.IsAbs(), "-IsAbs")
	fmt.Println(u.RequestURI(), "-RequestURI")
	fmt.Println(u.RequestURI()==urlstr)
	fmt.Println(u.Query())
	u.Query().Add("a1aa", "1232")
	u.Query().Set("aaa", "1232")
	fmt.Println(u.Query())
	fmt.Println("........................")
	fmt.Println(url.QueryUnescape(`keywords=__index%3Dgaxz*+__notimefilter%3Dtrue+person_id%3D%22sssss%22&from=0&size=10`))
}

func buildUrl(path string) string {
	u, err := url.Parse(path)
	if err != nil {
		fmt.Println(err.Error())
	}
	path = fmt.Sprint(u.Scheme, "://", u.Host, u.Path, "?")
	for k, _ := range u.Query() {
		v := u.Query().Get(k)
		if k == "startTime" || k == "endTime" {
			t, err := utils.ParseTime(v)
			if err != nil {
				v = "0"
			} else {
				v = fmt.Sprint(t.Unix() * 1000)
			}
		}
		path += fmt.Sprint("&", k, "=", v)

	}
	return path
}
