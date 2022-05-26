package ecs

import (
	"reflect"
	"strconv"
	"unsafe"
)

type Index uint32

type Storage[T any] interface {
	Init()
	GetValue(eid Index) *T
	SetValue(eid Index, v T)
	DelValue(eid Index)
}

type HashMapStorage[T any] struct {
	values map[Index]*T
}

func (h *HashMapStorage[T]) Init()                   { h.values = make(map[Index]*T) }
func (h *HashMapStorage[T]) Len() int                { return len(h.values) }
func (h *HashMapStorage[T]) GetValue(eid Index) *T   { return h.values[eid] }
func (h *HashMapStorage[T]) SetValue(eid Index, v T) { h.values[eid] = &v }
func (h *HashMapStorage[T]) DelValue(eid Index)      { delete(h.values, eid) }
func (h *HashMapStorage[T]) Range(f func(eid Index, value *T)) {
	for i, v := range h.values {
		f(i, v)
	}
}

type NullStorage[T any] struct{}

func (NullStorage[T]) Init() {
	var v T
	if size := unsafe.Sizeof(v); size != 0 {
		typeName := reflect.TypeOf(v).String()
		typeSize := strconv.Itoa(int(size))
		panic("NullStorage can only be used with ZST, " + typeName + " has size of " + typeSize)
	}
}
func (NullStorage[T]) GetValue(eid Index) *T   { return nil }
func (NullStorage[T]) SetValue(eid Index, v T) {}
func (NullStorage[T]) DelValue(eid Index)      {}

type MaskedStorage[T any] struct {
	BitSetLike
	Storage[T]
	Len int
}

func (m *MaskedStorage[T]) Init() {
	if m.BitSetLike == nil {
		m.BitSetLike = BitSet{make(map[Index]struct{})}
	}
	m.Storage.Init()
}
func (m *MaskedStorage[T]) GetValue(eid Index) *T {
	if m.Contains(eid) {
		return m.Storage.GetValue(eid)
	}
	return nil
}
func (m *MaskedStorage[T]) GetValueAny(eid Index) any { return m.GetValue(eid) }
func (m *MaskedStorage[T]) SetValue(eid Index, v T) {
	if !m.BitSetLike.Set(eid) {
		m.Len++
	}
	m.Storage.SetValue(eid, v)
}
func (m *MaskedStorage[T]) SetAny(eid Index, v any) { m.SetValue(eid, v.(T)) }
func (m *MaskedStorage[T]) DelValue(eid Index) {
	if m.BitSetLike.Unset(eid) {
		m.Len--
	}
	m.Storage.DelValue(eid)
}
