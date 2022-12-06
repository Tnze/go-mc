package maths

import (
	"math"
)

var NullVec3d = Vec3d{X: 0, Y: 0, Z: 0}

type Vec3d struct {
	X, Y, Z float32
}

func (v Vec3d) Add(vec3d Vec3d) Vec3d {
	return Vec3d{X: v.X + vec3d.X, Y: v.Y + vec3d.Y, Z: v.Z + vec3d.Z}
}

func (v Vec3d) AddScalar(scalar float32) Vec3d {
	return Vec3d{X: v.X + scalar, Y: v.Y + scalar, Z: v.Z + scalar}
}

func (v Vec3d) Sub(vec3d Vec3d) Vec3d {
	return Vec3d{X: v.X - vec3d.X, Y: v.Y - vec3d.Y, Z: v.Z - vec3d.Z}
}

func (v Vec3d) SubScalar(scalar float32) Vec3d {
	return Vec3d{X: v.X - scalar, Y: v.Y - scalar, Z: v.Z - scalar}
}

func (v Vec3d) Mul(vec3d Vec3d) Vec3d {
	return Vec3d{X: v.X * vec3d.X, Y: v.Y * vec3d.Y, Z: v.Z * vec3d.Z}
}

func (v Vec3d) MulScalar(scalar float32) Vec3d {
	return Vec3d{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v Vec3d) Div(vec3d Vec3d) Vec3d {
	return Vec3d{X: v.X / vec3d.X, Y: v.Y / vec3d.Y, Z: v.Z / vec3d.Z}
}

func (v Vec3d) DivScalar(scalar float32) Vec3d {
	return Vec3d{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}

func (v Vec3d) DistanceTo(vec3d Vec3d) float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X-vec3d.X), 2) + math.Pow(float64(v.Y-vec3d.Y), 2) + math.Pow(float64(v.Z-vec3d.Z), 2)))
}

func (v Vec3d) Center() Vec3d {
	return Vec3d{X: v.X + 0.5, Y: v.Y + 0.5, Z: v.Z + 0.5}
}

func (v Vec3d) Offset(x, y, z float32) Vec3d {
	return Vec3d{X: v.X + x, Y: v.Y + y, Z: v.Z + z}
}

func (v Vec3d) OffsetMul(x, y, z float32) Vec3d {
	return Vec3d{X: v.X * x, Y: v.Y * y, Z: v.Z * z}
}

func (v Vec3d) Floor() Vec3d {
	return Vec3d{X: float32(math.Floor(float64(v.X))), Y: float32(math.Floor(float64(v.Y))), Z: float32(math.Floor(float64(v.Z)))}
}

func (v Vec3d) Length() float32 {
	return float32(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3d) Normalize() Vec3d {
	length := v.Length()
	return Vec3d{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}

func (v Vec3d) ToChunkPos() [2]int32 {
	return [2]int32{int32(v.X) >> 4, int32(v.Z) >> 4}
}
