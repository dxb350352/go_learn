package main

import (
	"github.com/garyburd/redigo/redis"
	"fmt"
)

func main() {
	c, err := redis.Dial("tcp", "192.168.130.201:6379")
	if err!=nil{
		fmt.Println(err)
		return
	}
	r,err:=c.Do("SET","HELLO","WORLD")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(r)
	r,err=c.Do("GET","HELLO")
	if err!=nil{
		fmt.Println(err)
		return
	}
	fmt.Println(r)
}
