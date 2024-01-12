package main

import "fmt"

/*
在模板方法模式中，通常存在一个抽象类，它定义了算法的骨架和一些抽象方法。
子类通过实现这些抽象方法来提供具体的行为。
模板方法模式使得子类可以在不改变算法结构的情况下对算法进行扩展。
具体实现:

	1.定义一个接口，定义算法骨架
	2.定义一个具体struct，实现接口,并封装流程列表
	3.业务层继承具体struct，实现接口方法

适用于以下场景：

	1.当多个类具有类似的算法时，可以使用模板方法模式将这些算法提取到一个共同的抽象父类中，并且在子类中实现具体的差异化行为。
	2.当需要控制某个算法的流程，但是各个步骤的实现可能因应用需要而有所不同时，可以使用模板方法模式。
	3.当部分复杂的算法步骤可能会变化，但是整体算法的结构却很稳定时，可以使用模板方法模式。
*/
type Beverage interface {
	//煮开水
	BoilWater()
	//泡茶
	Brew()
	//倒入杯中
	PourInCup()
	//加酌料
	AddThings()
	//是否添加酌料
	WantAddSugar() bool
}

// 封装一套流程模板基类,让具体的制作流程继承且实现
type template struct {
	b Beverage
}

// 具体制作流程
func (t *template) MakeBeverage() {
	if t == nil {
		return
	}
	t.b.BoilWater()
	t.b.Brew()
	t.b.PourInCup()
	// 判断是否需要加糖
	if t.b.WantAddSugar() {
		t.b.AddThings()
	}
}

// 业务层
type MakeCoffee struct {
	template
}

func NewMakeCoffee() *MakeCoffee {
	//b 为Beverage，是MakeTea，这里需要给接口赋值，指向具体的子类对象
	//来触发b全部接口方法的多态特性。
	makeCoffe := &MakeCoffee{}
	makeCoffe.b = makeCoffe
	return makeCoffe
}

func (m *MakeCoffee) BoilWater() {
	fmt.Println("coffee-煮开水")
}
func (m *MakeCoffee) Brew() {
	fmt.Println("coffee-泡咖啡")
}
func (m *MakeCoffee) PourInCup() {
	fmt.Println("coffee-倒入杯中")
}
func (m *MakeCoffee) AddThings() {
	fmt.Println("coffee-加糖")
}
func (m *MakeCoffee) WantAddSugar() bool {
	return true
}

type MakeTea struct {
	template
}

func NewMakeTea() *MakeTea {
	makeTea := &MakeTea{}
	makeTea.b = makeTea
	return makeTea
}
func (m *MakeTea) BoilWater() {
	fmt.Println("茶-煮开水")
}
func (m *MakeTea) Brew() {
	fmt.Println("茶-泡茶")
}
func (m *MakeTea) PourInCup() {
	fmt.Println("茶-倒入杯中")
}
func (m *MakeTea) AddThings() {
	fmt.Println("茶-加糖")
}
func (m *MakeTea) WantAddSugar() bool {
	return false
}

func main() {
	makeCoffe := NewMakeCoffee()
	makeCoffe.MakeBeverage()

	makeTea := NewMakeTea()
	makeTea.MakeBeverage()
}

/*
优点:
(1) 在父类中形式化地定义一个算法，而由它的子类来实现细节的处理，在子类实现详细的处理算法时并不会改变算法中步骤的执行次序。
(2) 模板方法模式是一种代码复用技术，它在类库设计中尤为重要，它提取了类库中的公共行为，将公共行为放在父类中，而通过其子类来实现不同的行为，它鼓励我们恰当使用继承来实现代码复用。
(3) 可实现一种反向控制结构，通过子类覆盖父类的钩子方法来决定某一特定步骤是否需要执行。
(4) 在模板方法模式中可以通过子类来覆盖父类的基本方法，不同的子类可以提供基本方法的不同实现，更换和增加新的子类很方便，符合单一职责原则和开闭原则。
*/

/*
缺点：
需要为每一个基本方法的不同实现提供一个子类，如果父类中可变的基本方法太多，将会导致类的个数增加
系统更加庞大，设计也更加抽象。
*/
