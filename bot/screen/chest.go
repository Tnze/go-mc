package screen

import (
	"errors"

	"github.com/Tnze/go-mc/data/inventory"
)

type Chest struct {
	Type  inventory.InventoryID
	Slots []Slot
	Rows  int
}

func (c *Chest) onSetSlot(i int, slot Slot) error {
	if i < 0 || i >= len(c.Slots) {
		return errors.New("slot index out of bounds")
	}
	c.Slots[i] = slot
	return nil
}

func (c *Chest) onClose() error {
	return nil
}

func (c *Chest) Container() []Slot {
	return c.Slots[0 : c.Rows*9]
}

func (c *Chest) Main() []Slot {
	return c.Slots[c.Rows*9 : c.Rows*9+27]
}

func (c *Chest) Hotbar() []Slot {
	return c.Slots[c.Rows*9+27 : (c.Rows+4)*9]
}
