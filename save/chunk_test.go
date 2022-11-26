package save

import (
	"github.com/Tnze/go-mc/save/region"
	"testing"
)

func TestColumn(t *testing.T) {
	r, err := region.Open("testdata/region/r.0.0.mca")
	if err != nil {
		t.Fatal(err)
	}
	defer r.Close()

	var c Chunk
	data, err := r.ReadSector(0, 0)
	if err != nil {
		t.Fatal(err)
	}
	err = c.Load(data)
	if err != nil {
		t.Fatal(err)
	}

	t.Logf("%+v", c)
}

func BenchmarkColumn_Load(b *testing.B) {
	// Test how many times we load a chunk
	var c Chunk
	r, err := region.Open("testdata/region/r.-1.-1.mca")
	if err != nil {
		b.Fatal(err)
	}
	defer r.Close()

	for i := 0; i < b.N; i++ {
		x, z := (i%1024)/32, (i%1024)%32
		//x, z := rand.Intn(32), rand.Intn(32)
		if !r.ExistSector(x, z) {
			continue
		}

		data, err := r.ReadSector(x, z)
		if err != nil {
			b.Fatal(err)
		}

		err = c.Load(data)
		if err != nil {
			b.Fatal(err)
		}
	}
}
