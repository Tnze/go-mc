package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

var NullVec3d = Vec3d[float64]{X: 0, Y: 0, Z: 0}

type Vec3d[T constraints.Float] struct {
	X, Y, Z T
}

func (v Vec3d[T]) Add(vec3d Vec3d[T]) Vec3d[T] {
	return Vec3d[T]{X: v.X + vec3d.X, Y: v.Y + vec3d.Y, Z: v.Z + vec3d.Z}
}

func (v Vec3d[T]) AddScalar(scalar T) Vec3d[T] {
	return Vec3d[T]{X: v.X + scalar, Y: v.Y + scalar, Z: v.Z + scalar}
}

func (v Vec3d[T]) Sub(vec3d Vec3d[T]) Vec3d[T] {
	return Vec3d[T]{X: v.X - vec3d.X, Y: v.Y - vec3d.Y, Z: v.Z - vec3d.Z}
}

func (v Vec3d[T]) SubScalar(scalar T) Vec3d[T] {
	return Vec3d[T]{X: v.X - scalar, Y: v.Y - scalar, Z: v.Z - scalar}
}

func (v Vec3d[T]) Mul(vec3d Vec3d[T]) Vec3d[T] {
	return Vec3d[T]{X: v.X * vec3d.X, Y: v.Y * vec3d.Y, Z: v.Z * vec3d.Z}
}

func (v Vec3d[T]) MulScalar(scalar T) Vec3d[T] {
	return Vec3d[T]{X: v.X * scalar, Y: v.Y * scalar, Z: v.Z * scalar}
}

func (v Vec3d[T]) Div(vec3d Vec3d[T]) Vec3d[T] {
	return Vec3d[T]{X: v.X / vec3d.X, Y: v.Y / vec3d.Y, Z: v.Z / vec3d.Z}
}

func (v Vec3d[T]) DivScalar(scalar T) Vec3d[T] {
	return Vec3d[T]{X: v.X / scalar, Y: v.Y / scalar, Z: v.Z / scalar}
}

func (v Vec3d[T]) DistanceTo(vec3d Vec3d[T]) T {
	xDiff, yDiff, zDiff := v.X-vec3d.X, v.Y-vec3d.Y, v.Z-vec3d.Z
	return T(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff + zDiff*zDiff)))
}

func (v Vec3d[T]) Center() Vec3d[T] {
	return Vec3d[T]{X: v.X + 0.5, Y: v.Y + 0.5, Z: v.Z + 0.5}
}

func (v Vec3d[T]) Offset(x, y, z T) Vec3d[T] {
	return Vec3d[T]{X: v.X + x, Y: v.Y + y, Z: v.Z + z}
}

func (v Vec3d[T]) OffsetMul(x, y, z T) Vec3d[T] {
	return Vec3d[T]{X: v.X * x, Y: v.Y * y, Z: v.Z * z}
}

func (v Vec3d[T]) Floor() Vec3d[T] {
	return Vec3d[T]{X: T(math.Floor(float64(v.X))), Y: T(math.Floor(float64(v.Y))), Z: T(math.Floor(float64(v.Z)))}
}

func (v Vec3d[T]) Length() T {
	return T(math.Sqrt(float64(v.X*v.X + v.Y*v.Y + v.Z*v.Z)))
}

func (v Vec3d[T]) Normalize() Vec3d[T] {
	length := v.Length()
	return Vec3d[T]{X: v.X / length, Y: v.Y / length, Z: v.Z / length}
}

func (v Vec3d[T]) Spread() (T, T, T) {
	return v.X, v.Y, v.Z
}

func (v Vec3d[T]) ToChunkPos() [2]int32 {
	return [2]int32{int32(v.X) >> 4, int32(v.Z) >> 4}
}
