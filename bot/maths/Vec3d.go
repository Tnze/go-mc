package maths

import "math"

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
