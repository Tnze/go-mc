package packet

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
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

// Pack 打包一个数据包
func (p *Packet) Pack(w io.Writer, threshold int) error {
	if threshold >= 0 {
		return p.withCompression(w)
	} else {
		return p.withoutCompression(w)
	}
}

func (p *Packet) withoutCompression(w io.Writer) error {
	var buf [5]byte
	buffer := bytes.NewBuffer(buf[:0])
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

func (p *Packet) withCompression(w io.Writer) error {
	var buff bytes.Buffer
	zw := zlib.NewWriter(&buff)
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

	var dataLength bytes.Buffer
	n3, err := VarInt(int(n1) + n2).WriteTo(&dataLength)
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
	var length VarInt
	if _, err := length.ReadFrom(r); err != nil {
		return err
	}
	if length < 1 {
		return fmt.Errorf("packet length too short")
	}
	buf := make([]byte, length)
	if _, err := io.ReadFull(r, buf); err != nil {
		return fmt.Errorf("read content of packet fail: %w", err)
	}
	buffer := bytes.NewBuffer(buf)

	//解压数据
	if threshold >= 0 {
		if err := unCompress(buffer); err != nil {
			return err
		}
	}

	var packetID VarInt
	if _, err := packetID.ReadFrom(buffer); err != nil {
		return fmt.Errorf("read packet id fail: %v", err)
	}
	p.ID = int32(packetID)
	p.Data = buffer.Bytes()
	return nil
}

// unCompress 读取一个压缩的包
func unCompress(data *bytes.Buffer) error {
	reader := bytes.NewReader(data.Bytes())

	var sizeUncompressed VarInt
	if _, err := sizeUncompressed.ReadFrom(reader); err != nil {
		return err
	}

	var uncompressedData []byte
	if sizeUncompressed == 0 {
		uncompressedData = data.Bytes()[1:]
	} else { // != 0 means compressed, let's decompress
		uncompressedData = make([]byte, sizeUncompressed)
		r, err := zlib.NewReader(reader)
		if err != nil {
			return fmt.Errorf("decompress fail: %v", err)
		}
		defer r.Close()
		_, err = io.ReadFull(r, uncompressedData)
		if err != nil {
			return fmt.Errorf("decompress fail: %v", err)
		}
	}
	*data = *bytes.NewBuffer(uncompressedData)
	return nil
}
