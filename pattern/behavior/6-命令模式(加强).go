package main

import "fmt"

// 医生-命令接收者
type Doctorss struct{}

func (d *Doctorss) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctorss) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 抽象的命令
type Commands interface {
	Treat()
}

// 治疗眼睛的病单
type CommandTreatEyes struct {
	doctor *Doctorss
}

func (cmd *CommandTreatEyes) Treat() {
	cmd.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNoses struct {
	doctor *Doctorss
}

func (cmd *CommandTreatNoses) Treat() {
	cmd.doctor.treatNose()
}

// 护士-调用命令者
type Nurse struct {
	CmdList []Commands //收集的命令集合
}

// 发送病单，发送命令的方法
func (n *Nurse) Notify() {
	if n.CmdList == nil {
		return
	}

	for _, cmd := range n.CmdList {
		cmd.Treat() //执行病单绑定的命令(这里会调用病单已经绑定的医生的诊断方法)
	}
}

// 病人
func main() {
	//依赖病单，通过填写病单，让医生看病
	doctor := new(Doctorss)
	//治疗眼睛的病单
	cmdEye := CommandTreatEyes{doctor}
	//治疗鼻子的病单
	cmdNose := CommandTreatNoses{doctor}

	//护士
	nurse := new(Nurse)
	//收集管理病单
	nurse.CmdList = append(nurse.CmdList, &cmdEye)
	nurse.CmdList = append(nurse.CmdList, &cmdNose)

	//执行病单指令
	nurse.Notify()
}
