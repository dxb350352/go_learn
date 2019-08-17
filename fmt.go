package main
import (
	"fmt"
	"errors"
	"strings"
	"os"
	"github.com/sas/utils"
	"time"
	"crypto/rand"
	"encoding/hex"
	"reflect"
	"net/url"
	r "math/rand"
	"runtime"
)

func main() {
	ddddd(1,"2")
	fmt.Printf("err:%# v \n", errors.New("goerrors"))
	fmt.Println(strings.Index("123", "33"))
	ss := []int{1, 2, 3, 4, 5, 6, 7}
	fmt.Println(ss)
	//	remove(1, ss)
	remove(7, ss)

	var mode os.FileMode
	mode = 0775
	fmt.Println(mode | os.ModeDir)
	fmt.Println(os.ModePerm)

	paths := []string{"[123]"}
	fmt.Println(paths[0][1:len(paths[0]) - 1])
	paths = []string{"abc[123]"}
	fmt.Println(paths[0][0:strings.Index(paths[0], "[")])
	fmt.Println(paths[0][strings.Index(paths[0], "[") + 1:len(paths[0]) - 1])


	fmt.Println(strings.TrimSpace("   dfs fds "))

	fmt.Println(utils.EncryptPassword("2eb6e8d62818bd0cc3b34ce134cab93a", ""),"..............MD5")
	fmt.Println(utils.EncryptPassword("wondersoft_123456", ""),"..............MD5")

	f, _ := os.Open("E:\\GOPATH_GO\\src\\github.com\\sas\\sas\\conf\\patterns\\admin_应用_log4j.txt")
	file := f.Name()
	fmt.Println(file[strings.LastIndex(file, "\\") + 1:strings.LastIndex(file, ".")])

	m := make(map[string]string)
	k, re := m["ss"]
	fmt.Println(re, k)

	fmt.Println(time.ParseDuration("1h"))
	fmt.Println(30 * 24 * time.Hour)

	buffer := make([]byte, 32)
	aa, err := rand.Read(buffer)
	fmt.Println(aa, err, buffer)
	access_token := hex.EncodeToString(buffer)
	fmt.Println(access_token)

	fmt.Println(reflect.TypeOf(access_token[0]).Kind())

	urlstr := "http://www.sohu.com:934/dfsd"
	fmt.Println(getPath(urlstr), ".....urlstr")

	fmt.Println(runtime.GOOS, "...........goos")

	fi, err := os.Stat("E:\\GOPATH_GO\\src\\github.com\\sas")
	fmt.Println(fi.IsDir(), fi.Mode().IsDir())

	ssss := strings.TrimSuffix("fdsfds/////", "/")
	ssss = strings.TrimRight("///fds//fsa", "/")
	fmt.Println(ssss)

	var astr []string
	astr = append(astr, "11111")
	astr = append(astr, strings.Split("", ",")...)
	fmt.Println(astr,len(astr))


	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	fmt.Println(r.Intn(999999))
	var in int
	in =123
	fmt.Println(int64(in))
}

func getPath(full  string) string {
	urls, err := url.Parse(full)
	if err != nil {
		return ""
	}
	result := urls.Path
	fmt.Println(result,".....................pppp")
	if i := strings.Index(urls.Host, ":"); i != -1 {
		result = urls.Host[i:] + result
	}
	return result
}

func remove(r int, ss []int) {
	for i, v := range ss {
		if v == r {
			temp := append(ss[:i], ss[i + 1:]...)
			fmt.Println(i, ss[i], temp)
		}
	}
}

func ddddd(a int,b string,c...interface{}) {
	fmt.Println(a)
	fmt.Println(b)
	fmt.Println(len(c),"...........c")
}