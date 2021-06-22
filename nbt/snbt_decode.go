package nbt

import (
	"bytes"
	"math"
	"strconv"
	"strings"
)

type decodeState struct {
	data   []byte
	off    int // next read Offset in data
	opcode int // last read result
	scan   scanner
}

const phasePanicMsg = "SNBT decoder out of sync - data changing underfoot?"

func (e *Encoder) WriteSNBT(snbt string) error {
	d := decodeState{data: []byte(snbt)}
	d.scan.reset()
	return writeValue(e, &d, "")
}

func writeValue(e *Encoder, d *decodeState, tagName string) error {
	d.scanWhile(scanSkipSpace)
	switch d.opcode {
	case scanError:
		return d.error(d.scan.errContext)
	default:
		panic(phasePanicMsg)

	case scanBeginLiteral:
		start := d.readIndex()
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		literal := d.data[start:d.readIndex()]
		tagType, litVal := parseLiteral(literal)
		e.writeTag(tagType, tagName)
		return writeLiteralPayload(e, litVal)

	case scanBeginCompound:
		e.writeTag(TagCompound, tagName)
		return writeCompoundPayload(e, d)

	case scanBeginList:
		_, err := writeListOrArray(e, d, true, tagName)
		return err
	}
}

func writeLiteralPayload(e *Encoder, v interface{}) error {
	switch v.(type) {
	case string:
		str := v.(string)
		e.writeInt16(int16(len(str)))
		e.w.Write([]byte(str))
	case int8:
		e.w.Write([]byte{byte(v.(int8))})
	case int16:
		e.writeInt16(v.(int16))
	case int32:
		e.writeInt32(v.(int32))
	case int64:
		e.writeInt64(v.(int64))
	case float32:
		e.writeInt32(int32(math.Float32bits(v.(float32))))
	case float64:
		e.writeInt64(int64(math.Float64bits(v.(float64))))
	}
	return nil
}

func writeCompoundPayload(e *Encoder, d *decodeState) error {
	defer d.scanNext()
	for {
		d.scanWhile(scanSkipSpace)
		if d.opcode == scanEndValue {
			break
		}
		if d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		if d.opcode != scanBeginLiteral {
			panic(phasePanicMsg)
		}
		// read tag name
		start := d.readIndex()
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		var tagName string
		if tt, v := parseLiteral(d.data[start:d.readIndex()]); tt == TagString {
			tagName = v.(string)
		} else {
			tagName = string(d.data[start:d.readIndex()])
		}
		// read value
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		if d.opcode != scanCompoundTagName {
			panic(phasePanicMsg)
		}
		writeValue(e, d, tagName)

		// Next token must be , or }.
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		if d.opcode == scanEndValue {
			break
		}
		if d.opcode != scanCompoundValue {
			panic(phasePanicMsg)
		}
	}
	e.w.Write([]byte{TagEnd})
	return nil
}

func writeListOrArray(e *Encoder, d *decodeState, writeTag bool, tagName string) (tagType byte, err error) {
	d.scanWhile(scanSkipSpace)
	if d.opcode == scanEndValue { // ']', empty TAG_List
		e.writeListHeader(TagEnd, tagName, 0, writeTag)
		d.scanNext()
		return TagList, nil
	}

	// We don't know the length of the List,
	// so we read them into a buffer and count.
	var buf bytes.Buffer
	var count int
	e2 := NewEncoder(&buf)
	start := d.readIndex()

	switch d.opcode {
	case scanBeginLiteral:
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return TagList, d.error(d.scan.errContext)
		}
		literal := d.data[start:d.readIndex()]
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode == scanError {
			return tagType, d.error(d.scan.errContext)
		}
		if d.opcode == scanListType { // TAG_X_Array
			var elemType byte
			switch literal[0] {
			case 'B':
				tagType = TagByteArray
				elemType = TagByte
			case 'I':
				tagType = TagIntArray
				elemType = TagInt
			case 'L':
				tagType = TagLongArray
				elemType = TagLong
			default:
				return TagList, d.error("unknown Array type")
			}
			if writeTag {
				e.writeTag(tagType, tagName)
			}
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			d.scanWhile(scanSkipSpace)    // ;
			if d.opcode == scanEndValue { // ]
				// empty array
				e.writeInt32(0)
				break
			}
			for {
				if d.opcode == scanSkipSpace {
					d.scanWhile(scanSkipSpace)
				}
				if d.opcode != scanBeginLiteral {
					return tagType, d.error("not literal in Array")
				}
				start := d.readIndex()

				if d.scanWhile(scanContinue); d.opcode == scanError {
					return tagType, d.error(d.scan.errContext)
				}
				literal := d.data[start:d.readIndex()]
				tagType, litVal := parseLiteral(literal)
				if tagType != elemType {
					return tagType, d.error("unexpected element type in TAG_Array")
				}
				switch elemType {
				case TagByte:
					e2.w.Write([]byte{byte(litVal.(int8))})
				case TagInt:
					e2.writeInt32(litVal.(int32))
				case TagLong:
					e2.writeInt64(litVal.(int64))
				}
				count++

				if d.opcode == scanSkipSpace {
					d.scanWhile(scanSkipSpace)
				}
				if d.opcode == scanError {
					return tagType, d.error(d.scan.errContext)
				}
				if d.opcode == scanEndValue { // ]
					break
				}
				if d.opcode != scanListValue {
					panic(phasePanicMsg)
				}
				d.scanWhile(scanSkipSpace) // ,
			}
			e.writeInt32(int32(count))
			e.w.Write(buf.Bytes())
			break
		}
		if d.opcode != scanListValue { // TAG_List<TAG_String>
			panic(phasePanicMsg)
		}
		var tagType byte
		for {
			t, v := parseLiteral(literal)
			if tagType == 0 {
				tagType = t
			}
			if t != tagType {
				return TagList, d.error("different TagType in List")
			}
			writeLiteralPayload(e2, v)
			count++

			// read ',' or ']'
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode == scanError {
				return tagType, d.error(d.scan.errContext)
			}
			if d.opcode == scanEndValue {
				break
			}
			if d.opcode != scanListValue {
				panic(phasePanicMsg)
			}
			d.scanWhile(scanSkipSpace)
			start = d.readIndex()
			if d.scanWhile(scanContinue); d.opcode == scanError {
				return tagType, d.error(d.scan.errContext)
			}
			literal = d.data[start:d.readIndex()]
		}
		e.writeListHeader(tagType, tagName, count, writeTag)
		e.w.Write(buf.Bytes())
	case scanBeginList: // TAG_List<TAG_List>
		var elemType byte
		for {
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode != scanBeginList {
				return TagList, d.error("different TagType in List")
			}
			elemType, err = writeListOrArray(e2, d, false, "")
			if err != nil {
				return tagType, err
			}
			count++
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode == scanError {
				return tagType, d.error(d.scan.errContext)
			}
			// ',' or ']'
			if d.opcode == scanEndValue {
				break
			}
			if d.opcode != scanListValue {
				panic(phasePanicMsg)
			}
			// read '['
			d.scanNext()
		}
		e.writeListHeader(elemType, tagName, count, writeTag)
		e.w.Write(buf.Bytes())
	case scanBeginCompound: // TAG_List<TAG_Compound>
		for {
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode != scanBeginCompound {
				return TagList, d.error("different TagType in List")
			}
			writeCompoundPayload(e2, d)
			count++
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			// read ',' or ']'
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode == scanError {
				return tagType, d.error(d.scan.errContext)
			}
			if d.opcode == scanEndValue {
				break
			}
			if d.opcode != scanListValue {
				panic(phasePanicMsg)
			}
			// read '{'
			d.scanNext()
		}
		e.writeListHeader(TagCompound, tagName, count, writeTag)
		e.w.Write(buf.Bytes())
	}
	d.scanNext()
	return
}

// readIndex returns the position of the last byte read.
func (d *decodeState) readIndex() int {
	return d.off - 1
}

// scanNext processes the byte at d.data[d.off].
func (d *decodeState) scanNext() {
	if d.off < len(d.data) {
		d.opcode = d.scan.step(&d.scan, d.data[d.off])
		d.off++
	} else {
		//d.opcode = d.scan.eof()
		d.off = len(d.data) + 1 // mark processed EOF with len+1
	}
}

// scanWhile processes bytes in d.data[d.off:] until it
// receives a scan code not equal to op.
func (d *decodeState) scanWhile(op int) {
	s, data, i := &d.scan, d.data, d.off
	for i < len(data) {
		newOp := s.step(s, data[i])
		i++
		if newOp != op {
			d.opcode = newOp
			d.off = i
			return
		}
	}

	d.off = len(data) + 1 // mark processed EOF with len+1
	d.opcode = d.scan.eof()
}

// parseLiteral parse an SNBT literal, might be
// TAG_String, TAG_Int, TAG_Float, ... etc.
// so returned value is one of string, int32, float32 ...
func parseLiteral(literal []byte) (byte, interface{}) {
	switch literal[0] {
	case '"', '\'': // Quoted String
		var sb strings.Builder
		sb.Grow(len(literal) - 2)
		for i := 1; ; i++ {
			c := literal[i]
			switch c {
			case literal[0]:
				return TagString, sb.String()
			case '\\':
				i++
				c = literal[i]
			}
			sb.WriteByte(c)
		}
	default:
		strlen := len(literal)
		integer := true
		number := true
		unqstr := true
		var numberType byte

		for i, c := range literal {
			if isNumber(c) {
				continue
			} else if integer {
				if i == strlen-1 && i != 0 && isIntegerType(c) {
					numberType = c
					strlen--
				} else if i > 0 || i == 0 && c != '-' {
					integer = false
					if i == 0 || c != '.' {
						number = false
					}
				}
			} else if number {
				if i == strlen-1 && isFloatType(c) {
					numberType = c
				} else {
					number = false
				}
			} else if !isAllowedInUnquotedString(c) {
				unqstr = false
			}
		}
		if integer {
			num, err := strconv.ParseInt(string(literal[:strlen]), 10, 64)
			if err != nil {
				panic(err)
			}
			switch numberType {
			case 'B', 'b':
				return TagByte, int8(num)
			case 'S', 's':
				return TagShort, int16(num)
			default:
				return TagInt, int32(num)
			case 'L', 'l':
				return TagLong, num
			case 'F', 'f':
				return TagFloat, float32(num)
			case 'D', 'd':
				return TagDouble, float64(num)
			}
		} else if number {
			num, err := strconv.ParseFloat(string(literal[:strlen-1]), 64)
			if err != nil {
				panic(err)
			}
			switch numberType {
			case 'F', 'f':
				return TagFloat, float32(num)
			case 'D', 'd':
				fallthrough
			default:
				return TagDouble, num
			}
		} else if unqstr {
			return TagString, string(literal)
		}
	}
	panic(phasePanicMsg)
}

func (d *decodeState) error(msg string) *SyntaxError {
	return &SyntaxError{Message: msg, Offset: d.off}
}

func isIntegerType(c byte) bool {
	return isFloatType(c) ||
		c == 'B' || c == 'b' ||
		c == 's' || c == 'S' ||
		c == 'L' || c == 'l'
}

func isFloatType(c byte) bool {
	return c == 'F' || c == 'f' || c == 'D' || c == 'd'
}

type SyntaxError struct {
	Message string
	Offset  int
}

func (e *SyntaxError) Error() string { return e.Message }
