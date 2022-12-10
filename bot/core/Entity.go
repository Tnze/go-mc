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
	lastPosition        maths.Vec3d
	Position            maths.Vec3d
	Rotation            maths.Vec2d
	Motion              maths.Vec3d
	BoundingBox         AxisAlignedBB // TODO: Add bounding box
	Width, Height       float32
	invulnerableDamages []DamageSource
}

/*
SetSize

	@param width (float32) - the width to set
	@param height (float32) - the height to set
	@return none
*/
func (e *Entity) SetSize(width, height float32) {
	/*
		From net.minecraft.entity.Entity#setSize
	*/
	if width != e.Width || height != e.Height {
		f := e.Width
		e.Width = width
		e.Height = height

		if e.Width < f {
			d0 := width / 2.0
			e.BoundingBox = AxisAlignedBB{
				MinX: e.Position.X - d0,
				MinY: e.Position.Y,
				MinZ: e.Position.Z - d0,
				MaxX: e.Position.X + d0,
				MaxY: e.Position.Y + height,
				MaxZ: e.Position.Z + d0,
			}
		}

		aabb := e.BoundingBox
		e.BoundingBox = AxisAlignedBB{
			MinX: aabb.MinX,
			MinY: aabb.MinY,
			MinZ: aabb.MinZ,
			MaxX: aabb.MaxX + e.Width,
			MaxY: aabb.MaxY + e.Height,
			MaxZ: aabb.MaxZ + e.Width,
		}
	}
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
SetLastPosition

	@param position (maths.Vec3d) - the last position to set
	@return none
*/
func (e *Entity) SetLastPosition(position maths.Vec3d) {
	e.lastPosition = position
}

/*
GetLastPosition

	@param none
	@return maths.Vec3d - the last position
*/
func (e *Entity) GetLastPosition() maths.Vec3d {
	return e.lastPosition
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
SetMotion

	@param motion (maths.Vec3d) - the motion to set
	@return none
*/
func (e *Entity) SetMotion(motion maths.Vec3d) {
	e.Motion = motion
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
		e := EntityPlayer{
			EntityLiving: &EntityLiving{
				Entity: &Entity{
					Name:     entityType.Name,
					Type:     *entityType,
					ID:       EID,
					UUID:     EUUID,
					Position: maths.Vec3d{X: X, Y: Y, Z: Z},
					Rotation: maths.Vec2d{X: Yaw, Y: Pitch},
				},
			},
		}
		e.SetSize(entityType.Width, entityType.Height)
		return &e
	default:
		e := EntityLiving{
			Entity: &Entity{
				Name:     entityType.Name,
				Type:     *entityType,
				ID:       EID,
				UUID:     EUUID,
				Position: maths.Vec3d{X: X, Y: Y, Z: Z},
				Rotation: maths.Vec2d{X: Yaw, Y: Pitch},
			},
		}
		e.SetSize(entityType.Width, entityType.Height)
		return &e
	}
}
