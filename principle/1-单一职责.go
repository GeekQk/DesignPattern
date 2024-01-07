package main

import "fmt"

/*
单一职责原则（Simple Responsibility Principle, SRP）指不要存在一个以上导致类变更的原因。
假设有一个Class负责两个职责，一旦发生需求变更，修改其中一个职责的逻辑代码，有可能会导致另一个职责的功能发生故障。
分别用两个Class来实现两个职责，进行解耦。
总体来说就是一个Class、Interface、Method只负责一项职责。
*/

type ClothesShop struct{}

func (cs *ClothesShop) OnShop() {
	fmt.Println("休闲的装扮")
}

type ClothesWork struct{}

func (cw *ClothesWork) OnWork() {
	fmt.Println("工作的装扮")
}

func main() {
	//工作的时候
	cw := new(ClothesWork)
	cw.OnWork()

	//shopping的时候
	cs := new(ClothesShop)
	cs.OnShop()
}

/*
错误写法:
type Clothes struct{}

func (cs *Clothes) OnShop() {
	fmt.Println("休闲的装扮")
}

func (cs *Clothes) OnWork() {
	fmt.Println("工作的装扮")
}
func main() {
	cw := new(Clothes)
	cw.OnShop()
	cw.OnWork()
}
这种写法，在后期维护的时候，一旦需求变更，需要修改Clothes结构体的方法,而且Clothes结构体中存在两个职责,如果Clothes结构体中存在多个职责，那么Clothes结构体就会变得臃肿。
单一职责核心就是类对外暴露的接口，只负责一项职责,但是这种模式也有不好的情况,如果项目足够复杂，拆分的类就会很多,那么类与类的依赖关系就会很复杂,耦合度就会增高。
哎, 写代码的时候，不一定要遵循单一职责原则,面向对象追求的是高内聚低耦合,但是高内聚和低耦合不是绝对的,这里只是追求一个平衡。
说说缺点:
	一个职责的变化可能会削弱或者抑制这个类实现其他职责的能力；
	当客户端需要该对象的某一个职责时，不得不将其他不需要的职责全都包含进来，从而造成冗余代码或代码的浪费
	耦合度增高
*/
