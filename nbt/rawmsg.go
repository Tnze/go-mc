package nbt

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
)

// RawMessage stores the raw binary data of NBT.
// This is usable if you wanna store an unknown NBT data and parse it later.
// Notice that this struct doesn't store the tag name. To convert RawMessage to valid NBT binary value:
// Encoder.Encode(RawMessage, Name) = []byte{ Type (1 byte) | n (2 byte) | Name (n byte) | Data}.
type RawMessage struct {
	Type byte
	Data []byte
}

func (m RawMessage) TagType() byte {
	return m.Type
}

func (m RawMessage) Encode(w io.Writer) error {
	_, err := w.Write(m.Data)
	return err
}

func (m *RawMessage) Decode(tagType byte, r DecoderReader) error {
	if tagType == TagEnd {
		return ErrEND
	}
	buf := bytes.NewBuffer(m.Data[:0])
	tee := io.TeeReader(r, buf)
	err := NewDecoder(tee).rawRead(tagType)
	if err != nil {
		return err
	}
	m.Type = tagType
	m.Data = buf.Bytes()
	return nil
}

// String convert the data into the SNBT(Stringified NBT) format.
// The output is valid for using in in-game command.
func (m RawMessage) String() string {
	if m.Type == TagEnd {
		return "TagEnd"
	}
	var snbt StringifiedMessage
	var sb strings.Builder
	r := bytes.NewReader(m.Data)
	d := NewDecoder(r)
	err := snbt.encode(d, &sb, m.Type)
	if err != nil {
		return "Invalid"
	}
	return sb.String()
}

// Unmarshal decode the data into v.
func (m RawMessage) Unmarshal(v interface{}) error {
	d := NewDecoder(bytes.NewReader(m.Data))
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("nbt: non-pointer passed to Decode")
	}
	return d.unmarshal(val, m.Type)
}
