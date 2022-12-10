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
			return backtrace(path[1:])
		}

		neighbor := filterNeighbors(w, w.GetNeighbors(currentNode.point), endNode.point)
		if !contains(closedList, neighbor) {
			newNode := &Node{point: neighbor, parent: currentNode}
			newNode.g = currentNode.g + euclideanDistance(currentNode.point, newNode.point)
			newNode.h = euclideanDistance(newNode.point, endNode.point)
			newNode.f = newNode.g + newNode.h
			openList = append(openList, newNode)
		}
	}
	return make([]maths.Vec3d, 0)
}

func contains(list []*Node, node maths.Vec3d) bool {
	for _, n := range list {
		if n.point == node {
			return true
		}
	}
	return false
}

func backtrace(node []maths.Vec3d) []maths.Vec3d {
	for i := 0; i < len(node)/2; i++ {
		node[i], node[len(node)-i-1] = node[len(node)-i-1], node[i]
	}
	return node
}

func filterNeighbors(w *World, points []maths.Vec3d, end maths.Vec3d) maths.Vec3d {
	var closestPoint maths.Vec3d
	for _, point := range points {
		if euclideanDistance(point, end) < euclideanDistance(closestPoint, end) {
			closestPoint = point
		}
	}
	return closestPoint
}
