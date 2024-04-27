package chat

import (
	"errors"
	"io"
	"strconv"

	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ReadFrom decode Message in a Text component
func (m *Message) ReadFrom(r io.Reader) (int64, error) {
	var code pk.String
	n, err := code.ReadFrom(r)
	if err != nil {
		return n, err
	}
	err = nbt.Unmarshal([]byte(code), (*Message)(m))
	return n, err
}

// WriteTo encode Message into a Text component
func (m Message) WriteTo(w io.Writer) (int64, error) {
	code, err := nbt.Marshal(Message(m))
	if err != nil {
		panic(err)
	}
	return pk.String(code).WriteTo(w)
}

func (m Message) MarshalNBT() ([]byte, error) {
	if m.Translate != "" {
		return nbt.Marshal(translateMsg(m))
	} else {
		return nbt.Marshal(rawMsgStruct(m))
	}
}

func (m *Message) UnmarshalNBT(raw []byte) (err error) {
	if len(raw) == 0 {
		return io.EOF
	}
	switch raw[0] {
	case nbt.TagString:
		return nbt.Unmarshal(raw, &m.Text) // Unmarshal as jsonString
	case nbt.TagCompound:
		return nbt.Unmarshal(raw, (*rawMsgStruct)(m)) // Unmarshal as jsonMsg
	case nbt.TagList:
		return nbt.Unmarshal(raw, &m.Extra) // Unmarshal as []Message
	default:
		return errors.New("unknown chat message type: '" + strconv.FormatUint(uint64(raw[0]), 16) + "'")
	}
}
