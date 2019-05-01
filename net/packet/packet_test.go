package packet

import (
	"bytes"
	"testing"
)

var VarInts = []VarInt{0, 1, 2, 127, 128, 255, 2147483647, -1, -2147483648}

var PackedVarInts = [][]byte{
	[]byte{0x00},
	[]byte{0x01},
	[]byte{0x02},
	[]byte{0x7f},
	[]byte{0x80, 0x01},
	[]byte{0xff, 0x01},
	[]byte{0xff, 0xff, 0xff, 0xff, 0x07},
	[]byte{0xff, 0xff, 0xff, 0xff, 0x0f},
	[]byte{0x80, 0x80, 0x80, 0x80, 0x08},
}

func TestPackInt(t *testing.T) {
	for i, v := range VarInts {
		p := v.Encode()
		if !bytes.Equal(p, PackedVarInts[i]) {
			t.Errorf("pack int %d should be \"% x\", get \"% x\"", v, PackedVarInts[i], p)
		}
	}
}
func TestUnpackInt(t *testing.T) {
	for i, v := range PackedVarInts {
		var vi VarInt
		if err := vi.Decode(bytes.NewReader(v)); err != nil {
			t.Errorf("unpack \"% x\" error: %v", v, err)
		}
		if vi != VarInts[i] {
			t.Errorf("unpack \"% x\" should be %d, get %d", v, VarInts[i], vi)
		}
	}
}

func TestPositionPack(t *testing.T) {
	// x (-33554432 to 33554431), y (-2048 to 2047), z (-33554432 to 33554431)

	for x := -33554432; x < 33554432; x += 55443 {
		for y := -2048; y < 2048; y += 48 {
			for z := -33554432; z < 33554432; z += 55443 {
				var (
					pos1 Position
					pos2 = Position{x, y, z}
				)
				if err := pos1.Decode(bytes.NewReader(pos2.Encode())); err != nil {
					t.Errorf("Position decode fail: %v", err)
				}

				if pos1 != pos2 {
					t.Errorf("cannot pack %v", pos2)
				}
			}
		}
	}
}
