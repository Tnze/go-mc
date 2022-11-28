package maths

import (
	"math"
)

func ProjectPosition(rotation Vec2d, distance float32, head bool) Vec3d {
	x := distance * float32(math.Sin(ToRadians(float64(rotation.X)))*math.Cos(ToRadians(float64(rotation.Y))))
	y := float32(0)
	if rotation.Y > 0 {
		y = distance * float32(-math.Sin(ToRadians(float64(rotation.Y))))
	} else {
		y = distance * float32(math.Sin(ToRadians(float64(rotation.Y))))
	}
	z := distance * float32(math.Cos(ToRadians(float64(rotation.X)))*math.Cos(ToRadians(float64(rotation.X))))
	if head {
		y += 1.62
	}
	return Vec3d{X: x, Y: y, Z: z}
}

func GetRotationFromVector(vec Vec3d) Vec2d {
	xz := math.Hypot(float64(vec.X), float64(vec.Z))
	y := normalizeAngle(ToDegrees(math.Atan2(float64(vec.Z), float64(vec.X))) - 90)
	x := normalizeAngle(ToDegrees(-math.Atan2(float64(vec.Y), xz)))
	return Vec2d{X: float32(x), Y: float32(y)}
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
