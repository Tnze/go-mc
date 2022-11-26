package grids

import (
	"errors"
	. "github.com/Tnze/go-mc/data/slots"
)

type LargeChestGrid struct {
	Slots [89]Slot
}

func (c *LargeChestGrid) OnSetSlot(i int, s Slot) error {
	if i < 0 || i >= len(c.GetContentSlots()) {
		return errors.New("slot index out of bounds")
	}
	c.Slots[i] = s
	return nil
}

func (c *LargeChestGrid) OnClose() error {
	return nil
}

func (c *LargeChestGrid) GetContentSlots() []Slot { return c.Slots[:53] }

func (c *LargeChestGrid) GetInventorySlots() []Slot { return c.Slots[54:80] }

func (c *LargeChestGrid) GetHotbarSlots() []Slot { return c.Slots[81:] }
