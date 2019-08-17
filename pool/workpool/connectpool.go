package workpool

import (
	"errors"
	"io"
	"log"
	"sync"
)

//创建连接的接口
type createConection interface {
	CreateConnect() (io.Closer, error)
}

//一个pool
type ConectPool struct {
	closeChan chan io.Closer
	factory   createConection
	isClose   bool
	mutex     sync.Mutex
}

//新建一个pool
func NewConnectPool(poolSize int, factory createConection) (*ConectPool, error) {
	if poolSize < 0 {
		return nil, errors.New("pool size  < 0")
	}

	return &ConectPool{
		closeChan: make(chan io.Closer, poolSize),
		factory:   factory,
		isClose:   false,
	}, nil
}

//获取连接
func (cp *ConectPool) GetConnect() (io.Closer, error) {
	connect, ok := <-cp.closeChan
	if !ok {
		return nil, errors.New("closeChan is closed")
	}
	return connect,nil
	//select {
	//case connect, ok := <-cp.closeChan:
	//	if !ok {
	//		return nil, errors.New("closeChan is closed")
	//	}
	//	log.Println("fetch connect sucessed")
	//	return connect, nil
	//default:
	//	log.Println("fetch connect faild,create a new one")
	//	return cp.factory.CreateConnect()
	//}
}

//把一个连接放入池子
func (cp *ConectPool) PutConnect(aConnection io.Closer) {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()
	if cp.isClose == true {
		return
	}

	select {
	case cp.closeChan <- aConnection:
		log.Println("put connection into the closer!")
	default:
		log.Println("put faild.close connection!")
		aConnection.Close()
	}

}

//关闭连接
func (cp *ConectPool) Close() {
	cp.mutex.Lock()
	defer cp.mutex.Unlock()
	if cp.isClose != true {
		cp.isClose = true
	}
	close(cp.closeChan)
	for aConnect := range cp.closeChan {
		aConnect.Close()
	}
}
