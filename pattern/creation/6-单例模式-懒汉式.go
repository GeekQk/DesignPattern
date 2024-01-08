package main

import "fmt"

type singelton2 struct{}

var instance2 *singelton2

// 懒汉式:只有首次GetInstance()方法被调用，才会生成这个单例的实例
// 上面的“懒汉式”实现是非线程安全的设计方式,也就是如果多个线程或者协程同时首次调用GetInstance()方法有概率导致多个实例被创建，则违背了单例的设计初衷
func GetInstance2() *singelton2 {
	if instance2 == nil {
		instance2 = new(singelton2)
	}

	//接下来的GetInstance直接返回已经申请的实例即可
	return instance2
}

func (s *singelton2) SomeThings() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s := GetInstance2()
	s.SomeThings()

}
