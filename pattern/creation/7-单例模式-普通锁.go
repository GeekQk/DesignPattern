package main

import (
	"fmt"
	"sync"
)

// 定义锁
var lock sync.Mutex

type singelton3 struct{}

var instance3 *singelton3

func GetInstance3() *singelton3 {
	//为了线程安全，增加互斥
	lock.Lock()
	defer lock.Unlock()

	if instance3 == nil {
		instance3 = new(singelton3)
	}
	return instance3
}

func (s *singelton3) SomeThing3() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance3()
	s.SomeThing3()
}
