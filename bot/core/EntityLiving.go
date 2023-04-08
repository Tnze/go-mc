package core

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/effects"
	"github.com/Tnze/go-mc/data/item"
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
	ActivePotionEffects     []effects.EffectStatus
	dead                    bool
	OnGround                bool
}

func (e *EntityLiving) IsDead() bool {
	return e.health <= e.minHealth
}

func (e *EntityLiving) IsPotionActive(effect effects.EffectStatus) bool {
	for _, v := range e.ActivePotionEffects {
		if v == effect {
			return true
		}
	}
	return false
}

func (e *EntityLiving) IsInvulnerableTo(source DamageSource) bool {
	return e.Entity.IsInvulnerableTo(source)
}

/*func (e *EntityLiving) IsEntityInsideOpaqueBlock() bool {
	return e.Entity.IsEntityInsideOpaqueBlock()
}*/

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
