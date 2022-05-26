package ecs

type BitSetLike interface {
	Set(i Index) (old bool)
	Unset(i Index) (old bool)
	Contains(i Index) bool
	And(other BitSetLike) (result BitSetLike)
	AndNot(other BitSetLike) (result BitSetLike)
	Range(f func(eid Index))
}

type BitSet struct {
	// TODO: this is not a BitSet, I'm just testing
	values map[Index]struct{}
}

func (b BitSet) Set(i Index) (old bool) {
	_, old = b.values[i]
	b.values[i] = struct{}{}
	return
}

func (b BitSet) Unset(i Index) (old bool) {
	_, old = b.values[i]
	delete(b.values, i)
	return
}

func (b BitSet) Contains(i Index) bool {
	_, contains := b.values[i]
	return contains
}

func (b BitSet) And(other BitSetLike) BitSetLike {
	result := BitSet{values: make(map[Index]struct{})}
	for i := range b.values {
		if other.Contains(i) {
			result.values[i] = struct{}{}
		}
	}
	return result
}

func (b BitSet) AndNot(other BitSetLike) BitSetLike {
	result := BitSet{values: make(map[Index]struct{})}
	for i := range b.values {
		if !other.Contains(i) {
			result.values[i] = struct{}{}
		}
	}
	return result
}

func (b BitSet) Range(f func(eid Index)) {
	for i := range b.values {
		f(i)
	}
}
