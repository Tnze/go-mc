package nbt

import (
	"bytes"
	"testing"
)

func TestRawMessage_Encode(t *testing.T) {
	data := []byte{
		TagCompound, 0, 2, 'a', 'b',
		TagInt, 0, 3, 'K', 'e', 'y', 0, 0, 0, 12,
		TagString, 0, 5, 'V', 'a', 'l', 'u', 'e', 0, 4, 'T', 'n', 'z', 'e',
		TagEnd,
	}
	var container struct {
		Key   int32
		Value RawMessage
	}
	container.Key = 12
	container.Value.Type = TagString
	container.Value.Data = []byte{0, 4, 'T', 'n', 'z', 'e'}

	var buf bytes.Buffer
	if err := NewEncoder(&buf).Encode(container, "ab"); err != nil {
		t.Fatalf("Encode error: %v", err)
	} else if !bytes.Equal(data, buf.Bytes()) {
		t.Fatalf("Encode error: want %v, get: %v", data, buf.Bytes())
	}
}

func TestRawMessage_Decode(t *testing.T) {
	data := []byte{
		TagCompound, 0, 2, 'a', 'b',
		TagInt, 0, 3, 'K', 'e', 'y', 0, 0, 0, 12,
		TagString, 0, 5, 'V', 'a', 'l', 'u', 'e', 0, 4, 'T', 'n', 'z', 'e',
		TagList, 0, 4, 'L', 'i', 's', 't', TagCompound, 0, 0, 0, 2, 0, 0,
		TagEnd,
	}
	var container struct {
		Key   int32
		Value RawMessage
		List  RawMessage
	}

	if tag, err := NewDecoder(bytes.NewReader(data)).Decode(&container); err != nil {
		t.Fatal(tag)
	} else {
		if tag != "ab" {
			t.Fatalf("Decode tag name error: want %s, get: %s", "ab", tag)
		}
		if container.Key != 12 {
			t.Fatalf("Decode Key error: want %v, get: %v", 12, container.Key)
		}
		if !bytes.Equal(container.Value.Data, []byte{
			0, 4, 'T', 'n', 'z', 'e',
		}) {
			t.Fatalf("Decode Key error: get: %v", container.Value)
		}
		if !bytes.Equal(container.List.Data, []byte{
			TagCompound, 0, 0, 0, 2,
			0, 0,
		}) {
			t.Fatalf("Decode List error: get: %v", container.List)
		}
	}
}
