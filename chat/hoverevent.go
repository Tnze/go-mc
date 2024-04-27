package chat

// HoverEvent defines an event that occurs when this component hovered over.
type HoverEvent struct {
	Action   string  `json:"action" nbt:"action"`
	Contents any     `json:"contents" nbt:"contents"` // Didn't handled yet
	Value    Message `json:"value" nbt:"value"`       // Legacy
}

type HoverSub struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

// ShowText show the text to display.
func ShowText(text Message) *HoverEvent {
	return &HoverEvent{
		Action: "show_text",
		Value:  text,
	}
}

// ShowItem show the item to display.
// Item is encoded as the S-NBT format, nbt.StringifiedMessage could help.
// See: https://wiki.vg/Chat#:~:text=show_item,in%20red%20instead.
func ShowItem(item string) *HoverEvent {
	return &HoverEvent{
		Action: "show_item",
		Value:  Text(item),
	}
}

// ShowEntity show an entity describing by the S-NBT, nbt.StringifiedMessage could help.
// See: https://wiki.vg/Chat#:~:text=show_entity,given%20entity%20loaded.
func ShowEntity(entity string) *HoverEvent {
	return &HoverEvent{
		Action: "show_entity",
		Value:  Text(entity),
	}
}
