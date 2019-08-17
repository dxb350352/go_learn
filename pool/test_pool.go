package main

import (
	"fmt"
	"io"
	"log"
	"sync/atomic"
	"time"
	"testgo/pool/workpool"
	"math/rand"
)

//测试runner
type Arunner struct {
}

func sleep() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}
func (a *Arunner) Run(avar interface{}) (interface{}, error) {
	sleep()
	fmt.Println(avar)
	return nil, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	//tpool1()
	//tpool2()
	tpool3()
}

//创建一个connect
var connectId int32 = 0

type DbConnect struct {
	cid int32
}

func (db *DbConnect) getConnectId() int32 {
	return db.cid
}

func (db *DbConnect) Close() error {
	return nil
}

//创建连接
type DbConnectCreater struct{}

func (dbc *DbConnectCreater) CreateConnect() (io.Closer, error) {

	id := atomic.AddInt32(&connectId, 1)

	return &DbConnect{cid: id}, nil
}

func tpool3() {
	creater := &DbConnectCreater{}
	pool3, err := workpool.NewConnectPool(10, creater)
	if err != nil {
		log.Println("create faild")
	}
	defer pool3.Close()
	//放入10个连接
	for i := 0; i < 10; i++ {
		conect, err := creater.CreateConnect()
		if err == nil {
			pool3.PutConnect(conect)
		}
	}

	for i := 0; i < 100; i++ {
		go func(pool *workpool.ConectPool) {
			connect, err := pool.GetConnect()
			if err == nil {
				sleep()
				log.Println("connect_id=", connect.(*DbConnect).getConnectId())
				pool.PutConnect(connect)
			}
		}(pool3)
	}
	time.Sleep(5 * time.Second)
}

func tpool2() {
	//test pool 2
	pool2 := workpool.NewKworkpoolChan(5, 10)
	for i := 0; i < 100; i++ {
		pool2.AddRunner(workpool.NewWork(&Arunner{}, i))
	}
	pool2.Close()
}

func tpool1() {
	//test pool 1
	pool := workpool.NewKworkpool(20)
	pool.Start()
	for i := 0; i < 100; i++ {
		pool.AddRunner(workpool.NewWork(&Arunner{}, i))
	}
	pool.Close()
	time.Sleep(5 * time.Second)
}
