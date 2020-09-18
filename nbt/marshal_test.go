package nbt

import (
	"bytes"
	"io/ioutil"
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


// This test is for compliance with the "bigtest.dat" described in detail here:
// https://wiki.vg/NBT#bigtest.nbt
func TestMarshal_BigTest(t *testing.T) {
	byteValues := make([]byte, 1000)
	for n := 0; n < 1000; n++ {
		byteValues[n] = byte((n*n*255 + n*7) % 100)
	}

	type NestedCompound struct {
		Name  string  `nbt:"name"`
		Value float32 `nbt:"value"`
	}

	type NestedCompoundCont struct {
		Ham NestedCompound `nbt:"ham"`
		Egg NestedCompound `nbt:"egg"`
	}

	type ListCompound struct {
		Name      string `nbt:"name"`
		CreatedOn int64  `nbt:"created-on"`
	}

	val := struct {
		LongTest           int64              `nbt:"longTest"`
		ShortTest          int16              `nbt:"shortTest"`
		StringTest         string             `nbt:"stringTest"`
		FloatTest          float32            `nbt:"floatTest"`
		IntTest            int32              `nbt:"intTest"`
		NestedCompoundTest NestedCompoundCont `nbt:"nested compound test"`
		ListTestLong       []int64            `nbt:"listTest (long)" nbt_type:"noarray"`
		ListTestCompound   []ListCompound     `nbt:"listTest (compound)"`
		ByteTest           byte               `nbt:"byteTest"`
		ByteArrayTest      []byte             `nbt:"byteArrayTest (the first 1000 values of (n*n*255+n*7)%100, starting with n=0 (0, 62, 34, 16, 8, ...))"`
		DoubleTest         float64            `nbt:"doubleTest"`
	}{
		LongTest:   9223372036854775807,
		ShortTest:  32767,
		StringTest: "HELLO WORLD THIS IS A TEST STRING \xc3\x85\xc3\x84\xc3\x96!",
		FloatTest:  0.49823147058486938,
		IntTest:    2147483647,
		NestedCompoundTest: NestedCompoundCont{
			NestedCompound{"Hampus", 0.75},
			NestedCompound{"Eggbert", 0.5},
		},
		ListTestLong: []int64{11, 12, 13, 14, 15},
		ListTestCompound: []ListCompound{
			{"Compound tag #0", 1264099775885},
			{"Compound tag #1", 1264099775885},
		},
		ByteTest:      127,
		ByteArrayTest: byteValues,
		DoubleTest:    0.49312871321823148,
	}

	var b bytes.Buffer
	err := MarshalCompound(&b, val, "Level")
	if err != nil {
		t.Error(err)
	}

	want, err := ioutil.ReadFile("bigtest.nbt")
	if err != nil {
		t.Error(err)
	}

	err = ioutil.WriteFile("bigtest_got.nbt", b.Bytes(), 0644)
	if err != nil {
		t.Error(err)
	}
	if !bytes.Equal(b.Bytes(), want) {
		t.Errorf("got:\n[% 2x]\nwant:\n[% 2x]", b.Bytes(), want)
	}
}
