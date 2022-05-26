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

func SetResource[Res any](w *World, v Res) *Res {
	w.resources[reflect.TypeOf(v)] = &v
	return &v
}

func (w *World) Remove(resource any) any {
	t := reflect.ValueOf(resource).Type()
	resource = w.resources[t]
	delete(w.resources, t)
	return resource
}

func GetResource[Res any](w *World) *Res {
	var res Res
	t := reflect.TypeOf(res)
	if v, ok := w.resources[t]; ok {
		return v.(*Res)
	}
	panic("Resource " + t.Name() + " not found")
}

func (w *World) GetResourceRaw(t reflect.Type) any {
	v, _ := w.resources[t]
	return v
}

func GetComponent[T any](w *World) *MaskedStorage[T] {
	var value T
	t := reflect.ValueOf(value).Type()
	if res, ok := w.resources[t]; ok {
		return res.(*MaskedStorage[T])
	}
	panic("Component " + t.Name() + " not found")
}

// Register the component with the storage.
//
// Will be changed to func (w *World) Register[C Component]() after Go support it
func Register[T any, S Storage[T]](w *World) {
	var value T
	t := reflect.TypeOf(value)
	if _, ok := w.resources[t]; ok {
		panic("Component " + t.Name() + " already exist")
	}
	var storage S
	var storageInt Storage[T]
	storageType := reflect.TypeOf(storage)
	if storageType.Kind() == reflect.Pointer {
		storageInt = reflect.New(storageType.Elem()).Interface().(Storage[T])
	} else {
		storageInt = storage
	}
	ms := MaskedStorage[T]{Storage: storageInt}
	ms.Init()
	w.resources[t] = &ms
}

func (w *World) CreateEntity(components ...any) (i Index) {
	type Storage interface{ SetAny(Index, any) }
	eid := Index(atomic.AddUint32((*uint32)(&w.maxEID), 1))
	for _, c := range components {
		w.resources[reflect.TypeOf(c)].(Storage).SetAny(eid, c)
	}
	return eid
}

func (w *World) DeleteEntity(eid Index) {
	type Storage interface{ Del(eid Index) }
	for _, r := range w.resources {
		if c, ok := r.(Storage); ok {
			c.Del(eid)
		}
	}
}
