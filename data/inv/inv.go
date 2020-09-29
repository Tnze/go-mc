// Package inv maps window types to inventory slot information.
package inv

type Info struct {
	Name       string
	Start, End int // Player inventory
	Slots      int
}

func (i Info) PlayerInvStart() int {
	return i.Start
}

func (i Info) PlayerInvEnd() int {
	return i.End
}

func (i Info) HotbarIdx(place int) int {
	return i.End - (8 - place)
}

var ByType = map[int]Info{
	-1: Info{Name: "inventory", Start: 9, End: 44, Slots: 46},
	0:  Info{Name: "generic_9x1", Start: 1 * 9, End: 1*9 + 35, Slots: 1*9 + 36},
	1:  Info{Name: "generic_9x2", Start: 2 * 9, End: 2*9 + 35, Slots: 2*9 + 36},
	2:  Info{Name: "generic_9x3", Start: 3 * 9, End: 3*9 + 35, Slots: 3*9 + 36},
	3:  Info{Name: "generic_9x4", Start: 4 * 9, End: 4*9 + 35, Slots: 4*9 + 36},
	4:  Info{Name: "generic_9x5", Start: 5 * 9, End: 5*9 + 35, Slots: 5*9 + 36},
	5:  Info{Name: "generic_9x6", Start: 6 * 9, End: 6*9 + 35, Slots: 6*9 + 36},
}
