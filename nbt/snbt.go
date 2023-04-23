package nbt

import (
	"fmt"
	"io"
	"math"
	"strconv"
	"strings"
)

type StringifiedMessage string

func (m StringifiedMessage) TagType() byte {
	d := decodeState{data: []byte(m)}
	d.scan.reset()
	d.scanWhile(scanSkipSpace)
	switch d.opcode {
	default:
		return TagEnd

	case scanBeginLiteral:
		start := d.readIndex()
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return TagEnd
		}
		literal := d.data[start:d.readIndex()]
		tagType, _, _ := parseLiteral(literal)
		return tagType

	case scanBeginCompound:
		return TagCompound

	case scanBeginList:
		d.scanWhile(scanSkipSpace)
		if d.opcode == scanBeginLiteral {
			start := d.readIndex()
			if d.scanWhile(scanContinue); d.opcode == scanError {
				return TagEnd
			}
			literal := d.data[start:d.readIndex()]
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode == scanListType {
				switch literal[0] {
				case 'B':
					return TagByteArray
				case 'I':
					return TagIntArray
				case 'L':
					return TagLongArray
				}
			}
		}
		return TagList
	}
}

func (m StringifiedMessage) MarshalNBT(w io.Writer) error {
	d := decodeState{data: []byte(m)}
	d.scan.reset()
	return writeValue(NewEncoder(w), &d, false, "")
}

func (m *StringifiedMessage) UnmarshalNBT(tagType byte, r DecoderReader) error {
	if tagType == TagEnd {
		return ErrEND
	}
	var sb strings.Builder
	d := NewDecoder(r)
	err := m.encode(d, &sb, tagType)
	if err != nil {
		return err
	}
	*m = StringifiedMessage(sb.String())
	return nil
}

func (m *StringifiedMessage) encode(d *Decoder, sb *strings.Builder, tagType byte) error {
	switch tagType {
	default:
		return fmt.Errorf("unknown to read 0x%02x", tagType)
	case TagByte:
		b, err := d.r.ReadByte()
		sb.WriteString(strconv.FormatInt(int64(b), 10) + "B")
		return err
	case TagString:
		str, err := d.readString()
		writeEscapeStr(sb, str)
		return err
	case TagShort:
		s, err := d.readInt16()
		sb.WriteString(strconv.FormatInt(int64(s), 10) + "S")
		return err
	case TagInt:
		i, err := d.readInt32()
		sb.WriteString(strconv.FormatInt(int64(i), 10))
		return err
	case TagFloat:
		i, err := d.readInt32()
		f := float64(math.Float32frombits(uint32(i)))
		sb.WriteString(strconv.FormatFloat(f, 'f', 10, 32) + "F")
		return err
	case TagLong:
		i, err := d.readInt64()
		sb.WriteString(strconv.FormatInt(i, 10) + "L")
		return err
	case TagDouble:
		i, err := d.readInt64()
		f := math.Float64frombits(uint64(i))
		sb.WriteString(strconv.FormatFloat(f, 'f', 10, 64) + "D")
		return err
	case TagByteArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		first := true
		sb.WriteString("[B;")
		for i := int32(0); i < aryLen; i++ {
			b, err := d.r.ReadByte()
			if err != nil {
				return err
			}
			if first {
				first = false
			} else {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.FormatInt(int64(b), 10) + "B")
		}
		sb.WriteString("]")
	case TagIntArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		sb.WriteString("[I;")
		first := true
		for i := 0; i < int(aryLen); i++ {
			v, err := d.readInt32()
			if err != nil {
				return err
			}
			if first {
				first = false
			} else {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.FormatInt(int64(v), 10) + "I")
		}
		sb.WriteString("]")
	case TagLongArray:
		aryLen, err := d.readInt32()
		if err != nil {
			return err
		}
		first := true
		sb.WriteString("[L;")
		for i := 0; i < int(aryLen); i++ {
			v, err := d.readInt64()
			if err != nil {
				return err
			}
			if first {
				first = false
			} else {
				sb.WriteString(",")
			}
			sb.WriteString(strconv.FormatInt(v, 10) + "L")
		}
		sb.WriteString("]")
	case TagList:
		listType, err := d.r.ReadByte()
		if err != nil {
			return err
		}
		listLen, err := d.readInt32()
		if err != nil {
			return err
		}
		first := true
		sb.WriteString("[")
		for i := 0; i < int(listLen); i++ {
			if first {
				first = false
			} else {
				sb.WriteString(",")
			}
			if err := m.encode(d, sb, listType); err != nil {
				return err
			}
		}
		sb.WriteString("]")
	case TagCompound:
		first := true
		for {
			tt, tn, err := d.readTag()
			if err != nil {
				return err
			}
			if first {
				sb.WriteString("{")
				first = false
			} else if tt != TagEnd {
				sb.WriteString(",")
			}
			if tt == TagEnd {
				sb.WriteString("}")
				break
			}

			writeEscapeStr(sb, tn)
			sb.WriteString(":")
			err = m.encode(d, sb, tt)
			if err != nil {
				return err
			}
		}
	}
	return nil
}

func writeEscapeStr(sb *strings.Builder, str string) {
	for _, v := range []byte(str) {
		if !isAllowedInUnquotedString(v) {
			// need quote
			dc := strings.Count(str, `"`)
			sc := strings.Count(str, `'`)
			if dc > sc {
				sb.WriteString("'")
				if _, err := strings.NewReplacer(`'`, `\'`, `\`, `\\`).WriteString(sb, str); err != nil {
					panic(err)
				}
				sb.WriteString("'")
			} else {
				sb.WriteString(`"`)
				if _, err := strings.NewReplacer(`"`, `\"`, `\`, `\\`).WriteString(sb, str); err != nil {
					panic(err)
				}
				sb.WriteString(`"`)
			}
			return
		}
	}
	sb.WriteString(str)
}
