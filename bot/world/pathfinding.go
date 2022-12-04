package world

import (
	"github.com/Tnze/go-mc/bot/maths"
	"math"
)

type Node struct {
	point   maths.Vec3d // The point is a vector3d
	parent  *Node       // The parent node in the search tree
	f, g, h float64     // f, g, and h values used by the A* algorithm
}

func euclideanDistance(a, b maths.Vec3d) float64 {
	return math.Sqrt(math.Pow(float64(a.X-b.X), 2) + math.Pow(float64(a.Y-b.Y), 2) + math.Pow(float64(a.Z-b.Z), 2))
}

func (w *World) PathFind(start, end maths.Vec3d) []maths.Vec3d {
	startNode := &Node{point: start, f: 0, g: 0, h: 0}
	endNode := &Node{point: end, f: 0, g: 0, h: 0}
	openList := []*Node{startNode}
	var closedList []*Node

	for len(openList) > 0 {
		// Find the node with the lowest f value in the open list
		currentNode := openList[0]
		currentIndex := 0
		for i, node := range openList {
			if node.f < currentNode.f {
				currentNode = node
				currentIndex = i
			}
		}

		openList = append(openList[:currentIndex], openList[currentIndex+1:]...)
		closedList = append(closedList, currentNode)

		if currentNode.point == endNode.point {
			path := make([]maths.Vec3d, 0)
			current := currentNode
			for current != nil {
				path = append(path, current.point)
				current = current.parent
			}
			return backtrace(path)
		}

		neighbors := toNodes(w.GetNeighbors(currentNode.point), currentNode)

		for _, neighbor := range neighbors {
			if contains(closedList, neighbor) {
				continue
			} else {
				// If the neighbor is not in the open list, calculate the g, h, and f values
				neighbor.g = currentNode.g + euclideanDistance(currentNode.point, neighbor.point)
				neighbor.h = euclideanDistance(neighbor.point, endNode.point)
				neighbor.f = neighbor.g + neighbor.h
				openList = append(openList, neighbor)
			}
		}
	}
	return make([]maths.Vec3d, 0)
}

func contains(list []*Node, node *Node) bool {
	for _, n := range list {
		if node == n {
			return true
		}
	}
	return false
}

func toNodes(points []maths.Vec3d, current *Node) []*Node {
	nodes := make([]*Node, 0)
	for _, point := range points {
		nodes = append(nodes, &Node{point: point, parent: current})
	}
	return nodes
}

func backtrace(node []maths.Vec3d) []maths.Vec3d {
	for i := 0; i < len(node)/2; i++ {
		node[i], node[len(node)-i-1] = node[len(node)-i-1], node[i]
	}
	return node
}
