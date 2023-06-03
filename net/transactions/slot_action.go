package transactions

import (
	"github.com/Tnze/go-mc/data/slots"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type SlotAction struct {
	Slot    pk.Short
	Button  pk.Byte
	Mode    pk.VarInt
	Item    *slots.Slot
	Changed []*slots.Slot
}

func NewSlotAction(slot, button, mode int, cursor *slots.Slot, items ...*slots.Slot) *SlotAction {
	return &SlotAction{
		Slot:    pk.Short(slot),
		Button:  pk.Byte(button),
		Mode:    pk.VarInt(mode),
		Item:    cursor,
		Changed: items,
	}
}

func (s *SlotAction) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.Tuple{
		&s.Slot, &s.Button, &s.Mode,
	}.WriteTo(w)
	n0, err := pk.VarInt(len(s.Changed)).WriteTo(w)
	if err != nil {
		return n + n0, err
	}
	for _, v := range s.Changed {
		n1, err := v.Index.WriteTo(w)
		if err != nil {
			return n + n1, err
		}
		n2, err := v.WriteTo(w)
		if err != nil {
			return n + n1 + n2, err
		}
		n += n1 + n2
	}
	n3, err := s.Item.WriteTo(w)
	n += n0 + n3
	return
}

/*func (s *SlotAction) Validate() basic.Error {
	if s.Slot < 0 {
		return basic.NewError(basic.InvalidSlot, fmt.Sprintf("slot %d is less than 0", s.Slot))
	}
	if s.Button < 0 {
		return basic.NewError(basic.InvalidButton, fmt.Sprintf("button %d is less than 0", s.Button))
	}
	if s.Mode < 0 || s.Mode > 2 {
		return basic.NewError(basic.InvalidMode, fmt.Sprintf("mode %d is not in range [0, 2]", s.Mode))
	}
	return basic.NoError
}*/
