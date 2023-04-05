// Package chat implements Minecraft's chat message encoding system.
//
// The type Message is the Minecraft chat message. Can be encoded as JSON
// or net/packet.Field .
//
// It's very recommended that use SetLanguage before using Message.String or Message.ClearString,
// or the `github.com/Tnze/go-mc/data/en-us` will be used.
// Note: The package of data/lang/... will SetLanguage on theirs init() so you don't need to call by your self.
//
// Some of these docs is copied from https://wiki.vg/Chat.
package chat

import (
	"bytes"
	"encoding/json"
	"errors"
	"fmt"
	"io"
	"regexp"
	"strings"

	en_us "github.com/Tnze/go-mc/data/lang/en-us"
	pk "github.com/Tnze/go-mc/net/packet"
)

const (
	Chat = iota
	System
	GameInfo
	SayCommand
	MsgCommand
	TeamMsgCommand
	EmoteCommand
	TellrawCommand
)

// Colors
const (
	Black       = "black"
	DarkBlue    = "dark_blue"
	DarkGreen   = "dark_green"
	DarkAqua    = "dark_aqua"
	DarkRed     = "dark_red"
	DarkPurple  = "dark_purple"
	Gold        = "gold"
	Gray        = "gray"
	DarkGray    = "dark_gray"
	Blue        = "blue"
	Green       = "green"
	Aqua        = "aqua"
	Red         = "red"
	LightPurple = "light_purple"
	Yellow      = "yellow"
	White       = "white"
)

// Message is a message sent by other
type Message struct {
	Text string `json:"text"`

	Bold          bool `json:"bold,omitempty"`          // 粗体
	Italic        bool `json:"italic,omitempty"`        // 斜体
	UnderLined    bool `json:"underlined,omitempty"`    // 下划线
	StrikeThrough bool `json:"strikethrough,omitempty"` // 删除线
	Obfuscated    bool `json:"obfuscated,omitempty"`    // 随机
	// Font of the message, could be one of minecraft:uniform, minecraft:alt or minecraft:default
	// This option is only valid on 1.16+, otherwise the property is ignored.
	Font  string `json:"font,omitempty"`  // 字体
	Color string `json:"color,omitempty"` // 颜色

	// Insertion contains text to insert. Only used for messages in chat.
	// When shift is held, clicking the component inserts the given text
	// into the chat box at the cursor (potentially replacing selected text).
	Insertion  string      `json:"insertion,omitempty"`
	ClickEvent *ClickEvent `json:"clickEvent,omitempty"`
	HoverEvent *HoverEvent `json:"hoverEvent,omitempty"`

	Translate string    `json:"translate,omitempty"`
	With      []Message `json:"with,omitempty"`
	Extra     []Message `json:"extra,omitempty"`
}

// Same as Message, but "Text" is omitempty
type translateMsg struct {
	Text string `json:"text,omitempty"`

	Bold          bool `json:"bold,omitempty"`
	Italic        bool `json:"italic,omitempty"`
	UnderLined    bool `json:"underlined,omitempty"`
	StrikeThrough bool `json:"strikethrough,omitempty"`
	Obfuscated    bool `json:"obfuscated,omitempty"`

	Font  string `json:"font,omitempty"`
	Color string `json:"color,omitempty"`

	Insertion  string      `json:"insertion,omitempty"`
	ClickEvent *ClickEvent `json:"clickEvent,omitempty"`
	HoverEvent *HoverEvent `json:"hoverEvent,omitempty"`

	Translate string    `json:"translate"`
	With      []Message `json:"with,omitempty"`
	Extra     []Message `json:"extra,omitempty"`
}

type jsonMsg Message

func (m Message) MarshalJSON() ([]byte, error) {
	if m.Translate != "" {
		return json.Marshal(translateMsg(m))
	} else {
		return json.Marshal(jsonMsg(m))
	}
}

// UnmarshalJSON decode json to Message
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
		return json.Unmarshal(raw, (*jsonMsg)(m)) // Unmarshal as jsonMsg
	case '[':
		return json.Unmarshal(raw, &m.Extra) // Unmarshal as []Message
	default:
		return errors.New("unknown chat message type: '" + string(raw[0]) + "'")
	}
}

// ReadFrom decode Message in a ChatMsg packet
func (m *Message) ReadFrom(r io.Reader) (int64, error) {
	var code pk.String
	n, err := code.ReadFrom(r)
	if err != nil {
		return n, err
	}
	err = json.Unmarshal([]byte(code), m)
	return n, err
}

// WriteTo encode Message into a ChatMsg packet
func (m Message) WriteTo(w io.Writer) (int64, error) {
	code, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return pk.String(code).WriteTo(w)
}

// Append extra message to the end of the message and return the new one.
// The source message remains unchanged.
func (m Message) Append(extraMsg ...Message) Message {
	origLen := len(m.Extra)
	finalLen := origLen + len(extraMsg)
	var extra []Message
	if cap(m.Extra) < len(m.Extra)+len(extraMsg) {
		extra = make([]Message, finalLen)
		copy(extra, m.Extra)
	} else {
		extra = m.Extra[:finalLen]
	}
	copy(extra[origLen:], extraMsg)
	m.Extra = extra
	return m
}

func (m Message) SetColor(color string) Message {
	m.Color = color
	return m
}

func Text(str string) Message {
	return Message{Text: str}
}

func TranslateMsg(key string, with ...Message) (m Message) {
	m.Translate = key
	m.With = with
	return
}

var fmtCode = map[byte]string{
	'0': "30",
	'1': "34",
	'2': "32",
	'3': "36",
	'4': "31",
	'5': "35",
	'6': "33",
	'7': "37",
	'8': "90",
	'9': "94",
	'a': "92",
	'b': "96",
	'c': "91",
	'd': "95",
	'e': "93",
	'f': "97",

	// 'k':"",	//random
	'l': "1",
	'm': "9",
	'n': "4",
	'o': "3",
	'r': "0",
}

var colors = map[string]string{
	Black:       "30",
	DarkBlue:    "34",
	DarkGreen:   "32",
	DarkAqua:    "36",
	DarkRed:     "31",
	DarkPurple:  "35",
	Gold:        "33",
	Gray:        "37",
	DarkGray:    "90",
	Blue:        "94",
	Green:       "92",
	Aqua:        "96",
	Red:         "91",
	LightPurple: "95",
	Yellow:      "93",
	White:       "97",
}

// translateMap is the translation table.
// By default, it's en-us.
var translateMap = en_us.Map

// SetLanguage set the default language used by String() and ClearString().
func SetLanguage(trans map[string]string) {
	translateMap = trans
}

// ClearString return the message String without escape sequence for ansi color.
func (m Message) ClearString() string {
	var msg strings.Builder
	text, _ := TransCtrlSeq(m.Text, false)
	msg.WriteString(text)

	// handle translate
	if m.Translate != "" {
		args := make([]any, len(m.With))
		for i, v := range m.With {
			args[i] = v.ClearString()
		}

		_, _ = fmt.Fprintf(&msg, translateMap[m.Translate], args...)
	}

	if m.Extra != nil {
		for i := range m.Extra {
			msg.WriteString(m.Extra[i].ClearString())
		}
	}
	return msg.String()
}

// String return the message string with escape sequence for ansi color.
// To convert Translated Message to string, you must set
// On Windows, you may want print this string using github.com/mattn/go-colorable.
func (m Message) String() string {
	var msg, format strings.Builder
	if m.Bold {
		format.WriteString("1;")
	}
	if m.Italic {
		format.WriteString("3;")
	}
	if m.UnderLined {
		format.WriteString("4;")
	}
	if m.StrikeThrough {
		format.WriteString("9;")
	}
	if m.Color != "" {
		format.WriteString(colors[m.Color] + ";")
	}
	if format.Len() > 0 {
		msg.WriteString("\033[" + format.String()[:format.Len()-1] + "m")
	}

	text, ok := TransCtrlSeq(m.Text, true)
	msg.WriteString(text)

	// handle translate
	if m.Translate != "" {
		args := make([]any, len(m.With))
		for i, v := range m.With {
			args[i] = v
		}

		_, _ = fmt.Fprintf(&msg, translateMap[m.Translate], args...)
	}

	if m.Extra != nil {
		for i := range m.Extra {
			msg.WriteString(m.Extra[i].String())
		}
	}

	if format.Len() > 0 || ok {
		msg.WriteString("\033[0m")
	}
	return msg.String()
}

var fmtPat = regexp.MustCompile(`(?i)§[\dA-FK-OR]`)

// TransCtrlSeq will transform control sequences into ANSI code
// or simply filter them. Depends on the second argument.
// if the str contains control sequences, returned change=true.
func TransCtrlSeq(str string, ansi bool) (dst string, change bool) {
	dst = fmtPat.ReplaceAllStringFunc(
		str,
		func(str string) string {
			f, ok := fmtCode[str[2]]
			if ok {
				if ansi {
					change = true
					return "\033[" + f + "m" // enable, add ANSI code
				}
				return "" // disable, remove the § code
			}
			return str // not a § code
		},
	)
	return
}
