// Package net pack network connection for Minecraft.
package net

import (
	"bufio"
	"crypto/cipher"
	"io"
	"net"
	"time"

	pk "github.com/Tnze/go-mc/net/packet"
)

// A Listener is a minecraft Listener
type Listener struct{ net.Listener }

//ListenMC listen as TCP but Accept a mc Conn
func ListenMC(addr string) (*Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{l}, nil
}

//Accept a minecraft Conn
func (l Listener) Accept() (Conn, error) {
	conn, err := l.Listener.Accept()
	return Conn{
		Socket: conn,
		Reader: bufio.NewReader(conn),
		Writer: conn,
	}, err
}

//Conn is a minecraft Connection
type Conn struct {
	Socket net.Conn
	Reader interface {
		io.ByteReader
		io.Reader
	}
	io.Writer

	threshold int
}

// DialMC create a Minecraft connection
func DialMC(addr string) (*Conn, error) {
	conn, err := net.Dial("tcp", addr)
	return &Conn{
		Socket: conn,
		Reader: bufio.NewReader(conn),
		Writer: conn,
	}, err
}

// DialMCTimeout acts like DialMC but takes a timeout.
func DialMCTimeout(addr string, timeout time.Duration) (*Conn, error) {
	conn, err := net.DialTimeout("tcp", addr, timeout)
	return &Conn{
		Socket: conn,
		Reader: bufio.NewReader(conn),
		Writer: conn,
	}, err
}

// WrapConn warp an net.Conn to MC-Conn
// Helps you modify the connection process (eg. using DialContext).
func WrapConn(conn net.Conn) *Conn {
	return &Conn{
		Socket: conn,
		Reader: bufio.NewReader(conn),
		Writer: conn,
	}
}

//Close close the connection
func (c *Conn) Close() error { return c.Socket.Close() }

// ReadPacket read a Packet from Conn.
func (c *Conn) ReadPacket() (pk.Packet, error) {
	p, err := pk.RecvPacket(c.Reader, c.threshold > 0)
	if err != nil {
		return pk.Packet{}, err
	}
	return *p, err
}

//WritePacket write a Packet to Conn.
func (c *Conn) WritePacket(p pk.Packet) error {
	_, err := c.Write(p.Pack(c.threshold))
	return err
}

// SetCipher load the decode/encode stream to this Conn
func (c *Conn) SetCipher(ecoStream, decoStream cipher.Stream) {
	//加密连接
	c.Reader = bufio.NewReader(cipher.StreamReader{ //Set receiver for AES
		S: decoStream,
		R: c.Socket,
	})
	c.Writer = cipher.StreamWriter{
		S: ecoStream,
		W: c.Socket,
	}
}

// SetThreshold set threshold to Conn.
// The data packet with length longer then threshold
// will be compress when sending.
func (c *Conn) SetThreshold(t int) {
	c.threshold = t
}
