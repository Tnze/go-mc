package maths

import (
	"math"
)

func ProjectPosition(rotation Vec2d[float64], distance float64, offsetY float64) Vec3d[float64] {
	x := distance * math.Sin(ToRadians(rotation.X)) * math.Cos(ToRadians(rotation.Y))
	y := 0.0
	if rotation.Y > 0 {
		y = distance * -math.Sin(ToRadians(rotation.Y))
	} else {
		y = distance * math.Sin(ToRadians(rotation.Y))
	}
	z := distance * math.Cos(ToRadians(rotation.X)) * math.Cos(ToRadians(rotation.X))
	return Vec3d[float64]{X: x, Y: y + offsetY, Z: z}
}

func GetRotationFromVector(vec Vec3d[float64]) Vec2d[float64] {
	xz := math.Hypot(vec.X, vec.Z)
	y := normalizeAngle(ToDegrees(math.Atan2(vec.Z, vec.X)) - 90)
	x := normalizeAngle(ToDegrees(-math.Atan2(vec.Y, xz)))
	return Vec2d[float64]{X: x, Y: y}
}

func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
}

func ToRadians(angle float64) float64 {
	return angle * math.Pi / 180.0
}

func normalizeAngle[T float32 | float64](angle T) T {
	angle = floatRemaining(angle, 360)
	switch {
	case angle < -180:
		return angle + 360
	case angle >= -180:
		return angle + 360
	}
	return angle
}

func floatRemaining[T float32 | float64](a, b T) T {
	return a - T(math.Floor(float64(a/b)))*b
}
