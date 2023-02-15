// Package net pack network connection for Minecraft.
package net

import (
	"context"
	"crypto/cipher"
	"errors"
	"io"
	"net"
	"strconv"
	"time"

	pk "github.com/Tnze/go-mc/net/packet"
)

const DefaultPort = 25565

// A Listener is a minecraft Listener
type Listener struct{ net.Listener }

// ListenMC listen as TCP but Accept a mc Conn
func ListenMC(addr string) (*Listener, error) {
	l, err := net.Listen("tcp", addr)
	if err != nil {
		return nil, err
	}
	return &Listener{l}, nil
}

// Accept a minecraft Conn
func (l Listener) Accept() (Conn, error) {
	conn, err := l.Listener.Accept()
	return Conn{
		Socket:    conn,
		Reader:    conn,
		Writer:    conn,
		threshold: -1,
	}, err
}

// Conn is a minecraft Connection
type Conn struct {
	Socket net.Conn
	io.Reader
	io.Writer

	threshold int
}

var DefaultDialer = Dialer{}

// DialMC create a Minecraft connection
// Lookup SRV records only if port doesn't exist or equals to 0.
func DialMC(addr string) (*Conn, error) {
	return DefaultDialer.DialMCContext(context.Background(), addr)
}

// DialMCTimeout acts like DialMC but takes a timeout.
func DialMCTimeout(addr string, timeout time.Duration) (*Conn, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return DefaultDialer.DialMCContext(ctx, addr)
}

// MCDialer provide DialMCContext method, can be used to dial a minecraft server.
// [Dialer] is its default implementation, and support SRV lookup.
//
// Typically, if you want to use built-in proxies or custom dialer,
// you can hook go-mc/bot package by implement this interface.
// When implementing a custom MCDialer, SRV lookup is optional.
type MCDialer interface {
	// The DialMCContext dial TCP connection to a minecraft server, and warp the net.Conn by calling [WrapConn].
	DialMCContext(ctx context.Context, addr string) (*Conn, error)
}

// Dialer implements MCDialer interface.
//
// It can be easily convert from net.Dialer.
//
//	dialer := net.Dialer{}
//	mcDialer := (*Dialer)(&dialer)
type Dialer net.Dialer

func (d *Dialer) resolver() *net.Resolver {
	if d != nil && d.Resolver != nil {
		return d.Resolver
	}
	return net.DefaultResolver
}

func (d *Dialer) DialMCContext(ctx context.Context, addr string) (*Conn, error) {
	host, port, err := net.SplitHostPort(addr)
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host, port, err = addr, "", nil
		} else {
			return nil, err
		}
	}
	var addresses []string
	if port == "" {
		// We look up SRV only if the port is not specified
		_, srvRecords, err := d.resolver().LookupSRV(ctx, "minecraft", "tcp", host)
		if err == nil {
			for _, record := range srvRecords {
				addr := net.JoinHostPort(record.Target, strconv.Itoa(int(record.Port)))
				addresses = append(addresses, addr)
			}
		}
		// Whatever the SRV records is found,
		addr = net.JoinHostPort(addr, strconv.Itoa(DefaultPort))
	}
	addresses = append(addresses, addr)

	var firstErr error
	for i, addr := range addresses {
		select {
		case <-ctx.Done():
			return nil, context.Canceled
		default:
		}
		dialCtx := ctx
		if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
			partialDeadline, err := partialDeadline(time.Now(), deadline, len(addresses)-i)
			if err != nil {
				// Ran out of time.
				if firstErr == nil {
					firstErr = context.DeadlineExceeded
				}
				break
			}
			if partialDeadline.Before(deadline) {
				var cancel context.CancelFunc
				dialCtx, cancel = context.WithDeadline(ctx, partialDeadline)
				defer cancel()
			}
		}
		conn, err := (*net.Dialer)(d).DialContext(dialCtx, "tcp", addr)
		if err != nil {
			if firstErr == nil {
				firstErr = err
			}
			continue
		}
		return WrapConn(conn), nil
	}
	return nil, firstErr
}

// deadline returns the earliest of:
//   - now+Timeout
//   - d.Deadline
//   - the context's deadline
//
// Or zero, if none of Timeout, Deadline, or context's deadline is set.
//
// Copied from net/dial.go
func (d *Dialer) deadline(ctx context.Context, now time.Time) (earliest time.Time) {
	if d.Timeout != 0 { // including negative, for historical reasons
		earliest = now.Add(d.Timeout)
	}
	if d, ok := ctx.Deadline(); ok {
		earliest = minNonzeroTime(earliest, d)
	}
	return minNonzeroTime(earliest, d.Deadline)
}

// Copied from net/dial.go
func minNonzeroTime(a, b time.Time) time.Time {
	if a.IsZero() {
		return b
	}
	if b.IsZero() || a.Before(b) {
		return a
	}
	return b
}

// partialDeadline returns the deadline to use for a single address,
// when multiple addresses are pending.
//
// Copied from net/dial.go
func partialDeadline(now, deadline time.Time, addrsRemaining int) (time.Time, error) {
	if deadline.IsZero() {
		return deadline, nil
	}
	timeRemaining := deadline.Sub(now)
	if timeRemaining <= 0 {
		return time.Time{}, context.DeadlineExceeded
	}
	// Tentatively allocate equal time to each remaining address.
	timeout := timeRemaining / time.Duration(addrsRemaining)
	// If the time per address is too short, steal from the end of the list.
	const saneMinimum = 2 * time.Second
	if timeout < saneMinimum {
		if timeRemaining < saneMinimum {
			timeout = timeRemaining
		} else {
			timeout = saneMinimum
		}
	}
	return now.Add(timeout), nil
}

// WrapConn warp a net.Conn to MC-Conn
// Helps you modify the connection process (e.g. using DialContext).
func WrapConn(conn net.Conn) *Conn {
	return &Conn{
		Socket:    conn,
		Reader:    conn,
		Writer:    conn,
		threshold: -1,
	}
}

// Close the connection
func (c *Conn) Close() error { return c.Socket.Close() }

// ReadPacket read a Packet from Conn.
func (c *Conn) ReadPacket(p *pk.Packet) error {
	return p.UnPack(c.Reader, c.threshold)
}

// WritePacket write a Packet to Conn.
func (c *Conn) WritePacket(p pk.Packet) error {
	return p.Pack(c.Writer, c.threshold)
}

// SetCipher load the decode/encode stream to this Conn
func (c *Conn) SetCipher(ecoStream, decoStream cipher.Stream) {
	// 加密连接
	c.Reader = cipher.StreamReader{ // Set receiver for AES
		S: decoStream,
		R: c.Socket,
	}
	c.Writer = cipher.StreamWriter{
		S: ecoStream,
		W: c.Socket,
	}
}

// SetThreshold set threshold to Conn.
// The data packet with length equal or longer then threshold
// will be compressed when sending.
func (c *Conn) SetThreshold(t int) {
	c.threshold = t
}
