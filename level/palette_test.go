package level

import (
	"testing"
)

func TestPaletteResize(t *testing.T) {
	container := NewStatesPaletteContainer(16*16*16, 0)

	for i := 0; i < 4096; i++ {
		container.Set(i, BlocksState(i))
	}
	for i := 0; i < 4096; i++ {
		if container.Get(i) != BlocksState(i) {
			t.Errorf("Get Error, got: %v,but expect: %v", container.Get(i), BlocksState(i))
		}
	}
}
