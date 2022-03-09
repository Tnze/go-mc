package block

import (
	"bytes"
	"compress/zlib"
	_ "embed"
	"fmt"

	"github.com/Tnze/go-mc/nbt"
)

type Block interface {
	ID() string
}

// This file stores all possible block states into a TAG_List with zlib compressed.
//go:embed block_states.nbt
var blockStates []byte

var toStateID = make(map[Block]int)
var fromStateID []Block

func init() {
	regState := func(s Block) {
		if _, ok := toStateID[s]; ok {
			panic(fmt.Errorf("state %#v already exist", s))
		}
		toStateID[s] = len(fromStateID)
		fromStateID = append(fromStateID, s)
	}
	var states []struct {
		Name       string
		Properties nbt.RawMessage
	}
	// decompress
	z, err := zlib.NewReader(bytes.NewReader(blockStates))
	if err != nil {
		panic(err)
	}
	// decode all states
	if _, err = nbt.NewDecoder(z).Decode(&states); err != nil {
		panic(err)
	}
	for _, state := range states {
		block := fromID[state.Name]
		if state.Properties.Type != nbt.TagEnd {
			err := state.Properties.Unmarshal(&block)
			if err != nil {
				panic(err)
			}
		}
		regState(block)
	}
}

func FromStateID(stateID int) Block {
	return fromStateID[stateID]
}

func ToStateID(b Block) int {
	return toStateID[b]
}
