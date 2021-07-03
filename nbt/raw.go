package nbt

import (
	"bytes"
	"io"
)

type RawMessage []byte

func (m *RawMessage) Unmarshal(tagType byte, _ string, r DecoderReader) error {
	if tagType == TagEnd {
		return ErrEND
	}

	buf := bytes.NewBuffer((*m)[:0])
	tee := io.TeeReader(r, buf)
	err := NewDecoder(tee).rawRead(tagType)
	if err != nil {
		return err
	}
	*m = buf.Bytes()
	return nil
}
