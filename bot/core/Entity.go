package core

import (
	"fmt"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/entity"
	"github.com/google/uuid"
)

type Entity struct {
	Name                string
	Type                entity.TypeEntity
	ID                  int32
	UUID                uuid.UUID
	Position            maths.Vec3d
	Rotation            maths.Vec2d
	Motion              maths.Vec3d
	BoundingBox         AxisAlignedBB // TODO: Add bounding box
	invulnerableDamages []DamageSource
}

/*
SetPosition

	@param position (maths.Vec3d) - the position to set
	@return none
*/
func (e *Entity) SetPosition(position maths.Vec3d) {
	e.Position = position
}

/*
AddRelativePosition

	@param position (maths.Vec3d) - the position to add
	@return none
*/
func (e *Entity) AddRelativePosition(position maths.Vec3d) {
	fmt.Println("Before:", e.Position)
	e.SetPosition(e.Position.MulScalar(32).Sub(position).MulScalar(32).MulScalar(128))
	fmt.Println("After:", e.Position)
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

func NewEntity(
	EID int32,
	EUUID uuid.UUID,
	Type int32,
	X, Y, Z float32,
	Yaw, Pitch float32,
) interface{} {
	entityType := entity.TypeEntityByID[Type]
	switch *entityType {
	case entity.Player:
		return EntityPlayer{
			EntityLiving: &EntityLiving{
				Entity: &Entity{
					Name:     entityType.Name,
					Type:     *entityType,
					ID:       EID,
					UUID:     EUUID,
					Position: maths.Vec3d{X: X, Y: Y, Z: Z},
					Rotation: maths.Vec2d{X: Yaw, Y: Pitch},
					/*
						BoundingBox
					*/
				},
			},
		}
	default:
		return EntityLiving{
			Entity: &Entity{
				Name:     entityType.Name,
				Type:     *entityType,
				ID:       EID,
				UUID:     EUUID,
				Position: maths.Vec3d{X: X, Y: Y, Z: Z},
				Rotation: maths.Vec2d{X: Yaw, Y: Pitch},
				/*
					BoundingBox
				*/
			},
		}
	}
}
