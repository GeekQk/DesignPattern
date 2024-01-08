package main

import (
	"fmt"
	"sync"
)

// 线程安全
//1.锁机制实现 开销大
//2.原子操作实现
//3.Once机制实现

var once sync.Once

type singelton5 struct{}

var instance5 *singelton5

func GetInstance5() *singelton5 {
	once.Do(func() {
		instance5 = new(singelton5)
	})
	return instance5
}

func (s *singelton5) SomeThing5() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance5()
	s.SomeThing5()
}
