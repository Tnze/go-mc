package packet

import (
	"bytes"
	"testing"
)

var VarInts = []VarInt{0, 1, 2, 127, 128, 255, 2147483647, -1, -2147483648}

var PackedVarInts = [][]byte{
	{0x00},
	{0x01},
	{0x02},
	{0x7f},
	{0x80, 0x01},
	{0xff, 0x01},
	{0xff, 0xff, 0xff, 0xff, 0x07},
	{0xff, 0xff, 0xff, 0xff, 0x0f},
	{0x80, 0x80, 0x80, 0x80, 0x08},
}

func TestPackVarInt(t *testing.T) {
	for i, v := range VarInts {
		p := v.Encode()
		if !bytes.Equal(p, PackedVarInts[i]) {
			t.Errorf("pack int %d should be \"% x\", get \"% x\"", v, PackedVarInts[i], p)
		}
	}
}
func TestUnpackVarInt(t *testing.T) {
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

var VarLongs = []VarLong{0, 1, 2, 127, 128, 255, 2147483647, 9223372036854775807, -1, -2147483648, -9223372036854775808}

var PackedVarLongs = [][]byte{
	{0x00},
	{0x01},
	{0x02},
	{0x7f},
	{0x80, 0x01},
	{0xff, 0x01},
	{0xff, 0xff, 0xff, 0xff, 0x07},
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x7f},
	{0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0xff, 0x01},
	{0x80, 0x80, 0x80, 0x80, 0xf8, 0xff, 0xff, 0xff, 0xff, 0x01},
	{0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x80, 0x01},
}

func TestPackVarLong(t *testing.T) {
	for i, v := range VarLongs {
		p := v.Encode()
		if !bytes.Equal(p, PackedVarLongs[i]) {
			t.Errorf("pack long %d should be \"% x\", get \"% x\"", v, PackedVarLongs[i], p)
		}
	}
}
func TestUnpackVarLong(t *testing.T) {
	for i, v := range PackedVarLongs {
		var vi VarLong
		if err := vi.Decode(bytes.NewReader(v)); err != nil {
			t.Errorf("unpack \"% x\" error: %v", v, err)
		}
		if vi != VarLongs[i] {
			t.Errorf("unpack \"% x\" should be %d, get %d", v, VarLongs[i], vi)
		}
	}
}
