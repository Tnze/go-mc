package bvh

//
//import (
//	"container/heap"
//	"fmt"
//	"golang.org/x/exp/constraints"
//)
//
//type Node[Vec constraints.Signed | constraints.Float, B interface {
//	WithIn(Vec) bool
//	Union(B) B
//	Surface() float64
//}, V any] struct {
//	box      B
//	Value    V
//	parent   *Node[Vec, B, V]
//	children [2]*Node[Vec, B, V]
//	isLeaf   bool
//}
//
//func (n *Node[Vec, B, V]) findAnotherChild(not *Node[Vec, B, V]) *Node[Vec, B, V] {
//	if n.children[0] == not {
//		return n.children[1]
//	} else if n.children[1] == not {
//		return n.children[0]
//	}
//	panic("unreachable, please make sure the 'not' is the n's child")
//}
//
//func (n *Node[Vec, B, V]) findChildPointer(child *Node[Vec, B, V]) **Node[Vec, B, V] {
//	if n.children[0] == child {
//		return &n.children[0]
//	} else if n.children[1] == child {
//		return &n.children[1]
//	}
//	panic("unreachable, please make sure the 'not' is the n's child")
//}
//
//type Tree[I constraints.Signed | constraints.Float, B interface {
//	Union(B) B
//	Surface() I
//}, V any] struct {
//	root *Node[I, B, V]
//}
//
//func (t *Tree[Vec, B, V]) Insert(leaf B, value V) (n *Node[Vec, B, V]) {
//	n = &Node[Vec, B, V]{
//		box:      leaf,
//		Value:    value,
//		parent:   nil,
//		children: [2]*Node[Vec, B, V]{},
//		isLeaf:   true,
//	}
//	if t.root == nil {
//		t.root = n
//		return
//	}
//
//	// Stage 1: find the best sibling for the new leaf
//	sibling := t.root
//	bestCost := t.root.box.Union(leaf).Surface()
//	parentTo := &t.root // the parent's children pointer which point to the sibling
//	queue := searchHeap[Node[Vec, B, V]]{searchItem[Node[Vec, B, V]]{pointer: t.root, parentTo: &t.root}}
//
//	leafCost := leaf.Surface()
//	for queue.Len() > 0 {
//		p := heap.Pop(&queue).(searchItem[Node[Vec, B, V]])
//		// determine if node p has the best cost
//		mergeSurface := p.pointer.box.Union(leaf).Surface()
//		deltaCost := mergeSurface - p.pointer.box.Surface()
//		cost := p.inheritedCost + mergeSurface
//		if cost <= bestCost {
//			bestCost = cost
//			sibling = p.pointer
//			parentTo = p.parentTo
//		}
//		// determine if it is worthwhile to explore the children of node p.
//		inheritedCost := p.inheritedCost + deltaCost // lower bound
//		if !p.pointer.isLeaf && inheritedCost+leafCost < bestCost {
//			heap.Push(&queue, searchItem[Node[Vec, B, V]]{
//				pointer:       p.pointer.children[0],
//				parentTo:      &p.pointer.children[0],
//				inheritedCost: inheritedCost,
//			})
//			heap.Push(&queue, searchItem[Node[Vec, B, V]]{
//				pointer:       p.pointer.children[1],
//				parentTo:      &p.pointer.children[1],
//				inheritedCost: inheritedCost,
//			})
//		}
//	}
//
//	// Stage 2: create a new parent
//	*parentTo = &Node[Vec, B, V]{
//		box:      sibling.box.Union(leaf), // we will calculate in Stage3
//		parent:   sibling.parent,
//		children: [2]*Node[Vec, B, V]{sibling, n},
//		isLeaf:   false,
//	}
//	n.parent = *parentTo
//	sibling.parent = *parentTo
//
//	// Stage 3: walk back up the tree refitting AABBs
//	for p := *parentTo; p != nil; p = p.parent {
//		p.box = p.children[0].box.Union(p.children[1].box)
//		t.rotate(p)
//	}
//	return
//}
//
//func (t *Tree[Vec, B, V]) Delete(n *Node[Vec, B, V]) interface{} {
//	if n.parent == nil {
//		// n is the root
//		t.root = nil
//		return n.Value
//	}
//	sibling := n.parent.findAnotherChild(n)
//	grand := n.parent.parent
//	if grand == nil {
//		// n's parent is root
//		t.root = sibling
//		sibling.parent = nil
//	} else {
//		p := grand.findChildPointer(n.parent)
//		*p = sibling
//		sibling.parent = grand
//		for p := sibling.parent; p.parent != nil; p = p.parent {
//			p.box = p.children[0].box.Union(p.children[1].box)
//			t.rotate(p)
//		}
//	}
//	return n.Value
//}
//
//func (t *Tree[Vec, B, V]) rotate(n *Node[Vec, B, V]) {
//	if n.isLeaf || n.parent == nil {
//		return
//	}
//	// trying to swap n's sibling and children
//	sibling := n.parent.findAnotherChild(n)
//	current := n.box.Surface()
//	if n.children[1].box.Union(sibling.box).Surface() < current {
//		// swap n.children[0] and sibling
//		n.parent.children, n.children, n.children[0].parent, sibling.parent = [2]*Node[Vec, B, V]{n, n.children[0]}, [2]*Node[Vec, B, V]{sibling, n.children[1]}, n.parent, n
//		n.box = n.children[0].box.Union(n.children[1].box)
//	} else if n.children[0].box.Union(sibling.box).Surface() < current {
//		// swap n.children[1] and sibling
//		n.parent.children, n.children, n.children[1].parent, sibling.parent = [2]*Node[Vec, B, V]{n, n.children[1]}, [2]*Node[Vec, B, V]{sibling, n.children[0]}, n.parent, n
//		n.box = n.children[0].box.Union(n.children[1].box)
//	}
//}
//
////func lookupPoint[B interface {
////	Union(B) B
////	Surface() float64
////}, V any](n *Node[B, V], point Vec2, f func(v V)) {
////	if n == nil || !n.box.WithIn(point) {
////		return
////	}
////	if n.isLeaf {
////		f(n.Value)
////	} else {
////		lookupVec(n.children[0], point, f)
////		lookupVec(n.children[1], point, f)
////	}
////}
//
////
////func lookupAABB(n *Node, aabb AABB, f func(v interface{})) {
////	if n == nil || !n.box.Touch(aabb) {
////		return
////	}
////	if n.isLeaf {
////		f(n.Value)
////	} else {
////		lookupAABB(n.children[0], aabb, f)
////		lookupAABB(n.children[1], aabb, f)
////	}
////}
//
//type searchHeap[V any] []searchItem[V]
//type searchItem[V any] struct {
//	pointer       *V
//	parentTo      **V
//	inheritedCost float64
//}
//
//func (h searchHeap[V]) Len() int            { return len(h) }
//func (h searchHeap[V]) Less(i, j int) bool  { return h[i].inheritedCost < h[j].inheritedCost }
//func (h searchHeap[V]) Swap(i, j int)       { h[i], h[j] = h[j], h[i] }
//func (h *searchHeap[V]) Push(x interface{}) { *h = append(*h, x.(searchItem[V])) }
//func (h *searchHeap[V]) Pop() interface{} {
//	old := *h
//	n := len(old)
//	x := old[n-1]
//	*h = old[0 : n-1]
//	return x
//}
//
//func (t Tree[Vec, B, V]) String() string {
//	return t.root.String()
//}
//
//func (n *Node[Vec, B, V]) String() string {
//	if n.isLeaf {
//		return fmt.Sprint(n.Value)
//	} else {
//		return fmt.Sprintf("{%v, %v}", n.children[0], n.children[1])
//	}
//}
