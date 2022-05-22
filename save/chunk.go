package save

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"io"

	"github.com/Tnze/go-mc/nbt"
)

// Chunk is 16* chunk
type Chunk struct {
	DataVersion   int32
	XPos          int32          `nbt:"xPos"`
	YPos          int32          `nbt:"yPos"`
	ZPos          int32          `nbt:"zPos"`
	BlockEntities nbt.RawMessage `nbt:"block_entities"`
	Structures    nbt.RawMessage `nbt:"structures"`
	Heightmaps    struct {
		MotionBlocking         []uint64 `nbt:"MOTION_BLOCKING"`
		MotionBlockingNoLeaves []uint64 `nbt:"MOTION_BLOCKING_NO_LEAVES"`
		OceanFloor             []uint64 `nbt:"OCEAN_FLOOR"`
		WorldSurface           []uint64 `nbt:"WORLD_SURFACE"`
	}
	Sections []Section `nbt:"sections"`

	BlockTicks     nbt.RawMessage `nbt:"block_ticks"`
	FluidTicks     nbt.RawMessage `nbt:"fluid_ticks"`
	PostProcessing nbt.RawMessage
	InhabitedTime  int64
	IsLightOn      byte `nbt:"isLightOn"`
	LastUpdate     int64
	Status         string
}

type Section struct {
	Y           int8
	BlockStates struct {
		Palette []BlockState `nbt:"palette"`
		Data    []uint64     `nbt:"data"`
	} `nbt:"block_states"`
	Biomes struct {
		Palette []string `nbt:"palette"`
		Data    []uint64 `nbt:"data"`
	} `nbt:"biomes"`
	SkyLight   []byte
	BlockLight []byte
}

type BlockState struct {
	Name       string
	Properties nbt.RawMessage
}

// Load read column data from []byte
func (c *Chunk) Load(data []byte) (err error) {
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

func (c *Chunk) Data(compressingType byte) ([]byte, error) {
	var buff bytes.Buffer

	buff.WriteByte(compressingType)
	var w io.Writer
	switch compressingType {
	default:
		return nil, errors.New("unknown compression")
	case 1:
		w = gzip.NewWriter(&buff)
	case 2:
		w = zlib.NewWriter(&buff)
	}
	err := nbt.NewEncoder(w).Encode(c, "")
	return buff.Bytes(), err
}
