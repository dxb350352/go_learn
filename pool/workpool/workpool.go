package workpool

import (
	"container/list"
	"sync"
	"time"
)

//工作池
type Kworkpool struct {
	rmutex     sync.RWMutex
	mutex      sync.Mutex //写锁
	runnerList *list.List //任务列表
	poolSize   int        //启动的pool的个数
	flag       bool       //是否关闭
}

//新建一个池子
func NewKworkpool(size int) *Kworkpool {
	l := list.New()
	return &Kworkpool{
		poolSize:   size,
		flag:       true,
		runnerList: l,
	}
}

//添加数据
func (kl *Kworkpool) AddRunner(w *Work) {
	//缺点是可以无限添加
	if kl.flag {
		kl.runnerList.PushFront(w)
	}
}

func (kl *Kworkpool) Start() {
	if !kl.flag {
		kl.Close()
		return
	}
	kl.run()
}

//启动多个goroutine
func (kl *Kworkpool) run() {
	for i := 0; i < kl.poolSize; i++ {
		go kl.work()
	}
}

//实际的工作脚本运行pool
func (kl *Kworkpool) work() {
	for {
		//检测数据的长度
		kl.rmutex.RLock()
		listLen := kl.runnerList.Len()
		kl.rmutex.RUnlock()
		if listLen == 0 { //休眠100毫秒
			//已经关闭就结束程序,需要判断是否已经没有任务
			if !kl.flag {
				break
			}
			time.Sleep(time.Millisecond * 100)
			continue
		}
		kl.mutex.Lock()
		elem := kl.runnerList.Back()
		if elem == nil {
			kl.mutex.Unlock()
			continue
		}
		worker := kl.runnerList.Remove(elem).(*Work)
		kl.mutex.Unlock()
		worker.Runner.Run(worker.Args)
	}
}

func (kl *Kworkpool) Close() {
	kl.flag = false
}
