package main

import "fmt"

// 策略模式 定义一系列的算法，把它们一个个封装起来，并且使它们可相互替换。
// 策略模式把行为和环境分开，让它们独立变化。
// 策略模式让算法的使用和算法的实现分离，算法的实现可以独立变化。

// 武器策略-使用武器
type WeaponStrategy interface {
	useWeapon()
}

// 武器策略实现-使用AK47
type Ak47 struct{}

func (a *Ak47) useWeapon() {
	fmt.Println("使用AK47")
}

// 武器策略实现-使用刀
type Knife struct{}

func (k *Knife) useWeapon() {
	fmt.Println("使用刀")
}

type Hero struct {
	strategy WeaponStrategy
}

func (h *Hero) SetWeapon(s WeaponStrategy) {
	h.strategy = s
}

func (h *Hero) Fight() {
	h.strategy.useWeapon()
}

func main() {
	hero := Hero{}
	hero.SetWeapon(&Ak47{})
	hero.Fight()

	//更换策略
	hero.SetWeapon(&Knife{})
	hero.Fight()
}
