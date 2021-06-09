package nbt

import (
	"bytes"
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
