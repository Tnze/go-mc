package packet

import (
	"bytes"
	"compress/zlib"
	"fmt"
	"io"
)

// Packet define a net data package
type Packet struct {
	ID   byte
	Data []byte
}

//Marshal generate Packet with the ID and Fields
func Marshal(ID byte, fields ...FieldEncoder) (pk Packet) {
	pk.ID = ID

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
	data := []byte{p.ID}           //data
	data = append(data, p.Data...) //data

	if threshold > 0 { //是否启用了压缩
		if len(data) > threshold { //是否需要压缩
			Len := len(data)
			VarLen := VarInt(Len).Encode()
			data = Compress(data)

			pack = append(pack, VarInt(len(VarLen)+len(data)).Encode()...)
			pack = append(pack, VarLen...)
			pack = append(pack, data...)
		} else {
			pack = append(pack, VarInt(int32(len(data)+1)).Encode()...)
			pack = append(pack, 0x00)
			pack = append(pack, data...)
		}
	} else {
		pack = append(pack, VarInt(int32(len(data))).Encode()...) //len
		pack = append(pack, data...)
	}

	return
}

// RecvPacket recive a packet from server
func RecvPacket(r io.ByteReader, useZlib bool) (*Packet, error) {
	var len int
	for i := 0; i < 5; i++ { //读数据前的长度标记
		b, err := r.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("read len of packet fail: %v", err)
		}
		len |= (int(b&0x7F) << uint(7*i))
		if b&0x80 == 0 {
			break
		}
	}

	if len < 1 {
		return nil, fmt.Errorf("packet length too short")
	}

	data := make([]byte, len) //读包内容
	var err error
	for i := 0; i < len; i++ {
		data[i], err = r.ReadByte()
		if err != nil {
			return nil, fmt.Errorf("read content of packet fail: %v", err)
		}
	}

	//解压数据
	if useZlib {
		return UnCompress(data)
	}

	return &Packet{
		ID:   data[0],
		Data: data[1:],
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
		_, err = io.ReadFull(r, uncompressData)
		if err != nil {
			return nil, fmt.Errorf("decompress fail: %v", err)
		}
		r.Close()
	} else {
		uncompressData = data[1:]
	}
	return &Packet{
		ID:   uncompressData[0],
		Data: uncompressData[1:],
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
