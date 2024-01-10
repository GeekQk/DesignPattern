package main

import "fmt"

/*
代理模式建议新建一个与原服务对象接口相同的代理类， 然后更新应用以将代理对象传递给所有原始对象客户端。
代理类接收到客户端请求后会创建实际的服务对象， 并将所有工作委派给它。
举例：mysql-proxy 代理连接请求、nginx代理请求、tomcat 代理连接
subject（抽象主题角色）：真实主题与代理主题的共同接口。
RealSubject（真实主题角色）：定义了代理角色所代表的真实对象。
Proxy（代理主题角色）：含有对真实主题角色的引用，代理角色通常在将客户端调用传递给真是主题对象之前或者之后执行某些操作，而不是单纯返回真实的对象。
*/

type Good struct {
	Kind string //商品种类
	Fact bool   //商品真伪
}

// =============抽象层============
type Shopping interface {
	Buy(good *Good) //购买商品
}

// =============实现层============
type KoreaShopping struct{}

func (k KoreaShopping) Buy(good *Good) {
	fmt.Println("韩国购物,买了", good.Kind)
}

type ChinaShopping struct{}

func (c ChinaShopping) Buy(good *Good) {
	fmt.Println("中国购物,买了", good.Kind)
}

// =============代理层============
type ProxyShopping struct {
	shopping Shopping
}

func newProxy(shopping Shopping) *ProxyShopping {
	return &ProxyShopping{shopping}
}

func (p *ProxyShopping) Buy(good *Good) {
	shopStatus := p.watch(good)
	if shopStatus == true { //检查商品合法性
		p.shopping.Buy(good)
	}
	p.check(good, shopStatus)
}

func (p *ProxyShopping) watch(good *Good) bool {
	fmt.Println("检查商品：", good.Kind)
	if good.Fact == false {
		fmt.Println("商品不合法:", good.Kind)
	}
	return good.Fact
}

func (p *ProxyShopping) check(good *Good, status bool) {
	if status == false {
		fmt.Println("检查-商品不合法:", good.Kind)
	} else {
		fmt.Println("检查-商品合法:", good.Kind)
	}
}

func main() {
	//若没有代理模式 只能这么书写
	// good1 := &Good{Kind: "手机", Fact: true}
	// var shopping Shopping
	// shopping = &KoreaShopping{}
	// if good1.Fact == true {
	// 	shopping.Buy(good1)
	// }

	//代理模式
	good1 := &Good{Kind: "手机", Fact: true}
	good2 := &Good{Kind: "电脑", Fact: false}

	fmt.Println("============给韩国代理=============")
	var k_shop Shopping
	k_shop = new(KoreaShopping)
	proxy := newProxy(k_shop)
	proxy.Buy(good1)
	fmt.Println("=========================")
	proxy.Buy(good2)

	fmt.Println("============给中国代理=============")
	var c_shop Shopping
	c_shop = new(ChinaShopping)
	proxy = newProxy(c_shop)
	proxy.Buy(good1)
	fmt.Println("=========================")
	proxy.Buy(good2)

	// 代理模式的好处：
	// 1. 隐藏了真实对象， 客户端不需要知道代理的存在
	// 2. 代理对象可以控制真实对象的访问， 可以执行一些附加操作
	// 3. 代理对象可以充当一个中介， 客户端不需要直接访问真实对象， 而是通过代理对象间接访问真实对象
	// 4. 代理对象可以控制真实对象的生命周期， 比如在真实对象创建之前和销毁之后执行一些操作
	// 5. 代理对象可以充当一个保护层， 客户端不能直接访问真实对象， 只能通过代理对象间接访问真实对象

}
