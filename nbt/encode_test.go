package nbt

import (
	"bytes"
	"compress/gzip"
	"io"
	"math"
	"reflect"
	"testing"
)

func TestEncoder_Encode_intArray(t *testing.T) {
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

func TestEncoder_Encode_floatArray(t *testing.T) {
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

func TestEncoder_Encode_string(t *testing.T) {
	v := "Test"
	out := []byte{TagString, 0x00, 0x00, 0, 4,
		'T', 'e', 's', 't'}

	if data, err := Marshal(v); err != nil {
		t.Error(err)
	} else if !bytes.Equal(data, out) {
		t.Errorf("output binary not right: got % 02x, want % 02x ", data, out)
	}
}

func TestEncoder_Encode_interfaceArray(t *testing.T) {
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

func TestEncoder_Encode_structArray(t *testing.T) {
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

func TestEncoder_Encode_bigTest(t *testing.T) {
	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(MakeBigTestStruct(), "Level"); err != nil {
		t.Error(err)
	}

	rd, _ := gzip.NewReader(bytes.NewReader(bigTestData[:]))
	want, err := io.ReadAll(rd)
	if err != nil {
		t.Error(err)
	}

	if !bytes.Equal(buf.Bytes(), want) {
		t.Errorf("got:\n[% 2x]\nwant:\n[% 2x]", buf.Bytes(), want)
	}
}

func TestEncoder_Encode_map(t *testing.T) {
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

	if _, err := NewDecoder(bytes.NewReader(b)).Decode(&data); err != nil {
		t.Fatal(err)
	}
	if !reflect.DeepEqual(data.Tnze, v["Tnze"]) {
		t.Fatalf("Marshal map error: got: %q, want %q", data.Tnze, v["Tnze"])
	}
	if !reflect.DeepEqual(data.XXM, v["Xi_Xi_Mi"]) {
		t.Fatalf("Marshal map error: got: %#v, want %#v", data.XXM, v["Xi_Xi_Mi"])
	}
}

func TestRawMessage_Encode(t *testing.T) {
	data := []byte{
		TagCompound, 0, 2, 'a', 'b',
		TagInt, 0, 3, 'K', 'e', 'y', 0, 0, 0, 12,
		TagString, 0, 5, 'V', 'a', 'l', 'u', 'e', 0, 4, 'T', 'n', 'z', 'e',
		TagEnd,
	}
	var container struct {
		Key   int32
		Value RawMessage
	}
	container.Key = 12
	container.Value.Type = TagString
	container.Value.Data = []byte{0, 4, 'T', 'n', 'z', 'e'}

	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(container, "ab"); err != nil {
		t.Fatalf("Encode error: %v", err)
	} else if !bytes.Equal(data, buf.Bytes()) {
		t.Fatalf("Encode error: want %v, get: %v", data, buf.Bytes())
	}
}

func TestEncoder_Encode_interface(t *testing.T) {
	data := map[string]interface{}{
		"Key":   int32(12),
		"Value": "Tnze",
	}
	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(data, "ab"); err != nil {
		t.Fatalf("Encode error: %v", err)
	}

	var container struct {
		Key   int32
		Value string
	}
	if _, err := NewDecoder(&buf).Decode(&container); err != nil {
		t.Fatalf("Decode error: %v", err)
	}

	if container.Key != 12 || container.Value != "Tnze" {
		t.Fatalf("want: (%v, %v), but got (%v, %v)", 12, "Tnze", container.Key, container.Value)
	}
}
