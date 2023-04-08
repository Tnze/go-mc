package maths

import (
	"github.com/Tnze/go-mc/data/enums"
	"github.com/Tnze/go-mc/level/block"
)

type RayTraceResult struct {
	// The position of the ray trace
	Position Vec3d[float64]
	// The side of the block that was hit
	Side enums.EnumFacing
	// The block that was hit
	Block block.Block
}

func RayTraceBlocks(start, end Vec3d[float64]) []Vec3d[float64] {
	var result []Vec3d[float64]
	diff := end.Sub(start)
	distance := diff.Length()
	if distance == 0 {
		return result
	}
	for i := 0; i < int(distance); i++ {
		result = append(result, start.Add(diff.MulScalar(float64(i)/distance)))
	}
	return result
}
