package block

import (
	"os"
	"testing"

	"github.com/Tnze/go-mc/nbt"
)

func TestNewFromStateID(t *testing.T) {
	f, err := os.Open("testdata/block_states.nbt")
	if err != nil {
		t.Fatal(err)
	}
	defer f.Close()

	var states []struct {
		Name       string
		Properties nbt.RawMessage
	}

	_, err = nbt.NewDecoder(f).Decode(&states)
	if err != nil {
		t.Fatal(err)
	}

	for i, v := range states {
		state1 := NewFromStateID(i)
		if id := "minecraft:" + state1.ID(); id != v.Name {
			t.Errorf("StateID [%d] Name not match: %v != %v", i, id, v.Name)
		}
		state2 := state1
		if err := v.Properties.Unmarshal(&state2); err != nil {
			t.Errorf("Decode error: %v", err)
		}
		if state1 != state2 {
			t.Errorf("StateID [%d] Properties not match: %v != %v", i, state1, state2)
		}
	}
}
