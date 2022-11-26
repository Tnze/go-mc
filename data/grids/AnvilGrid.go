package grids

import (
	"errors"
	. "github.com/Tnze/go-mc/data/slots"
)

type AnvilGrid struct {
	Slots [38]Slot
}

func (g *AnvilGrid) OnSetSlot(i int, s Slot) error {
	if i < 0 || i >= len(g.GetContentSlots()) {
		return errors.New("slot index out of bounds")
	}
	g.Slots[i] = s
	return nil
}

func (g *AnvilGrid) OnClose() error {
	return nil
}

func (g *AnvilGrid) GetFirstInputSlot() *Slot { return &g.Slots[0] }

func (g *AnvilGrid) GetSecondInputSlot() *Slot { return &g.Slots[1] }

func (g *AnvilGrid) GetOutputSlot() *Slot { return &g.Slots[2] }

func (g *AnvilGrid) GetContentSlots() []Slot { return g.Slots[:3] }

func (g *AnvilGrid) GetInventorySlots() []Slot { return g.Slots[3:29] }

func (g *AnvilGrid) GetHotbarSlots() []Slot { return g.Slots[30:] }
