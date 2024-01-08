package main

import "fmt"

type Fruit struct {
}

func (f *Fruit) show(name string) {
	if name == "apple" {
		fmt.Println("我是苹果")
	} else if name == "banana" {
		fmt.Println("我是香蕉")
	} else if name == "pear" {
		fmt.Println("我是梨")
	}
}

func NewFruit(name string) *Fruit {
	fruit := &Fruit{}
	if name == "apple" {
		//苹果逻辑
	} else if name == "banana" {
		//香蕉逻辑
	} else if name == "pear" {
		//梨逻辑
	}
	return fruit
}

func main() {
	apple := NewFruit("apple")
	apple.show("apple")

	banana := NewFruit("banana")
	banana.show("banana")

	pear := NewFruit("pear")
	pear.show("pear")

}

/*
不难看出，Fruit类是一个“巨大的”类，在该类的设计中存在如下几个问题：
	(1) 在Fruit类中包含很多“if…else…”代码块，整个类的代码相当冗长，代码越长，阅读难度、维护难度和测试难度也越大；而且大量条件语句的存在还将影响系统的性能，程序在执行过程中需要做大量的条件判断。
	(2) Fruit类的职责过重，它负责初始化和显示所有的水果对象，将各种水果对象的初始化代码和显示代码集中在一个类中实现，违反了“单一职责原则”，不利于类的重用和维护；
	(3) 当需要增加新类型的水果时，必须修改Fruit类的构造函数NewFruit()和其他相关方法源代码，违反了“开闭原则”。
*/
