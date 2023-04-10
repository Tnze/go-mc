package core

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/effects"
	"github.com/Tnze/go-mc/data/enums"
	"github.com/Tnze/go-mc/data/item"
	"github.com/google/uuid"
)

var EyePosVec = maths.Vec3d[float64]{Y: 1.62}
var EyePos = 1.62

type EntityLiving struct {
	*Entity
	health                  float32
	minHealth               float32
	maxHealth               float32
	Food                    int32
	MaxFood                 int32
	Saturation              float32
	Absorption              float32
	ActiveItem              item.Item
	ActiveItemStackUseCount int32
	ActivePotionEffects     map[int32]*effects.EffectStatus
	dead                    bool
	OnGround                bool
	MoveStrafing            float32
	MoveForward             float32
	MoveVertical            float32
}

type EntityLivingInterface interface {
	EntityInterface
	GetHealth(absorption bool) float32
	SetHealth(health float32) bool
	GetEyePos() maths.Vec3d[float64]
	IsDead() bool
	IsPotionActive(effect effects.Effect) bool
	GetPotionEffect(effect effects.Effect) *effects.EffectStatus
	IsInvulnerableTo(source enums.DamageSource) bool
	//IsEntityInsideOpaqueBlock() bool
}

func (e *EntityLiving) GetHealth(absorption bool) float32 {
	if absorption {
		return e.health + e.Absorption
	}
	return e.health
}

func (e *EntityLiving) SetHealth(health float32) bool {
	e.health = health
	if e.IsDead() {
		return true
	}
	return false
}

func (e *EntityLiving) GetEyePos() maths.Vec3d[float64] {
	return e.Position.Add(EyePosVec)
}

func (e *EntityLiving) IsDead() bool {
	return e.health <= e.minHealth
}

func (e *EntityLiving) IsPotionActive(effect effects.Effect) bool {
	_, ok := e.ActivePotionEffects[effect.ID]
	return ok
}

func (e *EntityLiving) GetPotionEffect(effect effects.Effect) *effects.EffectStatus {
	return e.ActivePotionEffects[effect.ID]
}

func (e *EntityLiving) IsInvulnerableTo(source enums.DamageSource) bool {
	return e.Entity.IsInvulnerableTo(source)
}

func (e *EntityLiving) IsLivingEntity() bool {
	return true
}

func (e *EntityLiving) IsPlayerEntity() bool {
	return false
}

/*func (e *EntityLiving) IsEntityInsideOpaqueBlock() bool {
	return e.Entity.IsEntityInsideOpaqueBlock()
}*/

func NewEntityLiving(EID int32,
	EUUID uuid.UUID,
	Type int32,
	X, Y, Z float64,
	Yaw, Pitch float64,
) *EntityLiving {
	return &EntityLiving{
		Entity:              NewEntity(EID, EUUID, Type, X, Y, Z, Yaw, Pitch),
		ActivePotionEffects: make(map[int32]*effects.EffectStatus),
	}
}
