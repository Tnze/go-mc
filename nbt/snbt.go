package nbt

import "io"

type StringifiedNBT string

func (n StringifiedNBT) TagType() (tagType byte) {
	d := decodeState{data: []byte(n)}
	d.scan.reset()
	d.scanWhile(scanSkipSpace)
	switch d.opcode {
	case scanBeginLiteral:
		start := d.readIndex()
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return
		}
		literal := d.data[start:d.readIndex()]
		tagType, _ = parseLiteral(literal)

	case scanBeginCompound:
		tagType = TagCompound

	case scanBeginList:
		d.scanWhile(scanSkipSpace)
		if d.opcode == scanBeginLiteral {
			start := d.readIndex()
			if d.scanWhile(scanContinue); d.opcode == scanError {
				return
			}
			literal := d.data[start:d.readIndex()]
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode == scanListType {
				switch literal[0] {
				case 'B':
					tagType = TagByteArray
				case 'I':
					tagType = TagIntArray
				case 'L':
					tagType = TagLongArray
				}
			}
		} else {
			tagType = TagList
		}
	}
	return
}

func (n StringifiedNBT) Encode(w io.Writer) error {
	d := decodeState{data: []byte(n)}
	d.scan.reset()
	return writeValue(NewEncoder(w), &d, "")
}

//func (n *StringifiedNBT) Decode(tagType byte, r DecoderReader) error {
//}
