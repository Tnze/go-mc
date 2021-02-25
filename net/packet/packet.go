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
	pk.ID = id

	for _, v := range fields {
		pk.Data = append(pk.Data, v.Encode()...)
	}

	return
}

//Scan decode the packet and fill data into fields
func (p Packet) Scan(fields ...FieldDecoder) error {
	r := bytes.NewReader(p.Data)
	for _, v := range fields {
		err := v.Decode(r)
		if err != nil {
			return err
		}
	}
	return nil
}

// Pack 打包一个数据包
func (p *Packet) Pack(threshold int) (pack []byte) {
	d := append(VarInt(p.ID).Encode(), p.Data...)
	if threshold > 0 { //是否启用了压缩
		if len(d) > threshold { //是否需要压缩
			Len := len(d)
			VarLen := VarInt(Len).Encode()
			d = compress(d)

			pack = append(pack, VarInt(len(VarLen)+len(d)).Encode()...)
			pack = append(pack, VarLen...)
			pack = append(pack, d...)
		} else {
			pack = append(pack, VarInt(int32(len(d)+1)).Encode()...)
			pack = append(pack, 0x00)
			pack = append(pack, d...)
		}
	} else {
		pack = append(pack, VarInt(int32(len(d))).Encode()...) //len
		pack = append(pack, d...)
	}

	return
}

// UnPack in-place decompression a packet
func (p *Packet) UnPack(r DecodeReader, threshold int) error {
	var length VarInt
	if err := length.Decode(r); err != nil {
		return err
	}
	if length < 1 {
		return fmt.Errorf("packet length too short")
	}

	buf := bytes.NewBuffer(p.Data[:0])
	if _, err := io.CopyN(buf, r, int64(length)); err != nil {
		return fmt.Errorf("read content of packet fail: %w", err)
	}

	//解压数据
	if threshold > 0 {
		if err := unCompress(buf); err != nil {
			return err
		}
	}

	var packetID VarInt
	if err := packetID.Decode(buf); err != nil {
		return fmt.Errorf("read packet id fail: %v", err)
	}
	p.ID = int32(packetID)
	p.Data = buf.Bytes()
	return nil
}

// unCompress 读取一个压缩的包
func unCompress(data *bytes.Buffer) error {
	reader := bytes.NewReader(data.Bytes())

	var sizeUncompressed VarInt
	if err := sizeUncompressed.Decode(reader); err != nil {
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
func compress(data []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	if _, err := w.Write(data); err != nil {
		panic(err)
	}
	if err := w.Close(); err != nil {
		panic(err)
	}
	return b.Bytes()
}
