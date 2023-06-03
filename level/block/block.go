package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"github.com/Tnze/go-mc/bot/maths"
	"github.com/Tnze/go-mc/data/shapes"
	"github.com/Tnze/go-mc/level/block/states"
	"github.com/Tnze/go-mc/nbt"
	"math/bits"
)

type Block struct {
	*BlockProperty
	*StateHolder
	Name   string
	feeder StateFeeder[Block]
}

func NewBlock(name string, property *BlockProperty) *Block {
	return &Block{
		BlockProperty: property,
		StateHolder:   NewStateHolder(make(map[states.Property[any]]uint32)),
		Name:          name,
	}
}

func (b *Block) registerState(property states.Property[any], value any) *Block {
	// This is still not complete, some properties are not registered, please use ToStateID and StateList
	return b.SetValue(property, parseState(value))
}

func (b *Block) feedStates(feeder *StateFeeder[Block]) *Block {
	properties := make([]states.Property[any], 0, len(b.Properties))
	for k := range b.Properties {
		properties = append(properties, k)
	}
	feeder.FeedState(b.StateHolder, properties)
	b.feeder = *feeder
	//BitsPerBlock = bits.Len(uint(feeder.max))
	return b
}

func (b *Block) GetValue(property states.Property[any]) StateID {
	for k, v := range b.Properties {
		if k == property {
			fmt.Println("get value", property, v)
			stateid := b.StateHolder.GetValue(property, v)
			fmt.Println("stateid", stateid)
			return stateid
		}
	}
	return 0
}

func (b *Block) SetValue(property states.Property[any], value uint32) *Block {
	if !property.CanUpdate(value) {
		fmt.Println(fmt.Errorf("invalid value %v for property %v", value, property))
	}

	b.Properties[property] = value
	return b
}

func parseState(v any) uint32 {
	switch v.(type) {
	case bool:
		if v.(bool) {
			return 1
		} else {
			return 0
		}
	case int:
		return uint32(v.(int))
	case states.PropertiesEnum:
		return uint32(v.(states.PropertiesEnum).Value())
	default:
		panic(fmt.Errorf("invalid type %T for state value", v))
	}
}

func (b *Block) StateID() StateID {
	return b.GetValue(nil)
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

// This legacy code is still compatible with the current implementation of block states
// Because it's not complete and NOT RELIABLE, please use ToStateID and StateList
var (
	ToStateID map[*Block]StateID
	StateList []*Block
)

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
	ToStateID = make(map[*Block]StateID, len(states))
	StateList = make([]*Block, 0, len(states))
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
