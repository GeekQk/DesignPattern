package main

import "fmt"

/*
也称最少知道原则，一个类对自己依赖的类知道的越少越好。
两个类如果不必彼此同行，那么这两个类就不应当发生直接的相互作用，而是通过引入一个第三者发生剪接交互
迪米特法则初衷在于降低类之间的耦合。由于每个类尽量减少对其它类的依赖,因此。
很容易使得系统的功能模块独立,相互之间不存在(或很少有)依赖关系。
*/
// 迪米特法则：一个对象应该对其他对象有最少的了解。
// 迪米特法则又叫最少知道原则，一个类应该对自己需要耦合或调用的类知道得越少越好。
// 也就是说，对于被耦合的类来说，只应当了解和该类有关联关系的对象，而不要了解与它有关联关系的对象的内部细节。

// 员工
type Employee struct {
	name string
	id   int
}

// 部门
type Department struct {
	name      string
	employees []Employee
}

func (d *Department) PrintEmployeeNames() {
	for _, employee := range d.employees {
		fmt.Println(employee.name)
	}
}

func main() {
	d := &Department{"研发部", []Employee{{"张三", 1}, {"李四", 2}}}
	d.PrintEmployeeNames()

}

/*
在这个例子中，Department 类有一个 employees 字段，它是一个 Employee 类型的切片。
Department 类还有一个 PrintEmployeeNames 方法，它遍历 employees 切片并打印每个员工的名字。
这个例子符合迪米特原则，因为 Department 类只与它的直接朋友（即 Employee 类型）交互，而没有与陌生对象（即其他类型的对象）交互。
*/
