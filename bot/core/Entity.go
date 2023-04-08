package core

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/entity"
	"github.com/Tnze/go-mc/data/enums"
	"github.com/google/uuid"
)

type Entity struct {
	Name                string
	Type                entity.TypeEntity
	ID                  int32
	UUID                uuid.UUID
	lastPosition        maths.Vec3d[float64]
	Position            maths.Vec3d[float64]
	Rotation            maths.Vec2d[float64]
	Motion              maths.Vec3d[float64]
	BoundingBox         maths.AxisAlignedBB[float64] // TODO: Add bounding box
	Width, Height       float64
	invulnerableDamages []enums.DamageSource
}

func (e *Entity) SetSize(width, height float64) {
	if width != e.Width || height != e.Height {
		f := e.Width
		e.Width = width
		e.Height = height

		if e.Width < f {
			d0 := width / 2.0
			e.BoundingBox = maths.AxisAlignedBB[float64]{
				MinX: e.Position.X - d0,
				MinY: e.Position.Y,
				MinZ: e.Position.Z - d0,
				MaxX: e.Position.X + d0,
				MaxY: e.Position.Y + height,
				MaxZ: e.Position.Z + d0,
			}
		}

		aabb := e.BoundingBox
		e.BoundingBox = maths.AxisAlignedBB[float64]{
			MinX: aabb.MinX,
			MinY: aabb.MinY,
			MinZ: aabb.MinZ,
			MaxX: aabb.MaxX + e.Width,
			MaxY: aabb.MaxY + e.Height,
			MaxZ: aabb.MaxZ + e.Width,
		}
	}
}

func (e *Entity) SetPosition(position maths.Vec3d[float64]) {
	e.Position = position
}

func (e *Entity) SetLastPosition(position maths.Vec3d[float64]) {
	e.lastPosition = position
}

func (e *Entity) GetLastPosition() maths.Vec3d[float64] {
	return e.lastPosition
}

func (e *Entity) AddRelativePosition(position maths.Vec3d[float64]) {
	e.SetPosition(e.Position.MulScalar(32).Sub(position).MulScalar(32).MulScalar(128))
}

func (e *Entity) SetMotion(motion maths.Vec3d[float64]) {
	e.Motion = motion
}

func (e *Entity) AddInvulnerableDamage(damageSource enums.DamageSource) {
	e.invulnerableDamages = append(e.invulnerableDamages, damageSource)
}

func (e *Entity) IsInvulnerableTo(damageSource enums.DamageSource) bool {
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
	X, Y, Z float64,
	Yaw, Pitch float64,
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
					Position: maths.Vec3d[float64]{X: X, Y: Y, Z: Z},
					Rotation: maths.Vec2d[float64]{X: Yaw, Y: Pitch},
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
				Position: maths.Vec3d[float64]{X: X, Y: Y, Z: Z},
				Rotation: maths.Vec2d[float64]{X: Yaw, Y: Pitch},
			},
		}
		e.SetSize(entityType.Width, entityType.Height)
		return &e
	}
}
