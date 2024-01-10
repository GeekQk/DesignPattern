package main

import "fmt"

/*
装饰器模式（Decorator Pattern）允许向一个现有的对象添加新的功能，同时又不改变其结构。
这种类型的设计模式属于结构型模式，它是作为现有的类的一个包装。
装饰器模式通过将对象包装在装饰器类中，以便动态地修改其行为。
这种模式创建了一个装饰类，用来包装原有的类，并在保持类方法签名完整性的前提下，提供了额外的功能。
*/

// -----------------抽象层--------------------
type Phone interface {
	Show()
}

// -----------------抽象的装饰器--------------------
type Decoator struct {
	phone Phone
}

func (d *Decoator) Show() {
	d.phone.Show()
}

// -----------------实现层--------------------
type HuaWei struct {
}

func (h *HuaWei) Show() {
	fmt.Println("华为手机.....Show")
}

type Xiaomi struct {
}

func (x *Xiaomi) Show() {
	fmt.Println("小米手机.....Show")
}

// -----------------手机膜装饰器--------------------
func NewMoDecorator(phone Phone) Phone {
	return &MoDecorator{Decoator{phone: phone}}
}

type MoDecorator struct {
	Decoator
}

func (m *MoDecorator) Show() {
	m.phone.Show() // 调用被装饰的自己的方法
	fmt.Println("手机膜装饰.....Show")
}

// -----------------手机壳装饰器--------------------
func NewKeDecorator(phone Phone) Phone {
	return &KeDecorator{Decoator{phone: phone}}
}

type KeDecorator struct {
	Decoator
}

func (m *KeDecorator) Show() {
	m.phone.Show() // 调用被装饰的自己的方法
	fmt.Println("手机壳装饰.....Show")
}

func main() {
	huawei := &HuaWei{}
	huawei.Show()

	fmt.Println("-------------------")
	var moHuaweiMo Phone
	moHuaweiMo = NewMoDecorator(huawei)
	moHuaweiMo.Show()

	fmt.Println("-------------------")
	var keHuaWei Phone
	keHuaWei = NewKeDecorator(huawei)
	keHuaWei.Show()

	//装饰器模式，可以动态的扩展一个对象的功能。
	//装饰器模式通过组合来拓展一个对象的功能,这样方法得到扩充，而不是继承。

}
