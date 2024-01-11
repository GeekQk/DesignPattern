package main

import "fmt"

// 适配对象
type V5 interface {
	use5V()
}

// 需要兼容的对象
type V220 struct{}

func (v *V220) use22V() {
	fmt.Println("V220 电压充电")
}

// 适配器 需要把220v 适配成5V
func NewAdapter(v220 *V220) *Adapter {
	return &Adapter{V220: v220}
}

type Adapter struct {
	V220 *V220
}

func (a *Adapter) use5V() {
	fmt.Println("Adapter 适配了5V")
	a.V220.use22V()
}

// 业务层
type phone struct {
	v V5
}

func (p *phone) Charge() {
	fmt.Println("Phone 进行了充电")
	p.v.use5V()
}

func NewPhone(v V5) *phone {
	return &phone{v: v}
}

func main() {
	phone := NewPhone(NewAdapter(&V220{}))
	phone.Charge()
}
