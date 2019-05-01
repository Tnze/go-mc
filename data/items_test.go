package data

import "testing"

func TestItemsString(t *testing.T) {
	for i := 0; i < 789+1; i++ {
		t.Log(Solt{ID: i, Count: byte(i%64 + 1)})
	}
}
