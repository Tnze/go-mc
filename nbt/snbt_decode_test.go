package nbt

import (
	"bytes"
	"strings"
	"testing"
)

func TestEncoder_WriteSNBT(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)
	testCases := []struct {
		snbt string
		nbt  []byte
	}{
		{`10b`, []byte{1, 0, 0, 10}},
		{`12S`, []byte{2, 0, 0, 0, 12}},
		{`0`, []byte{3, 0, 0, 0, 0, 0, 0}},
		{`12L`, []byte{4, 0, 0, 0, 0, 0, 0, 0, 0, 0, 12}},

		{`""`, []byte{8, 0, 0, 0, 0}},
		{`'""' `, []byte{8, 0, 0, 0, 2, '"', '"'}},
		{`"ab\"c\""`, []byte{8, 0, 0, 0, 5, 'a', 'b', '"', 'c', '"'}},
		{` "1\\23"`, []byte{8, 0, 0, 0, 4, '1', '\\', '2', '3'}},

		{`{}`, []byte{10, 0, 0, 0}},
		{`{a:1b}`, []byte{10, 0, 0, 1, 0, 1, 'a', 1, 0}},
		{`{ a : 1b }`, []byte{10, 0, 0, 1, 0, 1, 'a', 1, 0}},
		{`{b:1,2:c}`, []byte{10, 0, 0, 3, 0, 1, 'b', 0, 0, 0, 1, 8, 0, 1, '2', 0, 1, 'c', 0}},
		{`{c:{d:{}}}`, []byte{10, 0, 0, 10, 0, 1, 'c', 10, 0, 1, 'd', 0, 0, 0}},
		{`{h:{},"i":{}}`, []byte{10, 0, 0, 10, 0, 1, 'h', 0, 10, 0, 1, 'i', 0, 0}},

		{`[]`, []byte{9, 0, 0, 0, 0, 0, 0, 0}},
		{`[1b,2b,3b]`, []byte{9, 0, 0, 1, 0, 0, 0, 3, 1, 2, 3}},
		{`[ 1b , 2b , 3b ]`, []byte{9, 0, 0, 1, 0, 0, 0, 3, 1, 2, 3}},
		{`[a,"b",'c']`, []byte{9, 0, 0, 8, 0, 0, 0, 3, 0, 1, 'a', 0, 1, 'b', 0, 1, 'c'}},
		{`[{},{a:1b},{}]`, []byte{9, 0, 0, 10, 0, 0, 0, 3, 0, 1, 0, 1, 'a', 1, 0, 0}},
		{`[ {  } , { a  : 1b  } , { } ] `, []byte{9, 0, 0, 10, 0, 0, 0, 3, 0, 1, 0, 1, 'a', 1, 0, 0}},
		{`[[],[]]`, []byte{9, 0, 0, 9, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 0, 0, 0}},

		{`[B; ]`, []byte{7, 0, 0, 0, 0, 0, 0}},
		{`[B; 1b ,2B,3B]`, []byte{7, 0, 0, 0, 0, 0, 3, 1, 2, 3}},
		{`[I;]`, []byte{11, 0, 0, 0, 0, 0, 0}},
		{`[I; 1, 2 ,3]`, []byte{11, 0, 0, 0, 0, 0, 3, 0, 0, 0, 1, 0, 0, 0, 2, 0, 0, 0, 3}},
		{`[L;]`, []byte{12, 0, 0, 0, 0, 0, 0}},
		{`[ L; 1L,2L,3L]`, []byte{12, 0, 0, 0, 0, 0, 3, 0, 0, 0, 0, 0, 0, 0, 1, 0, 0, 0, 0, 0, 0, 0, 2, 0, 0, 0, 0, 0, 0, 0, 3}},

		{`{d:[]}`, []byte{10, 0, 0, 9, 0, 1, 'd', 0, 0, 0, 0, 0, 0}},
		{`{e:[]}`, []byte{10, 0, 0, 9, 0, 1, 'e', 0, 0, 0, 0, 0, 0}},
		{`{f:[], g:[]}`, []byte{10, 0, 0, 9, 0, 1, 'f', 0, 0, 0, 0, 0, 9, 0, 1, 'g', 0, 0, 0, 0, 0, 0}},
	}
	for i := range testCases {
		buf.Reset()
		if err := e.WriteSNBT(testCases[i].snbt); err != nil {
			t.Errorf("Convert SNBT %q error: %v", testCases[i].snbt, err)
			continue
		}
		want := testCases[i].nbt
		got := buf.Bytes()
		if !bytes.Equal(want, got) {
			t.Errorf("Convert SNBT %q wrong:\nwant: % 02X\ngot:  % 02X", testCases[i].snbt, want, got)
		}
	}
}

func TestEncoder_WriteSNBT_bigTest(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)

	err := e.WriteSNBT(bigTestSNBT)
	if err != nil {
		t.Error(err)
	}
}

func BenchmarkEncoder_WriteSNBT_bigTest(b *testing.B) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)
	for i := 0; i < b.N; i++ {
		err := e.WriteSNBT(bigTestSNBT)
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
	err := e.WriteSNBT(strings.Repeat("[", 10001) + strings.Repeat("]", 10001))
	if err != nil {
		t.Error(err)
	}

	// Following code should return error instant of panic.
	buf.Reset()
	err = e.WriteSNBT(strings.Repeat("[", 10002) + strings.Repeat("]", 10002))
	if err == nil {
		t.Error("Exceeded the maximum depth of support, but no error was reported")
	}
	// Panic test
	buf.Reset()
	err = e.WriteSNBT(strings.Repeat("[", 20000) + strings.Repeat("]", 20000))
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
		if err := StringifiedMessage(v.snbt).Encode(&buff); err != nil {
			t.Errorf("Encode SNBT error: %v", err)
		}
		if !bytes.Equal(buff.Bytes(), v.data) {
			t.Errorf("Encode SNBT error: %q should encoded to %d, not %d", v.snbt, v.data, buff.Bytes())
		}
		buff.Reset()
	}
}
