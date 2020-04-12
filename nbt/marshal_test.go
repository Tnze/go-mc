package nbt

import (
	"bytes"
	"math"
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
		0x00, 0x00, 0x00, 0x00, // 0
		0xff, 0xff, 0xff, 0xf6, // -10
		0x00, 0x00, 0x00, 0x03, // 3
		TagEnd,
	}
	buf.Reset()
	if err := Marshal(&buf, v2); err != nil {
		t.Error(err)
	} else if !bytes.Equal(buf.Bytes(), out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", buf.Bytes(), out)
	}
}

func TestMarshal_FloatArray(t *testing.T) {
	// Test marshal pure Int array
	v := []float32{0.3, -100, float32(math.NaN())}
	out := []byte{TagList, 0x00, 0x00, TagFloat, 0, 0, 0, 3,
		0x3e, 0x99, 0x99, 0x9a, // 0.3
		0xc2, 0xc8, 0x00, 0x00, // -100
		0x7f, 0xc0, 0x00, 0x00, // NaN
	}
	var buf bytes.Buffer
	if err := Marshal(&buf, v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(buf.Bytes(), out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", buf.Bytes(), out)
	}
}
