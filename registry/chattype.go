package registry

import "github.com/Tnze/go-mc/chat"

type ChatType struct {
	Chat      chat.Decoration `nbt:"chat"`
	Narration chat.Decoration `nbt:"narration"`
}
