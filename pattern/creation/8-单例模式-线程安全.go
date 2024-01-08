package main

import (
	"fmt"
	"sync"
	"sync/atomic"
)

// 标记
var initialized uint32
var locks sync.Mutex

type singelton4 struct{}

var instance4 *singelton4

func GetInstance4() *singelton4 {
	//如果标记为被设置，直接返回，不加锁
	if atomic.LoadUint32(&initialized) == 1 {
		return instance4
	}

	//如果没有，则加锁申请
	locks.Lock()
	defer locks.Unlock()

	if initialized == 0 {
		instance4 = new(singelton4)
		//设置标记位
		atomic.StoreUint32(&initialized, 1)
	}

	return instance4
}

func (s *singelton4) SomeThing4() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance4()
	s.SomeThing4()
}
