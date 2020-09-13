package world

import (
	"bytes"
	"fmt"

	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

// DecodeChunkColumn decode the chunk data structure.
// If decoding went error, successful decoded data will be returned.
func DecodeChunkColumn(mask int32, data []byte) (*Chunk, error) {
	var c Chunk
	r := bytes.NewReader(data)
	for sectionY := 0; sectionY < 16; sectionY++ {
		// If the section's bit set in the mask
		if (mask & (1 << uint(sectionY))) != 0 {
			// read section
			sec, err := readSection(r)
			if err != nil {
				return &c, fmt.Errorf("read section[%d] error: %w", sectionY, err)
			}
			c.Sections[sectionY] = sec
		}
	}
	return &c, nil
}

func perBits(BitsPerBlock byte) int {
	switch {
	case BitsPerBlock <= 4:
		return 4
	case BitsPerBlock < 9:
		return int(BitsPerBlock)
	default:
		return block.BitsPerBlock
	}
}

func readSection(data pk.DecodeReader) (s Section, err error) {
	var BlockCount pk.Short
	if err := BlockCount.Decode(data); err != nil {
		return nil, fmt.Errorf("read block count error: %w", err)
	}
	var bpb pk.UnsignedByte
	if err := bpb.Decode(data); err != nil {
		return nil, fmt.Errorf("read bits per block error: %w", err)
	}
	// If bpb values greater than or equal to 9, use directSection.
	// Otherwise use paletteSection.
	var palettes []BlockStatus
	var palettesIndex map[BlockStatus]int
	if bpb < 9 {
		// read palettes
		var length pk.VarInt
		if err := length.Decode(data); err != nil {
			return nil, fmt.Errorf("read palettes length error: %w", err)
		}
		palettes = make([]BlockStatus, length)
		palettesIndex = make(map[BlockStatus]int, length)
		for i := 0; i < int(length); i++ {
			var v pk.VarInt
			if err := v.Decode(data); err != nil {
				return nil, fmt.Errorf("read palettes[%d] error: %w", i, err)
			}
			palettes[i] = BlockStatus(v)
			palettesIndex[BlockStatus(v)] = i
		}
	}

	// read data array
	var dataLen pk.VarInt
	if err := dataLen.Decode(data); err != nil {
		return nil, fmt.Errorf("read data array length error: %w", err)
	}
	if int(dataLen) < 16*16*16*int(bpb)/64 {
		return nil, fmt.Errorf("data length (%d) is not enough of given bpb (%d)", dataLen, bpb)
	}
	dataArray := make([]uint64, dataLen)
	for i := 0; i < int(dataLen); i++ {
		var v pk.Long
		if err := v.Decode(data); err != nil {
			return nil, fmt.Errorf("read dataArray[%d] error: %w", i, err)
		}
		dataArray[i] = uint64(v)
	}

	sec := directSection{bpb: perBits(byte(bpb)), data: dataArray}
	if bpb < 9 {
		return &paletteSection{
			palette:       palettes,
			palettesIndex: palettesIndex,
			directSection: sec,
		}, nil
	} else {
		return &sec, nil
	}
}

type directSection struct {
	bpb  int
	data []uint64
}

func (d *directSection) GetBlock(offset int) BlockStatus {
	offset *= d.bpb
	padding := offset % 64
	block := uint32(d.data[offset/64] >> padding)
	if padding > 64-d.bpb {
		l := 64 - padding
		block |= uint32(d.data[offset/64+1] << l)
	}
	return BlockStatus(block & (1<<d.bpb - 1)) // mask
}

func (d *directSection) SetBlock(offset int, s BlockStatus) {
	offset *= d.bpb
	padding := offset % 64
	mask := ^uint64((1<<d.bpb - 1) << padding)
	d.data[offset/64] = d.data[offset/64]&mask | uint64(s)<<padding
	if padding > 64-d.bpb {
		l := padding - (64 - d.bpb)
		const maxUint64 = 1<<64 - 1
		d.data[offset/64+1] = d.data[offset/64+1]&(maxUint64<<l) | uint64(s)>>(64-padding)
	}
}

func (d *directSection) CanContain(s BlockStatus) bool {
	return s <= (1<<d.bpb - 1)
}

func (d *directSection) clone(bpb int) *directSection {
	newSection := &directSection{
		bpb:  bpb,
		data: make([]uint64, 16*16*16*bpb/64),
	}
	for offset := 0; offset < 16*16*16; offset++ {
		newSection.SetBlock(offset, d.GetBlock(offset))
	}
	return newSection
}

type paletteSection struct {
	palette       []BlockStatus
	palettesIndex map[BlockStatus]int
	directSection
}

func (p *paletteSection) GetBlock(offset int) BlockStatus {
	v := p.directSection.GetBlock(offset)
	return p.palette[v]
}

func (p *paletteSection) SetBlock(offset int, s BlockStatus) {
	if i, ok := p.palettesIndex[s]; ok {
		p.directSection.SetBlock(offset, BlockStatus(i))
		return
	}
	i := len(p.palette)
	p.palette = append(p.palette, s)
	p.palettesIndex[s] = i
	if !p.directSection.CanContain(BlockStatus(i)) {
		// Increase the underlying directSection
		// Suppose that old bpb fit len(p.palette) before it appended.
		// So bpb+1 must enough for new len(p.palette).
		p.directSection = *p.directSection.clone(p.bpb + 1)
	}
	p.directSection.SetBlock(offset, BlockStatus(i))
}
