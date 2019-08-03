package region

import (
	"bytes"
	"compress/zlib"
	"github.com/Tnze/go-mc/nbt"
	"testing"
)

func TestIn(t *testing.T) {
	for i := -10000; i < 10000; i++ {
		getX, _ := In(i, i)
		want := i % 32
		if want < 0 {
			want += 32
		}

		if getX != want {
			t.Errorf("fail when convert cord: get %d, want %d", getX, want)
		}
	}
}

func TestReadRegion(t *testing.T) {
	r, err := OpenRegion("../testdata/region/r.0.-1.mca")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			s, err := r.ReadSector(In(i, j))
			if err != nil {
				continue
			}

			r, err := zlib.NewReader(bytes.NewReader(s[1:]))
			if err != nil {
				t.Error(err)
			}
			var b interface{}
			err = nbt.NewDecoder(r).Decode(&b)
			if err != nil {
				t.Error(err)
			}
			//t.Log(b)
		}
	}

}

func TestFindSpace(t *testing.T) {
	r := &Region{
		sectors: map[int32]bool{
			0: true, 1: true,
			2: false, 3: true,
			4: false, 5: false, 6: true,
			7: false, 8: false, 9: false, 10: true, 11: true,
			12: false, 13: false, 14: false, 15: false, 16: false, 17: true,
			18: false, 19: false, 20: false, 21: false, 22: false, 23: false, 24: false,
		},
	}

	for _, test := range []struct{ need, index int32 }{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 7},
		{4, 12},
		{5, 12},
		{6, 18},
	} {
		i := r.findSpace(test.need)
		if i != test.index {
			t.Errorf("scan sctors fail: get %d, want %d", i, test.index)
		}
	}
}
