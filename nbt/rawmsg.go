package nbt

import (
	"bytes"
	"io"
)

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
	m.Data = buf.Bytes()
	return nil
}
