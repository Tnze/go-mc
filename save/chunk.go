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
	DataVersion   int32
	XPos          int32          `nbt:"xPos"`
	YPos          int32          `nbt:"yPos"`
	ZPos          int32          `nbt:"zPos"`
	BlockEntities nbt.RawMessage `nbt:"block_entities"`
	Structures    nbt.RawMessage `nbt:"structures"`
	Heightmaps    struct {
		MotionBlocking         []int64 `nbt:"MOTION_BLOCKING"`
		MotionBlockingNoLeaves []int64 `nbt:"MOTION_BLOCKING_NO_LEAVES"`
		OceanFloor             []int64 `nbt:"OCEAN_FLOOR"`
		WorldSurface           []int64 `nbt:"WORLD_SURFACE"`
	}
	Sections []struct {
		Y           byte
		BlockStates struct {
			Palette []BlockState `nbt:"palette"`
			Data    []int64      `nbt:"data"`
		} `nbt:"block_states"`
		Biomes struct {
			Palette []string `nbt:"palette"`
			Data    []int64  `nbt:"data"`
		} `nbt:"biomes"`
		SkyLight   []byte
		BlockLight []byte
	} `nbt:"sections"`
}

type BlockState struct {
	Name       string
	Properties nbt.RawMessage
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
