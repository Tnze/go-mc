package bvh

import (
	"math"

	"golang.org/x/exp/constraints"
)

type AABB[I constraints.Signed | constraints.Float, V interface {
	Add(V) V
	Sub(V) V
	Max(V) V
	Min(V) V
	Less(V) bool
	More(V) bool
	Sum() I
}] struct{ Upper, Lower V }

func (aabb AABB[I, V]) WithIn(point V) bool {
	return aabb.Lower.Less(point) && aabb.Upper.More(point)
}

func (aabb AABB[I, V]) Touch(other AABB[I, V]) bool {
	return aabb.Lower.Less(other.Upper) && other.Lower.Less(aabb.Upper) &&
		aabb.Upper.More(other.Lower) && other.Upper.More(aabb.Lower)
}

func (aabb AABB[I, V]) Union(other AABB[I, V]) AABB[I, V] {
	return AABB[I, V]{Upper: aabb.Upper.Max(other.Upper), Lower: aabb.Lower.Min(other.Lower)}
}
func (aabb AABB[I, V]) Surface() I { return aabb.Upper.Sub(aabb.Lower).Sum() * 2 }

type Sphere[I constraints.Float, V interface {
	Add(V) V
	Sub(V) V
	Mul(I) V
	Max(V) V
	Min(V) V
	Less(V) bool
	More(V) bool
	Norm() I
	Sum() I
}] struct {
	Center V
	R      I
}

func (s Sphere[I, V]) WithIn(point V) bool {
	return s.Center.Sub(point).Norm() < s.R
}

func (s Sphere[I, V]) Touch(other Sphere[I, V]) bool {
	return s.Center.Sub(other.Center).Norm() < s.R+other.R
}

func (s Sphere[I, V]) Union(other Sphere[I, V]) Sphere[I, V] {
	d := other.Center.Sub(s.Center).Norm()
	r1r2d := (s.R - other.R) / d
	return Sphere[I, V]{
		Center: s.Center.Mul(1 + r1r2d).Add(other.Center.Mul(1 - r1r2d)),
		R:      d + s.R + other.R,
	}
}
func (s Sphere[I, V]) Surface() I { return 2 * math.Pi * s.R }
