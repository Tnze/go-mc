package grids

import (
	"errors"
	. "github.com/Tnze/go-mc/data/slots"
)

type DispenserGrid struct {
	Slots [44]Slot
}

func (g *DispenserGrid) OnSetSlot(i int, s Slot) error {
	if i < 0 || i >= len(g.GetContentSlots()) {
		return errors.New("slot index out of bounds")
	}
	g.Slots[i] = s
	return nil
}

func (g *DispenserGrid) OnClose() error {
	return nil
}

func (g *DispenserGrid) GetContentSlots() []Slot { return g.Slots[:8] }

func (g *DispenserGrid) GetInventorySlots() []Slot { return g.Slots[9:35] }

func (g *DispenserGrid) GetHotbarSlots() []Slot { return g.Slots[36:] }
