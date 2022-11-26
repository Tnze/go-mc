package chat

type Decoration struct {
	TranslationKey string   `nbt:"translation_key"`
	Parameters     []string `nbt:"parameters"`
	Style          struct {
		Bold          bool   `nbt:"bold"`
		Italic        bool   `nbt:"italic"`
		UnderLined    bool   `nbt:"underlined"`
		StrikeThrough bool   `nbt:"strikethrough"`
		Obfuscated    bool   `nbt:"obfuscated"`
		Color         string `nbt:"color"`
		Insertion     string `nbt:"insertion"`
		Font          string `nbt:"font"`
	} `nbt:"style"`
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
