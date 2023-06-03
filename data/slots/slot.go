package slots

import (
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type Slot struct {
	Index pk.Short // This is used with the transaction system
	ID    pk.VarInt
	Count pk.Byte
	NBT   nbt.RawMessage
}

func (s *Slot) GetIndex() int {
	return int(s.Index)
}

func (s *Slot) WriteTo(w io.Writer) (n int64, err error) {
	var present pk.Boolean = s.ID != 0 && s.Count != 0
	return pk.Tuple{
		present, pk.Opt{
			If: present,
			Value: pk.Tuple{
				&s.ID, &s.Count,
				pk.Opt{
					If:    s.NBT.Data == nil,
					Value: pk.Boolean(false),
					Else:  pk.NBT(&s.NBT),
				},
			},
		},
	}.WriteTo(w)
}

func (s *Slot) ReadFrom(r io.Reader) (n int64, err error) {
	var present pk.Boolean
	return pk.Tuple{
		&present, pk.Opt{
			If: &present,
			Value: pk.Tuple{
				&s.ID, &s.Count, pk.NBT(&s.NBT),
			},
		},
	}.ReadFrom(r)
}

type Container interface {
	GetInventorySlots() []Slot
	GetHotbarSlots() []Slot
	OnSetSlot(i int, s Slot) error
	OnClose() error
}

type ChangedSlots map[int]*Slot

func (c ChangedSlots) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.VarInt(len(c)).WriteTo(w)
	if err != nil {
		return
	}
	for i, v := range c {
		n1, err := pk.Short(i).WriteTo(w)
		if err != nil {
			return n + n1, err
		}
		n2, err := v.WriteTo(w)
		if err != nil {
			return n + n1 + n2, err
		}
		n += n1 + n2
	}
	return
}
