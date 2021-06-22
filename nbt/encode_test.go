package nbt

import (
	"bytes"
	"compress/gzip"
	"io"
	"math"
	"reflect"
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
	if data, err := Marshal(v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(data, out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", data, out)
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
	if data, err := Marshal(v2); err != nil {
		t.Error(err)
	} else if !bytes.Equal(data, out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", data, out)
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
	if data, err := Marshal(v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(data, out) {
		t.Errorf("output binary not right: get % 02x, want % 02x ", data, out)
	}
}

func TestMarshal_String(t *testing.T) {
	v := "Test"
	out := []byte{TagString, 0x00, 0x00, 0, 4,
		'T', 'e', 's', 't'}

	if data, err := Marshal(v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(data, out) {
		t.Errorf("output binary not right: got % 02x, want % 02x ", data, out)
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
			data, err := Marshal(tt.args)
			if err != nil {
				t.Error(err)
			} else if !bytes.Equal(data, tt.want) {
				t.Errorf("Marshal([]interface{}) got = % 02x, want % 02x", data, tt.want)
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
			data, err := Marshal(tt.args)
			if err != nil {
				t.Error(err)
			} else if !bytes.Equal(data, tt.want) {
				t.Errorf("Marshal([]struct{}) got = % 02x, want % 02x", data, tt.want)
				return
			}
		})
	}
}

func TestMarshal_bigTest(t *testing.T) {
	data, err := Marshal(MakeBigTestStruct(), "Level")
	if err != nil {
		t.Error(err)
	}

	rd, _ := gzip.NewReader(bytes.NewReader(bigTestData[:]))
	want, err := io.ReadAll(rd)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(data, want) {
		t.Errorf("got:\n[% 2x]\nwant:\n[% 2x]", data, want)
	}
}

func TestMarshal_map(t *testing.T) {
	v := map[string][]int32{
		"Tnze":     {1, 2, 3, 4, 5},
		"Xi_Xi_Mi": {0, 0, 4, 7, 2},
	}

	b, err := Marshal(v)
	if err != nil {
		t.Fatal(err)
	}

	var data struct {
		Tnze []int32
		XXM  []int32 `nbt:"Xi_Xi_Mi"`
	}

	if err := NewDecoder(bytes.NewReader(b)).Decode(&data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(data.Tnze, v["Tnze"]) {
		t.Fatalf("Marshal map error: got: %q, want %q", data.Tnze, v["Tnze"])
	}
	if !reflect.DeepEqual(data.XXM, v["Xi_Xi_Mi"]) {
		t.Fatalf("Marshal map error: got: %#v, want %#v", data.XXM, v["Xi_Xi_Mi"])
	}
}
