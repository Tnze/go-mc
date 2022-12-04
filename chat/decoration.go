package chat

import (
	"io"

	pk "github.com/Tnze/go-mc/net/packet"
)

type Decoration struct {
	TranslationKey string   `nbt:"translation_key"`
	Parameters     []string `nbt:"parameters"`
	Style          struct {
		Bold          bool   `nbt:"bold,omitempty"`
		Italic        bool   `nbt:"italic,omitempty"`
		UnderLined    bool   `nbt:"underlined,omitempty"`
		StrikeThrough bool   `nbt:"strikethrough,omitempty"`
		Obfuscated    bool   `nbt:"obfuscated,omitempty"`
		Color         string `nbt:"color,omitempty"`
		Insertion     string `nbt:"insertion,omitempty"`
		Font          string `nbt:"font,omitempty"`
	} `nbt:"style,omitempty"`
}

type Type struct {
	ID         int32
	SenderName Message
	TargetName *Message
}

func (t *Type) Decorate(content Message, d *Decoration) (msg Message) {
	with := make([]Message, len(d.Parameters))
	for i, para := range d.Parameters {
		switch para {
		case "sender":
			with[i] = t.SenderName
		case "target":
			with[i] = *t.TargetName
		case "content":
			with[i] = content
		default:
			with[i] = Text("<nil>")
		}
	}
	return Message{
		Translate: d.TranslationKey,
		With:      with,

		Bold:          d.Style.Bold,
		Italic:        d.Style.Italic,
		UnderLined:    d.Style.UnderLined,
		StrikeThrough: d.Style.StrikeThrough,
		Obfuscated:    d.Style.Obfuscated,
		Font:          d.Style.Font,
		Color:         d.Style.Color,
		Insertion:     d.Style.Insertion,
	}
}

func (t *Type) ReadFrom(r io.Reader) (n int64, err error) {
	var hasTargetName pk.Boolean
	n1, err := (*pk.VarInt)(&t.ID).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := t.SenderName.ReadFrom(r)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := hasTargetName.ReadFrom(r)
	if err != nil {
		return n1 + n2 + n3, err
	}
	if hasTargetName {
		t.TargetName = new(Message)
		n4, err := t.TargetName.ReadFrom(r)
		return n1 + n2 + n3 + n4, err
	}
	return n1 + n2 + n3, nil
}

func (t *Type) WriteTo(w io.Writer) (n int64, err error) {
	hasTargetName := pk.Boolean(t.TargetName != nil)
	n1, err := (*pk.VarInt)(&t.ID).WriteTo(w)
	if err != nil {
		return n1, err
	}
	n2, err := t.SenderName.WriteTo(w)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := hasTargetName.WriteTo(w)
	if err != nil {
		return n1 + n2 + n3, err
	}
	if hasTargetName {
		n4, err := t.TargetName.WriteTo(w)
		return n1 + n2 + n3 + n4, err
	}
	return n1 + n2 + n3, nil
}
