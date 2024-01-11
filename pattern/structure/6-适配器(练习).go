package main

import "fmt"

type Attack interface {
	Fight()
}

// 技能
type Magic struct{}

func (m *Magic) Fight() { fmt.Println("魔法攻击") }

type BaoJian struct{}

func (b *BaoJian) Fight() { fmt.Println("宝剑攻击") }

// 适配者
type PowerOff struct {
}

func (p *PowerOff) PowerOn()  { fmt.Println("电源开启") }
func (p *PowerOff) PowerOff() { fmt.Println("电源关闭") }

// 适配器 只要魔法攻击就会关闭电源
type MagicAdapter struct {
	powerof *PowerOff
}

func (m *MagicAdapter) Fight() {
	m.powerof.PowerOff()
}

func NewMagicAdapter(powerof *PowerOff) *MagicAdapter {
	return &MagicAdapter{powerof}
}

// 英雄和对应的技能
type Hero struct {
	Name   string
	attack Attack
}

func (h *Hero) Skill() {
	fmt.Println("英雄:", h.Name, "使用技能")
	h.attack.Fight()
	fmt.Println("=================")
}

func main() {
	hero1 := Hero{Name: "盖伦", attack: &Magic{}}
	hero1.Skill()

	hero2 := Hero{Name: "艾希", attack: &BaoJian{}}
	hero2.Skill()

	// 适配器 使用适配器包装了技能
	hero3 := Hero{Name: "艾希", attack: NewMagicAdapter(&PowerOff{})}
	hero3.Skill()
}

// 优点：
// (1) 将目标类和适配者类解耦，通过引入一个适配器类来重用现有的适配者类，无须修改原有结构。
// (2) 增加了类的透明性和复用性，将具体的业务实现过程封装在适配者类中，对于客户端类而言是透明的，而且提高了适配者的复用性，同一个适配者类可以在多个不同的系统中复用。
// (3) 灵活性和扩展性都非常好，可以很方便地更换适配器，也可以在不修改原有代码的基础上增加新的适配器类，完全符合“开闭原则”。
// 缺点:
// 适配器中置换适配者类的某些方法比较麻烦。
