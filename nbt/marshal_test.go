package nbt

import (
	"bytes"
	"reflect"
	"testing"
)

func TestMarshal(t *testing.T) {
	var (
		want = []byte{
			0x0A, 0, 0,
			0x08, 0, 4, 0x4e, 0x61, 0x6d, 0x65, 0, 4, 0x54, 0x6e, 0x7a, 0x65,
			0x01, 0x00, 0x08, 0x42, 0x79, 0x74, 0x65, 0x54, 0x65, 0x73, 0x74, 0xFF,

			0,
		}
		value struct {
			Name     string
			ByteTest byte
		}
	)
	value.Name = "Tnze"
	value.ByteTest = 0xFF

	var buf bytes.Buffer
	if err := Marshal(&buf, value); err != nil {
		t.Fatal(err)
	}

	gets := buf.Bytes()
	if !reflect.DeepEqual(gets, want) {
		t.Errorf("marshal wrong: get [% 02x], want [% 02x]", gets, want)
	}
}
