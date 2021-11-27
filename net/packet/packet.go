package packet

import (
	"bytes"
	"compress/zlib"
	"io"
	"sync"
)

// Packet define a net data package
type Packet struct {
	ID   int32
	Data []byte
}

//Marshal generate Packet with the ID and Fields
func Marshal(id int32, fields ...FieldEncoder) (pk Packet) {
	var pb Builder
	for _, v := range fields {
		pb.WriteField(v)
	}
	return pb.Packet(id)
}

//Scan decode the packet and fill data into fields
func (p Packet) Scan(fields ...FieldDecoder) error {
	r := bytes.NewReader(p.Data)
	for _, v := range fields {
		_, err := v.ReadFrom(r)
		if err != nil {
			return err
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
		return p.packWithCompression(w)
	} else {
		return p.packWithoutCompression(w)
	}
}

func (p *Packet) packWithoutCompression(w io.Writer) error {
	buffer := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buffer)
	n, err := VarInt(p.ID).WriteTo(buffer)
	if err != nil {
		panic(err)
	}
	// Length
	_, err = VarInt(int(n) + len(p.Data)).WriteTo(w)
	if err != nil {
		return err
	}
	// Packet ID
	_, err = buffer.WriteTo(w)
	if err != nil {
		return err
	}
	// Data
	_, err = w.Write(p.Data)
	if err != nil {
		return err
	}
	return nil
}

func (p *Packet) packWithCompression(w io.Writer) error {
	buff := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(buff)
	zw := zlib.NewWriter(buff)
	n1, err := VarInt(p.ID).WriteTo(zw)
	if err != nil {
		return err
	}
	n2, err := zw.Write(p.Data)
	if err != nil {
		return err
	}
	err = zw.Close()
	if err != nil {
		return err
	}

	dataLength := bufPool.Get().(*bytes.Buffer)
	defer bufPool.Put(dataLength)
	n3, err := VarInt(int(n1) + n2).WriteTo(dataLength)
	if err != nil {
		return err
	}

	// Packet Length
	_, err = VarInt(int(n3) + buff.Len()).WriteTo(w)
	if err != nil {
		return err
	}
	// Data Length
	_, err = dataLength.WriteTo(w)
	if err != nil {
		return err
	}
	// PacketID + Data
	_, err = buff.WriteTo(w)
	if err != nil {
		return err
	}
	return nil
}

// UnPack in-place decompression a packet
func (p *Packet) UnPack(r io.Reader, threshold int) error {
	if threshold >= 0 {
		return p.unpackWithCompression(r)
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

func (p *Packet) unpackWithCompression(r io.Reader) error {
	var PacketLength VarInt
	_, err := PacketLength.ReadFrom(r)
	if err != nil {
		return err
	}

	var DataLength VarInt
	n2, err := DataLength.ReadFrom(r)
	if err != nil {
		return err
	}

	var PacketID VarInt
	if DataLength != 0 {
		r, err = zlib.NewReader(r)
		if err != nil {
			return err
		}
		_, err = PacketID.ReadFrom(r)
		if err != nil {
			return err
		}
	} else {
		n3, err := PacketID.ReadFrom(r)
		if err != nil {
			return err
		}
		DataLength = PacketLength - VarInt(n2) - VarInt(n3)
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
