package save

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"github.com/Tnze/go-mc/nbt"
	"io"
)

// Column is 16* chunk
type Column struct {
	DataVersion int
	Level       struct {
		Heightmaps map[string][]int64
		Structures struct {
			References map[string][]int64
			Starts     map[string]struct {
				ID string `nbt:"id"`
			}
		}
		// Entities
		// LiquidTicks
		// PostProcessing
		Sections []Chunk
		// TileEntities
		// TileTicks
		InhabitedTime int64
		IsLightOn     byte `nbt:"isLightOn"`
		LastUpdate    int64
		Status        string
		PosX          int32 `nbt:"xPos"`
		PosZ          int32 `nbt:"zPos"`
		Biomes        []int32
	}
}

type Chunk struct {
	Palette     []Block
	Y           byte
	BlockLight  []byte
	BlockStates []int64
	SkyLight    []byte
}

type Block struct {
	Name       string
	Properties map[string]interface{}
}

// Load read column data from []byte
func (c *Column) Load(data []byte) (err error) {
	var r io.Reader = bytes.NewReader(data[1:])

	switch data[0] {
	default:
		err = errors.New("unknown compression")
	case 1:
		r, err = gzip.NewReader(r)
	case 2:
		r, err = zlib.NewReader(r)
	}

	if err != nil {
		return err
	}

	_, err = nbt.NewDecoder(r).Decode(c)
	return
}
