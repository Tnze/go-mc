package packet

import "bytes"

type Builder struct {
	buf bytes.Buffer
}

func (p *Builder) WriteField(fields ...FieldEncoder) {
	for _, f := range fields {
		_, err := f.WriteTo(&p.buf)
		if err != nil {
			panic(err)
		}
	}
}

func (p *Builder) Packet(id int32) Packet {
	return Packet{ID: id, Data: p.buf.Bytes()}
}
