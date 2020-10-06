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

func TestMarshal_String(t *testing.T) {
	v := "Test"
	out := []byte{TagString, 0x00, 0x00, 0, 4,
		'T', 'e', 's', 't'}

	var buf bytes.Buffer
	if err := Marshal(&buf, v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(buf.Bytes(), out) {
		t.Errorf("output binary not right: got % 02x, want % 02x ", buf.Bytes(), out)
	}
}

func TestMarshal_InterfaceArray(t *testing.T) {
	type Struct1 struct {
		Val int32
	}

	type Struct2 struct {
		Val float32
	}

	tests := []struct {
		name string
		args []interface{}
		want []byte
	}{
		{
			name: "Two element interface array",
			args: []interface{}{Struct1{3}, Struct2{0.3}},
			want: []byte{
				TagList, 0x00, 0x00 /*no name*/, TagCompound, 0, 0, 0, 2,
				// 1st element
				TagInt, 0x00, 0x03, 'V', 'a', 'l', 0x00, 0x00, 0x00, 0x03, // 3
				TagEnd,
				// 2nd element
				TagFloat, 0x00, 0x03, 'V', 'a', 'l', 0x3e, 0x99, 0x99, 0x9a, // 0.3
				TagEnd,
			},
		},
	}

	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Marshal(w, tt.args)
			if err != nil {
				t.Error(err)
			} else if !bytes.Equal(w.Bytes(), tt.want) {
				t.Errorf("Marshal([]interface{}) got = % 02x, want % 02x", w.Bytes(), tt.want)
				return
			}
		})
	}
}

func TestMarshal_StructArray(t *testing.T) {
	type Struct1 struct {
		Val int32
	}

	type Struct2 struct {
		T   int32
		Ele Struct1
	}

	type StructCont struct {
		V []Struct2
	}

	tests := []struct {
		name string
		args StructCont
		want []byte
	}{
		{
			name: "One element struct array",
			args: StructCont{[]Struct2{{3, Struct1{3}}, {-10, Struct1{-10}}}},
			want: []byte{
				TagCompound, 0x00, 0x00,
				TagList, 0x00, 0x01, 'V', TagCompound, 0, 0, 0, 2,
				// Struct2
				TagInt, 0x00, 0x01, 'T', 0x00, 0x00, 0x00, 0x03,
				TagCompound, 0x00, 0x03, 'E', 'l', 'e',
				TagInt, 0x00, 0x03, 'V', 'a', 'l', 0x00, 0x00, 0x00, 0x03, // 3
				TagEnd,
				TagEnd,
				// 2nd element
				TagInt, 0x00, 0x01, 'T', 0xff, 0xff, 0xff, 0xf6,
				TagCompound, 0x00, 0x03, 'E', 'l', 'e',
				TagInt, 0x00, 0x03, 'V', 'a', 'l', 0xff, 0xff, 0xff, 0xf6, // -10
				TagEnd,
				TagEnd,
				TagEnd,
			},
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			w := &bytes.Buffer{}
			err := Marshal(w, tt.args)
			if err != nil {
				t.Error(err)
			} else if !bytes.Equal(w.Bytes(), tt.want) {
				t.Errorf("Marshal([]struct{}) got = % 02x, want % 02x", w.Bytes(), tt.want)
				return
			}
		})
	}
}
