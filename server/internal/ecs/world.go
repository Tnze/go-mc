package ecs

import (
	"reflect"
	"sync/atomic"
)

type World struct {
	resources map[reflect.Type]any
	maxEID    Index
}

func NewWorld() *World {
	return &World{resources: make(map[reflect.Type]any)}
}

func (w *World) Insert(resource any) {
	t := reflect.ValueOf(resource).Type()
	w.resources[t] = resource
}

func (w *World) Remove(resource any) any {
	t := reflect.ValueOf(resource).Type()
	resource = w.resources[t]
	delete(w.resources, t)
	return resource
}

func (w *World) GetResource(resource any) any {
	if resource == nil {
		return nil
	}
	t := reflect.ValueOf(resource).Type()
	v, _ := w.resources[t]
	return v
}

func (w *World) GetResourceRaw(t reflect.Type) any {
	v, _ := w.resources[t]
	return v
}

func Register[T any](w *World, component T) {
	t := reflect.TypeOf(component)
	s := NewHashMapStorage[T]()
	w.resources[t] = s
}

func (w *World) CreateEntity(components ...any) (i Index) {
	i = Index(atomic.AddUint32((*uint32)(&w.maxEID), 1))
	for _, c := range components {
		v := reflect.ValueOf(c)
		t := v.Type()
		storage := w.resources[t].(Storage)
		storage.Insert(w.maxEID, c)
	}
	return
}
