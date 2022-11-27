package core

import (
	"github.com/Tnze/go-mc/bot/maths"
)

type EnumFacing int8

const (
	DOWN  EnumFacing = 0
	UP    EnumFacing = 1
	NORTH EnumFacing = 2
	SOUTH EnumFacing = 3
	WEST  EnumFacing = 4
	EAST  EnumFacing = 5
)

var EnumFacingValues = []EnumFacing{DOWN, UP, NORTH, SOUTH, WEST, EAST}

func (f EnumFacing) Vector() (v maths.Vec3d) {
	switch f {
	case DOWN:
		v = maths.Vec3d{Y: -1}
	case UP:
		v = maths.Vec3d{Y: 1}
	case NORTH:
		v = maths.Vec3d{Z: -1}
	case SOUTH:
		v = maths.Vec3d{Z: 1}
	case WEST:
		v = maths.Vec3d{X: -1}
	case EAST:
		v = maths.Vec3d{X: 1}
	}
	return
}

func GetClosestFacing(eyePos, blockPos maths.Vec3d) EnumFacing {
	var closest EnumFacing
	var minDiff float32
	for _, side := range GetVisibleSides(eyePos, blockPos) {
		diff := eyePos.DistanceTo(blockPos.Add(side.Vector().Center()))
		if minDiff == 0 || diff < minDiff {
			minDiff = diff
			closest = side
		}
	}
	return closest
}

func GetVisibleSides(eyePos, blockPos maths.Vec3d) []EnumFacing {
	var sides []EnumFacing
	blockCenter := blockPos.Center()
	axis := checkAxis(eyePos.X-blockCenter.X, WEST)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Y-blockCenter.Y, DOWN)
	if axis != -1 {
		sides = append(sides, axis)
	}
	axis = checkAxis(eyePos.Z-blockCenter.Z, NORTH)
	if axis != -1 {
		sides = append(sides, axis)
	}
	return sides
}

func (f EnumFacing) GetOpposite() EnumFacing {
	return EnumFacingValues[(f+3)%6]
}

func checkAxis(diff float32, negativeSide EnumFacing) EnumFacing {
	if diff < -0.5 {
		return negativeSide
	} else if diff > 0.5 {
		return negativeSide.GetOpposite()
	} else {
		return -1
	}
}
