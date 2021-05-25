package nbt

import (
	"math"
	"strconv"
	"strings"
)

type decodeState struct {
	data   []byte
	off    int // next read offset in data
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
	default:
		panic(phasePanicMsg)
	case scanBeginLiteral:
		return writeLiteral(e, d, tagName)
	case scanBeginCompound:
		return writeCompound(e, d, tagName)
	case scanBeginList:
		panic("not implemented")
	}
}

func writeLiteral(e *Encoder, d *decodeState, tagName string) error {
	start := d.readIndex()
	d.scanWhile(scanContinue)
	literal := d.data[start:d.readIndex()]
	switch v := parseLiteral(literal); v.(type) {
	case string:
		str := v.(string)
		e.writeTag(TagString, tagName)
		e.writeInt16(int16(len(str)))
		e.w.Write([]byte(str))
	case int8:
		e.writeTag(TagByte, tagName)
		e.w.Write([]byte{byte(v.(int8))})
	case int16:
		e.writeTag(TagShort, tagName)
		e.writeInt16(v.(int16))
	case int32:
		e.writeTag(TagInt, tagName)
		e.writeInt32(v.(int32))
	case int64:
		e.writeTag(TagLong, tagName)
		e.writeInt64(v.(int64))
	case float32:
		e.writeTag(TagFloat, tagName)
		e.writeInt32(int32(math.Float32bits(v.(float32))))
	case float64:
		e.writeTag(TagDouble, tagName)
		e.writeInt64(int64(math.Float64bits(v.(float64))))
	}
	return nil
}

func writeCompound(e *Encoder, d *decodeState, tagName string) error {
	e.writeTag(TagCompound, tagName)
	for {
		d.scanWhile(scanSkipSpace)
		if d.opcode == scanEndValue {
			break
		}
		if d.opcode != scanBeginLiteral {
			panic(phasePanicMsg)
		}
		// read tag name
		start := d.readIndex()
		d.scanWhile(scanContinue)
		tagName := string(d.data[start:d.readIndex()])
		// read value
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
		}
		if d.opcode != scanCompoundTagName {
			panic(phasePanicMsg)
		}
		writeValue(e, d, tagName)

		// Next token must be , or }.
		if d.opcode == scanSkipSpace {
			d.scanWhile(scanSkipSpace)
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

// readIndex returns the position of the last byte read.
func (d *decodeState) readIndex() int {
	return d.off - 1
}

// scanNext processes the byte at d.data[d.off].
func (d *decodeState) scanNext() {
	if d.off < len(d.data) {
		d.opcode = d.scan.step(d.data[d.off])
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
		newOp := s.step(data[i])
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
func parseLiteral(literal []byte) interface{} {
	switch literal[0] {
	case '"', '\'': // Quoted String
		var sb strings.Builder
		sb.Grow(len(literal) - 2)
		for i := 1; ; i++ {
			c := literal[i]
			switch c {
			case literal[0]:
				return sb.String()
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
				if i == strlen-1 && isIntegerType(c) {
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
				return int8(num)
			case 'S', 's':
				return int16(num)
			default:
				return int32(num)
			case 'L', 'l':
				return num
			case 'F', 'f':
				return float32(num)
			case 'D', 'd':
				return float64(num)
			}
		} else if number {
			num, err := strconv.ParseFloat(string(literal[:strlen-1]), 64)
			if err != nil {
				panic(err)
			}
			switch numberType {
			case 'F', 'f':
				return float32(num)
			case 'D', 'd':
				fallthrough
			default:
				return num
			}
		} else if unqstr {
			return string(literal)
		}
	}
	panic(phasePanicMsg)
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
