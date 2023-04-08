package maths

import (
	"golang.org/x/exp/constraints"
	"math"
)

type AxisAlignedBB[T constraints.Float] struct {
	MinX, MinY, MinZ,
	MaxX, MaxY, MaxZ T
}

func (a AxisAlignedBB[T]) Contract(x, y, z T) AxisAlignedBB[T] {
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
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Expand(x, y, z T) AxisAlignedBB[T] {
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
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Grow(x, y, z T) AxisAlignedBB[T] {
	d0 := a.MinX - x
	d1 := a.MinY - y
	d2 := a.MinZ - z
	d3 := a.MaxX + x
	d4 := a.MaxY + y
	d5 := a.MaxZ + z
	return AxisAlignedBB[T]{MinX: d0, MinY: d1, MinZ: d2, MaxX: d3, MaxY: d4, MaxZ: d5}
}

func (a AxisAlignedBB[T]) Offset(x, y, z T) AxisAlignedBB[T] {
	return AxisAlignedBB[T]{MinX: a.MinX + x, MinY: a.MinY + y, MinZ: a.MinZ + z, MaxX: a.MaxX + x, MaxY: a.MaxY + y, MaxZ: a.MaxZ + z}
}

func (a AxisAlignedBB[T]) Intersect(other AxisAlignedBB[T]) AxisAlignedBB[T] {
	d0 := float32(math.Max(float64(a.MinX), float64(other.MinX)))
	d1 := float32(math.Min(float64(a.MinY), float64(other.MinY)))
	d2 := float32(math.Min(float64(a.MinZ), float64(other.MinZ)))
	d3 := float32(math.Min(float64(a.MaxX), float64(other.MaxX)))
	d4 := float32(math.Min(float64(a.MaxY), float64(other.MaxY)))
	d5 := float32(math.Min(float64(a.MaxZ), float64(other.MaxZ)))
	return AxisAlignedBB[T]{MinX: T(d0), MinY: T(d1), MinZ: T(d2), MaxX: T(d3), MaxY: T(d4), MaxZ: T(d5)}
}

func (a AxisAlignedBB[T]) IntersectsWith(other AxisAlignedBB[T]) bool {
	return a.MinX < other.MaxX && a.MaxX > other.MinX && a.MinY < other.MaxY && a.MaxY > other.MinY && a.MinZ < other.MaxZ && a.MaxZ > other.MinZ
}

func (a AxisAlignedBB[T]) CollideX(other AxisAlignedBB[T], x T) T {
	if other.MinY >= a.MaxY && other.MaxY <= a.MinY && other.MinZ >= a.MaxZ && other.MaxZ <= a.MinZ {
		if x > 0.0 && other.MaxX <= a.MinX {
			d0 := a.MinX - other.MaxX
			if d0 < x {
				x = d0
			}
		} else if x < 0.0 && other.MinX >= a.MaxX {
			d0 := a.MaxX - other.MinX
			if d0 > x {
				x = d0
			}
		}
	}
	return x
}

func (a AxisAlignedBB[T]) CollideY(other AxisAlignedBB[T], y T) T {
	if other.MinX >= a.MaxX && other.MaxX <= a.MinX && other.MinZ >= a.MaxZ && other.MaxZ <= a.MinZ {
		if y > 0.0 && other.MaxY <= a.MinY {
			d0 := a.MinY - other.MaxY
			if d0 < y {
				y = d0
			}
		} else if y < 0.0 && other.MinY >= a.MaxY {
			d0 := a.MaxY - other.MinY
			if d0 > y {
				y = d0
			}
		}
	}
	return y
}

func (a AxisAlignedBB[T]) CollideZ(other AxisAlignedBB[T], z T) T {
	if other.MinX >= a.MaxX && other.MaxX <= a.MinX && other.MinY >= a.MaxY && other.MaxY <= a.MinY {
		if z > 0.0 && other.MaxZ <= a.MinZ {
			d0 := a.MinZ - other.MaxZ
			if d0 < z {
				z = d0
			}
		} else if z < 0.0 && other.MinZ >= a.MaxZ {
			d0 := a.MaxZ - other.MinZ
			if d0 > z {
				z = d0
			}
		}
	}
	return z
}
