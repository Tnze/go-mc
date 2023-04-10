package core

import (
	. "github.com/Tnze/go-mc/data/slots"
	"github.com/google/uuid"
)

type EntityPlayer struct {
	*EntityLiving
	Cursor      *Slot
	Screens     map[int]Container
	DisplayName string
	expBar      float32
	TotalExp    int32
	Level       int32
}

type EntityPlayerInterface interface {
	EntityLivingInterface
	GetDisplayName() string
	GetCursor() *Slot
	GetScreens() map[int]Container
	GetExp() (float32, int32, int32)
	SetExp(bar float32, exp, total int32)
}

func (e *EntityPlayer) GetDisplayName() string {
	return e.DisplayName
}

func (e *EntityPlayer) GetCursor() *Slot {
	return e.Cursor
}

func (e *EntityPlayer) GetScreens() map[int]Container {
	return e.Screens
}

func (e *EntityPlayer) GetExp() (float32, int32, int32) {
	return e.expBar, e.TotalExp, e.Level
}

func (e *EntityPlayer) SetExp(bar float32, exp, total int32) {
	e.expBar, e.TotalExp, e.Level = bar, total, exp
}

func (e *EntityPlayer) IsPlayerEntity() bool {
	return true
}

func NewEntityPlayer(EID int32,
	EUUID uuid.UUID,
	Type int32,
	X, Y, Z float64,
	Yaw, Pitch float64,
) *EntityPlayer {
	return &EntityPlayer{
		EntityLiving: NewEntityLiving(EID, EUUID, Type, X, Y, Z, Yaw, Pitch),
		Screens:      make(map[int]Container),
	}
}
