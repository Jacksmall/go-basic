package combination

import "errors"

type UnitInterface interface {
	addUnit(unit Unit)
	removeUnit(unit Unit) (bool, error)
	bombardStrength() uint
}

type Unit struct{}

type Army struct {
	units []Unit
}

func (a Army) addUnit(unit Unit) {
	if isExist := find(a.units, unit); !isExist {
		a.units = append(a.units, unit)
	}
}

func (a Army) removeUnit(unit Unit) (bool, error) {
	removeSlice(a.units, unit)
	return true, nil
}

func (a Army) bombardStrength() uint {
	return 4
}

type Archer struct{}

func (ac Archer) addUnit(unit Unit) {}
func (ac Archer) removeUnit(unit Unit) (bool, error) {
	return false, errors.New("archer is leaf")
}
func (ac Archer) bombardStrength() uint {
	return 0
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
