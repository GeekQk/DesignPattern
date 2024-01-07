package main

import "fmt"

/*
接口隔离原则:

	一个接口应该只提供一种对外功能，不应该把所有操作都封装到一个接口中去。

合成复用原则:

	它要求在软件构建过程中，尽量使用组合/聚合的方式
	而不是使用继承,因为组合/聚合可以使系统更加灵活，降低类与类之间的耦合度,继承是静态的，组合是动态的。
*/
type Cat struct{}

func (c *Cat) Eat() {
	fmt.Println("父类小猫吃饭")
}

// 给小猫添加一个 可以睡觉的方法 （使用继承来实现）
type CatB struct {
	Cat
}

func (cb *CatB) Sleep() {
	fmt.Println("小猫睡觉")
}

// 给小猫添加一个 可以睡觉的方法 （使用组合的方式）
type CatC struct {
	C *Cat
}

// 组合的方式，可以复用父类的方法
// 组合本质就是把父类作为属性存在 更加灵活
func (cc *CatC) EatOne() {
	cc.C.Eat() //调用父类的Eat方法
	fmt.Println("子类小猫吃饭")
}

// 把父类作为属性存在 更加灵活,但是调用需要先初始化父类
func (cc *CatC) Eat(c *Cat) {
	c.Eat()
}

func (cc *CatC) Sleep() {
	fmt.Println("子类小猫睡觉")
}

func main() {
	/*
		父类
		c := &Cat{}
		c.Eat()
		//子类通过继承增加的功能，可以正常使用
		cb := new(CatB)
		cb.Eat()
		cb.Sleep()
	*/

	//通过组合增加的功能，可以正常使用
	cc := new(CatC)
	cc.C = new(Cat)
	cc.C.Eat()
	cc.Sleep()

}

/*
组合的好处是可以复用某个方法
而不是继承，要把全部方法继承过来,会比较危险
组合代价比较小，继承代价比较大
推荐使用组合的方式
*/
