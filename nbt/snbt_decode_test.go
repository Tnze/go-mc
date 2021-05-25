package nbt

import (
	"bytes"
	"testing"
)

func TestEncoder_WriteSNBT(t *testing.T) {
	var buf bytes.Buffer
	e := NewEncoder(&buf)
	if err := e.WriteSNBT(`{ abc: a123}`); err != nil {
		t.Fatal(err)
	}
	t.Log(buf.Bytes())

}
