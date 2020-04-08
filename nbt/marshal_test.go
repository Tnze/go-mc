package nbt

import (
	"bytes"
	"testing"
)

func TestMarshal_IntArray(t *testing.T) {
	// Test marshal pure Int array
	v := []int32{0, -10, 3}
	out := []byte{TagIntArray, 0x00, 0x00, 0, 0, 0, 3,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xf6,
		0x00, 0x00, 0x00, 0x03,
	}
	var buf bytes.Buffer
	if err := Marshal(&buf, v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(buf.Bytes(), out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", buf.Bytes(), out)
	}

	// Test marshal in a struct
	v2 := struct {
		Ary []int32 `nbt:"ary"`
	}{[]int32{0, -10, 3}}
	out = []byte{TagCompound, 0x00, 0x00,
		TagIntArray, 0x00, 0x03, 'a', 'r', 'y', 0, 0, 0, 3,
		0x00, 0x00, 0x00, 0x00,
		0xff, 0xff, 0xff, 0xf6,
		0x00, 0x00, 0x00, 0x03,
		TagEnd,
	}
	buf.Reset()
	if err := Marshal(&buf, v2); err != nil {
		t.Error(err)
	} else if !bytes.Equal(buf.Bytes(), out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", buf.Bytes(), out)
	}
}
