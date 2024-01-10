package main

import "fmt"

// 代理模式
// 王婆代理潘金莲 做操作
// 本质就是不暴露下层对象 而是通过代理对象来操作下层对象

type BeautyWoMan interface {
	//对男人抛媚眼
	MakeEyesToMan()
	//和男人共度春宵
	HappyWithMan()
}

type PanJInLian struct{}

func (p *PanJInLian) MakeEyesToMan() {
	fmt.Println("潘金莲 抛媚眼")
}
func (p *PanJInLian) HappyWithMan() {
	fmt.Println("潘金莲 和男人共度春宵")
}

type WangPo struct {
	wowan BeautyWoMan
}

func (w *WangPo) MakeEyesToMan() {
	w.wowan.MakeEyesToMan()
}

func (w *WangPo) HappyWithMan() {
	w.wowan.HappyWithMan()
}

func NewProxy(wowan BeautyWoMan) *WangPo {
	return &WangPo{wowan}
}

func main() {
	wangpo := NewProxy(&PanJInLian{})
	wangpo.MakeEyesToMan()
	wangpo.HappyWithMan()
}
