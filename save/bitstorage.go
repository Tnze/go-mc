package save

import (
	"fmt"
	"math"
)

// This implement the format since Minecraft 1.16
type BitStorage struct {
	data []uint64
	mask uint64

	bits, size    int
	valuesPerLong int

	divideMul   int
	divideAdd   int
	divideShift int
}

// NewBitStorage create a new BitStorage, // TODO: document
func NewBitStorage(bits, size int, arrl []uint64) (b *BitStorage) {
	var _MAGIC = [...]int{
		-1, -1, 0,
		math.MinInt32, 0, 0,
		1431655765, 1431655765, 0,
		math.MinInt32, 0, 1,
		858993459, 858993459, 0,
		715827882, 715827882, 0,
		613566756, 613566756, 0,
		math.MinInt32, 0, 2,
		477218588, 477218588, 0,
		429496729, 429496729, 0,
		390451572, 390451572, 0,
		357913941, 357913941, 0,
		330382099, 330382099, 0,
		306783378, 306783378, 0,
		286331153, 286331153, 0,
		math.MinInt32, 0, 3,
		252645135, 252645135, 0,
		238609294, 238609294, 0,
		226050910, 226050910, 0,
		214748364, 214748364, 0,
		204522252, 204522252, 0,
		195225786, 195225786, 0,
		186737708, 186737708, 0,
		178956970, 178956970, 0,
		171798691, 171798691, 0,
		165191049, 165191049, 0,
		159072862, 159072862, 0,
		153391689, 153391689, 0,
		148102320, 148102320, 0,
		143165576, 143165576, 0,
		138547332, 138547332, 0,
		math.MinInt32, 0, 4,
		130150524, 130150524, 0,
		126322567, 126322567, 0,
		122713351, 122713351, 0,
		119304647, 119304647, 0,
		116080197, 116080197, 0,
		113025455, 113025455, 0,
		110127366, 110127366, 0,
		107374182, 107374182, 0,
		104755299, 104755299, 0,
		102261126, 102261126, 0,
		99882960, 99882960, 0,
		97612893, 97612893, 0,
		95443717, 95443717, 0,
		93368854, 93368854, 0,
		91382282, 91382282, 0,
		89478485, 89478485, 0,
		87652393, 87652393, 0,
		85899345, 85899345, 0,
		84215045, 84215045, 0,
		82595524, 82595524, 0,
		81037118, 81037118, 0,
		79536431, 79536431, 0,
		78090314, 78090314, 0,
		76695844, 76695844, 0,
		75350303, 75350303, 0,
		74051160, 74051160, 0,
		72796055, 72796055, 0,
		71582788, 71582788, 0,
		70409299, 70409299, 0,
		69273666, 69273666, 0,
		68174084, 68174084, 0,
		math.MinInt32, 0, 5,
	}

	n3 := 3 * (64/bits - 1)
	b = &BitStorage{
		mask:          1<<bits - 1,
		bits:          bits,
		size:          size,
		valuesPerLong: 64 / bits,
		divideMul:     _MAGIC[n3+0],
		divideAdd:     _MAGIC[n3+1],
		divideShift:   _MAGIC[n3+2],
	}
	dataLen := (size + b.valuesPerLong - 1) / b.valuesPerLong
	if arrl != nil {
		if len(arrl) != dataLen {
			panic(fmt.Errorf("invalid length given for storage, got: %d but expected: %d", len(arrl), dataLen))
		}
		b.data = arrl
	} else {
		b.data = make([]uint64, dataLen)
	}
	return
}

func (b *BitStorage) cellIndex(n int) int {
	return int((uint64(uint32(n))*uint64(b.divideMul) + uint64(b.divideAdd)) >> 32 >> b.divideShift)
}

func (b *BitStorage) Swap(i, v int) (old int) {
	if i < 0 || i > b.size-1 ||
		v < 0 || uint64(v) > b.mask {
		panic("out of bounds")
	}
	c := b.cellIndex(i)
	l := b.data[c]
	offset := uint64((i - c*b.valuesPerLong) * b.bits)
	old = int(l >> offset & b.mask)
	b.data[c] = l&(b.mask<<offset^math.MaxUint64) | (uint64(v)&b.mask)<<offset
	return
}

func (b *BitStorage) Set(i, v int) {
	if i < 0 || i > b.size-1 ||
		v < 0 || uint64(v) > b.mask {
		panic("out of bounds")
	}
	c := b.cellIndex(i)
	l := b.data[c]
	offset := (i - c*b.valuesPerLong) * b.bits
	b.data[c] = l&(b.mask<<offset^math.MaxUint64) | (uint64(v)&b.mask)<<offset
}

func (b *BitStorage) Get(i int) int {
	if i < 0 || i > b.size-1 {
		panic("out of bounds")
	}
	c := b.cellIndex(i)
	l := b.data[c]
	offset := (i - c*b.valuesPerLong) * b.bits
	return int(l >> offset & b.mask)
}
