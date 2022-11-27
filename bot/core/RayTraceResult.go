package core

import (
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/level/block"
)

type RayTraceResult struct {
	// The position of the ray trace
	Position maths.Vec3d
	// The side of the block that was hit
	Side EnumFacing
	// The block that was hit
	Block block.Block
	// The entity that was hit
	Entity Entity // May be empty
}
