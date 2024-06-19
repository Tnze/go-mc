package nbt

import (
	_ "embed"
	"testing"
)

func TestSNBT_checkScanCode(t *testing.T) {
	// t.SkipNow()
	var s scanner
	s.reset()
	for _, c := range []byte(`{a:[B;]}`) {
		t.Logf("[%c] - %d", c, s.step(&s, c))
	}
	t.Logf("[%c] - %d", ' ', s.eof())
}

func TestSNBT_number(t *testing.T) {
	goods := []string{
		"0", "1234567890", "3.1415926",
		"-0", "-1234567890", "-3.1415926",
		"255B", "1234s", "6666L",
		"314F", "3.14f", "3.14159265358979323846264D",
	}
	var s scanner
	scan := func(str string) bool {
		s.reset()
		for _, c := range []byte(str) {
			res := s.step(&s, c)
			if res == scanError {
				return false
			}
		}
		return true
	}
	for _, str := range goods {
		if scan(str) == false {
			t.Errorf("scan valid data %q error: %v", str, s.errContext)
		}
	}
}

//go:embed testdata/bigTest_test.snbt
var bigTestSNBT string

//go:embed testdata/1-dimension_codec.snbt
var dimensionCodecSNBT string

//go:embed testdata/58f6356e-b30c-4811-8bfc-d72a9ee99e73.dat.snbt
var tnzePlayerDat string

//go:embed testdata/level.dat.snbt
var tnzeLevelDat string

func TestSNBT_compound(t *testing.T) {
	goods := []string{
		`{}`, `{name:3.14f}`, `{ "name" : 12345 }`,
		`{ abc: { }}`, `{ "a b\"c": {}, def: 12345}`,
		`{ ghi: [], klm: 1}`,
		`{T: 1.2E3d, U: 1.2e-3D, V: 12e3d, W: -1.2E3F }`,
		bigTestSNBT,
		dimensionCodecSNBT,
		tnzePlayerDat,
		tnzeLevelDat,
	}
	var s scanner
	for _, str := range goods {
		s.reset()
		for i, c := range []byte(str) {
			res := s.step(&s, c)
			if res == scanError {
				t.Errorf("scan valid data %q error: %v at [%d]", str[:i], s.errContext, i)
				break
			}
		}
	}
}

func TestSNBT_list(t *testing.T) {
	goods := []string{
		`[]`, `[a, 'b', "c", d]`, // List of string
		`[{}, {}, {"a\"b":520}]`, // List of Compound
		`[B,C,D]`, `[L, "abc"]`,  // List of string (like array)
		`[B; 01B, 02B, 3B, 10B, 127B]`, // Array
		`[I;]`, `[B;   ]`,              // Empty array
		`[I; 1, 2, 3, -1, -2, -3]`, // Int array with negtive numbers
		`[L; 123L, -123L]`,         // Long array with negtive numbers
		`{a:[],b:[B;]}`,            // List or Array in TagCompound
	}
	var s scanner
	scan := func(str string) bool {
		s.reset()
		for _, c := range []byte(str) {
			res := s.step(&s, c)
			if res == scanError {
				return false
			}
		}
		return true
	}
	for _, str := range goods {
		if scan(str) == false {
			t.Errorf("scan valid data %q error: %v", str, s.errContext)
		}
	}
}

func BenchmarkSNBT_bigTest(b *testing.B) {
	var s scanner
	for i := 0; i < b.N; i++ {
		s.reset()
		for _, c := range []byte(bigTestSNBT) {
			res := s.step(&s, c)
			if res == scanError {
				b.Errorf("scan valid data %q error: %v at [%d]", bigTestSNBT[:i], s.errContext, i)
				break
			}
		}
	}
}
