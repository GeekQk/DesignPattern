package main

/*
外观模式（Facade Pattern）是一种常用的软件设计模式，它为子系统中的一组接口提供了一个统一的高级接口，从而使得子系统更容易使用。
外观模式定义了一个高层接口，这个接口使得客户端或调用方可以更加方便地使用子系统中的一组接口，而无需了解子系统的内部实现细节。
外观模式的主要目的是简化客户端与子系统之间的交互，并将客户端与子系统的实现细节隔离开来。
通过提供一个简单而统一的接口，外观模式隐藏了子系统的复杂性，并减少了客户端与子系统之间的耦合。
在外观模式中，通常存在一个外观类（Facade），它实现了客户端所需的接口，并将客户端的请求转发给相应的子系统类进行处理。
外观类对客户端提供了一个简单而易于使用的接口，同时隐藏了子系统的具体实现细节。
*/
import "fmt"

type SubSystemA struct{}

func (sa *SubSystemA) MethodA() {
	fmt.Println("子系统方法A")
}

type SubSystemB struct{}

func (sb *SubSystemB) MethodB() {
	fmt.Println("子系统方法B")
}

type SubSystemC struct{}

func (sc *SubSystemC) MethodC() {
	fmt.Println("子系统方法C")
}

type SubSystemD struct{}

func (sd *SubSystemD) MethodD() {
	fmt.Println("子系统方法D")
}

// 外观模式，提供了一个外观类， 简化成一个简单的接口供使用
type Facade struct {
	a *SubSystemA
	b *SubSystemB
	c *SubSystemC
	d *SubSystemD
}

func (f *Facade) MethodOne() {
	f.a.MethodA()
	f.b.MethodB()
}

func (f *Facade) MethodTwo() {
	f.c.MethodC()
	f.d.MethodD()
}

func main() {
	//如果不用外观模式实现MethodA() 和 MethodB()
	sa := new(SubSystemA)
	sa.MethodA()
	sb := new(SubSystemB)
	sb.MethodB()

	fmt.Println("-----------")
	//使用外观模式
	f := Facade{
		a: new(SubSystemA),
		b: new(SubSystemB),
		c: new(SubSystemC),
		d: new(SubSystemD),
	}

	//调用外观包裹方法
	f.MethodTwo()
}
