package level

import (
	"math/rand"
	"testing"

	"github.com/Tnze/go-mc/level/block"
)

func TestPaletteContainer_seq(t *testing.T) {
	container := NewStatesPaletteContainer(4096, 0)
	for i := 0; i < 4096; i++ {
		container.Set(i, BlocksState(i))
	}
	for i := 0; i < 4096; i++ {
		if container.Get(i) != BlocksState(i) {
			t.Errorf("Get Error, got: %v, but expect: %v", container.Get(i), BlocksState(i))
		}
	}
}

func TestPaletteContainer_rand(t *testing.T) {
	data := make([]BlocksState, 4096)
	for i := range data {
		data[i] = BlocksState(rand.Intn(1 << block.BitsPerBlock))
	}
	container := NewStatesPaletteContainer(4096, 0)
	for i, v := range data {
		container.Set(i, v)
	}
	for i, v := range data {
		if v2 := container.Get(i); v != v2 {
			t.Errorf("value not match: got %v, except: %v", v2, v)
		}
	}
}

func BenchmarkPaletteContainer(b *testing.B) {
	data := make([]BlocksState, 4096)
	for i := range data {
		data[i] = BlocksState(rand.Intn(1 << block.BitsPerBlock))
	}
	rand.Shuffle(len(data), func(i, j int) { data[i], data[j] = data[j], data[i] })
	container := NewStatesPaletteContainer(4096, 0)
	b.ResetTimer()
	b.Run("Set", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			index := i % 4096
			container.Set(index, data[index])
		}
	})
	b.Run("Get", func(b *testing.B) {
		for i := 0; i < b.N; i++ {
			index := i % 4096
			container.Get(index)
		}
	})
}
