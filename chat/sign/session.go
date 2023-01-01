package sign

import (
	"crypto/sha256"
	"encoding/binary"
	"io"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/yggdrasil/user"
)

type Message struct {
	Prev      Prev
	Signature *Signature
	*MessageBody
	Unsigned *chat.Message
	FilterMask
}

type Prev struct {
	Index   int
	Sender  uuid.UUID
	Session uuid.UUID
}

type Session struct {
	SessionID uuid.UUID
	PublicKey user.PublicKey

	valid   bool
	lastMsg *Message
}

func (s Session) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := pk.UUID(s.SessionID).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := s.PublicKey.WriteTo(w)
	return n1 + n2, err
}

func (s *Session) ReadFrom(r io.Reader) (n int64, err error) {
	n1, err := ((*pk.UUID)(&s.SessionID)).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := s.PublicKey.ReadFrom(r)
	return n1 + n2, err
}

func (s *Session) InitValidate() {
	s.valid = true
	s.lastMsg = nil
}

func (s *Session) VerifyAndUpdate(msg *Message) bool {
	s.valid = s.valid && s.verifyHash(msg) && s.verifyChain(msg)
	if s.valid {
		s.lastMsg = msg
		return true
	}
	return false
}

func (s *Session) verifyHash(msg *Message) bool {
	h := sha256.New()
	// 1
	_ = binary.Write(h, binary.BigEndian, int32(1))
	// Prev
	_, _ = h.Write(msg.Prev.Sender[:])
	_, _ = h.Write(msg.Prev.Session[:])
	_ = binary.Write(h, binary.BigEndian, msg.Prev.Index)
	// Body
	_ = binary.Write(h, binary.BigEndian, msg.Salt)
	_ = binary.Write(h, binary.BigEndian, msg.Timestamp.Unix())
	content := []byte(msg.PlainMsg)
	_ = binary.Write(h, binary.BigEndian, int32(len(content)))
	_, _ = h.Write(content)
	// Body.LastSeen
	_ = binary.Write(h, binary.BigEndian, int32(len(msg.LastSeen)))
	for _, v := range msg.LastSeen {
		_, _ = h.Write((*v)[:])
	}
	return s.PublicKey.VerifyMessage(h.Sum(nil), msg.Signature[:]) == nil
}

func (s *Session) verifyChain(msg *Message) bool {
	return s.lastMsg != nil && (msg.Prev.Index < s.lastMsg.Prev.Index || msg.Prev.Sender != s.lastMsg.Prev.Sender || msg.Prev.Session != s.lastMsg.Prev.Session)
}
