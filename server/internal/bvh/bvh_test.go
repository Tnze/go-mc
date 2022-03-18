package bvh

//
//import (
//	"math/rand"
//	"testing"
//)
//
//func TestTree2_Insert(t *testing.T) {
//	aabbs := []AABB[Vec2[float64]]{
//		{Upper: Vec2[float64]{1, 1}, Lower: Vec2[float64]{0, 0}},
//		{Upper: Vec2[float64]{2, 1}, Lower: Vec2[float64]{1, 0}},
//		{Upper: Vec2[float64]{11, 1}, Lower: Vec2[float64]{10, 0}},
//		{Upper: Vec2[float64]{12, 1}, Lower: Vec2[float64]{11, 0}},
//		{Upper: Vec2[float64]{101, 1}, Lower: Vec2[float64]{100, 0}},
//		{Upper: Vec2[float64]{102, 1}, Lower: Vec2[float64]{101, 0}},
//		{Upper: Vec2[float64]{111, 1}, Lower: Vec2[float64]{110, 0}},
//		{Upper: Vec2[float64]{112, 1}, Lower: Vec2[float64]{111, 0}},
//		{Upper: Vec2[float64]{1, 1}, Lower: Vec2[float64]{-1, -1}},
//	}
//	var bvh Tree[float64, AABB[Vec2[float64]], int]
//	for i, aabb := range aabbs {
//		bvh.Insert(aabb, i)
//		// visualize
//		t.Log(bvh)
//	}
//	//bvh.FindVec(Vec2{0.5, 0.5}, func(v interface{}) {
//	//	t.Logf("find! %v", v)
//	//})
//}
//
//func TestTree2_FindVec(t *testing.T) {
//	aabbs := []AABB[Vec2[float64]]{
//		{Upper: Vec2[float64]{2, 2}, Lower: Vec2[float64]{-1, -1}},
//		{Upper: Vec2[float64]{2, 1}, Lower: Vec2[float64]{-1, -2}},
//		{Upper: Vec2[float64]{1, 1}, Lower: Vec2[float64]{-2, -2}},
//		{Upper: Vec2[float64]{1, 2}, Lower: Vec2[float64]{-2, -1}},
//	}
//	var bvh Tree[float64, AABB[Vec2[float64]], int]
//	for i, aabb := range aabbs {
//		bvh.Insert(aabb, i)
//		// visualize
//		t.Log(bvh)
//	}
//	//findVec := func(vec Vec2) (list []interface{}) {
//	//	bvh.FindVec(vec, func(v interface{}) { list = append(list, v) })
//	//	return
//	//}
//	//t.Log(findVec(Vec2{0, 0}))
//	//t.Log(findVec(Vec2{1.5, 0}))
//	//t.Log(findVec(Vec2{1.5, 1.5}))
//	//t.Log(findVec(Vec2{-1.5, 0}))
//	//
//	//findAABB := func(aabb AABB2) (list []interface{}) {
//	//	bvh.FindAABB(aabb, func(v interface{}) { list = append(list, v) })
//	//	return
//	//}
//	//t.Log(findAABB(AABB2{Upper: Vec2{1, 1}, Lower: Vec2{-1, -1}}))
//	//t.Log(findAABB(AABB2{Upper: Vec2{3, 3}, Lower: Vec2{1.5, 1.5}}))
//	//t.Log(findAABB(AABB2{Upper: Vec2{-1.5, 0.5}, Lower: Vec2{-2.5, -0.5}}))
//}
//
//func TestTree2_Insert_rotation(t *testing.T) {
//	var bvh Tree[float64, AABB[Vec2[float64]], int]
//	for i := 0; i < 5; i++ {
//		bvh.Insert(AABB[Vec2[float64]]{
//			Upper: Vec2[float64]{float64(i), float64(i)},
//			Lower: Vec2[float64]{float64(0), float64(0)},
//		}, i)
//	}
//}
//
//func BenchmarkTree2_Insert_random(b *testing.B) {
//	const size = 25
//	// generate test cases
//	aabbs := make([]AABB[Vec2[float64]], b.N)
//	poses := make([]Vec2[float64], b.N)
//	for i := range aabbs {
//		poses[i] = Vec2[float64]{rand.Float64() * 1e4, rand.Float64() * 1e4}
//		aabbs[i] = AABB[Vec2[float64]]{
//			Upper: Vec2[float64]{poses[i][0] + size, poses[i][0] + size},
//			Lower: Vec2[float64]{poses[i][0] - size, poses[i][0] - size},
//		}
//	}
//	b.ResetTimer()
//
//	var bvh Tree[float64, AABB[Vec2[float64]], any]
//	for _, v := range aabbs {
//		bvh.Insert(v, nil)
//	}
//	//for _, v := range poses {
//	//	bvh.FindVec(v, func(interface{}) {})
//	//}
//}
//
//func BenchmarkTree2_Insert_sorted1(b *testing.B) {
//	// generate test cases
//	var bvh Tree[float64, AABB[Vec2[float64]], int]
//	upper := Vec2[float64]{float64(b.N), float64(b.N)}
//	for i := 0; i < b.N; i++ {
//		bvh.Insert(AABB[Vec2[float64]]{
//			Upper: upper,
//			Lower: Vec2[float64]{float64(i), float64(i)},
//		}, i)
//	}
//}
//
//func BenchmarkTree2_Insert_sorted2(b *testing.B) {
//	// generate test cases
//	var bvh Tree[float64, AABB[Vec2[float64]], int]
//	for i := 0; i < b.N; i++ {
//		bvh.Insert(AABB[Vec2[float64]]{
//			Upper: Vec2[float64]{float64(i), float64(i)},
//			Lower: Vec2[float64]{0, 0},
//		}, i)
//	}
//}
//
//func BenchmarkTree2_Delete_random(b *testing.B) {
//	const size = 25
//	// generate test cases
//	aabbs := make([]AABB[Vec2[float64]], b.N)
//	poses := make([]Vec2[float64], b.N)
//	nodes := make([]*Node[AABB[Vec2[float64]], any], b.N)
//	for i := range aabbs {
//		poses[i] = Vec2[float64]{rand.Float64() * 1e4, rand.Float64() * 1e4}
//		aabbs[i] = AABB[Vec2[float64]]{
//			Upper: Vec2[float64]{poses[i][0] + size, poses[i][0] + size},
//			Lower: Vec2[float64]{poses[i][0] - size, poses[i][0] - size},
//		}
//	}
//	b.ResetTimer()
//
//	var bvh Tree[float64, AABB[Vec2[float64]], any]
//	for i, v := range aabbs {
//		nodes[i] = bvh.Insert(v, nil)
//	}
//
//	b.StopTimer()
//	rand.Shuffle(b.N, func(i, j int) {
//		nodes[i], nodes[j] = nodes[j], nodes[i]
//	})
//	b.StartTimer()
//
//	for _, v := range nodes {
//		bvh.Delete(v)
//	}
//}
