package world

// bitArray implements a bitfield array where values are packed into uint64
// values. If the next value does not fit into remaining space, the remaining
// space of a uint64 is unused.
type bitArray struct {
	width          uint // bit width of each value
	valsPerElement uint // number of values which fit into a single uint64.

	data []uint64
}

// Size returns the number of elements that can fit into the bit array.
func (b *bitArray) Size() int {
	return int(b.valsPerElement) * len(b.data)
}

func (b *bitArray) Set(idx, val uint) {
	var (
		arrayIdx = idx / b.valsPerElement
		startBit = (idx % b.valsPerElement) * b.width
		mask     = ^uint64((1<<b.width - 1) << startBit) // set for all bits except target
	)
	b.data[arrayIdx] = (b.data[arrayIdx] & mask) | uint64(val<<startBit)
}

func (b *bitArray) Get(idx uint) uint {
	var (
		arrayIdx = idx / b.valsPerElement
		offset   = (idx % b.valsPerElement) * b.width
		mask     = uint64((1<<b.width - 1) << offset) // set for just the target
	)
	return uint(b.data[arrayIdx]&mask) >> offset
}

func valsPerBitArrayElement(bitsPerValue uint) uint {
	return uint(64 / bitsPerValue)
}
