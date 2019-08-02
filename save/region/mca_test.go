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

func TestSectorsFinder(t *testing.T) {
	sectors := []byte{
		1, 1,
		0, 1, //2
		0, 0, 1, //4
		0, 0, 0, 1, 1, //7
		0, 0, 0, 0, 0, 1, //12
		0, 0, 0, 0, 0, 0, 0, //18
	}

	scan := func(need int) (index int) {
		for i := 0; i < need; i++ {
			if sectors[index+i] != 0 {
				index += i + 1
				i = -1 // 0 for next loop
			}
		}
		return
	}

	for _, test := range []struct{ need, index int }{
		{0, 0},
		{1, 2},
		{2, 4},
		{3, 7},
		{4, 12},
		{5, 12},
		{6, 18},
	} {
		i := scan(test.need)
		if i != test.index {
			t.Errorf("scan sctors fail: get %d, want %d", i, test.index)
		}
	}
}
