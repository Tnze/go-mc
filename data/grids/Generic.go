package grids

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Generic struct {
	Width, Height int          // Width and Height of the grid. Used to determine the offset of the non-content slots.
	Slots         []slots.Slot // Will be initialized in InitGenericContainer
}

func InitGenericContainer(size, width, height int) *Generic {
	return &Generic{
		Width:  width,
		Height: height,
		Slots:  make([]slots.Slot, size),
	}
}

func (g *Generic) OnClose() basic.Error {
	return basic.Error{Err: basic.NoError, Info: nil}
}

/* Slot data */

func (g *Generic) getSize() int                  { return g.Width * g.Height }
func (g *Generic) GetContentSlots() []slots.Slot { return g.Slots[:g.getSize()] }
func (g *Generic) GetInventorySlots() []slots.Slot {
	return g.Slots[g.getSize() : g.getSize()+35]
}
func (g *Generic) GetHotbarSlots() []slots.Slot {
	return g.Slots[len(g.Slots)-9:]
}

/* Getter & Setter */

func (g *Generic) GetSlot(i int) *slots.Slot { return &g.Slots[i] }
func (g *Generic) SetSlot(i int, s slots.Slot) basic.Error {
	if i < 0 || i >= len(g.Slots) {
		return basic.Error{Err: basic.OutOfBound, Info: fmt.Errorf("slot index %d out of bounds. maximum index is %d", i, len(g.Slots)-1)}
	}
	g.Slots[i] = s
	return basic.Error{Err: basic.NoError, Info: nil}
}

// TODO: Iterator for slots

func (g *Generic) GetInventorySlot(id int) *slots.Slot {
	for i := range g.GetInventorySlots() {
		if g.Slots[i].ID == pk.VarInt(id) {
			return &g.Slots[i]
		}
	}
	return nil
}
func (g *Generic) GetHotbarSlot(id int) *slots.Slot {
	for i := range g.GetHotbarSlots() {
		if g.Slots[i].ID == pk.VarInt(id) {
			return &g.Slots[i]
		}
	}
	return nil
}
