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

func GetVectorFromRotation(rotation Vec2d[float64]) Vec3d[float64] {
	f := math.Cos(ToRadians(-rotation.Y)*0.017453292 - math.Pi)
	f1 := math.Sin(ToRadians(-rotation.Y)*0.017453292 - math.Pi)
	f2 := -math.Cos(ToRadians(-rotation.X) * 0.017453292)
	f3 := math.Sin(ToRadians(-rotation.X) * 0.017453292)
	return Vec3d[float64]{X: f1 * f2, Y: f3, Z: f * f2}
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
