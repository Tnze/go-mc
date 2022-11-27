package maths

import (
	"math"
)

func RayTraceBlocks(rotation Vec2d, position Vec3d, maxDistance float32) []Vec3d {
	var result []Vec3d
	for i := float32(0); i < maxDistance; i += 0.005 {
		x := i * float32(math.Sin(ToRadians(float64(rotation.X)))*math.Cos(ToRadians(float64(rotation.Y))))
		y := float32(0)
		if rotation.Y > 0 {
			y = i * float32(-math.Sin(ToRadians(float64(rotation.Y))))
		} else {
			y = i * float32(math.Sin(ToRadians(float64(rotation.Y))))
		}
		z := i * float32(math.Cos(ToRadians(float64(rotation.X)))*math.Cos(ToRadians(float64(rotation.X))))
		result = append(result, Vec3d{X: position.X + x, Y: position.Y + y, Z: position.Z + z})
	}
	return result
}
