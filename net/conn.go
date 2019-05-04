package net

import (
	"bufio"
	"crypto/cipher"
	"io"
	"net"

	pk "github.com/Tnze/go-mc/net/packet"
)

type Conn struct {
	socket net.Conn
	io.ByteReader
	io.Writer

	threshold int
}

func DialMC(addr string) (conn *Conn, err error) {
	conn = new(Conn)
	conn.socket, err = net.Dial("tcp", addr)
	if err != nil {
		return
	}

	conn.ByteReader = bufio.NewReader(conn.socket)
	conn.Writer = conn.socket

	return
}

func (c *Conn) ReadPacket() (pk.Packet, error) {
	p, err := pk.RecvPacket(c.ByteReader, c.threshold > 0)
	if err != nil {
		return pk.Packet{}, err
	}
	return *p, err
}

func (c *Conn) WritePacket(p pk.Packet) error {
	_, err := c.Write(p.Pack(c.threshold))
	return err
}

func (c *Conn) SetCipher(encoStream, decoStream cipher.Stream) {
	//加密连接
	c.ByteReader = bufio.NewReader(cipher.StreamReader{ //Set reciver for AES
		S: decoStream,
		R: c.socket,
	})
	c.Writer = cipher.StreamWriter{
		S: encoStream,
		W: c.socket,
	}
}

func (c *Conn) SetThreshold(t int) {
	c.threshold = t
}
