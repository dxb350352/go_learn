package list

import (
	"sync"
	"container/list"
	"time"
	"log"
	"fmt"
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

type Kworkpool struct {
	rmutex   sync.RWMutex
	mutex    sync.Mutex
	runList  *list.List
	poolSize int
	flag     bool
}

func NewKworkpool(size int) *Kworkpool {
	l := list.New()
	return &Kworkpool{
		poolSize: size,
		flag:     true,
		runList:  l,
	}
}

func (this *Kworkpool) AddRunner(w *Work) {
	if this.flag {
		this.runList.PushBack(w)
	}
}

func (this *Kworkpool) Start() {
	if this.flag {
		this.run()
	}
}

func (this *Kworkpool) run() {
	for i := 0; i < this.poolSize; i++ {
		go this.work()
	}
}

func (this *Kworkpool) work() {
	for {
		this.rmutex.RLock()
		listen := this.runList.Len()
		this.rmutex.RUnlock()
		if listen == 0 {
			if !this.flag {
				break
			}
			time.Sleep(time.Millisecond * 100)
			continue
		}
		this.mutex.Lock()
		elem := this.runList.Front()
		if elem == nil {
			this.mutex.Unlock()
			continue
		}
		worker := this.runList.Remove(elem).(*Work)
		this.mutex.Unlock()
		worker.Runnable.Run(worker.Args)
	}
}

func (this *Kworkpool) Close() {
	this.flag = false
}
func init() {
	log.SetFlags(log.LstdFlags)
}

///////////////////////////////////////////////////////
type TestRun struct {
}

func (t *TestRun) Run(it interface{}) (interface{}, error) {
	i := it.(int)
	time.Sleep(time.Millisecond * 100)
	fmt.Println(i)
	return nil, nil
}

func main() {
	pool := NewKworkpool(2)
	pool.Start()
	for i := 0; i < 100; i++ {
		pool.AddRunner(NewWork(&TestRun{}, i))
	}
	time.Sleep(time.Hour)
}
