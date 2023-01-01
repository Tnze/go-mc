package sign

import (
	"io"
	"time"

	"github.com/google/uuid"

	pk "github.com/Tnze/go-mc/net/packet"
)

type PackedMessageBody struct {
	PlainMsg  string
	Timestamp time.Time
	Salt      int64
	LastSeen  []PackedSignature
}

func (m *PackedMessageBody) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.String(m.PlainMsg),
		pk.Long(m.Timestamp.UnixMilli()),
		pk.Long(m.Salt),
		pk.Array(m.LastSeen),
	}.WriteTo(w)
}

func (m *PackedMessageBody) ReadFrom(r io.Reader) (n int64, err error) {
	var timestamp pk.Long
	n, err = pk.Tuple{
		(*pk.String)(&m.PlainMsg),
		&timestamp,
		(*pk.Long)(&m.Salt),
		pk.Array(&m.LastSeen),
	}.ReadFrom(r)
	m.Timestamp = time.UnixMilli(int64(timestamp))
	return
}

type HistoryMessage struct {
	Sender    uuid.UUID
	Signature []byte
}

func (p *HistoryMessage) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = (*pk.UUID)(&p.Sender).ReadFrom(r)
	if err != nil {
		return
	}
	n2, err := (*pk.ByteArray)(&p.Signature).ReadFrom(r)
	return n + n2, err
}

func (p HistoryMessage) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.UUID(p.Sender).WriteTo(w)
	if err != nil {
		return
	}
	n2, err := pk.ByteArray(p.Signature).WriteTo(w)
	return n + n2, err
}

type HistoryUpdate struct {
	Offset       pk.VarInt
	Acknowledged pk.FixedBitSet // n == 20
}

func (h HistoryUpdate) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{h.Offset, h.Acknowledged}.WriteTo(w)
}

func (h *HistoryUpdate) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{&h.Offset, &h.Acknowledged}.ReadFrom(r)
}

type Signature [256]byte

func (s Signature) WriteTo(w io.Writer) (n int64, err error) {
	n2, err := w.Write(s[:])
	return int64(n2), err
}

func (s *Signature) ReadFrom(r io.Reader) (n int64, err error) {
	n2, err := r.Read(s[:])
	return int64(n2), err
}

type PackedSignature struct {
	ID int32
	*Signature
}

func (p PackedSignature) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := pk.VarInt(p.ID + 1).WriteTo(w)
	if err != nil {
		return n1, err
	}
	if p.Signature != nil {
		n2, err := w.Write(p.Signature[:])
		return n1 + int64(n2), err
	}
	return n1, err
}

func (p PackedSignature) ReadFrom(r io.Reader) (n int64, err error) {
	n1, err := (*pk.VarInt)(&p.ID).ReadFrom(r)
	if err != nil {
		return n1, err
	}

	if p.ID == -1 {
		if p.Signature == nil {
			p.Signature = new(Signature)
		}
		n2, err := r.Read(p.Signature[:])
		return n1 + int64(n2), err
	} else {
		p.Signature = nil
		return n1, err
	}
}

type FilterMask struct {
	Type byte
	Mask pk.BitSet
}

func (f *FilterMask) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.VarInt(f.Type).WriteTo(w)
	if err != nil {
		return
	}
	if f.Type == 2 {
		var n1 int64
		n1, err = f.Mask.WriteTo(w)
		n += n1
	}
	return
}

func (f *FilterMask) ReadFrom(r io.Reader) (n int64, err error) {
	var Type pk.VarInt
	if n, err = Type.ReadFrom(r); err != nil {
		return
	}
	f.Type = byte(Type)
	if f.Type == 2 {
		var n1 int64
		n1, err = f.Mask.ReadFrom(r)
		n += n1
	}
	return
}
