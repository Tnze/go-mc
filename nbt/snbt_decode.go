package nbt

import (
	"fmt"
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
	d.scanWhile(scanSkipSpace)
	switch d.opcode {
	default:
		panic(phasePanicMsg)

	case scanBeginLiteral:
		return writeLiteral(e, &d)
	case scanBeginCompound:
		panic("not implemented")

	case scanBeginList:
		panic("not implemented")
	}
	return nil
}

func writeLiteral(e *Encoder, d *decodeState) error {
	start := d.readIndex()
	d.scanNext()
	d.scanWhile(scanContinue)
	literal := d.data[start:d.readIndex()]
	fmt.Printf("%d %d [%d]- %q\n", start, d.off, d.opcode, literal)

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

func writeCompound(e *Encoder, d *decodeState) error {
	e.writeTag(TagCompound, "")
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
