package ecs

import "testing"

func Test_common(t *testing.T) {
	// W
	w := NewWorld()
	// C
	type pos [2]int
	type vel [2]int
	Register(w, pos{})
	Register(w, vel{})
	// E
	w.CreateEntity(pos{0, 0})
	w.CreateEntity(vel{1, 2})
	w.CreateEntity(pos{1, 2}, vel{2, 0})
	// S
	s1 := FuncSystem(func(p pos) {
		t.Log("system 1", p)
	})
	s2 := FuncSystem(func(p pos, v vel) {
		t.Log("system 2", p, v)
	})
	// Run
	s1.Update(w)
	s2.Update(w)
}
