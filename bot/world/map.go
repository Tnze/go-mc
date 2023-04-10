package world

import (
	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"io"
)

type Map struct {
	MapID       int64
	Scale       int8
	Locked      bool
	TrackingPos bool
	IconCount   int32
	Icons       []MapIcon
	Columns     uint8
	Rows        uint8
	X           int8
	Z           int8
	Length      int32
	Data        []byte
}

func (m *Map) ReadFrom(r io.Reader) (int64, error) {
	n, err := pk.Tuple{
		(*pk.VarLong)(&m.MapID),
		(*pk.Byte)(&m.Scale),
		(*pk.Boolean)(&m.Locked),
		(*pk.Boolean)(&m.TrackingPos),
		(*pk.VarInt)(&m.IconCount),
	}.ReadFrom(r)

	if err != nil {
		return n, err
	}

	m.Icons = make([]MapIcon, m.IconCount)
	for i := range m.Icons {
		n1, err := m.Icons[i].ReadFrom(r)
		n += n1
		if err != nil {
			return n, err
		}
	}

	n1, err := (*pk.UnsignedByte)(&m.Columns).ReadFrom(r)
	n += n1

	n2, err := pk.Opt{
		If: m.Columns > 0,
		Value: pk.Tuple{
			(*pk.UnsignedByte)(&m.Rows),
			(*pk.Byte)(&m.X),
			(*pk.Byte)(&m.Z),
			(*pk.VarInt)(&m.Length),
		},
	}.ReadFrom(r)
	n += n2

	m.Data = make([]byte, m.Length)
	n3, err := (*pk.ByteArray)(&m.Data).ReadFrom(r)
	n += n3

	return n, err
}

type MapIcon struct {
	Type      int32
	X, Z      int8
	Direction int8
	HasName   bool
	Name      chat.Message
}

func (m *MapIcon) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		(*pk.VarInt)(&m.Type),
		(*pk.Byte)(&m.X),
		(*pk.Byte)(&m.Z),
		(*pk.Byte)(&m.Direction),
		(*pk.Boolean)(&m.HasName),
		pk.Opt{
			If: m.HasName,
			Value: pk.Tuple{
				&m.Name,
			},
		},
	}.ReadFrom(r)
}
