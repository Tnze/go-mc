package world

import (
	"github.com/Tnze/go-mc/nbt"
	// "fmt"
	pk "github.com/Tnze/go-mc/net/packet"
	// "io"
)

func UnpackChunkDataPacket(p pk.Packet, hasSkyLight bool) (c *Chunk, x, z int, err error) {
	var (
		X, Z           pk.Int
		FullChunk      pk.Boolean
		PrimaryBitMask pk.VarInt
		Heightmaps     struct{}
		Data           chunkData
		BlockEntities  blockEntities
	)

	p.Scan(&X, &Z, &FullChunk, &PrimaryBitMask, pk.NBT{V: &Heightmaps}, &Data, &BlockEntities)

	//解析区块数据
	cc, err := readChunkColumn(bool(FullChunk), int32(PrimaryBitMask), []byte(Data), hasSkyLight)
	if err != nil {
		panic(err)
	}
	return cc, int(X), int(Z), err
}

type chunkData []byte
type blockEntities []blockEntitie
type blockEntitie struct {
}

func (c *chunkData) Decode(r pk.DecodeReader) error {
	var Size pk.VarInt
	if err := Size.Decode(r); err != nil {
		return err
	}
	*c = make([]byte, Size)
	if _, err := r.Read(*c); err != nil {
		return err
	}
	return nil
}

func (b *blockEntities) Decode(r pk.DecodeReader) error {
	var NumberofBlockEntities pk.VarInt
	if err := NumberofBlockEntities.Decode(r); err != nil {
		return err
	}
	*b = make(blockEntities, NumberofBlockEntities)
	decoder := nbt.NewDecoder(r)
	for i := 0; i < int(NumberofBlockEntities); i++ {
		if err := decoder.Decode(&(*b)[i]); err != nil {
			return err
		}
	}
	return nil
}

func readChunkColumn(isFull bool, mask int32, data []byte, hasSkyLight bool) (*Chunk, error) {
	var c Chunk
	// for sectionY := 0; sectionY < 16; sectionY++ {
	// 	if (mask & (1 << uint(sectionY))) != 0 { // Is the given bit set in the mask?
	// 		BitsPerBlock, err := data.ReadByte()
	// 		if err != nil {
	// 			return nil, fmt.Errorf("read BitsPerBlock fail: %v", err)
	// 		}
	// 		//读调色板
	// 		var palette []uint
	// 		if BitsPerBlock < 9 {
	// 			length, err := pk.UnpackVarInt(data)
	// 			if err != nil {
	// 				return nil, fmt.Errorf("read palette (id len) fail: %v", err)
	// 			}
	// 			palette = make([]uint, length)

	// 			for id := uint(0); id < uint(length); id++ {
	// 				stateID, err := pk.UnpackVarInt(data)
	// 				if err != nil {
	// 					return nil, fmt.Errorf("read palette (id) fail: %v", err)
	// 				}

	// 				palette[id] = uint(stateID)
	// 			}
	// 		}

	// 		//Section数据
	// 		DataArrayLength, err := pk.UnpackVarInt(data)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("read DataArrayLength fail: %v", err)
	// 		}

	// 		DataArray := make([]int64, DataArrayLength)
	// 		for i := 0; i < int(DataArrayLength); i++ {
	// 			DataArray[i], err = pk.UnpackInt64(data)
	// 			if err != nil {
	// 				return nil, fmt.Errorf("read DataArray fail: %v", err)
	// 			}
	// 		}
	// 		//用数据填充区块
	// 		fillSection(&c.sections[sectionY], perBits(BitsPerBlock), DataArray, palette)

	// 		//throw BlockLight data
	// 		_, err = pk.ReadNBytes(data, 2048)
	// 		if err != nil {
	// 			return nil, fmt.Errorf("read BlockLight fail: %v", err)
	// 		}

	// 		if hasSkyLight {
	// 			//throw SkyLight data
	// 			_, err = pk.ReadNBytes(data, 2048)
	// 			if err != nil {
	// 				return nil, fmt.Errorf("read SkyLight fail: %v", err)
	// 			}
	// 		}
	// 	}
	// }
	// if isFull { //need recive Biomes datas
	// 	_, err := pk.ReadNBytes(data, 256*4)
	// 	if err != nil {
	// 		return nil, fmt.Errorf("read Biomes fail: %v", err)
	// 	}
	// }

	// fmt.Println(c)
	return &c, nil
}

// const defaultBitsPerBlock = 14

// func perBits(BitsPerBlock byte) uint {
// 	switch {
// 	case BitsPerBlock <= 4:
// 		return 4
// 	case BitsPerBlock < 9:
// 		return uint(BitsPerBlock)
// 	default:
// 		return defaultBitsPerBlock
// 	}
// }

// func fillSection(s *Section, bpb uint, DataArray []int64, palette []uint) {
// 	mask := uint(1<<bpb - 1)
// 	for n := 0; n < 16*16*16; n++ {
// 		offset := uint(n * int(bpb))
// 		data := uint(DataArray[offset/64])
// 		data >>= offset % 64
// 		if offset%64 > 64-bpb {
// 			l := bpb + offset%64 - 64
// 			data &= uint(DataArray[offset/64+1] << l)
// 		}
// 		data &= mask

// 		if bpb < 9 {
// 			s.blocks[n%16][n/(16*16)][n%(16*16)/16].id = palette[data]
// 		} else {
// 			s.blocks[n%16][n/(16*16)][n%(16*16)/16].id = data
// 		}
// 	}
// }
