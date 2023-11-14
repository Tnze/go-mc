package nbt

import (
	"bytes"
	"io"
	"strings"
	"testing"
)

type testCase struct {
	snbt    StringifiedMessage
	tagType byte
	data    []byte
}

var testCases = []testCase{
	{`10b`, TagByte, []byte{1, 0, 0, 10}},
	{`12S`, TagShort, []byte{2, 0, 0, 0, 12}},
	{`0`, TagInt, []byte{3, 0, 0, 0, 0, 0, 0}},
	{`12L`, TagLong, []byte{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12}},

	{`""`, TagString, []byte{8, 0, 0, 0, 0}},
	{`'""' `, TagString, []byte{8, 0, 0, 0, 2, '"', '"'}},
	{`"ab\"c\""`, TagString, []byte{8, 0, 0, 0, 5, 'a', 'b', '"', 'c', '"'}},
	{` "1\\23"`, TagString, []byte{8, 0, 0, 0, 4, '1', '\\', '2', '3'}},

	{`{}`, TagCompound, []byte{10, 0, 0, 0}},
	{`{a:1b}`, TagCompound, []byte{10, 0, 0, 1, 0, 1, 'a', 1, 0}},
	{`{ a : 1b }`, TagCompound, []byte{10, 0, 0, 1, 0, 1, 'a', 1, 0}},
	{`{b:1,2:c}`, TagCompound, []byte{10, 0, 0, 3, 0, 1, 'b', 0, 0, 0, 1, 8, 0, 1, '2', 0, 1, 'c', 0}},
	{`{c:{d:{}}}`, TagCompound, []byte{10, 0, 0, 10, 0, 1, 'c', 10, 0, 1, 'd', 0, 0, 0}},
	{`{h:{},"i":{}}`, TagCompound, []byte{10, 0, 0, 10, 0, 1, 'h', 0, 10, 0, 1, 'i', 0, 0}},

	{`[]`, TagList, []byte{9, 0, 0, 0, 0, 0, 0, 0}},
	{`[1b,2b,3b]`, TagList, []byte{9, 0, 0, 1, 0, 0, 0, 3, 1, 2, 3}},
	{`[ 1b , 2b , 3b ]`, TagList, []byte{9, 0, 0, 1, 0, 0, 0, 3, 1, 2, 3}},
	{`[a,"b",'c']`, TagList, []byte{9, 0, 0, 8, 0, 0, 0, 3, 0, 1, 'a', 0, 1, 'b', 0, 1, 'c'}},
	{`[{},{a:1b},{}]`, TagList, []byte{9, 0, 0, 10, 0, 0, 0, 3, 0, 1, 0, 1, 'a', 1, 0, 0}},
	{`[ {  } , { a  : 1b  } , { } ] `, TagList, []byte{9, 0, 0, 10, 0, 0, 0, 3, 0, 1, 0, 1, 'a', 1, 0, 0}},
	{`[[],[]]`, TagList, []byte{9, 0, 0, 9, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},

	{`[B; ]`, TagByteArray, []byte{7, 0, 0, 0, 0, 0, 0}},
	{`[B; 1b ,2B,3B]`, TagByteArray, []byte{7, 0, 0, 0, 0, 0, 3, 1, 2, 3}},
	{`[I;]`, TagIntArray, []byte{11, 0, 0, 0, 0, 0, 0}},
	{`[I; -1, 0, 1, 2 ,3]`, TagIntArray, []byte{11, 0, 0, 0, 0, 0, 5, 255, 255, 255, 255, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3}},
	{`[L;]`, TagLongArray, []byte{12, 0, 0, 0, 0, 0, 0}},
	{`[ L; 1L,2L,3L]`, TagLongArray, []byte{12, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3}},
	{`{a:[B;]}`, TagCompound, []byte{10, 0, 0, 7, 0, 1, 'a', 0, 0, 0, 0, 0}},
	{`{a:[B;1b,2B,3B]}`, TagCompound, []byte{10, 0, 0, 7, 0, 1, 'a', 0, 0, 0, 3, 1, 2, 3, 0}},

	{`{d:[]}`, TagCompound, []byte{10, 0, 0, 9, 0, 1, 'd', 0, 0, 0, 0, 0, 0}},
	{`{e:[]}`, TagCompound, []byte{10, 0, 0, 9, 0, 1, 'e', 0, 0, 0, 0, 0, 0}},
	{`{f:[], g:[]}`, TagCompound, []byte{10, 0, 0, 9, 0, 1, 'f', 0, 0, 0, 0, 0, 9, 0, 1, 'g', 0, 0, 0, 0, 0, 0}},
	{`{a:[{b:3B}]}`, TagCompound, []byte{10, 0, 0, 9, 0, 1, 'a', 10, 0, 0, 0, 1, 1, 0, 1, 'b', 3, 0, 0}},

	// issue#121
	{`{a:[b],c:0B}`, TagCompound, []byte{10, 0, 0, 9, 0, 1, 'a', 8, 0, 0, 0, 1, 0, 1, 'b', 1, 0, 1, 'c', 0, 0}},
}

func TestStringifiedMessage_TagType(t *testing.T) {
	for i := range testCases {
		got := testCases[i].snbt.TagType()
		if want := testCases[i].tagType; got != want {
			t.Errorf("TagType assert for %s error: want % 02X, got % 02X", testCases[i].snbt, want, got)
		}
	}
}

func TestEncoder_writeSNBT(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)
	for i := range testCases {
		buf.Reset()
		if err := e.Encode(testCases[i].snbt, ""); err != nil {
			t.Errorf("Convert SNBT %q error: %v", testCases[i].snbt, err)
			continue
		}
		want := testCases[i].data
		got := buf.Bytes()
		if !bytes.Equal(want, got) {
			t.Errorf("Convert SNBT %q wrong:\nwant: % 02X\ngot:  % 02X", testCases[i].snbt, want, got)
		}
	}
}

func TestEncoder_ErrorInput(t *testing.T) {
	for _, s := range []string{
		"][",
		"[I; 1, b, 3]",
		"[I; 1, -2a, 3]",
		"[I; 1, -2/I, 3]",
	} {
		err := NewEncoder(io.Discard).Encode(StringifiedMessage(s), "")
		if err == nil {
			t.Errorf("String %q is not valid SNBT, expeted to got a error", s)
		}
	}
}

func TestEncoder_WriteSNBT_bigTest(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)

	err := e.Encode(StringifiedMessage(bigTestSNBT), "")
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkEncoder_WriteSNBT_bigTest(b *testing.B) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)
	for i := 0; i < b.N; i++ {
		err := e.Encode(StringifiedMessage(bigTestSNBT), "")
		if err != nil {
			b.Fatal(err)
		}
		buf.Reset()
	}
}

func Test_WriteSNBT_nestingList(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)

	// Our maximum supported nesting depth is 10000.
	// The nesting depth of 10001 is 10000
	s := strings.Repeat("[", 10001) + strings.Repeat("]", 10001)
	err := e.Encode(StringifiedMessage(s), "")
	if err != nil {
		t.Error(err)
	}

	// Following code should return error instant of panic.
	buf.Reset()
	s = strings.Repeat("[", 10002) + strings.Repeat("]", 10002)
	err = e.Encode(StringifiedMessage(s), "")
	if err == nil {
		t.Error("Exceeded the maximum depth of support, but no error was reported")
	}
	// Panic test
	buf.Reset()
	s = strings.Repeat("[", 20000) + strings.Repeat("]", 20000)
	err = e.Encode(StringifiedMessage(s), "")
	if err == nil {
		t.Error("Exceeded the maximum depth of support, but no error was reported")
	}
}

func TestStringifiedNBT_TagType(t *testing.T) {
	for _, v := range []struct {
		snbt string
		Type byte
	}{
		{`123B`, TagByte},
		{`123`, TagInt},
		{`[]`, TagList},
		{`[{}, {}]`, TagList},
		{`[B;]`, TagByteArray},
		{`[I;]`, TagIntArray},
		{`[L;]`, TagLongArray},
		{`{abc:123B}`, TagCompound},
	} {
		if T := StringifiedMessage(v.snbt).TagType(); T != v.Type {
			t.Errorf("Parse SNBT TagType error: %s is %d, not %d", v.snbt, v.Type, T)
		}
	}
}

func TestStringifiedMessage_Encode(t *testing.T) {
	var buff bytes.Buffer
	for _, v := range []struct {
		snbt string
		data []byte
	}{
		{`123B`, []byte{123}},
		{`[B; 1B, 2B, 3B]`, []byte{0, 0, 0, 3, 1, 2, 3}},
		{`[{},{}]`, []byte{TagCompound, 0, 0, 0, 2, 0, 0}},
	} {
		if err := StringifiedMessage(v.snbt).MarshalNBT(&buff); err != nil {
			t.Errorf("Encode SNBT error: %v", err)
		}
		if !bytes.Equal(buff.Bytes(), v.data) {
			t.Errorf("Encode SNBT error: %q should encoded to %d, not %d", v.snbt, v.data, buff.Bytes())
		}
		buff.Reset()
	}
}
