package physics

import (
	"github.com/Tnze/go-mc/bot/core"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/bot/world"
)

type ExplosionStrength float64

const (
	EndCrystal     ExplosionStrength = 6.0
	ChargedCreeper ExplosionStrength = 6.0
	Bed            ExplosionStrength = 5.0
	TNT            ExplosionStrength = 4.0
	Creeper        ExplosionStrength = 3.0
	WitherSkull    ExplosionStrength = 1.0
	Fireball       ExplosionStrength = 1.0
)

func (e ExplosionStrength) GetExplosionRadius() float64 {
	return float64(1.3 * (e / 0.225) * 0.3)
}

type Explosion struct {
	// The position of the explosion
	Position maths.Vec3d[float64]
	// Strength of the explosion
	Strength ExplosionStrength
}

func (e Explosion) GetAffectedEntities() []maths.Vec3d[float64] { return nil }

func (e Explosion) CalculateDamage(pos maths.Vec3d[float64], entity core.EntityLiving, world *world.World, explosionType ExplosionStrength) float64 {
	radius := explosionType.GetExplosionRadius()
	distance := entity.Position.DistanceTo(pos) / radius

	blockDensity := world.GetBlockDensity(pos, entity.BoundingBox)
	v := (1.0 - distance) * blockDensity
	damage := (v*v+v)*8*explosionType.GetExplosionRadius() + 1

	return damage // TODO: Add armor and resistance
}
