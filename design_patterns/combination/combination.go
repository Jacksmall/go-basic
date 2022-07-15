/**
 * TODO 组合模式
 */
package main

import "fmt"

type Unit struct {
	bombandStrength uint // 战斗强度
}

// ======= region start =======
type Army struct {
	units []Unit
}

func NewArmy() *Army {
	return &Army{
		units: make([]Unit, 0),
	}
}

func (a *Army) addUnit(u Unit) {
	if isExist := find(a.units, u); !isExist {
		a.units = append(a.units, u)
	}
}

func (a *Army) removeUnit(u Unit) {
	a.units = removeSlice(a.units, u)
}

// ======= end region =======

type Archer struct {
	Unit
}

type Tomp struct {
	Unit
}

func removeSlice(x []Unit, r Unit) []Unit {
	for i := 0; i < len(x); i++ {
		if x[i] == r {
			x = append(x[:i], x[i+1:]...)
			i--
		}
	}
	return x
}

func find(x []Unit, r Unit) bool {
	for _, v := range x {
		if v == r {
			return true
		}
	}
	return false
}

func main() {
	archer := Archer{Unit: Unit{bombandStrength: 10}}
	tomp := Tomp{Unit: Unit{bombandStrength: 5}}
	army := NewArmy()
	army.addUnit(archer.Unit)
	army.addUnit(tomp.Unit)
	army.removeUnit(tomp.Unit)
	sum := 0
	for _, v := range army.units {
		sum += int(v.bombandStrength)
	}
	fmt.Println(sum)
}
