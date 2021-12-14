package save

import (
	pk "github.com/Tnze/go-mc/net/packet"
	"math/bits"
	"reflect"
	"testing"
)

var data = []uint64{0x0020863148418841, 0x01018A7260F68C87}
var want = []int{1, 2, 2, 3, 4, 4, 5, 6, 6, 4, 8, 0, 7, 4, 3, 13, 15, 16, 9, 14, 10, 12, 0, 2}

func TestBitStorage_Get(t *testing.T) {
	bs := NewBitStorage(5, 24, data)
	for i := 0; i < 24; i++ {
		if got := bs.Get(i); got != want[i] {
			t.Errorf("Decode error, got: %d but expected: %d", got, want[i])
		}
	}
}

func TestBitStorage_Set(t *testing.T) {
	bs := NewBitStorage(5, 24, nil)
	for i := 0; i < 24; i++ {
		bs.Set(i, want[i])
	}
	if !reflect.DeepEqual(bs.data, data) {
		t.Errorf("Encode error, got %v but expected: %v", bs.data, data)
	}
}

func ExampleNewBitStorage_heightmaps() {
	// Create a BitStorage
	bs := NewBitStorage(bits.Len(256), 16*16, nil)
	// Fill your data
	for i := 0; i < 16; i++ {
		for j := 0; j < 16; j++ {
			bs.Set(i*16+j, 0)
		}
	}
	// Encode as NBT, and this is ready for packet.Marshal
	type HeightMaps struct {
		MotionBlocking []uint64 `nbt:"MOTION_BLOCKING"`
	}
	_ = pk.NBT(HeightMaps{bs.Longs()})
}
