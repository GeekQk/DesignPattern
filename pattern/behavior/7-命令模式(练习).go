package main

// 联想路边撸串烧烤场景， 有烤羊肉，烧鸡翅命令，有烤串师傅，和服务员MM。根据命令模式，设计烤串场景。
//先抽象结构 烤羊肉 烤鸡翅 烤串
//分层结构 服务员(命令触发者) 师傅

import "fmt"

type Cooker struct{}

func (c *Cooker) MakeChicken() {
	fmt.Println("烤串师傅烤了鸡肉串儿")
}

func (c *Cooker) MakeChuaner() {
	fmt.Println("烤串师傅烤了羊肉串儿")
}

// 抽象的命令
type CommandList interface {
	Make()
}

// 烤鸡肉串
type CommandCookChicken struct {
	cooker *Cooker
}

func (cmd *CommandCookChicken) Make() {
	cmd.cooker.MakeChicken()
}

// 烤羊肉串
type CommandCookChuaner struct {
	cooker *Cooker
}

func (cmd *CommandCookChuaner) Make() {
	cmd.cooker.MakeChuaner()
}

type WaiterMM struct {
	CmdList []CommandList //收集的命令集合
}

func (w *WaiterMM) Notify() {
	if w.CmdList == nil {
		return
	}

	for _, cmd := range w.CmdList {
		cmd.Make()
	}
}

func main() {
	cooker := new(Cooker)
	cmdChicken := CommandCookChicken{cooker}
	cmdChuaner := CommandCookChuaner{cooker}

	mm := new(WaiterMM)
	mm.CmdList = append(mm.CmdList, &cmdChicken)
	mm.CmdList = append(mm.CmdList, &cmdChuaner)

	mm.Notify()
}
