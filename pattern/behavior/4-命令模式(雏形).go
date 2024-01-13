package main

import "fmt"

/*
命令模式是一种设计模式，它封装了一个请求作为一个对象，从而让你使用不同的请求把客户端与服务端操作解耦。命令模式通常用于实现撤销（undo）和重做（redo）功能，以及实现事务操作。

在命令模式中，通常会涉及以下角色：

接收者（Receiver）：执行命令的对象。它知道如何实施与执行一个请求相关的操作。
命令（Command）：将请求封装成一个对象，从而可以用不同的请求参数化客户端，对请求排队或者记录请求日志，以及支持可撤销的操作。
调用者（Invoker）：要求该命令执行这个请求。通常会持有命令对象，可以持有很多的命令对象。这个是客户端真正触发命令并要求命令执行相应操作的地方，也就是说相当于使用命令对象的入口。
客户端（Client）：创建具体的命令对象，并且设置命令对象的接收者。
*/
type Doctors struct{}

func (d *Doctors) treatEye() {
	fmt.Println("医生治疗眼睛")
}

func (d *Doctors) treatNose() {
	fmt.Println("医生治疗鼻子")
}

// 治疗眼睛的病单
type CommandTreatEye struct {
	doctor *Doctors
}

func (cmd *CommandTreatEye) Treat() {
	cmd.doctor.treatEye()
}

// 治疗鼻子的病单
type CommandTreatNose struct {
	doctor *Doctors
}

func (cmd *CommandTreatNose) Treat() {
	cmd.doctor.treatNose()
}

// 病人
func main() {
	//依赖病单，通过填写病单，让医生看病
	//治疗眼睛的病单
	doctor := new(Doctors)
	cmdEye := CommandTreatEye{doctor: doctor}
	cmdEye.Treat() //通过病单来让医生看病

	cmdNose := CommandTreatNose{doctor: doctor}
	cmdNose.Treat() //通过病单来让医生看病

	/*
		病单和医生解耦
		这样只要有病单，就可以通过病单来让医生看病
	*/
}
