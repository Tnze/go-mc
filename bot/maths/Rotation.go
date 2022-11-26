package maths

import "math"

func GetRotationFromVector(vec Vec3d) Vec2d {
	xz := math.Hypot(float64(vec.X), float64(vec.Z))
	y := normalizeAngle(ToDegrees(math.Atan2(float64(vec.Z), float64(vec.X))) - 90)
	x := normalizeAngle(ToDegrees(-math.Atan2(float64(vec.Y), xz)))
	return Vec2d{X: float32(x), Y: float32(y)}
}

func ToDegrees(angle float64) float64 {
	return angle * 180.0 / math.Pi
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
