package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

// JsonMessage is Message, unless when it is going to be Json instead of NBT
type JsonMessage Message

// ReadFrom decode Message in a JSON Text component
func (m *JsonMessage) ReadFrom(r io.Reader) (int64, error) {
	var code pk.String
	n, err := code.ReadFrom(r)
	if err != nil {
		return n, err
	}
	err = json.Unmarshal([]byte(code), (*Message)(m))
	return n, err
}

// WriteTo encode Message into a JSON Text component
func (m JsonMessage) WriteTo(w io.Writer) (int64, error) {
	code, err := json.Marshal(Message(m))
	if err != nil {
		panic(err)
	}
	return pk.String(code).WriteTo(w)
}

func (m Message) MarshalJSON() ([]byte, error) {
	if m.Translate != "" {
		return json.Marshal(translateMsg(m))
	} else {
		return json.Marshal(rawMsgStruct(m))
	}
}

func (m *Message) UnmarshalJSON(raw []byte) (err error) {
	raw = bytes.TrimSpace(raw)
	if len(raw) == 0 {
		return io.EOF
	}
	// The right way to distinguish JSON String and Object
	// is to look up the first character.
	switch raw[0] {
	case '"':
		return json.Unmarshal(raw, &m.Text) // Unmarshal as jsonString
	case '{':
		return json.Unmarshal(raw, (*rawMsgStruct)(m)) // Unmarshal as jsonMsg
	case '[':
		return json.Unmarshal(raw, &m.Extra) // Unmarshal as []Message
	default:
		return errors.New("unknown chat message type: '" + string(raw[0]) + "'")
	}
}
