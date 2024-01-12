package main

import "fmt"

// 把灯的开和关封装成命令
type Light struct {
	state string
}

func (l *Light) On() {
	l.state = "on"
	fmt.Println("Light is on")
}

func (l *Light) Off() {
	l.state = "off"
	fmt.Println("Light is off")
}

func (l *Light) GetState() string {
	return l.state
}

// Command
type Command interface {
	Execute()
}

// 开灯命令
type LightOnCommand struct {
	light *Light
}

func (c *LightOnCommand) Execute() {
	c.light.On()
}

// 关灯命令
type LightOffCommand struct {
	light *Light
}

func (c *LightOffCommand) Execute() {
	c.light.Off()
}

// Invoker
type RemoteControl struct {
	command Command
}

func (r *RemoteControl) SetCommand(command Command) {
	r.command = command
}

func (r *RemoteControl) PressButton() {
	r.command.Execute()
}

// Client
func main() {
	light := &Light{}
	onCommand := &LightOnCommand{light: light}
	offCommand := &LightOffCommand{light: light}
	remoteControl := &RemoteControl{}

	remoteControl.SetCommand(onCommand)
	remoteControl.PressButton() // Output: Light is on

	remoteControl.SetCommand(offCommand)
	remoteControl.PressButton() // Output: Light is off
}

/*
我们有一个Light接收者，它有一个状态（开或关）。
我们有两个命令：LightOnCommand和LightOffCommand，它们分别用于打开和关闭灯。
RemoteControl是调用者，它持有一个命令并执行它。
在客户端代码中，我们创建了一个灯、两个命令和一个遥控器，然后设置遥控器以执行打开和关闭灯的命令。
*/
