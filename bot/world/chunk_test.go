package world

import (
	"math/rand"
	"testing"

	"github.com/Tnze/go-mc/data/block"
)

func newDirectSection(bpb int) Section {
	return &directSection{
		bpb:  bpb,
		data: make([]uint64, 16*16*16*bpb/64),
	}
}

func TestDirectSection(t *testing.T) {
	for bpb := 4; bpb <= block.BitsPerBlock; bpb++ {
		testSection(newDirectSection(bpb), bpb)(t)
	}
}

func TestDirectSection_clone(t *testing.T) {
	s := newDirectSection(9)
	dataset := randData(9)
	for i := 0; i < 16*16*16; i++ {
		s.SetBlock(i, dataset[i])
	}
	s = s.(*directSection).clone(block.BitsPerBlock)
	for i := 0; i < 16*16*16; i++ {
		if s := s.GetBlock(i); dataset[i] != s {
			t.Fatalf("direct section error: want: %v, get %v", dataset[i], s)
		}
	}
}

func TestPaletteSection(t *testing.T) {
	t.Run("Correctness", testSection(&paletteSection{
		palettesIndex: make(map[BlockStatus]int),
		directSection: *(newDirectSection(7).(*directSection)),
	}, 7))
	t.Run("AutomaticExpansion", testSection(&paletteSection{
		palettesIndex: make(map[BlockStatus]int),
		directSection: *(newDirectSection(4).(*directSection)),
	}, 9))
}

func testSection(s Section, bpb int) func(t *testing.T) {
	return func(t *testing.T) {
		for _, dataset := range [][16 * 16 * 16]BlockStatus{secData(bpb), randData(bpb)} {
			for i := 0; i < 16*16*16; i++ {
				s.SetBlock(i, dataset[i])
			}
			for i := 0; i < 16*16*16; i++ {
				if v := s.GetBlock(i); dataset[i] != v {
					t.Fatalf("direct section error: want: %v, get %v", dataset[i], v)
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
