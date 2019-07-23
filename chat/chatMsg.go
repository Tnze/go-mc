package chat

import (
	"encoding/json"
	"fmt"
	"io"
	"regexp"
	"strings"

	"github.com/Tnze/go-mc/data"
	pk "github.com/Tnze/go-mc/net/packet"
)

//Message is a message sent by other
type Message jsonChat

type jsonChat struct {
	Text string `json:"text,omitempty"`

	Bold          bool   `json:"bold,omitempty"`          //粗体
	Italic        bool   `json:"Italic,omitempty"`        //斜体
	UnderLined    bool   `json:"underlined,omitempty"`    //下划线
	StrikeThrough bool   `json:"strikethrough,omitempty"` //删除线
	Obfuscated    bool   `json:"obfuscated,omitempty"`    //随机
	Color         string `json:"color,omitempty"`

	Translate string            `json:"translate,omitempty"`
	With      []json.RawMessage `json:"with,omitempty"` // How can go handle an JSON array with Object and String?
	Extra     []jsonChat        `json:"extra,omitempty"`
}

//UnmarshalJSON decode json to Message
func (m *Message) UnmarshalJSON(jsonMsg []byte) (err error) {
	if jsonMsg[0] == '"' {
		err = json.Unmarshal(jsonMsg, &m.Text) //Unmarshal as jsonString
	} else {
		err = json.Unmarshal(jsonMsg, (*jsonChat)(m)) //Unmarshal as jsonChat
	}
	return
}

//Decode for a ChatMsg packet
func (m *Message) Decode(r pk.DecodeReader) error {
	var Len pk.VarInt
	if err := Len.Decode(r); err != nil {
		return err
	}

	return json.NewDecoder(io.LimitReader(r, int64(Len))).Decode(m)
}

//Encode for a ChatMsg packet
func (m Message) Encode() []byte {
	code, err := json.Marshal(m)
	if err != nil {
		panic(err)
	}
	return code
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
	"black":        "30",
	"dark_blue":    "34",
	"dark_green":   "32",
	"dark_aqua":    "36",
	"dark_red":     "31",
	"dark_purple":  "35",
	"gold":         "33",
	"gray":         "37",
	"dark_gray":    "90",
	"blue":         "94",
	"green":        "92",
	"aqua":         "96",
	"red":          "91",
	"light_purple": "95",
	"yellow":       "93",
	"white":        "97",
}

// ClearString return the message without escape sequence for ansi color.
func (m Message) ClearString() string {
	var msg strings.Builder
	text, _ := trans(m.Text, false)
	msg.WriteString(text)

	//handle translate
	if m.Translate != "" {
		args := make([]interface{}, len(m.With))
		for i, v := range m.With {
			var arg Message
			_ = arg.UnmarshalJSON(v) //ignore error
			args[i] = arg.ClearString()
		}

		_, _ = fmt.Fprintf(&msg, data.EnUs[m.Translate], args...)
	}

	if m.Extra != nil {
		for i := range m.Extra {
			msg.WriteString(Message(m.Extra[i]).ClearString())
		}
	}
	return msg.String()
}

// String return the message string with escape sequence for ansi color.
// On windows, you may want print this string using
// github.com/mattn/go-colorable.
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

	text, ok := trans(m.Text, true)
	msg.WriteString(text)

	//handle translate
	if m.Translate != "" {
		args := make([]interface{}, len(m.With))
		for i, v := range m.With {
			var arg Message
			_ = arg.UnmarshalJSON(v) //ignore error
			args[i] = arg
		}

		_, _ = fmt.Fprintf(&msg, data.EnUs[m.Translate], args...)
	}

	if m.Extra != nil {
		for i := range m.Extra {
			msg.WriteString(Message(m.Extra[i]).String())
		}
	}

	if format.Len() > 0 || ok {
		msg.WriteString("\033[0m")
	}
	return msg.String()
}

var fmtPat = regexp.MustCompile("(?i)§[0-9A-FK-OR]")

func trans(str string, ansi bool) (dst string, change bool) {
	dst = fmtPat.ReplaceAllStringFunc(
		str,
		func(str string) string {
			f, ok := fmtCode[str[2]]
			if ok {
				if ansi {
					change = true
					return "\033[" + f + "m" // enable, add ANSI code
				}
				return "" //disable, remove the § code
			}
			return str //not a § code
		},
	)
	return
}
