package bvh

import "testing"

func TestAABB_WithIn(t *testing.T) {
	aabb := AABB[float64, Vec2[float64]]{
		Upper: Vec2[float64]{2, 2},
		Lower: Vec2[float64]{-1, -1},
	}
	if !aabb.WithIn(Vec2[float64]{0, 0}) {
		panic("(0, 0) should included")
	}
	if aabb.WithIn(Vec2[float64]{-2, -2}) {
		panic("(-2, -2) shouldn't included")
	}

	aabb2 := AABB[int, Vec3[int]]{
		Upper: Vec3[int]{1, 1, 1},
		Lower: Vec3[int]{-1, -1, -1},
	}
	if !aabb2.WithIn(Vec3[int]{0, 0, 0}) {
		panic("(0, 0, 0) should included")
	}
	if aabb2.WithIn(Vec3[int]{-2, -2, 0}) {
		panic("(-2, -2, 0) shouldn't included")
	}

	sphere := Sphere[float64, Vec2[float64]]{
		Center: Vec2[float64]{0, 0},
		R:      1.0,
	}
	if !sphere.WithIn(Vec2[float64]{0, 0}) {
		t.Errorf("(0,0) is in")
	}
	if sphere.WithIn(Vec2[float64]{1, 1}) {
		t.Errorf("(1,1) isn't in")
	}
}
