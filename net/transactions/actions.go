package transactions

import (
	"github.com/Tnze/go-mc/bot/screen"
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

func LeftClick(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.LeftClick.Int(), 0, &slots.Slot{Index: item.Index})).
		Build()
}

func DoubleClick(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.LeftClick.Int(), 6, &slots.Slot{Index: item.Index})).
		Build()
}

func RightClick(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.RightClick.Int(), 0, &slots.Slot{Index: item.Index})).
		Build()
}

func Drop(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.Drop.Int(), 4, &slots.Slot{Index: item.Index})).
		Build()
}

func DropAll(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.ControlDrop.Int(), 4, &slots.Slot{Index: item.Index})).
		Build()
}

func Swap(item1 *slots.Slot, item2 *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item1.GetIndex(), screen.LeftClick.Int(), 0, item1, &slots.Slot{Index: item2.Index})).
		AddAction(NewSlotAction(item2.GetIndex(), screen.LeftClick.Int(), 0, item1, item2)).                           // move item1 to item2
		AddAction(NewSlotAction(item1.GetIndex(), screen.LeftClick.Int(), 0, item2, &slots.Slot{Index: item2.Index})). // move item2 to item1
		AddAction(NewSlotAction(-999, screen.LeftClick.Int(), 4, &slots.Slot{}, &slots.Slot{})).                       // exit
		Build()
}

func SwapWithHotbar(item *slots.Slot, hotbarIndex int) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.LeftClick.Int(), 0, item, &slots.Slot{Index: pk.Short(hotbarIndex)})).
		AddAction(NewSlotAction(hotbarIndex, screen.LeftClick.Int(), 0, item, &slots.Slot{Index: item.Index})).
		Build()
}

func SwapWithOffhand(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.SwapHand.Int(), 2, item, &slots.Slot{Index: 40})).
		Build()
}

func QuickMove(item *slots.Slot) *Transaction {
	return NewTransactionBuilder().
		AddAction(NewSlotAction(item.GetIndex(), screen.ShiftLeftClick.Int(), 1, &slots.Slot{Index: item.Index})).
		Build()
}
