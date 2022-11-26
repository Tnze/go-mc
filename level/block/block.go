package block

import (
	"bytes"
	"compress/gzip"
	_ "embed"
	"fmt"
	"math/bits"

	"github.com/Tnze/go-mc/nbt"
)

type Block interface {
	ID() string
}

// This file stores all possible block states into a TAG_List with gzip compressed.
//
//go:embed block_states.nbt
var blockStates []byte

var ToStateID map[Block]StateID
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
		if _, ok := ToStateID[block]; ok {
			panic(fmt.Errorf("state %#v already exist", block))
		}
		ToStateID[block] = StateID(len(StateList))
		StateList = append(StateList, block)
	}
	BitsPerBlock = bits.Len(uint(len(StateList)))
}
