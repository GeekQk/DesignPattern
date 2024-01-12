package main

import "fmt"

// 定义一个抽象类
type AbstractClass interface {
	AddRouter()
	Dispatch()
	Hook()
}

// 封装一个基于抽象类的流程
type BaseClass struct {
	t AbstractClass
}

func (b *BaseClass) Run() {
	b.t.AddRouter()
	b.t.Dispatch()
	b.t.Hook()
}

// 框架入口
type FrameWork struct {
	BaseClass
}

func NewFrameWork() *FrameWork {
	frameWork := new(FrameWork)
	frameWork.t = frameWork
	return frameWork
}

func (f *FrameWork) AddRouter() {
	fmt.Println("框架路由加载完成")
}
func (f *FrameWork) Dispatch() {
	fmt.Println("框架路由分发完成")
}
func (f *FrameWork) Hook() {
	fmt.Println("框架路由钩子完成")
}

func main() {
	frameWork := NewFrameWork()
	frameWork.Run()
}
