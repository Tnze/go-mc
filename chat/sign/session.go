package sign

import (
	"io"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/yggdrasil/user"
)

type Message struct {
	Link struct {
		Index   int
		Sender  uuid.UUID
		Session uuid.UUID
	}
	Signature []byte
	PackedMessageBody
	Unsigned *chat.Message
	FilterMask
}

type Session struct {
	sessionID uuid.UUID
	publicKey user.PublicKey
}

func (s Session) WriteTo(w io.Writer) (n int64, err error) {
	n1, err := pk.UUID(s.sessionID).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := s.publicKey.WriteTo(w)
	return n1 + n2, err
}

func (s *Session) ReadFrom(r io.Reader) (n int64, err error) {
	n1, err := ((*pk.UUID)(&s.sessionID)).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := s.publicKey.ReadFrom(r)
	return n1 + n2, err
}

func (s *Session) Update(msg Message) bool {
	panic("todo")
}
