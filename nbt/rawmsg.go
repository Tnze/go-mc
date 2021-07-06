package nbt

import (
	"bytes"
	"errors"
	"io"
	"reflect"
	"strings"
)

// RawMessage stores the raw binary data of NBT.
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

func (m RawMessage) Unmarshal(v interface{}) error {
	d := NewDecoder(bytes.NewReader(m.Data))
	val := reflect.ValueOf(v)
	if val.Kind() != reflect.Ptr {
		return errors.New("nbt: non-pointer passed to Decode")
	}
	return d.unmarshal(val, m.Type)
}
