package packet

import (
	"bytes"
	"compress/zlib"
	"encoding/hex"
	"fmt"
	"io"
	"os"
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
	rr := io.TeeReader(r, hex.Dumper(os.Stdout))
	for _, v := range fields {
		_, err := v.ReadFrom(rr)
		if err != nil {
			return err
		}
	}
	return nil
}

// Pack 打包一个数据包
func (p *Packet) Pack(w io.Writer, threshold int) error {
	var content bytes.Buffer
	if _, err := VarInt(p.ID).WriteTo(&content); err != nil {
		panic(err)
	}
	if _, err := content.Write(p.Data); err != nil {
		panic(err)
	}
	if threshold > 0 { //是否启用了压缩
		Len := content.Len()
		var VarLen bytes.Buffer
		if _, err := VarInt(Len).WriteTo(&VarLen); err != nil {
			panic(err)
		}
		if _, err := VarInt(VarLen.Len() + Len).WriteTo(w); err != nil {
			return err
		}
		if Len > threshold { //是否需要压缩
			compress(&content)
		}
		if _, err := VarLen.WriteTo(w); err != nil {
			return err
		}
		if _, err := content.WriteTo(w); err != nil {
			return err
		}
	} else {
		if _, err := VarInt(content.Len()).WriteTo(w); err != nil {
			return err
		}
		if _, err := content.WriteTo(w); err != nil {
			return err
		}
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
	if threshold > 0 {
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
	if sizeUncompressed != 0 { // != 0 means compressed, let's decompress
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
	} else {
		uncompressedData = data.Bytes()[1:]
	}
	*data = *bytes.NewBuffer(uncompressedData)
	return nil
}

// compress 压缩数据
func compress(data *bytes.Buffer) {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := data.WriteTo(w); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	*data = b
	return
}
