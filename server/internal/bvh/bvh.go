package bvh

import (
	"container/heap"
	"fmt"

	"golang.org/x/exp/constraints"
)

type Node[I constraints.Float, B interface {
	Union(B) B
	Surface() I
}, V any] struct {
	Box      B
	Value    V
	parent   *Node[I, B, V]
	children [2]*Node[I, B, V]
	isLeaf   bool
}

func (n *Node[I, B, V]) findAnotherChild(not *Node[I, B, V]) *Node[I, B, V] {
	if n.children[0] == not {
		return n.children[1]
	} else if n.children[1] == not {
		return n.children[0]
	}
	panic("unreachable, please make sure the 'not' is the n's child")
}

func (n *Node[I, B, V]) findChildPointer(child *Node[I, B, V]) **Node[I, B, V] {
	if n.children[0] == child {
		return &n.children[0]
	} else if n.children[1] == child {
		return &n.children[1]
	}
	panic("unreachable, please make sure the 'not' is the n's child")
}

func (n *Node[I, B, V]) each(test func(bound B) bool, foreach func(n *Node[I, B, V]) bool) bool {
	if n == nil {
		return true
	}
	if n.isLeaf {
		return !test(n.Box) || foreach(n)
	} else {
		return n.children[0].each(test, foreach) && n.children[1].each(test, foreach)
	}
}

type Tree[I constraints.Float, B interface {
	Union(B) B
	Surface() I
}, V any] struct {
	root *Node[I, B, V]
}

func (t *Tree[I, B, V]) Insert(leaf B, value V) (n *Node[I, B, V]) {
	n = &Node[I, B, V]{
		Box:      leaf,
		Value:    value,
		parent:   nil,
		children: [2]*Node[I, B, V]{nil, nil},
		isLeaf:   true,
	}
	if t.root == nil {
		t.root = n
		return
	}

	// Stage 1: find the best sibling for the new leaf
	sibling := t.root
	bestCost := t.root.Box.Union(leaf).Surface()
	parentTo := &t.root // the parent's children pointer which point to the sibling

	var queue searchHeap[I, Node[I, B, V]]
	queue.Push(searchItem[I, Node[I, B, V]]{pointer: t.root, parentTo: &t.root})

	leafCost := leaf.Surface()
	for queue.Len() > 0 {
		p := heap.Pop(&queue).(searchItem[I, Node[I, B, V]])
		// determine if node p has the best cost
		mergeSurface := p.pointer.Box.Union(leaf).Surface()
		deltaCost := mergeSurface - p.pointer.Box.Surface()
		cost := p.inheritedCost + mergeSurface
		if cost <= bestCost {
			bestCost = cost
			sibling = p.pointer
			parentTo = p.parentTo
		}
		// determine if it is worthwhile to explore the children of node p.
		inheritedCost := p.inheritedCost + deltaCost // lower bound
		if !p.pointer.isLeaf && inheritedCost+leafCost < bestCost {
			heap.Push(&queue, searchItem[I, Node[I, B, V]]{
				pointer:       p.pointer.children[0],
				parentTo:      &p.pointer.children[0],
				inheritedCost: inheritedCost,
			})
			heap.Push(&queue, searchItem[I, Node[I, B, V]]{
				pointer:       p.pointer.children[1],
				parentTo:      &p.pointer.children[1],
				inheritedCost: inheritedCost,
			})
		}
	}

	// Stage 2: create a new parent
	*parentTo = &Node[I, B, V]{
		Box:      sibling.Box.Union(leaf), // we will calculate in Stage3
		parent:   sibling.parent,
		children: [2]*Node[I, B, V]{sibling, n},
		isLeaf:   false,
	}
	n.parent = *parentTo
	sibling.parent = *parentTo

	// Stage 3: walk back up the tree refitting AABBs
	for p := *parentTo; p != nil; p = p.parent {
		p.Box = p.children[0].Box.Union(p.children[1].Box)
		t.rotate(p)
	}
	return
}

func (t *Tree[I, B, V]) Delete(n *Node[I, B, V]) any {
	if n.parent == nil {
		// n is the root
		t.root = nil
		return n.Value
	}
	sibling := n.parent.findAnotherChild(n)
	grand := n.parent.parent
	if grand == nil {
		// n's parent is root
		t.root = sibling
		sibling.parent = nil
	} else {
		p := grand.findChildPointer(n.parent)
		*p = sibling
		sibling.parent = grand
		for p := sibling.parent; p.parent != nil; p = p.parent {
			p.Box = p.children[0].Box.Union(p.children[1].Box)
			t.rotate(p)
		}
	}
	return n.Value
}

func (t *Tree[I, B, V]) rotate(n *Node[I, B, V]) {
	if n.isLeaf || n.parent == nil {
		return
	}
	// trying to swap n's sibling and children
	sibling := n.parent.findAnotherChild(n)
	current := n.Box.Surface()
	if n.children[1].Box.Union(sibling.Box).Surface() < current {
		// swap n.children[0] and sibling
		t1 := [2]*Node[I, B, V]{n, n.children[0]}
		t2 := [2]*Node[I, B, V]{sibling, n.children[1]}
		n.parent.children, n.children, n.children[0].parent, sibling.parent = t1, t2, n.parent, n
		n.Box = n.children[0].Box.Union(n.children[1].Box)
	} else if n.children[0].Box.Union(sibling.Box).Surface() < current {
		// swap n.children[1] and sibling
		t1 := [2]*Node[I, B, V]{n, n.children[1]}
		t2 := [2]*Node[I, B, V]{sibling, n.children[0]}
		n.parent.children, n.children, n.children[1].parent, sibling.parent = t1, t2, n.parent, n
		n.Box = n.children[0].Box.Union(n.children[1].Box)
	}
}

func (t *Tree[I, B, V]) Find(test func(bound B) bool, foreach func(n *Node[I, B, V]) bool) {
	t.root.each(test, foreach)
}

func (t Tree[I, B, V]) String() string {
	return t.root.String()
}

func (n *Node[I, B, V]) String() string {
	if n.isLeaf {
		return fmt.Sprint(n.Value)
	} else {
		return fmt.Sprintf("{%v, %v}", n.children[0], n.children[1])
	}
}

func TouchPoint[Vec any, B interface{ WithIn(Vec) bool }](point Vec) func(bound B) bool {
	return func(bound B) bool {
		return bound.WithIn(point)
	}
}

func TouchBound[B interface{ Touch(B) bool }](other B) func(bound B) bool {
	return func(bound B) bool {
		return bound.Touch(other)
	}
}

type (
	searchHeap[I constraints.Float, V any] []searchItem[I, V]
	searchItem[I constraints.Float, V any] struct {
		pointer       *V
		parentTo      **V
		inheritedCost I
	}
)

func (h searchHeap[I, V]) Len() int           { return len(h) }
func (h searchHeap[I, V]) Less(i, j int) bool { return h[i].inheritedCost < h[j].inheritedCost }
func (h searchHeap[I, V]) Swap(i, j int)      { h[i], h[j] = h[j], h[i] }
func (h *searchHeap[I, V]) Push(x any)        { *h = append(*h, x.(searchItem[I, V])) }
func (h *searchHeap[I, V]) Pop() any {
	old := *h
	n := len(old)
	x := old[n-1]
	*h = old[0 : n-1]
	return x
}
