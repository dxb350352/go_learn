package main

import (
	"fmt"
	"github.com/mozillazg/go-pinyin"
	"sort"
)

var a = pinyin.NewArgs()

func main() {
	hans := "行人"
	a.Heteronym = true

	fmt.Println(GetAllGroups(hans))
}

func GetAllGroups(str string) []string {
	m := map[string]bool{}
	arr := pinyin.Pinyin(str, a)
	for i, v := range arr {
		if i == 0 {
			for _, vv := range v {
				m[vv] = true
			}
		} else {
			mm := map[string]bool{}
			for kk, _ := range m {
				for _, vv := range v {
					mm[kk+","+vv] = true
				}
			}
			m = mm
		}
		fmt.Println(v, len(v), m)
	}
	var result []string
	for k, _ := range m {
		result = append(result, k)
	}
	sort.Strings(result)
	return result
}
