package world

import (
	"github.com/Tnze/go-mc/data"
	"math/rand"
	"testing"
)

func newDirectSection(bpb int) *directSection {
	return &directSection{
		bpb:  bpb,
		data: make([]uint64, 16*16*16*bpb/64),
	}
}

func TestDirectSection(t *testing.T) {
	for bpb := 3; bpb <= data.BitsPerBlock; bpb++ {
		s := newDirectSection(bpb)
		for _, dataset := range [][16 * 16 * 16]BlockStatus{secData(bpb), randData(bpb)} {
			for i := 0; i < 16*16*16; i++ {
				s.SetBlock(i%16, i/16%16, i/16/16, dataset[i])
			}
			for i := 0; i < 16*16*16; i++ {
				if s := s.GetBlock(i%16, i/16%16, i/16/16); dataset[i] != s {
					t.Fatalf("direct section error: want: %v, get %v", dataset[i], s)
				}
			}
		}
	}
}

func secData(bpb int) (data [16 * 16 * 16]BlockStatus) {
	mask := 1<<bpb - 1
	var v int
	for i := range data {
		data[i] = BlockStatus(v)
		v = (v + 1) & mask
	}
	return
}

func randData(bpb int) (data [16 * 16 * 16]BlockStatus) {
	data = secData(bpb)
	rand.Shuffle(len(data), func(i, j int) {
		data[i], data[j] = data[j], data[i]
	})
	return
}
