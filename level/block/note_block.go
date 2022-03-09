package block

type NoteBlock struct {
	Instrument string
	Note       string
	Powered    string
}

func (NoteBlock) ID() string {
	return "minecraft:note_block"
}
