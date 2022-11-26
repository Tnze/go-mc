package maths

import "math"

var NullVec2d = Vec2d{X: 0, Y: 0}

type Vec2d struct {
	X, Y float32
}

func (v Vec2d) Add(v2 Vec2d) Vec2d {
	return Vec2d{X: v.X + v2.X, Y: v.Y + v2.Y}
}

func (v Vec2d) AddScalar(s float32) Vec2d {
	return Vec2d{X: v.X + s, Y: v.Y + s}
}

func (v Vec2d) Sub(v2 Vec2d) Vec2d {
	return Vec2d{X: v.X - v2.X, Y: v.Y - v2.Y}
}

func (v Vec2d) SubScalar(s float32) Vec2d {
	return Vec2d{X: v.X - s, Y: v.Y - s}
}

func (v Vec2d) Mul(v2 Vec2d) Vec2d {
	return Vec2d{X: v.X * v2.X, Y: v.Y * v2.Y}
}

func (v Vec2d) MulScalar(s float32) Vec2d {
	return Vec2d{X: v.X * s, Y: v.Y * s}
}

func (v Vec2d) Div(v2 Vec2d) Vec2d {
	return Vec2d{X: v.X / v2.X, Y: v.Y / v2.Y}
}

func (v Vec2d) DivScalar(s float32) Vec2d {
	return Vec2d{X: v.X / s, Y: v.Y / s}
}

func (v Vec2d) DistanceTo(v2 Vec2d) float32 {
	return float32(math.Sqrt(math.Pow(float64(v.X-v2.X), 2) + math.Pow(float64(v.Y-v2.Y), 2)))
}
