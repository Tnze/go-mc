package nbt

type Unmarshaler interface {
	Unmarshal(tagType byte, tagName string, r DecoderReader) error
}

//type Marshaler interface{
//	Marshal()
//}
