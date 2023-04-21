package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"github.com/Tnze/go-mc/data/shapes"
	"math/bits"

	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/nbt"
)

type Block struct {
	*BlockProperty
	Name string
}

func NewBlock(name string, property *BlockProperty) *Block {
	return &Block{
		BlockProperty: property,
		Name:          name,
	}
}

func (b *Block) StateID() StateID {
	return ToStateID[*b]
}

func (b *Block) Is(other *Block) bool {
	return b.Name == other.Name && b.StateID() == other.StateID()
}

func (b *Block) IsAir() bool {
	return b.BlockProperty.IsAir
}

func (b *Block) IsLiquid() bool {
	return b.Is(Water) || b.Is(Lava)
}

func (b *Block) GetCollisionBox() maths.AxisAlignedBB[float64] {
	aabb := shapes.GetShape(b.Name, int(b.StateID()))
	return maths.AxisAlignedBB[float64]{
		MinX: aabb[0], MinY: aabb[1], MinZ: aabb[2],
		MaxX: aabb[3], MaxY: aabb[4], MaxZ: aabb[5],
	}
}

// This file stores all possible block states into a TAG_List with gzip compressed.
//
//go:embed block_states.nbt
var blockStates []byte

var ToStateID map[*Block]StateID
var StateList []Block

// BitsPerBlock indicates how many bits are needed to represent all possible
// block states. This value is used to determine the size of the global palette.
var BitsPerBlock int

type StateID int
type State struct {
	Name       string
	Properties nbt.RawMessage
}

func init() {
	var states []State
	// decompress
	z, err := gzip.NewReader(bytes.NewReader(blockStates))
	if err != nil {
		panic(err)
	}
	// decode all states
	if _, err = nbt.NewDecoder(z).Decode(&states); err != nil {
		panic(err)
	}
	ToStateID = make(map[Block]StateID, len(states))
	StateList = make([]Block, 0, len(states))
	for _, state := range states {
		block := FromID[state.Name]
		if state.Properties.Type != nbt.TagEnd {
			err := state.Properties.Unmarshal(&block)
			if err != nil {
				panic(err)
			}
		}
		ToStateID[block] = StateID(len(StateList))
		StateList = append(StateList, block)
	}
	BitsPerBlock = bits.Len(uint(len(StateList)))
}
