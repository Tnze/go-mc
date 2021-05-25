package nbt

import (
	_ "embed"
	"testing"
)

func TestSNBT_checkScanCode(t *testing.T) {
	t.SkipNow()
	var s scanner
	s.reset()
	for _, c := range []byte(`{ "a b\"c": {}, def: 12345, 'gh"i': 0.123f}`) {
		t.Logf("[%c] - %d", c, s.step(c))
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
			res := s.step(c)
			if res == scanError {
				return false
			}
		}
		return true
	}
	for _, str := range goods {
		if scan(str) == false {
			t.Errorf("scan valid data %q error: %v", str, s.err)
		}
	}
}

//go:embed bigTest_test.snbt
var bigTest string

func TestSNBT_compound(t *testing.T) {
	goods := []string{
		`{}`, `{name:3.14f}`, `{ "name" : 12345 }`,
		`{ abc: { }}`, `{ "a b\"c": {}, def: 12345}`,
		bigTest,
	}
	var s scanner
	for _, str := range goods {
		s.reset()
		for i, c := range []byte(str) {
			res := s.step(c)
			if res == scanError {
				t.Errorf("scan valid data %q error: %v at [%d]", str[:i], s.err, i)
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
	}
	var s scanner
	scan := func(str string) bool {
		s.reset()
		for _, c := range []byte(str) {
			res := s.step(c)
			if res == scanError {
				return false
			}
		}
		return true
	}
	for _, str := range goods {
		if scan(str) == false {
			t.Errorf("scan valid data %q error: %v", str, s.err)
		}
	}
}
