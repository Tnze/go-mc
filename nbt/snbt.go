package nbt

type StringifiedNBT struct {
	Name    string
	Content string
}

func (n *StringifiedNBT) Decode(tagType byte, tagName string, r DecoderReader) error {
	panic("unimplemented")
}
