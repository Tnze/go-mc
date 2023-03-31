package packet

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
	"sync"
)

const MaxDataLength = 2097152

// Packet define a net data package
type Packet struct {
	ID   int32
	Data []byte
}

// Marshal generate Packet with the ID and Fields
func Marshal[ID ~int32 | int](id ID, fields ...FieldEncoder) (pk Packet) {
	var pb Builder
	for _, v := range fields {
		pb.WriteField(v)
	}
	return pb.Packet(int32(id))
}

// Scan decode the packet and fill data into fields
func (p Packet) Scan(fields ...FieldDecoder) error {
	r := bytes.NewReader(p.Data)
	for i, v := range fields {
		_, err := v.ReadFrom(r)
		if err != nil {
			return fmt.Errorf("scanning packet field[%d] error: %w", i, err)
		}
	}
	return nil
}

var bufPool = sync.Pool{
	New: func() interface{} {
		return new(bytes.Buffer)
	},
}

// Pack 打包一个数据包
func (p *Packet) Pack(w io.Writer, threshold int) error {
	if threshold >= 0 {
		return p.packWithCompression(w, threshold)
	} else {
		return p.packWithoutCompression(w)
	}
}

func (p *Packet) packWithoutCompression(w io.Writer) error {
	buffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buffer)

	// Pre-allocate room at the front of the packet for the length field
	buffer.Reset()
	buffer.Write([]byte{0, 0, 0})

	VarInt(p.ID).WriteTo(buffer)
	buffer.Write(p.Data)

	// Write length at front
	payloadLen := VarInt(buffer.Len() - 3)
	varIntOffset := 3 - payloadLen.Len()
	payloadLen.WriteToBytes(buffer.Bytes()[varIntOffset:])

	_, err := w.Write(buffer.Bytes()[varIntOffset:])
	return err
}

func (p *Packet) packWithCompression(w io.Writer, threshold int) error {
	buff := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buff)
	// Allocate room for the 'packet length' and 'data length' fields. Each can take up to 3 bytes
	buff.Reset()
	buff.Write([]byte{0, 0, 0, 0, 0, 0})

	var writeStart int

	if len(p.Data) < threshold {
		VarInt(p.ID).WriteTo(buff)
		buff.Write(p.Data)
		// Packet is below compression threshold so 'data length' is 0
		// Front of the packet is already initialized to 0, so just decrement the offset
		writeStart = 5
	} else {
		zw := zlib.NewWriter(buff)
		varIntLen, _ := VarInt(p.ID).WriteTo(zw)
		zw.Write(p.Data)

		err := zw.Close()
		if err != nil {
			return err
		}

		// Write 'data length' before ID + payload
		uncompressedLen := VarInt(varIntLen + int64(len(p.Data)))
		writeStart = 6 - uncompressedLen.Len()
		uncompressedLen.WriteToBytes(buff.Bytes()[writeStart:])
	}

	// Write 'packet length' before all other fields
	packetLen := VarInt(buff.Len() - writeStart)
	start := writeStart - packetLen.Len()
	VarInt(packetLen).WriteToBytes(buff.Bytes()[start:])

	_, err := w.Write(buff.Bytes()[start:])
	return err
}

// UnPack in-place decompression a packet
func (p *Packet) UnPack(r io.Reader, threshold int) error {
	if threshold >= 0 {
		return p.unpackWithCompression(r, threshold)
	} else {
		return p.unpackWithoutCompression(r)
	}
}

func (p *Packet) unpackWithoutCompression(r io.Reader) error {
	var Length VarInt
	_, err := Length.ReadFrom(r)
	if err != nil {
		return err
	}

	var PacketID VarInt
	n, err := PacketID.ReadFrom(r)
	if err != nil {
		return err
	}
	p.ID = int32(PacketID)

	lengthOfData := int(Length) - int(n)
	if lengthOfData < 0 || lengthOfData > MaxDataLength {
		return fmt.Errorf("uncompressed packet error: length is %d", lengthOfData)
	}
	if cap(p.Data) < lengthOfData {
		p.Data = make([]byte, lengthOfData)
	} else {
		p.Data = p.Data[:lengthOfData]
	}
	_, err = io.ReadFull(r, p.Data)
	if err != nil {
		return err
	}
	return nil
}

func (p *Packet) unpackWithCompression(r io.Reader, threshold int) error {
	var PacketLength VarInt
	_, err := PacketLength.ReadFrom(r)
	if err != nil {
		return err
	}

	buff := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buff)
	buff.Reset()

	_, err = io.CopyN(buff, r, int64(PacketLength))
	if err != nil {
		return err
	}
	r = bytes.NewReader(buff.Bytes())

	var DataLength VarInt
	n2, err := DataLength.ReadFrom(r)
	if err != nil {
		return err
	}

	var PacketID VarInt
	if DataLength != 0 {
		if int(DataLength) < threshold {
			return fmt.Errorf("compressed packet error: size of %d is below threshold of %d", DataLength, threshold)
		}
		if DataLength > MaxDataLength {
			return fmt.Errorf("compressed packet error: size of %d is larger than protocol maximum of %d", DataLength, MaxDataLength)
		}
		zr, err := zlib.NewReader(r)
		if err != nil {
			return err
		}
		defer zr.Close()
		r = zr
		n3, err := PacketID.ReadFrom(r)
		if err != nil {
			return err
		}
		DataLength -= VarInt(n3)
	} else {
		n3, err := PacketID.ReadFrom(r)
		if err != nil {
			return err
		}
		DataLength = VarInt(int64(PacketLength) - n2 - n3)
	}
	if cap(p.Data) < int(DataLength) {
		p.Data = make([]byte, DataLength)
	} else {
		p.Data = p.Data[:DataLength]
	}
	p.ID = int32(PacketID)
	_, err = io.ReadFull(r, p.Data)
	if err != nil {
		return err
	}
	return nil
}
