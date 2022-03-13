package bvh

import (
	"fmt"
	"math/rand"
	"strings"
	"testing"
)

func TestTree2_Insert(t *testing.T) {
	aabbs := []AABB2{
		{Upper: Vec2{1, 1}, Lower: Vec2{0, 0}},
		{Upper: Vec2{2, 1}, Lower: Vec2{1, 0}},
		{Upper: Vec2{11, 1}, Lower: Vec2{10, 0}},
		{Upper: Vec2{12, 1}, Lower: Vec2{11, 0}},
		{Upper: Vec2{101, 1}, Lower: Vec2{100, 0}},
		{Upper: Vec2{102, 1}, Lower: Vec2{101, 0}},
		{Upper: Vec2{111, 1}, Lower: Vec2{110, 0}},
		{Upper: Vec2{112, 1}, Lower: Vec2{111, 0}},
		{Upper: Vec2{1, 1}, Lower: Vec2{-1, -1}},
	}
	var bvh Tree2
	for _, aabb := range aabbs {
		bvh.Insert(aabb)
		// visualize
		var sb strings.Builder
		toString(&sb, bvh.root)
		t.Log(sb.String())
	}
	bvh.Find(Vec2{0.5, 0.5}, func(n *Node2) {
		t.Logf("find! %v", n.box)
	})
}

func TestTree2_Insert2_notLinkedTable(t *testing.T) {
	const items = 1000
	var bvh Tree2
	for i := 0; i < items; i++ {
		pos := Vec2{0, float64(i)}
		bvh.Insert(AABB2{
			Upper: Vec2{pos[0] + 1, pos[1] + 1},
			Lower: Vec2{pos[0], pos[1]},
		})
	}
	//calc depth
	if depth := lookupDepth(bvh.root); depth > items/2 {
		t.Errorf("the bvh is unbalanced: depth %d", depth)
	} else {
		t.Logf("the depth of bvh with %d element is %d", items, depth)
	}
}

func toString(sb *strings.Builder, n *Node2) {
	if n.isLeaf {
		_, _ = fmt.Fprintf(sb, "(%v,%v)", n.box.Upper[1], n.box.Upper[0])
		return
	}
	sb.WriteByte('{')
	v1 := n.children[0]
	if v1 != nil {
		toString(sb, v1)
	}
	v2 := n.children[1]
	if v2 != nil {
		if v1 != nil {
			sb.WriteString(", ")
		}
		toString(sb, v2)
	}
	sb.WriteByte('}')
}

func lookupDepth(n *Node2) int {
	depth := 0
	for _, child := range n.children {
		if child != nil {
			subdepth := lookupDepth(child)
			if subdepth > depth {
				depth = subdepth
			}
		}
	}
	return depth + 1 // add itself
}

func BenchmarkTree2_Insert(b *testing.B) {
	const size = 25
	// generate test cases
	aabbs := make([]AABB2, b.N)
	poses := make([]Vec2, b.N)
	for i := range aabbs {
		poses[i] = Vec2{rand.Float64() * 1e4, rand.Float64() * 1e4}
		aabbs[i] = AABB2{
			Upper: Vec2{poses[i][0] + size, poses[i][0] + size},
			Lower: Vec2{poses[i][0] - size, poses[i][0] - size},
		}
	}
	b.ResetTimer()

	var bvh Tree2
	for _, v := range aabbs {
		bvh.Insert(v)
	}
	for _, v := range poses {
		bvh.Find(v, func(n *Node2) {})
	}
}
