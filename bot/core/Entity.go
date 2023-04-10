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
	BoundingBox         maths.AxisAlignedBB[float64]
	Width, Height       float64
	dataManager         map[int32]interface{}
	invulnerableDamages []enums.DamageSource
}

type EntityInterface interface {
	GetName() string
	GetType() entity.TypeEntity
	GetID() int32
	GetUUID() uuid.UUID
	GetPosition() maths.Vec3d[float64]
	GetRotation() maths.Vec2d[float64]
	GetMotion() maths.Vec3d[float64]
	GetBoundingBox() maths.AxisAlignedBB[float64]
	GetWidth() float64
	GetHeight() float64
	GetDataManager() map[int32]interface{}
	GetInvulnerableDamages() []enums.DamageSource
	SetPosition(x, y, z float64)
	SetRotation(yaw, pitch float64)
	SetMotion(x, y, z float64)
	SetSize(width, height float64)
	IsInvulnerableTo(source enums.DamageSource) bool
	IsLivingEntity() bool
	IsPlayerEntity() bool
	//IsEntityInsideOpaqueBlock() bool
}

func (e *Entity) GetName() string {
	return e.Name
}

func (e *Entity) GetType() entity.TypeEntity {
	return e.Type
}

func (e *Entity) GetID() int32 {
	return e.ID
}

func (e *Entity) GetUUID() uuid.UUID {
	return e.UUID
}

func (e *Entity) GetPosition() maths.Vec3d[float64] {
	return e.Position
}

func (e *Entity) GetRotation() maths.Vec2d[float64] {
	return e.Rotation
}

func (e *Entity) GetMotion() maths.Vec3d[float64] {
	return e.Motion
}

func (e *Entity) GetBoundingBox() maths.AxisAlignedBB[float64] {
	return e.BoundingBox

}

func (e *Entity) GetWidth() float64 {
	return e.Width
}

func (e *Entity) GetHeight() float64 {
	return e.Height
}

func (e *Entity) GetDataManager() map[int32]interface{} {
	return e.dataManager
}

func (e *Entity) GetInvulnerableDamages() []enums.DamageSource {
	return e.invulnerableDamages
}

func (e *Entity) IsInvulnerableTo(damageSource enums.DamageSource) bool {
	for _, v := range e.invulnerableDamages {
		if v == damageSource {
			return true
		}
	}
	return false
}

func (e *Entity) IsLivingEntity() bool {
	return false
}

func (e *Entity) IsPlayerEntity() bool {
	return false
}

func (e *Entity) SetPosition(x, y, z float64) {
	e.lastPosition = e.Position
	e.Position = maths.Vec3d[float64]{X: x, Y: y, Z: z}
}

func (e *Entity) SetRotation(yaw, pitch float64) {
	e.Rotation = maths.Vec2d[float64]{X: yaw, Y: pitch}
}

func (e *Entity) SetMotion(x, y, z float64) {
	e.Motion = maths.Vec3d[float64]{X: x, Y: y, Z: z}
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

func (e *Entity) AddRelativePosition(position maths.Vec3d[float64]) {
	e.SetPosition(e.Position.MulScalar(32).Sub(position).MulScalar(32).MulScalar(128).Spread())
}

func (e *Entity) AddInvulnerableDamage(damageSource enums.DamageSource) {
	e.invulnerableDamages = append(e.invulnerableDamages, damageSource)
}

func NewEntity(
	EID int32,
	EUUID uuid.UUID,
	Type int32,
	X, Y, Z float64,
	Yaw, Pitch float64,
) *Entity {
	entityType := entity.TypeEntityByID[Type]
	e := &Entity{
		Name:     entityType.Name,
		Type:     *entityType,
		ID:       EID,
		UUID:     EUUID,
		Position: maths.Vec3d[float64]{X: X, Y: Y, Z: Z},
		Rotation: maths.Vec2d[float64]{X: Yaw, Y: Pitch},
	}
	e.SetSize(entityType.Width, entityType.Height)
	return e
}
