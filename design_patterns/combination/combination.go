/**
 * composite design_patterns
 * 将战斗单元不停的组合，射手是一个战斗单元，加农炮是一个战斗单元，两者都是局部对象
 * 军队是多个战斗单元组成的，力量是所有战斗单元之和，多个军队可以组合为一个军队，是组合对象
 * 可以合并军队，也可以移除军队中的某个军队
 */
package main

import "fmt"

// UnitComponent 战斗单元Component
type UnitComponent interface {
	bombandStrength() uint
}

// Archer 射手leaf
type Archer struct{}

func (ar Archer) bombandStrength() uint {
	return 4
}

// LandCanor 加农炮leaf
type LandCanor struct{}

func (lc LandCanor) bombandStrength() uint {
	return 10
}

// Army 军队->composite组合对象, 可以将不同的军队组合
type Army struct {
	units []UnitComponent
}

func NewArmy() *Army {
	return &Army{}
}

func (ay *Army) addUnit(unit UnitComponent) {
	ay.units = append(ay.units, unit)
}

func (ay *Army) removeUnit(unit UnitComponent) {
	for k, v := range ay.units {
		if v == unit {
			ay.units = append(ay.units[:k], ay.units[k+1:]...)
		}
	}
}

func (ay *Army) bombandStrength() uint {
	sum := 0
	for _, v := range ay.units {
		sum += int(v.bombandStrength())
	}
	return uint(sum)
}

func main() {
	unit1 := &Archer{}
	unit2 := &LandCanor{}

	army := NewArmy()
	army.addUnit(unit1)
	army.addUnit(unit2)
	army.removeUnit(unit1)

	fmt.Println(army.bombandStrength())

	army2 := NewArmy()
	army2.addUnit(unit1)
	army2.addUnit(army)

	fmt.Println(army2.bombandStrength())
}
