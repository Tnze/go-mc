package world

import (
	"encoding/binary"
	"encoding/hex"
	"testing"
)

func TestBitArrayBasic(t *testing.T) {
	a := bitArray{
		width:            5,
		valuesPerElement: valuesPerBitArrayElement(5),
		data:             make([]uint64, 5),
	}

	if got, want := a.Size(), 12*5; got != want {
		t.Errorf("size = %d, want %d", got, want)
	}

	a.Set(0, 4)
	if v := a.Get(0); v != 4 {
		t.Errorf("v[0] = %d, want 4", v)
	}
	a.Set(12, 8)
	if v := a.Get(12); v != 8 {
		t.Errorf("v[12] = %d, want 8", v)
	}
}

func TestBitArrayHardcoded(t *testing.T) {
	d1, _ := hex.DecodeString("0020863148418841")
	d2, _ := hex.DecodeString("01018A7260F68C87")

	a := bitArray{
		width:            5,
		valuesPerElement: valuesPerBitArrayElement(5),
		data:             []uint64{binary.BigEndian.Uint64(d1), binary.BigEndian.Uint64(d2)},
	}

	if got, want := a.Size(), 12*2; got != want {
		t.Errorf("size = %d, want %d", got, want)
	}

	want := []uint{1, 2, 2, 3, 4, 4, 5, 6, 6, 4, 8, 0, 7, 4, 3, 13, 15, 16, 9, 14, 10, 12, 0, 2}
	for idx, want := range want {
		if got := a.Get(uint(idx)); got != want {
			t.Errorf("v[%d] = %d, want %d", idx, got, want)
		}
	}
}
