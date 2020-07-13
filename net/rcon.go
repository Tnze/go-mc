package net

import (
	"bytes"
	"encoding/binary"
	"errors"
	"fmt"
	"math/rand"
	"net"
)

const MaxRCONPackageSize = 4096

func DialRCON(addr string, password string) (client RCONClientConn, err error) {
	c := &RCONConn{ReqID: rand.Int31()}
	client = c

	c.Conn, err = net.Dial("tcp", addr)
	if err != nil {
		err = fmt.Errorf("connect fail: %v", err)
		return
	}

	//Login
	err = c.WritePacket(c.ReqID, 3, password)
	if err != nil {
		err = fmt.Errorf("login fail: %v", err)
		return
	}

	//Login resp
	r, _, _, err := c.ReadPacket()
	if err != nil {
		err = fmt.Errorf("read login resp fail: %v", err)
		return
	}

	if r == c.ReqID {
		err = nil
	} else if r == -1 {
		err = errors.New("login fail")
	} else {
		err = errors.New("req id not match")
	}

	return
}

type RCONConn struct {
	net.Conn
	ReqID int32
}

func (r *RCONConn) ReadPacket() (RequestID, Type int32, Payload string, err error) {
	//read packet length
	var Length int32
	err = binary.Read(r, binary.LittleEndian, &Length)
	if err != nil {
		err = fmt.Errorf("read packet length fail: %v", err)
		return
	}

	//check length
	if Length < 4+4+0+2 {
		err = errors.New("packet too short")
		return
	}
	if Length > MaxRCONPackageSize {
		err = errors.New("packet too large")
		return
	}

	//read packet data
	buf := make([]byte, Length)
	err = binary.Read(r, binary.LittleEndian, &buf)
	if err != nil {
		err = fmt.Errorf("read packet body fail: %v", err)
		return
	}
	RequestID = int32(binary.LittleEndian.Uint32(buf[:4]))
	Type = int32(binary.LittleEndian.Uint32(buf[4:8]))
	Payload = string(buf[8 : Length-2])

	return
}

func (r *RCONConn) WritePacket(RequestID, Type int32, Payload string) error {
	buf := new(bytes.Buffer)

	for _, v := range []interface{}{
		int32(4 + 4 + len(Payload) + 2), //Length
		RequestID,                       //Request ID
		Type,                            //Type
		[]byte(Payload),                 //Payload
		[]byte{0, 0},                    //pad
	} {
		err := binary.Write(buf, binary.LittleEndian, v)
		if err != nil {
			return err
		}
	}

	_, err := r.Write(buf.Bytes())
	return err
}

func (r *RCONConn) Cmd(cmd string) error {
	err := r.WritePacket(r.ReqID, 2, cmd)
	return err
}

func (r *RCONConn) Resp() (resp string, err error) {
	var ReqID, Type int32
	ReqID, Type, resp, err = r.ReadPacket()
	if err != nil {
		return
	}

	if ReqID != r.ReqID {
		err = errors.New("req ID not match")
	} else if Type != 0 {
		err = fmt.Errorf("packet type wrong: %d", Type)
	}

	return
}

func (r *RCONConn) AcceptLogin(password string) error {
	R, T, P, err := r.ReadPacket()
	if err != nil {
		return err
	}

	r.ReqID = R

	//Check packet type
	if T != 3 {
		return fmt.Errorf("not a login packet: %d", T)
	}

	if P != password {
		err = r.WritePacket(-1, 2, "")
		if err != nil {
			return err
		}
		return errors.New("password wrong")
	}

	err = r.WritePacket(R, 2, "")
	if err != nil {
		return err
	}

	return nil
}

func (r *RCONConn) AcceptCmd() (string, error) {
	R, T, P, err := r.ReadPacket()
	if err != nil {
		return P, err
	}

	r.ReqID = R

	//Check packet type
	if T != 2 {
		return P, fmt.Errorf("not a command packet: %d", T)
	}

	return P, nil
}

func (r *RCONConn) RespCmd(resp string) error {
	return r.WritePacket(r.ReqID, 0, resp)
}

type RCONClientConn interface {
	Cmd(cmd string) error
	Resp() (resp string, err error)
	Close() error
}

type RCONServerConn interface {
	AcceptLogin(password string) error
	AcceptCmd() (cmd string, err error)
	RespCmd(resp string) error
	Close() error
}

func ListenRCON(addr string) (*RCONListener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}

	return &RCONListener{Listener: l}, nil
}

type RCONListener struct{ net.Listener }

func (r *RCONListener) Accept() (RCONServerConn, error) {
	conn, err := r.Listener.Accept()
	if err != nil {
		return nil, err
	}

	return &RCONConn{Conn: conn}, nil
}
