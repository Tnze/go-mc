package ecs

type Index uint32

type Storage interface {
	Get(eid Index) any
	Insert(eid Index, v any)
	Remove(eid Index) any
	BitSet() BitSet
}

type HashMapStorage[T any] struct {
	keys   BitSet
	values map[Index]T
}

func NewHashMapStorage[T any]() *HashMapStorage[T] {
	return &HashMapStorage[T]{
		keys:   BitSet{values: make(map[Index]struct{})},
		values: make(map[Index]T),
	}
}
func (h *HashMapStorage[T]) Get(eid Index) any { return h.values[eid] }
func (h *HashMapStorage[T]) Insert(eid Index, v any) {
	h.keys.Set(eid)
	h.values[eid] = v.(T)
}
func (h *HashMapStorage[T]) Remove(eid Index) any {
	h.keys.Unset(eid)
	v := h.values[eid]
	delete(h.values, eid)
	return v
}
func (h *HashMapStorage[T]) BitSet() BitSet { return h.keys }
