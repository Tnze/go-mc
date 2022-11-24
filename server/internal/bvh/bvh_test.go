package bvh

import (
	"math/rand"
	"testing"
)

func TestTree2_Insert(t *testing.T) {
	aabbs := []AABB[float64, Vec2[float64]]{
		{Upper: Vec2[float64]{1, 1}, Lower: Vec2[float64]{0, 0}},
		{Upper: Vec2[float64]{2, 1}, Lower: Vec2[float64]{1, 0}},
		{Upper: Vec2[float64]{11, 1}, Lower: Vec2[float64]{10, 0}},
		{Upper: Vec2[float64]{12, 1}, Lower: Vec2[float64]{11, 0}},
		{Upper: Vec2[float64]{101, 1}, Lower: Vec2[float64]{100, 0}},
		{Upper: Vec2[float64]{102, 1}, Lower: Vec2[float64]{101, 0}},
		{Upper: Vec2[float64]{111, 1}, Lower: Vec2[float64]{110, 0}},
		{Upper: Vec2[float64]{112, 1}, Lower: Vec2[float64]{111, 0}},
		{Upper: Vec2[float64]{1, 1}, Lower: Vec2[float64]{-1, -1}},
	}
	var bvh Tree[float64, AABB[float64, Vec2[float64]], int]
	for i, aabb := range aabbs {
		bvh.Insert(aabb, i)
		// visualize
		t.Log(bvh)
	}
	bvh.Find(TouchPoint[Vec2[float64], AABB[float64, Vec2[float64]]](Vec2[float64]{0.5, 0.5}), func(n *Node[float64, AABB[float64, Vec2[float64]], int]) bool {
		t.Logf("find! %v", n.Value)
		return true
	})
}

func TestTree2_Find_vec(t *testing.T) {
	type Vec2d = Vec2[float64]
	type AABBVec2d = AABB[float64, Vec2d]
	type TreeAABBVec2di = Tree[float64, AABBVec2d, int]

	aabbs := []AABBVec2d{
		{Upper: Vec2d{2, 2}, Lower: Vec2d{-1, -1}},
		{Upper: Vec2d{2, 1}, Lower: Vec2d{-1, -2}},
		{Upper: Vec2d{1, 1}, Lower: Vec2d{-2, -2}},
		{Upper: Vec2d{1, 2}, Lower: Vec2d{-2, -1}},
	}
	var bvh TreeAABBVec2di
	for i, aabb := range aabbs {
		bvh.Insert(aabb, i)
		t.Log(bvh)
	}
	find := func(test func(bound AABBVec2d) bool) []int {
		var result []int
		bvh.Find(test, func(n *Node[float64, AABBVec2d, int]) bool {
			result = append(result, n.Value)
			return true
		})
		return result
	}
	t.Log(find(TouchPoint[Vec2d, AABBVec2d](Vec2d{0, 0})))
	t.Log(find(TouchPoint[Vec2d, AABBVec2d](Vec2d{1.5, 0})))
	t.Log(find(TouchPoint[Vec2d, AABBVec2d](Vec2d{1.5, 1.5})))
	t.Log(find(TouchPoint[Vec2d, AABBVec2d](Vec2d{-1.5, 0})))

	t.Log(find(TouchBound(AABBVec2d{Upper: Vec2d{1, 1}, Lower: Vec2d{-1, -1}})))
	t.Log(find(TouchBound(AABBVec2d{Upper: Vec2d{1, 1}, Lower: Vec2d{1.5, 1.5}})))
	t.Log(find(TouchBound(AABBVec2d{Upper: Vec2d{-1.5, 0.5}, Lower: Vec2d{-2.5, -0.5}})))
}

func BenchmarkTree_Insert(b *testing.B) {
	type Vec2d = Vec2[float64]
	type AABBVec2d = AABB[float64, Vec2d]
	type TreeAABBVec2da = Tree[float64, AABBVec2d, any]

	const size = 25
	// generate test cases
	aabbs := make([]AABBVec2d, b.N)
	poses := make([]Vec2d, b.N)
	for i := range aabbs {
		poses[i] = Vec2d{rand.Float64() * 1e4, rand.Float64() * 1e4}
		aabbs[i] = AABBVec2d{
			Upper: Vec2d{poses[i][0] + size, poses[i][0] + size},
			Lower: Vec2d{poses[i][0] - size, poses[i][0] - size},
		}
	}
	b.ResetTimer()

	var bvh TreeAABBVec2da
	for _, v := range aabbs {
		bvh.Insert(v, nil)
	}
}

func BenchmarkTree2_Find_random(b *testing.B) {
	type Vec2d = Vec2[float64]
	type AABBVec2d = AABB[float64, Vec2d]
	type TreeAABBVec2da = Tree[float64, AABBVec2d, any]

	const size = 25
	// generate test cases
	aabbs := make([]AABBVec2d, b.N)
	poses := make([]Vec2d, b.N)
	for i := range aabbs {
		poses[i] = Vec2d{rand.Float64() * 1e4, rand.Float64() * 1e4}
		aabbs[i] = AABBVec2d{
			Upper: Vec2d{poses[i][0] + size, poses[i][0] + size},
			Lower: Vec2d{poses[i][0] - size, poses[i][0] - size},
		}
	}
	var bvh TreeAABBVec2da
	for _, v := range aabbs {
		bvh.Insert(v, nil)
	}
	b.ResetTimer()

	for _, v := range poses {
		bvh.Find(TouchPoint[Vec2d, AABBVec2d](v), func(n *Node[float64, AABBVec2d, any]) bool { return true })
	}
}

func BenchmarkTree2_Delete_random(b *testing.B) {
	const size = 25
	// generate test cases
	aabbs := make([]AABB[float64, Vec2[float64]], b.N)
	poses := make([]Vec2[float64], b.N)
	nodes := make([]*Node[float64, AABB[float64, Vec2[float64]], any], b.N)
	for i := range aabbs {
		poses[i] = Vec2[float64]{rand.Float64() * 1e4, rand.Float64() * 1e4}
		aabbs[i] = AABB[float64, Vec2[float64]]{
			Upper: Vec2[float64]{poses[i][0] + size, poses[i][0] + size},
			Lower: Vec2[float64]{poses[i][0] - size, poses[i][0] - size},
		}
	}
	b.ResetTimer()

	var bvh Tree[float64, AABB[float64, Vec2[float64]], any]
	for i, v := range aabbs {
		nodes[i] = bvh.Insert(v, nil)
	}

	b.StopTimer()
	rand.Shuffle(b.N, func(i, j int) {
		nodes[i], nodes[j] = nodes[j], nodes[i]
	})
	b.StartTimer()

	for _, v := range nodes {
		bvh.Delete(v)
	}
}
