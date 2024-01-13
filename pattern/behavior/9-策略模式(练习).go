package main

import "fmt"

type SaleRule interface {
	UseSale()
}

type ProductA struct {
}

func (p *ProductA) UseSale() {
	fmt.Println("ProductA 商品促销-8折")
}

type ProductB struct{}

func (p *ProductB) UseSale() {
	fmt.Println("ProductB 商品促销-6折")
}

type Person struct {
	sale SaleRule
}

func (p *Person) SetSale(s SaleRule) {
	p.sale = s
}
func (p *Person) Buy() {
	p.sale.UseSale()
}

func main() {
	pA := &ProductA{}
	p := &Person{}
	p.SetSale(pA)
	p.Buy()

	pB := &ProductB{}
	p.SetSale(pB)
	p.Buy()
}

/*
优点：
(1) 策略模式提供了对“开闭原则”的完美支持，用户可以在不修改原有系统的基础上选择算法或行为，也可以灵活地增加新的算法或行为。
(2) 使用策略模式可以避免多重条件选择语句。多重条件选择语句不易维护，它把采取哪一种算法或行为的逻辑与算法或行为本身的实现逻辑混合在一起，将它们全部硬编码(Hard Coding)在一个庞大的多重条件选择语句中，比直接继承环境类的办法还要原始和落后。
(3) 策略模式提供了一种算法的复用机制。由于将算法单独提取出来封装在策略类中，因此不同的环境类可以方便地复用这些策略类。

缺点：
(1) 客户端必须知道所有的策略类，并自行决定使用哪一个策略类。这就意味着客户端必须理解这些算法的区别，以便适时选择恰当的算法。换言之，策略模式只适用于客户端知道所有的算法或行为的情况。
(2) 策略模式将造成系统产生很多具体策略类，任何细小的变化都将导致系统要增加一个新的具体策略类。
*/
