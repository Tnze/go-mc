package save

import (
	"fmt"
	"math"
)

// BitStorage implement the compacted data array used in chunk storage.
// https://wiki.vg/Chunk_Format
// This implement the format since Minecraft 1.16
type BitStorage struct {
	data []uint64
	mask uint64

	bits, size    int
	valuesPerLong int
}

// NewBitStorage create a new BitStorage.
// bits is the number of bits per value.
// size is the number of values.
// arrl is optional data for initializing.
// It's length must match the bits and size if it's not nil.
func NewBitStorage(bits, size int, arrl []uint64) (b *BitStorage) {
	b = &BitStorage{
		mask:          1<<bits - 1,
		bits:          bits,
		size:          size,
		valuesPerLong: 64 / bits,
	}
	dataLen := (size + b.valuesPerLong - 1) / b.valuesPerLong
	if arrl != nil {
		if len(arrl) != dataLen {
			panic(initBitStorageErr{ArrlLen: len(arrl), WantLen: dataLen})
		}
		b.data = arrl
	} else {
		b.data = make([]uint64, dataLen)
	}
	return
}

type initBitStorageErr struct {
	ArrlLen int
	WantLen int
}

func (i initBitStorageErr) Error() string {
	return fmt.Sprintf("invalid length given for storage, got: %d but expected: %d", i.ArrlLen, i.WantLen)
}

func (b *BitStorage) calcIndex(n int) (c, o int) {
	c = n / b.valuesPerLong
	o = (n - c*b.valuesPerLong) * b.bits
	return
}

// Swap sets v into [i], and return the previous [i] value.
func (b *BitStorage) Swap(i, v int) (old int) {
	if i < 0 || i > b.size-1 ||
		v < 0 || uint64(v) > b.mask {
		panic("out of bounds")
	}
	c, offset := b.calcIndex(i)
	l := b.data[c]
	old = int(l >> offset & b.mask)
	b.data[c] = l&(b.mask<<offset^math.MaxUint64) | (uint64(v)&b.mask)<<offset
	return
}

// Set sets v into [i]
func (b *BitStorage) Set(i, v int) {
	if i < 0 || i > b.size-1 ||
		v < 0 || uint64(v) > b.mask {
		panic("out of bounds")
	}
	c, offset := b.calcIndex(i)
	l := b.data[c]
	b.data[c] = l&(b.mask<<offset^math.MaxUint64) | (uint64(v)&b.mask)<<offset
}

// Get gets [i] value.
func (b *BitStorage) Get(i int) int {
	if i < 0 || i > b.size-1 {
		panic("out of bounds")
	}
	c, offset := b.calcIndex(i)
	l := b.data[c]
	return int(l >> offset & b.mask)
}
