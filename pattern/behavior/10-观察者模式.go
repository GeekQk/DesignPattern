package main

/*
观察者模式是用于建立一种对象与对象之间的依赖关系，一个对象发生改变时将自动通知其他对象，其他对象将相应作出反应。
在观察者模式中，发生改变的对象称为观察目标，而被通知的对象称为观察者，
一个观察目标可以对应多个观察者，而且这些观察者之间可以没有任何相互联系，
可以根据需要增加和删除观察者，使得系统更易于扩展。
*/

// 定义观察者接口
type Listener interface {
	OnTeackerCommint()
}

type Notify interface {
	AddListener(listener Listener)
	ReMoveListener(listener Listener)
	NotifyAll()
}

type StuZhang3 struct {
	badthing string
}

func (s *StuZhang3) OnTeackerCommint() {
	println("StuZhang3停止:" + s.badthing)
}

func (s *StuZhang3) DoBadThing() {
	println("StuZhang3在:" + s.badthing)
}

type StuWang5 struct {
	badthing string
}

func (s *StuWang5) OnTeackerCommint() {
	println("StuWang5停止:" + s.badthing)
}

func (s *StuWang5) DoBadThing() {
	println("StuWang5在:" + s.badthing)
}

// 统治者 班长
type ClassMonitor struct {
	ListenerList []Listener
}

func (c *ClassMonitor) AddListener(listener Listener) {
	c.ListenerList = append(c.ListenerList, listener)
}

func (c *ClassMonitor) ReMoveListener(listener Listener) {
	for i, v := range c.ListenerList {
		if v == listener {
			c.ListenerList = append(c.ListenerList[:i], c.ListenerList[i+1:]...)
			break
		}
	}
}

func (c *ClassMonitor) NotifyAll() {
	println("班长通知所有同学")
	for _, v := range c.ListenerList {
		v.OnTeackerCommint()
	}
}

func main() {
	zhang3 := &StuZhang3{badthing: "玩手机"}
	wang5 := &StuWang5{badthing: "大闹"}
	zhang3.DoBadThing()
	wang5.DoBadThing()
	// 班长通知所有同学
	classMonitor := &ClassMonitor{}
	classMonitor.AddListener(zhang3)
	classMonitor.AddListener(wang5)
	classMonitor.NotifyAll()
}

/*
优点：
(1) 观察者模式可以实现表示层和数据逻辑层的分离，定义了稳定的消息更新传递机制，并抽象了更新接口，使得可以有各种各样不同的表示层充当具体观察者角色。
(2) 观察者模式在观察目标和观察者之间建立一个抽象的耦合。观察目标只需要维持一个抽象观察者的集合，无须了解其具体观察者。由于观察目标和观察者没有紧密地耦合在一起，因此它们可以属于不同的抽象化层次。
(3) 观察者模式支持广播通信，观察目标会向所有已注册的观察者对象发送通知，简化了一对多系统设计的难度。
(4) 观察者模式满足“开闭原则”的要求，增加新的具体观察者无须修改原有系统代码，在具体观察者与观察目标之间不存在关联关系的情况下，增加新的观察目标也很方便。

缺点：
(1) 如果一个观察目标对象有很多直接和间接观察者，将所有的观察者都通知到会花费很多时间。
(2) 如果在观察者和观察目标之间存在循环依赖，观察目标会触发它们之间进行循环调用，可能导致系统崩溃。
(3) 观察者模式没有相应的机制让观察者知道所观察的目标对象是怎么发生变化的，而仅仅只是知道观察目标发生了变化。

适用场景
(1) 一个抽象模型有两个方面，其中一个方面依赖于另一个方面，将这两个方面封装在独立的对象中使它们可以各自独立地改变和复用。
(2) 一个对象的改变将导致一个或多个其他对象也发生改变，而并不知道具体有多少对象将发生改变，也不知道这些对象是谁。
(3) 需要在系统中创建一个触发链，A对象的行为将影响B对象，B对象的行为将影响C对象……，可以使用观察者模式创建一种链式触发机制。
*/
