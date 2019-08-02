package region

import (
	"bytes"
	"compress/zlib"
	"io/ioutil"
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

	s, err := r.ReadSector(In(31, 0))
	if err != nil {
		t.Fatal(err)
	}

	reader, err := zlib.NewReader(bytes.NewReader(s[1:]))
	if err != nil {
		t.Fatal(err)
	}

	s, err = ioutil.ReadAll(reader)
	if err != nil {
		t.Fatal(err)
	}

	t.Log(s)
}
