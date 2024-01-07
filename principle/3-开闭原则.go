package main

import "fmt"

/*
类的改动是通过增加代码进行的，而不是修改源代码。
添加一个新的功能，应该是通过在已有代码基础上扩展代码(新增模块、类、方法、属性等)
而非修改已有代码(修改模块、类、方法、属性等)的方式来完成。
*/
// 抽象的银行业务员
type AbstractBanker interface {
	DoBusi() //抽象的处理业务接口
}

// 存款的业务员
type SaveBanker struct{}

func (sb *SaveBanker) DoBusi() {
	fmt.Println("进行了存款")
}

// 转账的业务员
type TransferBanker struct{}

func (tb *TransferBanker) DoBusi() {
	fmt.Println("进行了转账")
}

// 支付的业务员
type PayBanker struct{}

func (pb *PayBanker) DoBusi() {
	fmt.Println("进行了支付")
}

func main() {
	//进行存款
	sb := &SaveBanker{}
	sb.DoBusi()

	//进行转账
	tb := &TransferBanker{}
	tb.DoBusi()

	//进行支付
	pb := &PayBanker{}
	pb.DoBusi()
}

/*
通过把操作抽象出一层，通过接口来调用,只需要抽象出一层接口层
这样就可以实现开闭原则，对扩展开放，对修改关闭。
抽象的存款、转账、支付操作
另外可以拓展成架构层,基于抽象层,进行业务封装

//实现架构层(基于抽象层进行业务封装-针对interface接口进行封装)
func BankerBusiness(banker AbstractBanker) {
	//通过接口来向下调用，(多态现象)
	banker.DoBusi()
}
func main() {
	//进行存款
	BankerBusiness(&SaveBanker{})

	//进行存款
	BankerBusiness(&TransferBanker{})

	//进行存款
	BankerBusiness(&PayBanker{})
}
*/
