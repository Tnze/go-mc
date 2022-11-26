package core

import (
	. "github.com/Tnze/go-mc/data/slots"
	"github.com/google/uuid"
)

type EntityPlayer struct {
	*EntityLiving
	UUID        uuid.UUID
	DisplayName string
	Cursor      *Slot
	Screens     map[int]Container
	expBar      float32
	TotalExp    int32
	Level       int32
}

func (e *EntityPlayer) SetExp(bar float32, exp, total int32) {
	e.expBar, e.TotalExp, e.Level = bar, total, exp
}
