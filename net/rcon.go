package net

import (
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"net"
)

func DialRCON(addr string, password string) (c *RCONConn, err error) {
	c = &RCONConn{reqID: rand.Int31()}

	c.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		return
	}

	//Login
	err = c.WritePacket(3, password)
	if err != nil {
		return
	}

	t, p, err := c.ReadPacket()
	if err != nil {
		return
	}
	fmt.Print(t, p)

	return
}

type RCONConn struct {
	net.Conn
	reqID int32
}

func (c *RCONConn) ReadPacket() (Type int32, Payload string, err error) {
	//read packet length
	var Length int32
	err = binary.Read(c, binary.LittleEndian, &Length)
	if err != nil {
		return
	}

	//read packet data
	buf := make([]byte, Length)
	err = binary.Read(c, binary.LittleEndian, &buf)
	if err != nil {
		return
	}

	//check length
	if Length < 4+4+0+2 {
		err = errors.New("packet too short")
		return
	}

	RequestID := int32(binary.LittleEndian.Uint32(buf[:4]))
	Type = int32(binary.LittleEndian.Uint32(buf[4:8]))
	Payload = string(buf[8 : Length-2])

	if RequestID == -1 {
		err = errors.New("login fail")
	} else if RequestID != c.reqID {
		err = errors.New("request ID not match")
	}

	return
}

func (c *RCONConn) WritePacket(Type int32, Payload string) error {
	err := binary.Write(c, binary.LittleEndian, []interface{}{
		int32(4 + 4 + len(Payload) + 2), //Length
		c.reqID,                         //Request ID
		Type,                            //Type
		[]byte(Payload),                 //Payload
		[2]byte{0, 0},                   //pad
	})

	return err
}
