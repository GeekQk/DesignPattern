package main

import "fmt"

/*
子类对象能够替换程序(program)中父类对象出现的任何地方，并且保证原来程序的逻辑行为(behavior)不变及正确性不被破坏。
子类在设计的时候，要遵守父类的行为约定(或者叫协议)。
*/

type Notify interface {
	Send()
}

type Message struct {
}

func (message *Message) Send() {
	fmt.Println("message send")
}

type SMS struct {
}

func (message *SMS) Send() {
	fmt.Println("sms  send")
}

// LetDo 是一个函数，接受一个Notify类型的参数m，并执行以下操作：
// 1. 检查m是否为*SMS类型，如果是，则打印"sms send Before"，并调用v的Send方法。
// 2. 检查m是否为*Message类型，如果是，则打印"sms send Before"，并调用v的Send方法。
func LetDo(m Notify) {
	if v, ok := m.(*SMS); ok {
		fmt.Println("sms send Before")
		v.Send()
	}

	if v, ok := m.(*Message); ok {
		fmt.Println("sms send Before")
		v.Send()
	}
}

func main() {
	msg := &Message{}
	LetDo(msg)
	sms := &SMS{}
	LetDo(sms)
}

/*
通过封装了一层接口层 实现了里式替换原则
本质有点类似适配器模式 无损切换
*/
