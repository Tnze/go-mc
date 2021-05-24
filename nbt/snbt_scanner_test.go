package nbt

import "testing"

func TestSNBT(t *testing.T) {
	var s scanner
	for _, str := range []string{
		"0", "1234567890", "3.1415926",
		"255B", "1234s", "6666L",
		"314F", "3.14f", "3.14159265358979323846264D",
	} {
		s.reset()
		var scanCodes []int
		for _, c := range []byte(str) {
			res := s.step(c)
			if res == scanError {
				t.Errorf("scan error")
			}
			scanCodes = append(scanCodes, res)
		}
		t.Logf("scancodes: %v", scanCodes)
	}
}
