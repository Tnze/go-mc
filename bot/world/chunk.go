package world

import (
	"bytes"
	"fmt"

	"github.com/Tnze/go-mc/data/block"
	pk "github.com/Tnze/go-mc/net/packet"
)

const maxPaletteBits = 8

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

func perBits(bpb byte) uint {
	switch {
	case bpb <= 4:
		return 4
	case bpb <= maxPaletteBits:
		return uint(bpb)
	default:
		return uint(block.BitsPerBlock)
	}
}

func readSection(data pk.DecodeReader) (s Section, err error) {
	var nonAirBlockCount pk.Short
	if err := nonAirBlockCount.Decode(data); err != nil {
		return nil, fmt.Errorf("block count: %w", err)
	}
	var bpb pk.UnsignedByte
	if err := bpb.Decode(data); err != nil {
		return nil, fmt.Errorf("bits per block: %w", err)
	}
	// If bpb values greater than or equal to 9, use directSection.
	// Otherwise use paletteSection.
	var palettes []BlockStatus
	var palettesIndex map[BlockStatus]int
	if bpb <= maxPaletteBits {
		// read palettes
		var length pk.VarInt
		if err := length.Decode(data); err != nil {
			return nil, fmt.Errorf("palette length: %w", err)
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

	width := perBits(byte(bpb))
	sec := directSection{
		bitArray{
			width:          width,
			valsPerElement: valsPerBitArrayElement(width),
			data:           dataArray,
		},
	}
	if bpb <= maxPaletteBits {
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
	bitArray
}

func (d *directSection) GetBlock(offset uint) BlockStatus {
	return BlockStatus(d.Get(offset))
}

func (d *directSection) SetBlock(offset uint, s BlockStatus) {
	d.Set(offset, uint(s))
}

func (d *directSection) CanContain(s BlockStatus) bool {
	return s <= (1<<d.width - 1)
}

func (d *directSection) clone(bpb uint) *directSection {
	out := newSectionWithSize(bpb)
	for offset := uint(0); offset < 16*16*16; offset++ {
		out.SetBlock(offset, d.GetBlock(offset))
	}
	return out
}

func newSectionWithSize(bpb uint) *directSection {
	valsPerElement := valsPerBitArrayElement(bpb)
	return &directSection{
		bitArray{
			width:          bpb,
			valsPerElement: valsPerElement,
			data:           make([]uint64, 16*16*16/valsPerElement),
		},
	}
}

type paletteSection struct {
	palette       []BlockStatus
	palettesIndex map[BlockStatus]int
	directSection
}

func (p *paletteSection) GetBlock(offset uint) BlockStatus {
	v := p.directSection.GetBlock(offset)
	return p.palette[v]
}

func (p *paletteSection) SetBlock(offset uint, s BlockStatus) {
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
		p.directSection = *p.directSection.clone(p.width + 1)
	}
	p.directSection.SetBlock(offset, BlockStatus(i))
}
