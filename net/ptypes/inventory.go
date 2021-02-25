// Package ptypes implements encoding and decoding for high-level packets.
package ptypes

import (
	"errors"
	"io"

	"github.com/Tnze/go-mc/bot/world/entity"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
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

func (p *WindowItems) ReadFrom(r io.Reader) (int64, error) {
	var count pk.Short
	return pk.Tuple{
		&p.WindowID,
		&count,
		&pk.Ary{
			Len: &count,
			Ary: []entity.Slot{},
		},
	}.ReadFrom(r)
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

type ConfirmTransaction struct {
	WindowID pk.Byte
	ActionID pk.Short
	Accepted pk.Boolean
}

func (p *ConfirmTransaction) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.WindowID, &p.ActionID, &p.Accepted)
}

func (p ConfirmTransaction) Encode() pk.Packet {
	return pk.Marshal(
		packetid.TransactionServerbound,
		p.WindowID,
		p.ActionID,
		p.Accepted,
	)
}
