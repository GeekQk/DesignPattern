package main

import "fmt"

/*
简单工厂模式（Simple Factory Pattern）又称为静态工厂方法（Static Factory Method）模式，它属于类创建型模式。
在简单工厂模式中，可以根据参数的不同返回不同类的实例。简单工厂模式专门定义一个类来负责创建其他类的实例，被创建
的实例通常都具有共同的父类。
简单工厂模式的要点有三个：
工厂类：简单工厂模式的核心，它负责实现创建所有实例的内部逻辑。工厂类可以被外界直接调用，创建所需的产品对象。
抽象产品（AbstractProduct）角色：简单工厂模式所创建的所有对象的父类，它负责描述所有实例所共有的公共接口。
具体产品（Concrete Product）角色：简单工厂模式所创建的具体实例对象。
*/

//核心:业务逻辑层只会和工厂模块进行依赖,这样业务逻辑层将不再关心Fruit类是具体怎么创建基础对象的。
// ======= 抽象层 =========

// 水果类(抽象接口)
type Fruits interface {
	Show() //接口的某方法
}

// ======= 基础类模块 =========

type Apple struct{}

func (apple *Apple) Show() {
	fmt.Println("我是苹果")
}

type Banana struct{}

func (banana *Banana) Show() {
	fmt.Println("我是香蕉")
}

type Pear struct{}

func (pear *Pear) Show() {
	fmt.Println("我是梨子")
}

// ========= 工厂模块  =========
// 一个工厂， 有一个生产水果的机器，返回一个抽象水果的指针
type Factory struct{}

func (fac *Factory) CreateFruit(kind string) Fruits {
	var fruit Fruits

	if kind == "apple" {
		fruit = new(Apple)
	} else if kind == "banana" {
		fruit = new(Banana)
	} else if kind == "pear" {
		fruit = new(Pear)
	}
	return fruit
}

// ==========业务逻辑层==============
func main() {
	//工厂类的本质是返回一个抽象的接口,具体实现由工厂类来决定
	factory := new(Factory)

	//工厂类返回一个抽象的接口,具体实现由工厂类来决定
	apple := factory.CreateFruit("apple")
	apple.Show()

	banana := factory.CreateFruit("banana")
	banana.Show()

	pear := factory.CreateFruit("pear")
	pear.Show()
}
