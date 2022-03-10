package block

import (
	"bytes"
	"compress/zlib"
	_ "embed"
	"fmt"
	"math/bits"

	"github.com/Tnze/go-mc/nbt"
)

type Block interface {
	ID() string
}

// This file stores all possible block states into a TAG_List with zlib compressed.
//go:embed block_states.nbt
var blockStates []byte

var toStateID map[Block]int
var fromStateID []Block

// BitsPerBlock indicates how many bits are needed to represent all possible
// block states. This value is used to determine the size of the global palette.
var BitsPerBlock int

type State struct {
	Name       string
	Properties nbt.RawMessage
}

func init() {
	var states []State
	// decompress
	z, err := zlib.NewReader(bytes.NewReader(blockStates))
	if err != nil {
		panic(err)
	}
	// decode all states
	if _, err = nbt.NewDecoder(z).Decode(&states); err != nil {
		panic(err)
	}
	toStateID = make(map[Block]int, len(states))
	fromStateID = make([]Block, 0, len(states))
	for _, state := range states {
		block := fromID[state.Name]
		if state.Properties.Type != nbt.TagEnd {
			err := state.Properties.Unmarshal(&block)
			if err != nil {
				panic(err)
			}
		}
		if _, ok := toStateID[block]; ok {
			panic(fmt.Errorf("state %#v already exist", block))
		}
		toStateID[block] = len(fromStateID)
		fromStateID = append(fromStateID, block)
	}
	BitsPerBlock = bits.Len(uint(len(fromStateID)))
}

func FromStateID(stateID int) (b Block, ok bool) {
	if stateID >= 0 && stateID < len(fromStateID) {
		b = fromStateID[stateID]
		ok = true
	}
	return
}

func DefaultBlock(id string) (b Block, ok bool) {
	b, ok = fromID[id]
	return
}

func ToStateID(b Block) (i int, ok bool) {
	i, ok = toStateID[b]
	return
}
