package core

import "github.com/Tnze/go-mc/bot/maths"

type Entity struct {
	Name                string
	EntityID            int32
	Position            maths.Vec3d
	Rotation            maths.Vec2d
	invulnerableDamages []DamageSource
}

/*
AddInvulnerableDamage

	@param damageSource (DamageSource) - the damage source to add
	@return none
*/
func (e *Entity) AddInvulnerableDamage(damageSource DamageSource) {
	e.invulnerableDamages = append(e.invulnerableDamages, damageSource)
}

/*
IsInvulnerableTo

	@param damageSource (DamageSource) - the damage source to check
	@return bool - if the entity is invulnerable to the damage source
*/
func (e *Entity) IsInvulnerableTo(damageSource DamageSource) bool {
	for _, v := range e.invulnerableDamages {
		if v == damageSource {
			return true
		}
	}
	return false
}
