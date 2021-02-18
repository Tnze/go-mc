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
			d = Compress(d)

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

// RecvPacket receive a packet from server
func RecvPacket(r DecodeReader, useZlib bool) (*Packet, error) {
	var length VarInt
	if err := length.Decode(r); err != nil {
		return nil, err
	}
	if length < 1 {
		return nil, fmt.Errorf("packet length too short")
	}

	d := make([]byte, length) // read packet content
	if _, err := io.ReadFull(r, d); err != nil {
		return nil, fmt.Errorf("read content of packet fail: %v", err)
	}

	//解压数据
	if useZlib {
		return UnCompress(d)
	}

	buf := bytes.NewBuffer(d)
	var packetID VarInt
	if err := packetID.Decode(buf); err != nil {
		return nil, fmt.Errorf("read packet id fail: %v", err)
	}
	return &Packet{
		ID:   int32(packetID),
		Data: buf.Bytes(),
	}, nil
}

// UnCompress 读取一个压缩的包
func UnCompress(data []byte) (*Packet, error) {
	reader := bytes.NewReader(data)

	var sizeUncompressed VarInt
	if err := sizeUncompressed.Decode(reader); err != nil {
		return nil, err
	}

	uncompressData := make([]byte, sizeUncompressed)
	if sizeUncompressed != 0 { // != 0 means compressed, let's decompress
		r, err := zlib.NewReader(reader)
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
		defer r.Close()
		_, err = io.ReadFull(r, uncompressData)
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
	} else {
		uncompressData = data[1:]
	}
	buf := bytes.NewBuffer(uncompressData)
	var packetID VarInt
	if err := packetID.Decode(buf); err != nil {
		return nil, fmt.Errorf("read packet id fail: %v", err)
	}
	return &Packet{
		ID:   int32(packetID),
		Data: buf.Bytes(),
	}, nil
}

// Compress 压缩数据
func Compress(data []byte) []byte {
	var b bytes.Buffer
	w := zlib.NewWriter(&b)
	w.Write(data)
	w.Close()
	return b.Bytes()
}
