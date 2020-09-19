package entity

import (
	"github.com/Tnze/go-mc/data"
	"github.com/Tnze/go-mc/data/entity"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

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
	ItemID  int32
	Count   int8
	NBT     interface{}
}

//Decode implement packet.FieldDecoder interface
func (s *Slot) Decode(r pk.DecodeReader) error {
	if err := (*pk.Boolean)(&s.Present).Decode(r); err != nil {
		return err
	}
	if s.Present {
		if err := (*pk.VarInt)(&s.ItemID).Decode(r); err != nil {
			return err
		}
		if err := (*pk.Byte)(&s.Count).Decode(r); err != nil {
			return err
		}
		if err := nbt.NewDecoder(r).Decode(&s.NBT); err != nil {
			return err
		}
	}
	return nil
}

func (s Slot) String() string {
	return data.ItemNameByID[s.ItemID]
}
