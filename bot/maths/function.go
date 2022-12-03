package maths

import (
	"math"
)

/* Minecraft related functions */

func CalculateFallingTime(dst float64) float64 {
	return -2.383*
		0.00000000000001*math.Pow(dst, 6) +
		3.107*
			0.00000000001*math.Pow(dst, 5) -
		1.587*
			0.00000001*math.Pow(dst, 4) +
		4.035*
			0.000001*math.Pow(dst, 3) -
		0.000546*math.Pow(dst, 2) +
		0.0546*dst +
		0.2993 + 0.1
	// This is probably not the most accurate formula for calculating falling time
}

func CalculateFallVelocity(ticks float32) Vec3d {
	return Vec3d{X: 0, Y: -0.08 * ticks, Z: 0}
}
