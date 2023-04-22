package save

import (
	"path/filepath"
	"testing"

	"github.com/Tnze/go-mc/save/region"
)

func TestColumn(t *testing.T) {
	files, err := filepath.Glob("testdata/region/r.*.*.mca")
	if err != nil {
		return
	}
	for _, filename := range files {
		r, err := region.Open(filename)
		if err != nil {
			t.Fatal(err)
		}

		for x := 0; x < 32; x++ {
			for z := 0; z < 32; z++ {
				if !r.ExistSector(x, z) {
					continue
				}

				data, err := r.ReadSector(x, z)
				if err != nil {
					t.Fatalf("read %s sec (%d, %d) fail: %v", filepath.Base(filename), x, z, err)
				}

				var c Chunk
				err = c.Load(data)
				if err != nil {
					t.Fatalf("read %s sec (%d, %d) fail: %v", filepath.Base(filename), x, z, err)
				}
			}
		}

		err = r.Close()
		if err != nil {
			t.Fatal(err)
		}
	}
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
		// x, z := rand.Intn(32), rand.Intn(32)
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
