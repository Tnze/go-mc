package ecs

import (
	"golang.org/x/exp/constraints"
	"math/bits"
	"unsafe"
)

const uintsize = uint(unsafe.Sizeof(BitSet{}.values[0]))

type BitSet struct {
	// TODO: this is not a BitSet, I'm just testing
	values []uint
}

func (b *BitSet) Set(i Index) (old bool) {
	index := uint(i) / uintsize
	offset := uint(i) % uintsize
	if index >= uint(len(b.values)) {
		if index < uint(cap(b.values)) {
			b.values = b.values[:index]
		} else {
			newValues := make([]uint, index+1)
			copy(newValues, b.values)
			b.values = newValues
		}
	}
	v := &b.values[index]
	mask := uint(1 << offset)
	old = *v&mask != 0
	*v |= mask
	return
}

func (b *BitSet) Unset(i Index) (old bool) {
	index := uint(i) / uintsize
	offset := uint(i) % uintsize
	if index < uint(len(b.values)) {
		v := &b.values[index]
		mask := uint(1 << offset)
		old = *v&mask != 0
		*v &= ^mask
	}
	return
}

func (b *BitSet) Contains(i Index) bool {
	index := uint(i) / uintsize
	offset := uint(i) % uintsize
	return index < uint(len(b.values)) && b.values[index]&(1<<offset) != 0
}

func (b *BitSet) And(other *BitSet) *BitSet {
	result := BitSet{values: make([]uint, min(len(b.values), len(other.values)))}
	for i := range b.values {
		result.values[i] = b.values[i] & other.values[i]
	}
	return &result
}

func (b *BitSet) AndNot(other BitSet) *BitSet {
	result := BitSet{values: make([]uint, max(len(b.values), len(other.values)))}
	for i := range b.values {
		result.values[i] = b.values[i] & ^other.values[i]
	}
	return &result
}

func (b *BitSet) Range(f func(eid Index)) {
	for i, v := range b.values {
		base := int(unsafe.Sizeof(v)) * i
		for v != 0 {
			p := bits.TrailingZeros(v)
			f(Index(base + p))
			v ^= 1 << p
		}
	}
}

func max[T constraints.Integer](a, b T) T {
	if a > b {
		return a
	}
	return b
}

func min[T constraints.Integer](a, b T) T {
	if a < b {
		return a
	}
	return b
}
