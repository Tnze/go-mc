package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

var NullVec2d = Vec2d[float64]{X: 0, Y: 0}

type Vec2d[T constraints.Float] struct {
	X, Y T
	// Pitch, Yaw
}

func (v Vec2d[T]) Add(v2 Vec2d[T]) Vec2d[T] {
	return Vec2d[T]{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2d[T]) AddScalar(s T) Vec2d[T] {
	return Vec2d[T]{X: v.X + s, Y: v.Y + s}
}

func (v Vec2d[T]) Sub(v2 Vec2d[T]) Vec2d[T] {
	return Vec2d[T]{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2d[T]) SubScalar(s T) Vec2d[T] {
	return Vec2d[T]{X: v.X - s, Y: v.Y - s}
}

func (v Vec2d[T]) Mul(v2 Vec2d[T]) Vec2d[T] {
	return Vec2d[T]{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vec2d[T]) MulScalar(s T) Vec2d[T] {
	return Vec2d[T]{X: v.X * s, Y: v.Y * s}
}

func (v Vec2d[T]) Div(v2 Vec2d[T]) Vec2d[T] {
	return Vec2d[T]{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v Vec2d[T]) DivScalar(s T) Vec2d[T] {
	return Vec2d[T]{X: v.X / s, Y: v.Y / s}
}

func (v Vec2d[T]) DistanceTo(v2 Vec2d[T]) T {
	xDiff, yDiff := v.X-v2.X, v.Y-v2.Y
	return T(math.Sqrt(float64(xDiff*xDiff + yDiff*yDiff)))
}
