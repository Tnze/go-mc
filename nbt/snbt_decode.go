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

func writeValue(e *Encoder, d *decodeState, ifWriteTag bool, tagName string) error {
	d.scanWhile(scanSkipSpace)
	switch d.opcode {
	case scanError:
		return d.error(d.scan.errContext)

	case scanBeginLiteral:
		start := d.readIndex()
		if d.scanWhile(scanContinue); d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		literal := d.data[start:d.readIndex()]
		tagType, litVal, err := parseLiteral(literal)
		if err != nil {
			return err
		}
		if ifWriteTag {
			if err := writeTag(e.w, tagType, tagName); err != nil {
				return err
			}
		}
		return writeLiteralPayload(e, litVal)

	case scanBeginCompound:
		if ifWriteTag {
			if err := writeTag(e.w, TagCompound, tagName); err != nil {
				return err
			}
		}
		return writeCompoundPayload(e, d)

	case scanBeginList:
		_, err := writeListOrArray(e, d, ifWriteTag, tagName)
		return err

	default:
		panic(phasePanicMsg)
	}
}

func writeLiteralPayload(e *Encoder, v any) (err error) {
	switch v := v.(type) {
	case string:
		err = writeInt16(e.w, int16(len(v)))
		if err != nil {
			return
		}
		_, err = e.w.Write([]byte(v))
	case int8:
		_, err = e.w.Write([]byte{byte(v)})
	case int16:
		err = writeInt16(e.w, v)
	case int32:
		err = writeInt32(e.w, v)
	case int64:
		err = writeInt64(e.w, v)
	case float32:
		err = writeInt32(e.w, int32(math.Float32bits(v)))
	case float64:
		err = writeInt64(e.w, int64(math.Float64bits(v)))
	}
	return
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
		if tt, v, err := parseLiteral(d.data[start:d.readIndex()]); err != nil {
			return err
		} else if tt == TagString {
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

		if err := writeValue(e, d, true, tagName); err != nil {
			return err
		}

		// The next token must be , or }.
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
	_, err := e.w.Write([]byte{TagEnd})
	return err
}

func writeListOrArray(e *Encoder, d *decodeState, ifWriteTag bool, tagName string) (tagType byte, err error) {
	d.scanWhile(scanSkipSpace)
	if d.opcode == scanEndValue { // ']', empty TAG_List
		if ifWriteTag {
			err = writeTag(e.w, TagList, tagName)
			if err != nil {
				return tagType, err
			}
		}
		err = e.writeListHeader(TagEnd, 0)
		d.scanNext()
		return TagList, err
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
			if ifWriteTag {
				err = writeTag(e.w, tagType, tagName)
				if err != nil {
					return tagType, err
				}
			}
			err = writeArray(e, e2, d, elemType, &count, &buf)
			if err != nil {
				return tagType, err
			}
			break
		}
		if ifWriteTag {
			err = writeTag(e.w, TagList, tagName)
			if err != nil {
				return tagType, err
			}
		}
		if d.opcode != scanListValue && d.opcode != scanEndValue { // TAG_List<TAG_String>
			panic(phasePanicMsg)
		}
		var tagType byte
		for {
			t, v, err := parseLiteral(literal)
			if err != nil {
				return tagType, err
			}
			if tagType == 0 {
				tagType = t
			}
			if t != tagType {
				return TagList, d.error("different TagType in List")
			}
			err = writeLiteralPayload(e2, v)
			if err != nil {
				return tagType, err
			}
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

		if err := e.writeListHeader(tagType, count); err != nil {
			return tagType, err
		}
		if _, err := e.w.Write(buf.Bytes()); err != nil {
			return tagType, err
		}
	case scanBeginList: // TAG_List<TAG_List>
		if ifWriteTag {
			err = writeTag(e.w, TagList, tagName)
			if err != nil {
				return tagType, err
			}
		}
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

		if err = e.writeListHeader(elemType, count); err != nil {
			return
		}
		if _, err = e.w.Write(buf.Bytes()); err != nil {
			return
		}
	case scanBeginCompound: // TAG_List<TAG_Compound>
		if ifWriteTag {
			err = writeTag(e.w, TagList, tagName)
			if err != nil {
				return tagType, err
			}
		}
		for {
			if d.opcode == scanSkipSpace {
				d.scanWhile(scanSkipSpace)
			}
			if d.opcode != scanBeginCompound {
				return TagList, d.error("different TagType in List")
			}

			if err = writeCompoundPayload(e2, d); err != nil {
				return
			}
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

		if err = e.writeListHeader(TagCompound, count); err != nil {
			return
		}

		if _, err = e.w.Write(buf.Bytes()); err != nil {
			return
		}
	}
	d.scanNext()
	return
}

func writeArray(e, e2 *Encoder, d *decodeState, elemType byte, count *int, buf *bytes.Buffer) (err error) {
	if d.opcode == scanSkipSpace {
		d.scanWhile(scanSkipSpace)
	}
	d.scanWhile(scanSkipSpace)    // ;
	if d.opcode == scanEndValue { // ]
		// empty array
		if err = writeInt32(e.w, 0); err != nil {
			return
		}
		return
	}
	for {
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode != scanBeginLiteral {
			return d.error("not literal in Array")
		}
		start := d.readIndex()

		if d.scanWhile(scanContinue); d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		literal := d.data[start:d.readIndex()]
		var subType byte
		var litVal any
		subType, litVal, err = parseLiteral(literal)
		if err != nil {
			return err
		}
		if subType != elemType {
			err = d.error("unexpected element type in TAG_Array")
			return
		}
		switch elemType {
		case TagByte:
			_, err = e2.w.Write([]byte{byte(litVal.(int8))})
		case TagInt:
			err = writeInt32(e2.w, litVal.(int32))
		case TagLong:
			err = writeInt64(e2.w, litVal.(int64))
		}
		if err != nil {
			return
		}
		*count++

		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode == scanError {
			return d.error(d.scan.errContext)
		}
		if d.opcode == scanEndValue { // ]
			break
		}
		if d.opcode != scanListValue {
			panic(phasePanicMsg)
		}
		d.scanWhile(scanSkipSpace) // ,
	}

	if err = writeInt32(e.w, int32(*count)); err != nil {
		return err
	}
	_, err = e.w.Write(buf.Bytes())
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
		// d.opcode = d.scan.eof()
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
// so the returned value is one of string, int32, float32 ...
func parseLiteral(literal []byte) (byte, any, error) {
	switch literal[0] {
	case '"', '\'': // Quoted String
		var sb strings.Builder
		sb.Grow(len(literal) - 2)
		for i := 1; ; i++ {
			c := literal[i]
			switch c {
			case literal[0]:
				return TagString, sb.String(), nil
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
		hasExp := false
		afterExp := false
		unqstr := true
		var numberType byte

		for i, c := range literal {
			if isNumber(c) {
				continue
			} else if integer {
				if i == strlen-1 && i != 0 && isIntegerType(c) {
					numberType = c
					strlen--
				} else if i > 0 || i == 0 && c != '-' && c != '+' {
					integer = false
					if i == 0 || c != '.' {
						number = false
					}
				}
			} else if number {
				if hasExp && !afterExp && c == '-' || c == '+' {
					afterExp = true
				} else if c == 'E' || c == 'e' {
					hasExp = true
				} else if i == strlen-1 && isFloatType(c) {
					numberType = c
				} else {
					number = false
				}
			} else if !isAllowedInUnquotedString(c) {
				unqstr = false
			}
		}
		if integer {
			switch numberType {
			case 'B', 'b':
				num, err := strconv.ParseInt(string(literal[:strlen]), 10, 8)
				return TagByte, int8(num), err
			case 'S', 's':
				num, err := strconv.ParseInt(string(literal[:strlen]), 10, 16)
				return TagShort, int16(num), err
			case 'I', 'i', 0:
				num, err := strconv.ParseInt(string(literal[:strlen]), 10, 32)
				return TagInt, int32(num), err
			case 'L', 'l':
				num, err := strconv.ParseInt(string(literal[:strlen]), 10, 64)
				return TagLong, num, err
			case 'F', 'f':
				num, err := strconv.ParseFloat(string(literal[:strlen]), 32)
				return TagFloat, float32(num), err
			case 'D', 'd':
				num, err := strconv.ParseFloat(string(literal[:strlen]), 64)
				return TagDouble, num, err
			}
		} else if number {
			switch numberType {
			case 'F', 'f':
				num, err := strconv.ParseFloat(string(literal[:strlen-1]), 64)
				return TagFloat, float32(num), err
			case 'D', 'd':
				fallthrough
			default:
				num, err := strconv.ParseFloat(string(literal[:strlen-1]), 64)
				return TagDouble, num, err
			}
		} else if unqstr {
			return TagString, string(literal), nil
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
