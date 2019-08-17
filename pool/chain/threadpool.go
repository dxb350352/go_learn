package main

import (
	"sync"
	"time"
	"log"
	"fmt"
	"math/rand"
)

type Runnable interface {
	Run(interface{}) (interface{}, error)
}

type Work struct {
	Runnable Runnable
	Args     interface{}
}

func NewWork(run Runnable, args interface{}) *Work {
	return &Work{run, args}
}

type KworkpoolChan struct {
	wgroup   sync.WaitGroup
	works    chan *Work
	poolSize int
	flag     bool
}

func NewKworkpoolChan(size int) *KworkpoolChan {
	pool := &KworkpoolChan{
		poolSize: size,
		flag:     true,
		works:    make(chan *Work, size),
	}
	pool.Start()
	return pool
}

func (this *KworkpoolChan) AddRunner(w *Work) {
	if this.flag {
		start := time.Now().UnixNano()
		this.works <- w
		fmt.Println("add", getTime(), w.Args, time.Now().UnixNano()-start, len(this.works))
	}
	sleep()
}

func (this *KworkpoolChan) Start() {
	if this.flag {
		this.wgroup.Add(this.poolSize)
		for i := 0; i < this.poolSize; i++ {
			go func(num int) {
				for runWorker := range this.works {
					runWorker.Runnable.Run(runWorker.Args)
					fmt.Println("run", getTime(), num, runWorker.Args)
				}
				this.wgroup.Done()
			}(i)
		}
	}
}

func (this *KworkpoolChan) Close() {
	this.flag = false
	close(this.works)
	this.wgroup.Wait()
}
func init() {
	log.SetFlags(log.LstdFlags)
}

///////////////////////////////////////////////////////
type TestRun struct {
}

func sleep() {
	time.Sleep(time.Millisecond * time.Duration(rand.Intn(100)))
}

func getTime() int64 {
	result := time.Now().UnixNano()
	result = result - result/1e10*1e10
	return result / 1e6
}
func (t *TestRun) Run(it interface{}) (interface{}, error) {
	//i := it.(int)
	sleep()
	//fmt.Println(i)
	return nil, nil
}

func main() {
	rand.Seed(time.Now().UnixNano())
	pool := NewKworkpoolChan(1)
	for i := 0; i < 100; i++ {
		pool.AddRunner(NewWork(&TestRun{}, i))
	}
	pool.Close()
}
