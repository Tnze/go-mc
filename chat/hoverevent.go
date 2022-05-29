package chat

// HoverEvent defines an event that occurs when this component hovered over.
type HoverEvent struct {
	Action string     `json:"action"`
	Value  []HoverSub `json:"value"` // The issue i found is the fact the json is different now, it wasnt parsing properly because Message was invalid, its a listed dict.
}

type HoverSub struct {
	Color string `json:"color"`
	Text  string `json:"text"`
}

// ShowText show the text to display.
func ShowText(text []HoverSub) *HoverEvent {
	return &HoverEvent{
		Action: "show_text",
		Value:  text,
	}
}

// ShowItem show the item to display.
// Item is encoded as the S-NBT format, nbt.StringifiedMessage could help.
// See: https://wiki.vg/Chat#:~:text=show_item,in%20red%20instead.
func ShowItem(item string) *HoverEvent {
	T := Text(item)
	return &HoverEvent{
		Action: "show_item",
		Value: []HoverSub{
			{
				Color: T.Color,
				Text:  T.Text,
			},
		},
	}
}

// ShowEntity show an entity describing by the S-NBT, nbt.StringifiedMessage could help.
// See: https://wiki.vg/Chat#:~:text=show_entity,given%20entity%20loaded.
func ShowEntity(entity string) *HoverEvent {
	T := Text(entity)
	return &HoverEvent{
		Action: "show_entity",
		Value: []HoverSub{
			{
				Color: T.Color,
				Text:  T.Color,
			},
		},
	}
}
