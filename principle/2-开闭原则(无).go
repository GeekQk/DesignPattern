package main

import "fmt"

/*
开闭原则:
添加一个新的功能，应该是通过在已有代码基础上扩展代码(新增模块、类、方法、属性等)
而非修改已有代码(修改模块、类、方法、属性等)的方式来完成。
*/

// 我们要写一个类,Banker银行业务员
type Banker struct {
}

// 存款业务
func (this *Banker) Save() {
	fmt.Println("进行了 存款业务...")
}

// 转账业务
func (this *Banker) Transfer() {
	fmt.Println("进行了 转账业务...")
}

// 支付业务
func (this *Banker) Pay() {
	fmt.Println("进行了 支付业务...")
}

func (this *Banker) Fund() {

}

func main() {
	banker := &Banker{}

	banker.Save()
	banker.Transfer()
	banker.Pay()
}

/*
代码很简单，就是一个银行业务员，他可能拥有很多的业务，比如Save()存款、Transfer()转账、Pay()支付等。
那么如果这个业务员模块只有这几个方法还好，但是随着我们的程序写的越来越复杂，银行业务员可能就要增加方法，会导致业务员模块越来越臃肿。
当我们去给Banker添加新的业务的时候，会直接修改原有的Banker代码，那么Banker模块的功能会越来越多，出现问题的几率也就越来越大
*/
