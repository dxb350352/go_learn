package main

import "fmt"

func main() {
	m := map[int]map[int]int{
		1:map[int]int{11:111, 1111:11111},
		2:map[int]int{22:222, 2222:22222},
		3:map[int]int{33:333, 3333:33333},
	}
	for k,v:=range m{
		for kk,vv:=range v{
			if vv%2==0 && vv>333{
				fmt.Println(k,kk)
				delete(m[k],kk)
				if len(m[k])==0{
					delete(m,k)
				}
			}
		}
	}
	fmt.Println(m)
}
