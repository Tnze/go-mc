package screen

import (
	"errors"
	item2 "github.com/Tnze/go-mc/data/item"
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Inventory struct {
	Slots [46]slots.Slot
}

func (inv *Inventory) onClose() error {
	return nil
}

func (inv *Inventory) OnSetSlot(i int, s slots.Slot) error {
	if i < 0 || i >= len(inv.Slots) {
		return errors.New("slot index out of bounds")
	}
	inv.Slots[i] = s
	return nil
}

func (inv *Inventory) CraftingOutput() *slots.Slot { return &inv.Slots[0] }
func (inv *Inventory) CraftingInput() []slots.Slot { return inv.Slots[1 : 1+4] }

// Armor returns to the armor section of the Inventory.
// The length is 4, which are head, chest, legs and feet.
func (inv *Inventory) Armor() []slots.Slot  { return inv.Slots[5 : 5+4] }
func (inv *Inventory) Main() []slots.Slot   { return inv.Slots[9 : 9+3*9] }
func (inv *Inventory) Hotbar() []slots.Slot { return inv.Slots[36 : 36+9] }
func (inv *Inventory) Offhand() *slots.Slot { return &inv.Slots[45] }

/*
GetItemSlotById returns the slot of the item in the inventory.

	@param itemID the item ID
	@return the slot of the item, -1 if not found
*/
func (inv *Inventory) GetItemSlotById(itemID int32) int {
	for i, slot := range inv.Slots {
		if slot.ID == pk.VarInt(itemID) {
			return i
		}
	}
	return -1
}

/*
GetItemSlotByName returns the slot of the item in the inventory.

	@param itemName the item name
	@return the slot of the item, -1 if not found
*/
func (inv *Inventory) GetItemSlotByName(itemName string) int {
	if item, ok := item2.ByName[itemName]; ok {
		return inv.GetItemSlotById(int32(item.ID))
	} else {
		return -1
	}
}

/*
GetHotbarSlotById returns the slot of the item in the hotbar.

	@param ID the slot ID
	@return the slot of the item, -1 if not found
*/
func (inv *Inventory) GetHotbarSlotById(ID uint8) (slots.Slot, error) {
	if ID > 8 {
		return slots.Slot{}, errors.New("slot ID out of bounds")
	}
	return inv.Hotbar()[ID], nil
}
