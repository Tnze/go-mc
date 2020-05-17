package world

import (
	"bytes"
	"fmt"
	// "io"
	"github.com/Tnze/go-mc/data"
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

func perBits(BitsPerBlock byte) uint32 {
	switch {
	case BitsPerBlock <= 4:
		return 4
	case BitsPerBlock < 9:
		return uint32(BitsPerBlock)
	default:
		return uint32(data.BitsPerBlock) // DefaultBitsPerBlock
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
	var palettes []uint32
	if bpb < 9 {
		// read palettes
		var length pk.VarInt
		if err := length.Decode(data); err != nil {
			return nil, fmt.Errorf("read palettes length error: %w", err)
		}
		palettes = make([]uint32, length)
		for i := 0; i < int(length); i++ {
			var v pk.VarInt
			if err := v.Decode(data); err != nil {
				return nil, fmt.Errorf("read palettes[%d] error: %w", i, err)
			}
			palettes[i] = uint32(v)
		}
	}

	// read data array
	var dataLen pk.VarInt
	if err := dataLen.Decode(data); err != nil {
		return nil, fmt.Errorf("read data array length error: %w", err)
	}
	dataArray := make([]int64, dataLen)
	for i := 0; i < int(dataLen); i++ {
		var v pk.Long
		if err := v.Decode(data); err != nil {
			return nil, fmt.Errorf("read dataArray[%d] error: %w", i, err)
		}
		dataArray[i] = int64(v)
	}

	sec := directSection{bpb: perBits(byte(bpb)), data: dataArray}
	if bpb < 9 {
		return &paletteSection{palette: palettes, directSection: sec}, nil
	} else {
		return &sec, nil
	}
}

type directSection struct {
	bpb  uint32
	data []int64
}

func (d *directSection) GetBlock(x, y, z int) BlockStatus {
	// According to wiki.vg: Data Array is given for each block with increasing x coordinates,
	// within rows of increasing z coordinates, within layers of increasing y coordinates.
	// So offset equals to ( x*16^0 + z*16^1 + y*16^2 )*(bits per block).
	offset := uint32(x+z*16+y*16*16) * d.bpb
	block := uint32(d.data[offset/64])
	block >>= offset % 64
	if offset%64 > 64-d.bpb {
		l := 64 - offset%64
		block |= uint32(d.data[offset/64+1] << l)
	}
	return BlockStatus(block & (1<<d.bpb - 1)) // mask
}

type paletteSection struct {
	palette []uint32
	directSection
}

func (p *paletteSection) GetBlock(x, y, z int) BlockStatus {
	v := p.directSection.GetBlock(x, y, z)
	return BlockStatus(p.palette[v])
}
