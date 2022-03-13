package bvh

import (
	"container/heap"
	"math"
)

type Vec2 [2]float64

func (v Vec2) Add(other Vec2) Vec2 { return Vec2{v[0] + other[0], v[1] + other[1]} }
func (v Vec2) Sub(other Vec2) Vec2 { return Vec2{v[0] - other[0], v[1] - other[1]} }
func (v Vec2) Max(other Vec2) Vec2 { return Vec2{math.Max(v[0], other[0]), math.Max(v[1], other[1])} }
func (v Vec2) Min(other Vec2) Vec2 { return Vec2{math.Min(v[0], other[0]), math.Min(v[1], other[1])} }

type AABB2 struct{ Upper, Lower Vec2 }

func (aabb AABB2) WithIn(point Vec2) bool {
	return aabb.Lower[0] < point[0] && point[0] < aabb.Upper[0] &&
		aabb.Lower[1] < point[1] && point[1] < aabb.Upper[1]
}

func (aabb AABB2) Union(other AABB2) AABB2 {
	return AABB2{
		Upper: aabb.Upper.Max(other.Upper),
		Lower: aabb.Lower.Min(other.Lower),
	}
}

func (aabb AABB2) Surface() float64 {
	d := aabb.Upper.Sub(aabb.Lower)
	return 2 * (d[0] + d[1])
}

type Node2 struct {
	box      AABB2
	Value    interface{}
	parent   *Node2
	children [2]*Node2
	isLeaf   bool
}

func (n *Node2) findAnotherChild(not *Node2) *Node2 {
	if n.isLeaf {
		return nil
	} else if n.children[0] == not {
		return n.children[1]
	} else if n.children[1] == not {
		return n.children[0]
	}
	panic("unreachable, please make sure the 'not' is the n's child")
}

type Tree2 struct {
	root *Node2
}

func (t *Tree2) Insert(leaf AABB2) (n *Node2) {
	n = &Node2{
		box:      leaf,
		parent:   nil,
		children: [2]*Node2{},
		isLeaf:   true,
	}
	if t.root == nil {
		t.root = n
		return
	}

	// Stage 1: find the best sibling for the new leaf
	sibling := t.root
	bestCost := t.root.box.Union(leaf).Surface()
	parentTo := &t.root // the parent's children pointer which point to the sibling
	queue := searchHeap{searchItem{pointer: t.root, parentTo: &t.root}}

	leafCost := leaf.Surface()
	for len(queue) > 0 {
		p := heap.Pop(&queue).(searchItem)
		// determine if node p has the best cost
		mergeSurface := p.pointer.box.Union(leaf).Surface()
		deltaCost := mergeSurface - p.pointer.box.Surface()
		cost := p.inheritedCost + mergeSurface
		if cost < bestCost {
			bestCost = cost
			sibling = p.pointer
			parentTo = p.parentTo
		}
		// determine if it is worthwhile to explore the children of node p.
		inheritedCost := p.inheritedCost + deltaCost // lower bound
		if inheritedCost+leafCost < bestCost {
			if p.pointer.children[0] != nil {
				heap.Push(&queue, searchItem{
					pointer:       p.pointer.children[0],
					parentTo:      &p.pointer.children[0],
					inheritedCost: inheritedCost,
				})
			}
			if p.pointer.children[1] != nil {
				heap.Push(&queue, searchItem{
					pointer:       p.pointer.children[1],
					parentTo:      &p.pointer.children[1],
					inheritedCost: inheritedCost,
				})
			}
		}
	}

	// Stage 2: create a new parent
	*parentTo = &Node2{
		box:      sibling.box.Union(leaf), // we will calculate in Stage3
		parent:   sibling.parent,
		children: [2]*Node2{sibling, n},
		isLeaf:   false,
	}
	n.parent = *parentTo
	sibling.parent = *parentTo

	// Stage 3: walk back up the tree refitting AABBs
	for p := *parentTo; p.parent != nil; p = p.parent {
		p.box = p.children[0].box.Union(p.children[1].box)
		//TODO: t.rotate(p)
	}
	return
}

func (t *Tree2) Find(point Vec2, f func(*Node2)) {
	t.root.lookup(point, f)
}

func (n *Node2) lookup(point Vec2, f func(node2 *Node2)) {
	if n == nil || !n.box.WithIn(point) {
		return
	}
	if n.isLeaf {
		f(n)
	} else {
		n.children[0].lookup(point, f)
		n.children[1].lookup(point, f)
	}
}

type searchHeap []searchItem
type searchItem struct {
	pointer       *Node2
	parentTo      **Node2
	inheritedCost float64
}

func (h searchHeap) Len() int            { return len(h) }
func (h searchHeap) Less(i, j int) bool  { return h[i].inheritedCost < h[j].inheritedCost }
func (h searchHeap) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
func (h *searchHeap) Push(x interface{}) { *h = append(*h, x.(searchItem)) }
func (h *searchHeap) Pop() interface{} {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
