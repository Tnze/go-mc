package entity

import (
	"io"

	"github.com/Tnze/go-mc/data/entity"
	item "github.com/Tnze/go-mc/data/item"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

// BlockEntity describes the representation of a tile entity at a position.
type BlockEntity struct {
	ID string `nbt:"id"`

	// global co-ordinates
	X int `nbt:"x"`
	Y int `nbt:"y"`
	Z int `nbt:"z"`

	// sign-specific.
	Color string `nbt:"color"`
	Text1 string `nbt:"Text1"`
	Text2 string `nbt:"Text2"`
	Text3 string `nbt:"Text3"`
	Text4 string `nbt:"Text4"`
}

//Entity represents an instance of an entity.
type Entity struct {
	ID   int32
	Data int32
	Base *entity.Entity

	UUID uuid.UUID

	X, Y, Z          float64
	Pitch, Yaw       int8
	VelX, VelY, VelZ int16
	OnGround         bool

	IsLiving  bool
	HeadPitch int8
}

// The Slot data structure is how Minecraft represents an item and its associated data in the Minecraft Protocol
type Slot struct {
	Present bool
	ItemID  item.ID
	Count   int8
	NBT     pk.NBT
}

type SlotNBT struct {
	data interface{}
}

//Decode implement packet.FieldDecoder interface
func (s *Slot) ReadFrom(r io.Reader) (int64, error) {
	var itemID pk.VarInt
	n, err := pk.Tuple{
		(*pk.Boolean)(&s.Present),
		pk.Opt{
			Has: (*pk.Boolean)(&s.Present),
			Field: pk.Tuple{
				&itemID,
				(*pk.Byte)(&s.Count),
				&s.NBT,
			},
		},
	}.ReadFrom(r)
	if err != nil {
		return n, err
	}
	s.ItemID = item.ID(itemID)
	return n, nil
}

func (s Slot) WriteTo(w io.Writer) (int64, error) {
	return pk.Tuple{
		pk.Boolean(s.Present),
		pk.Opt{
			Has: (*pk.Boolean)(&s.Present),
			Field: pk.Tuple{
				pk.VarInt(s.ItemID),
				pk.Byte(s.Count),
				s.NBT,
			},
		},
	}.WriteTo(w)
}

func (s Slot) String() string {
	return item.ByID[s.ItemID].DisplayName
}
