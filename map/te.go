package main

import (
	"regexp"
	"fmt"
)

var dNumExp = "[0-9〇零一二三四五六七八九十]"
var doorNumExp = regexp.MustCompile(`(` + dNumExp + `{1,3}(([栋幢\-#])|(号楼)))?(` + dNumExp + `{1,3}((单元)|-))?(` + dNumExp + `{1,4}[楼层\-])?` + dNumExp + `{1,5}[号室]?`)

func main() {
	str := "1号1栋1单元1楼"
	arr := doorNumExp.FindAllStringSubmatch(str, -1)
	var key string
	for _, v := range arr {
		fmt.Println(v)
		if v[1] == "" &&v[5] == "" &&v[8] == "" {
			continue
		}
		key = v[0]
		break
	}
	fmt.Println(key)
}
