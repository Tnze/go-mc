package nbt

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

	switch literal[0] {
	case '"', '\'': // TAG_String
		str := literal // TODO: Parse string
		e.writeTag(TagString, "")
		e.writeInt16(int16(len(str)))
		e.w.Write(str)

	default:
		e.w.Write(literal) // TODO: Parse other literal
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
		d.scanWhile(scanSkipSpace)
		writeLiteral(e, d, tagName)

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
