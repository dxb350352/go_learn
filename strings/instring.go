package main

import (
	"fmt"
)

func main() {
	fmt.Println(InString("bc", "abc"))
}

func InString(src, target string) bool {
	srcrune := []rune(src)
	targetrune := []rune(target)
	lsrc := len(srcrune)
	ltarget := len(targetrune)
	if lsrc > ltarget {
		return false
	}
	var i, j int
	for ; i < lsrc; i++ {
		var flag bool
		for ; j < ltarget; j++ {
			if srcrune[i] == targetrune[j] {
				flag = true
				j++
				//break了没有j++
				break
			}
		}
		if !flag {
			return false
		}
	}
	return i == lsrc
}
