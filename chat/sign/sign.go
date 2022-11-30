package sign

import (
	"crypto/rand"
	"crypto/sha256"
	"encoding/binary"
	"encoding/json"
	"io"
	"time"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

type MessageHeader struct {
	PrevSignature []byte
	Sender        uuid.UUID
}

func (m *MessageHeader) WriteTo(w io.Writer) (n int64, err error) {
	hasSignature := pk.Boolean(len(m.PrevSignature) > 0)
	n, err = hasSignature.WriteTo(w)
	if err != nil {
		return
	}
	if hasSignature {
		n2, err := pk.ByteArray(m.PrevSignature).WriteTo(w)
		n += n2
		if err != nil {
			return n, err
		}
	}
	n3, err := pk.UUID(m.Sender).WriteTo(w)
	return n + n3, err
}

func (m *MessageHeader) ReadFrom(r io.Reader) (n int64, err error) {
	var hasSignature pk.Boolean
	n, err = hasSignature.ReadFrom(r)
	if err != nil {
		return
	}
	if hasSignature {
		n2, err := (*pk.ByteArray)(&m.PrevSignature).ReadFrom(r)
		n += n2
		if err != nil {
			return n, err
		}
	}
	n3, err := (*pk.UUID)(&m.Sender).ReadFrom(r)
	return n + n3, err
}

type MessageBody struct {
	PlainMsg     string
	DecoratedMsg json.RawMessage
	Timestamp    time.Time
	Salt         int64
	History      []HistoryMessage
}

func (m *MessageBody) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.String(m.PlainMsg),
		pk.Boolean(m.DecoratedMsg != nil),
		pk.Opt{
			Has:   m.DecoratedMsg != nil,
			Field: m.DecoratedMsg,
		},
		pk.Long(m.Timestamp.UnixMilli()),
		pk.Long(m.Salt),
		pk.Array(&m.History),
	}.WriteTo(w)
}

func (m *MessageBody) ReadFrom(r io.Reader) (n int64, err error) {
	var hasContent pk.Boolean
	var timestamp pk.Long
	n, err = pk.Tuple{
		(*pk.String)(&m.PlainMsg),
		&hasContent,
		pk.Opt{
			Has:   &hasContent,
			Field: (*pk.ByteArray)(&m.DecoratedMsg),
		},
		&timestamp,
		(*pk.Long)(&m.Salt),
		pk.Array(&m.History),
	}.ReadFrom(r)
	m.Timestamp = time.UnixMilli(int64(timestamp))
	return
}

func (m *MessageBody) Hash() []byte {
	hash := sha256.New()
	binary.Write(hash, binary.BigEndian, m.Salt)
	binary.Write(hash, binary.BigEndian, m.Timestamp.Second())
	hash.Write([]byte(m.PlainMsg))
	hash.Write([]byte{70})
	if m.DecoratedMsg != nil {
		decorated, _ := m.DecoratedMsg.MarshalJSON()
		hash.Write(decorated)
	}
	for _, v := range m.History {
		hash.Write([]byte{70})
		hash.Write(v.Sender[:])
		hash.Write(v.Signature)
	}
	return hash.Sum(nil)
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

type PlayerMessage struct {
	MessageHeader
	MessageSignature []byte
	MessageBody
	UnsignedContent *chat.Message
	FilterMask
}

func (msg *PlayerMessage) ReadFrom(r io.Reader) (n int64, err error) {
	var hasUnsignedContent pk.Boolean
	return pk.Tuple{
		&msg.MessageHeader,
		(*pk.ByteArray)(&msg.MessageSignature),
		&msg.MessageBody,
		&hasUnsignedContent,
		pk.Opt{
			Has: &hasUnsignedContent,
			Field: func() pk.FieldDecoder {
				msg.UnsignedContent = new(chat.Message)
				return msg.UnsignedContent
			},
		},
		&msg.FilterMask,
	}.ReadFrom(r)
}

func (msg *PlayerMessage) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		&msg.MessageHeader,
		pk.ByteArray(msg.MessageSignature),
		&msg.MessageBody,
		pk.Boolean(msg.UnsignedContent != nil),
		pk.Opt{
			Has:   msg.UnsignedContent != nil,
			Field: msg.UnsignedContent,
		},
		&msg.FilterMask,
	}.WriteTo(w)
}

func genSalt() (salt int64) {
	err := binary.Read(rand.Reader, binary.BigEndian, &salt)
	if err != nil {
		panic(err)
	}
	return
}

func Unsigned(id uuid.UUID, plain string, content *chat.Message) (msg PlayerMessage) {
	return PlayerMessage{
		MessageHeader: MessageHeader{
			PrevSignature: nil,
			Sender:        id,
		},
		MessageSignature: []byte{},
		MessageBody: MessageBody{
			PlainMsg:     plain,
			DecoratedMsg: nil,
			Timestamp:    time.Now(),
			Salt:         genSalt(),
			History:      nil,
		},
		UnsignedContent: nil,
		FilterMask:      FilterMask{Type: 0},
	}
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

func (p *HistoryMessage) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.UUID(p.Sender).WriteTo(w)
	if err != nil {
		return
	}
	n2, err := pk.ByteArray(p.Signature).WriteTo(w)
	return n + n2, err
}
