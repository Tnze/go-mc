package core

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/effects"
	"github.com/Tnze/go-mc/data/item"
)

var EyePosVec = maths.Vec3d{Y: 1.62}
var EyePos = float32(1.62)

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
	ActivePotionEffects     []effects.Effect
	dead                    bool
}

/*
IsDead

	@return bool - if the entity health is less than the minimum health
*/
func (e *EntityLiving) IsDead() bool {
	return e.health <= e.minHealth
}

/*
IsPotionActive

	@param effect (effects.Effect) - the effect to check
	@return bool - if the entity has the effect
*/
func (e *EntityLiving) IsPotionActive(effect effects.Effect) bool {
	for _, v := range e.ActivePotionEffects {
		if v == effect {
			return true
		}
	}
	return false
}

/*
IsInvulnerableTo

	@param damageSource (DamageSource) - the damage source to check
	@return bool - if the entity is invulnerable to the damage source
*/
func (e *EntityLiving) IsInvulnerableTo(source DamageSource) bool {
	return e.Entity.IsInvulnerableTo(source)
}

/*
IsEntityInsideOpaqueBlock
	@return bool - if the entity is inside an opaque block
*/
/*func (e *EntityLiving) IsEntityInsideOpaqueBlock() bool {
	return e.Entity.IsEntityInsideOpaqueBlock()
}*/

/*
GetHealth

	@param absorption (bool) - if true, returns the total health (health + absorption)
	@return float64 - the health of the entity
*/
func (e *EntityLiving) GetHealth(absorption bool) float32 {
	if absorption {
		return e.health + e.Absorption
	}
	return e.health
}

/*
SetHealth

	@param health (float32) - the new health
	@return bool - if the player should respawn
*/
func (e *EntityLiving) SetHealth(health float32) bool {
	e.health = health
	if e.IsDead() {
		return true
	}
	return false
}

/*
GetEyePos

	@param partialTicks (float32) - the partial ticks
	@return Vec3d - the position of the entity's eyes
*/
func (e *EntityLiving) GetEyePos() maths.Vec3d {
	return e.Position.Add(EyePosVec)
}
