package world

import (
	"bytes"
	"fmt"
	// "io"
	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

//DecodeChunkColumn decode the chunk data structure
func DecodeChunkColumn(mask int32, data []byte) (*Chunk, error) {
	var c Chunk
	r := bytes.NewReader(data)
	for sectionY := 0; sectionY < 16; sectionY++ {
		if (mask & (1 << uint(sectionY))) == 0 { // Is the given bit set in the mask?
			continue
		}
		var (
			BlockCount   pk.Short
			BitsPerBlock pk.Byte
		)
		if err := BlockCount.Decode(r); err != nil {
			return nil, err
		}
		if err := BitsPerBlock.Decode(r); err != nil {
			return nil, err
		}
		//读调色板
		var palette []uint
		if BitsPerBlock < 9 {
			var length pk.VarInt
			if err := length.Decode(r); err != nil {
				return nil, fmt.Errorf("read palette (id len) fail: %v", err)
			}
			palette = make([]uint, length)

			for id := uint(0); id < uint(length); id++ {
				var stateID pk.VarInt
				if err := stateID.Decode(r); err != nil {
					return nil, fmt.Errorf("read palette (id) fail: %v", err)
				}

				palette[id] = uint(stateID)
			}
		}

		//Section数据
		var DataArrayLength pk.VarInt
		if err := DataArrayLength.Decode(r); err != nil {
			return nil, fmt.Errorf("read DataArrayLength fail: %v", err)
		}

		DataArray := make([]int64, DataArrayLength)
		for i := 0; i < int(DataArrayLength); i++ {
			if err := (*pk.Long)(&DataArray[i]).Decode(r); err != nil {
				return nil, fmt.Errorf("read DataArray fail: %v", err)
			}
		}
		//用数据填充区块
		fillSection(&c.sections[sectionY], perBits(byte(BitsPerBlock)), DataArray, palette)
	}

	return &c, nil
}

func perBits(BitsPerBlock byte) uint {
	switch {
	case BitsPerBlock <= 4:
		return 4
	case BitsPerBlock < 9:
		return uint(BitsPerBlock)
	default:
		return uint(data.BitsPerBlock) // DefaultBitsPerBlock
	}
}

func fillSection(s *Section, bpb uint, DataArray []int64, palette []uint) {
	mask := uint(1<<bpb - 1)
	for n := 0; n < 16*16*16; n++ {
		offset := uint(n * int(bpb))
		data := uint(DataArray[offset/64])
		data >>= offset % 64
		if offset%64 > 64-bpb {
			l := bpb + offset%64 - 64
			data &= uint(DataArray[offset/64+1] << l)
		}
		data &= mask

		if bpb < 9 {
			s.blocks[n%16][n/(16*16)][n%(16*16)/16].id = palette[data]
		} else {
			s.blocks[n%16][n/(16*16)][n%(16*16)/16].id = data
		}
	}
}
