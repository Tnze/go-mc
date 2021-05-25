package nbt

import "testing"

func TestSNBT_number(t *testing.T) {
	goods := []string{
		"0", "1234567890", "3.1415926",
		"-0", "-1234567890", "-3.1415926",
		"255B", "1234s", "6666L",
		"314F", "3.14f", "3.14159265358979323846264D",
	}
	bads := []string{
		".0", "1234.5678.90",
		"25-5B", "1234.s",
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
	for _, str := range bads {
		if scan(str) {
			t.Errorf("scan invalid data %q success", str)
		}
	}
}

func TestSNBT_compound(t *testing.T) {
	goods := []string{
		`{}`, `{name:3.14f}`, `{ "name" : 12345 }`,
		`{ abc: {}}`,
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
