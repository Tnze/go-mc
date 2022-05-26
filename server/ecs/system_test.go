package ecs

type PositionComponent struct {
	X, Y int
}

type MySystem1 struct {
	*PositionComponent
}

func (s *MySystem1) Update(w *World) {

}
