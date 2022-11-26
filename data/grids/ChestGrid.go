package grids

import (
	"errors"
	. "github.com/Tnze/go-mc/data/slots"
)

type ChestGrid struct {
	Slots [62]Slot
}

func (c *ChestGrid) OnSetSlot(i int, s Slot) error {
	if i < 0 || i >= len(c.GetContentSlots()) {
		return errors.New("slot index out of bounds")
	}
	c.Slots[i] = s
	return nil
}

func (c *ChestGrid) OnClose() error {
	return nil
}

func (c *ChestGrid) GetContentSlots() []Slot { return c.Slots[:26] }

func (c *ChestGrid) GetInventorySlots() []Slot { return c.Slots[27:53] }

func (c *ChestGrid) GetHotbarSlots() []Slot { return c.Slots[54:] }
