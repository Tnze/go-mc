package packet

import "bytes"

type Builder struct {
	buf bytes.Buffer
}

func (p *Builder) WriteField(fields ...FieldEncoder) {
	for _, f := range fields {
		p.buf.Write(f.Encode())
	}
}

func (p *Builder) Packet(id int32) Packet {
	return Packet{ID: id, Data: p.buf.Bytes()}
}
