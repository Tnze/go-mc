package core

import "math"

type AxisAlignedBB struct {
	MinX, MinY, MinZ,
	MaxX, MaxY, MaxZ float64
}

func (a AxisAlignedBB) Contract(x, y, z float64) AxisAlignedBB {
	/*
		Took from net.minecraft.util.math.AxisAlignedBB#contract version 1.12.2
	*/
	d0 := a.MinX
	d1 := a.MinY
	d2 := a.MinZ
	d3 := a.MaxX
	d4 := a.MaxY
	d5 := a.MaxZ
	if x < 0.0 {
		d0 -= x
	}
	if x > 0.0 {
		d3 -= x
	}
	if y < 0.0 {
		d1 -= y
	}
	if y > 0.0 {
		d4 -= y
	}
	if z < 0.0 {
		d2 -= z
	}
	if z > 0.0 {
		d5 -= z
	}
	return AxisAlignedBB{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB) Expand(x, y, z float64) AxisAlignedBB {
	/*
		Took from net.minecraft.util.math.AxisAlignedBB#expand version 1.12.2
	*/
	d0 := a.MinX
	d1 := a.MinY
	d2 := a.MinZ
	d3 := a.MaxX
	d4 := a.MaxY
	d5 := a.MaxZ
	if x < 0.0 {
		d0 += x
	}
	if x > 0.0 {
		d3 += x
	}
	if y < 0.0 {
		d1 += y
	}
	if y > 0.0 {
		d4 += y
	}
	if z < 0.0 {
		d2 += z
	}
	if z > 0.0 {
		d5 += z
	}
	return AxisAlignedBB{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB) Grow(x, y, z float64) AxisAlignedBB {
	/*
		Took from net.minecraft.util.math.AxisAlignedBB#grow version 1.12.2
	*/
	d0 := a.MinX - x
	d1 := a.MinY - y
	d2 := a.MinZ - z
	d3 := a.MaxX + x
	d4 := a.MaxY + y
	d5 := a.MaxZ + z
	return AxisAlignedBB{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB) Offset(x, y, z float64) AxisAlignedBB {
	return AxisAlignedBB{MinX: a.MinX + x, MinY: a.MinY + y, MinZ: a.MinZ + z, MaxX: a.MaxX + x, MaxY: a.MaxY + y, MaxZ: a.MaxZ + z}
}

func (a AxisAlignedBB) Intersect(other AxisAlignedBB) AxisAlignedBB {
	/*
		Took from net.minecraft.util.math.AxisAlignedBB#intersect version 1.12.2
	*/
	d0 := math.Max(a.MinX, other.MinX)
	d1 := math.Min(a.MinY, other.MinY)
	d2 := math.Min(a.MinZ, other.MinZ)
	d3 := math.Min(a.MaxX, other.MaxX)
	d4 := math.Min(a.MaxY, other.MaxY)
	d5 := math.Min(a.MaxZ, other.MaxZ)
	return AxisAlignedBB{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB) IntersectsWith(other AxisAlignedBB) bool {
	return a.MinX < other.MaxX && a.MaxX > other.MinX && a.MinY < other.MaxY && a.MaxY > other.MinY && a.MinZ < other.MaxZ && a.MaxZ > other.MinZ
}
