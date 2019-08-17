package workpool

import (
	"sync"
)

//工作池
type KworkpoolChan struct {
	agroup      sync.WaitGroup
	works       chan *Work //一个chan
	processSize int        //启动的work pool的个数
	flag        bool       //是否关闭
}

//新建一个池子
func NewKworkpoolChan(poolSize, processSize int) *KworkpoolChan {
	pool := &KworkpoolChan{
		works:       make(chan *Work, poolSize),
		processSize: processSize,
		flag:        true,
	}

	pool.Start()
	return pool
}

//添加数据
func (kl *KworkpoolChan) AddRunner(w *Work) {
	//发送数据到chan
	//超出poolSize会卡住
	if kl.flag {
		kl.works <- w
	}
}

func (kl *KworkpoolChan) Start() {
	if !kl.flag {
		kl.Close()
		return
	}
	kl.agroup.Add(kl.processSize)

	for i := 0; i < kl.processSize; i++ {
		go func(num int) {
			//每个线程随机取得worker
			//当close了k1.works才会break出来
			for runWorker := range kl.works {
				runWorker.Runner.Run(runWorker.Args)
			}
			kl.agroup.Done()
		}(i)
	}

}

func (kl *KworkpoolChan) Close() {
	if kl.flag != false {
		kl.flag = false
	}
	close(kl.works)
	//等待所有线程完成任务
	kl.agroup.Wait()
}
