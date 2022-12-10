package transactions

import (
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

func SwitchSlot(n1 int32, item1 *slots.Slot, n2 int32, item2 *slots.Slot) []*SlotAction {
	return []*SlotAction{
		NewSlotAction(n1, int32(screen.LeftClick), 0, item1, &slots.Slot{Index: pk.Short(n1)}), // clear slot
		NewSlotAction(n2, int32(screen.LeftClick), 0, item1, item2),                            // move item1 to item2
		NewSlotAction(n1, int32(screen.LeftClick), 0, item2, &slots.Slot{Index: pk.Short(n2)}), // move item2 to item1
		NewSlotAction(-999, int32(screen.LeftClick), 4, &slots.Slot{}, &slots.Slot{}),          // exit
	}
}
