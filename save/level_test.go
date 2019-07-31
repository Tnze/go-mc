package save

import (
	"compress/gzip"
	"os"
	"testing"
)

func TestLevel(t *testing.T) {
	f, err := os.Open("testdata/level.dat")
	if err != nil {
		t.Fatal(err)
	}

	r, err := gzip.NewReader(f)
	if err != nil {
		t.Fatal(err)
	}

	data, err := ReadLevel(r)
	if err != nil {
		t.Fatal(err)
	}

	//want := PlayerData{
	//	Pos:    [3]float64{-41.5, 65, -89.5},
	//	Motion: [3]float64{0, -0.0784000015258789, 0},
	//	Rotation: [2]float32{0,0},
	//}

	t.Logf("%+v", data)
	//if data != want {
	//	t.Errorf("player data parse error: get %v, want %v", data, want)
	//}
}
