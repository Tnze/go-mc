package screen

import "errors"

type Inventory struct {
	Slots [46]Slot
}

func (inv *Inventory) onClose() error {
	return nil
}

func (inv *Inventory) onSetSlot(i int, s Slot) error {
	if i < 0 || i >= len(inv.Slots) {
		return errors.New("slot index out of bounds")
	}
	inv.Slots[i] = s
	return nil
}

func (inv *Inventory) CraftingOutput() *Slot { return &inv.Slots[0] }
func (inv *Inventory) CraftingInput() []Slot { return inv.Slots[1 : 1+4] }

// Armor returns to the armor section of the Inventory.
// The length is 4, which are head, chest, legs and feet.
func (inv *Inventory) Armor() []Slot  { return inv.Slots[5 : 5+4] }
func (inv *Inventory) Main() []Slot   { return inv.Slots[9 : 9+3*9] }
func (inv *Inventory) Hotbar() []Slot { return inv.Slots[36 : 36+9] }
func (inv *Inventory) Offhand() *Slot { return &inv.Slots[45] }
