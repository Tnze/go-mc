package nbt

import (
	"bytes"
	"testing"
)

func TestStringifiedMessage_Decode(t *testing.T) {
	data := []byte{
		TagCompound, 0, 2, 'a', 'b',
		TagInt, 0, 3, 'K', 'e', 'y', 0, 0, 0, 12,
		TagString, 0, 5, 'V', 'a', 'l', 'u', 'e', 0, 5, 'T', 'n', ' ', 'z', 'e',
		TagList, 0, 4, 'L', 'i', 's', 't', TagCompound, 0, 0, 0, 2, 0, 0,
		TagEnd,
	}
	var container struct {
		Key   int32
		Value StringifiedMessage
		List  StringifiedMessage
	}

	if tag, err := NewDecoder(bytes.NewReader(data)).Decode(&container); err != nil {
		t.Fatal(tag, err)
	} else {
		if tag != "ab" {
			t.Fatalf("UnmarshalNBT tag name error: want %s, get: %s", "ab", tag)
		}
		if container.Key != 12 {
			t.Fatalf("UnmarshalNBT Key error: want %v, get: %v", 12, container.Key)
		}
		if container.Value != `"Tn ze"` {
			t.Fatalf("UnmarshalNBT Key error: get: %v", container.Value)
		}
		if container.List != "[{},{}]" {
			t.Fatalf("UnmarshalNBT List error: get: %v", container.List)
		}
	}
}
