// Package ptypes implements encoding and decoding for high-level packets.
package ptypes

import (
	"bytes"
	"errors"

	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// SetSlot is a clientbound packet which configures an inventory slot.
// A window ID of -1 represents the cursor, and a window ID of 0 represents
// the players inventory.
type SetSlot struct {
	WindowID pk.Byte
	Slot     pk.Short
	SlotData entity.Slot
}

func (p *SetSlot) Decode(pkt pk.Packet) error {
	if err := pkt.Scan(&p.WindowID, &p.Slot, &p.SlotData); err != nil && !errors.Is(err, nbt.ErrEND) {
		return err
	}
	return nil
}

// WindowItems is a clientbound packet describing the contents of multiple
// inventory slots in a window/inventory.
type WindowItems struct {
	WindowID pk.Byte
	Slots    []entity.Slot
}

func (p *WindowItems) Decode(pkt pk.Packet) error {
	r := bytes.NewReader(pkt.Data)
	if err := p.WindowID.Decode(r); err != nil {
		return err
	}

	var count pk.Short
	if err := count.Decode(r); err != nil {
		return err
	}

	p.Slots = make([]entity.Slot, int(count))
	for i := 0; i < int(count); i++ {
		if err := p.Slots[i].Decode(r); err != nil && !errors.Is(err, nbt.ErrEND) {
			return err
		}
	}
	return nil
}

// OpenWindow is a clientbound packet which opens an inventory.
type OpenWindow struct {
	WindowID   pk.VarInt
	WindowType pk.VarInt
	Title      chat.Message
}

func (p *OpenWindow) Decode(pkt pk.Packet) error {
	if err := pkt.Scan(&p.WindowID, &p.WindowType, &p.Title); err != nil && !errors.Is(err, nbt.ErrEND) {
		return err
	}
	return nil
}
