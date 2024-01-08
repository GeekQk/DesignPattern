package main

import "fmt"

/*

饿汉式含义是，在初始化单例唯一指针的时候，就已经提前开辟好了一个对象，申请了内存。
饿汉式的好处是，不会出现线程并发创建，导致多个单例的出现，但是缺点是如果这个单例对象在业务逻辑没有被使用，也会客观的创建一块内存对象。


Singleton（单例）：在单例类的内部实现只生成一个实例，同时它提供一个静态的getInstance()工厂方法，让客户可以访问它的唯一实例；
为了防止在外部对其实例化，将其构造函数设计为私有；
在单例类内部定义了一个Singleton类型的静态对象，作为外部共享的唯一实例。

单例模式要解决的问题是：
保证一个类永远只能有一个对象，且该对象的功能依然能被其他模块使用。
*/

/*
单例模式的三个要点：
	一是某个类只能有一个实例；
	二是它必须自行创建这个实例；
	三是它必须自行向整个系统提供这个实例。
*/

/*
	保证一个类永远只能有一个对象
*/

// 1、保证这个类非公有化，外界不能通过这个类直接创建一个对象,那么这个类就应该变得非公有访问 类名称首字母要小写
type singelton struct{}

// 2、但是还要有一个指针可以指向这个唯一对象，但是这个指针永远不能改变方向,Golang中没有常指针概念，所以只能通过将这个指针私有化不让外部模块访问
var instance *singelton = new(singelton)

// 3、如果全部为私有化，那么外部模块将永远无法访问到这个类和对象，
//
//	所以需要对外提供一个方法来获取这个唯一实例对象
//	注意：这个方法是否可以定义为singelton的一个成员方法呢？
//	答案是不能，因为如果为成员方法就必须要先访问对象、再访问函数
//	但是类和对象目前都已经私有化，外界无法访问，所以这个方法一定是一个全局普通函数
func GetInstance() *singelton {
	return instance
}

func (s *singelton) SomeThing() {
	fmt.Println("单例对象的某方法")
}

func main() {
	s1 := GetInstance()
	s1.SomeThing()

	s2 := GetInstance()
	if s1 == s2 {
		fmt.Println("s1 == s2")
	}
}
