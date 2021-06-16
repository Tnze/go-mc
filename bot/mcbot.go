// Package bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at examples/ .
package bot

import (
	"context"
	"errors"
	"net"
	"strconv"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ProtocolVersion , the protocol version number of minecraft net protocol
const ProtocolVersion = 754
const DefaultPort = 25565

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err error) {
	return c.join(&net.Dialer{}, addr)
}

// JoinServerWithDialer is similar to JoinServer but using a Dialer.
func (c *Client) JoinServerWithDialer(d *net.Dialer, addr string) (err error) {
	return c.join(d, addr)
}

// parseAddress will lookup SRV records for the address
func parseAddress(r *net.Resolver, addr string) (string, error) {
	var port uint16
	var addrErr *net.AddrError
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		if errors.As(err, &addrErr) {
			host, port = addr, DefaultPort
		} else {
			return "", err
		}
	} else {
		if portInt, err := strconv.ParseUint(portStr, 10, 16); err != nil {
			port = DefaultPort
		} else {
			port = uint16(portInt)
		}
	}

	_, srvs, err := r.LookupSRV(context.TODO(), "minecraft", "tcp", host)
	if err != nil && len(srvs) > 0 {
		host, port = srvs[0].Target, srvs[0].Port
	}

	return net.JoinHostPort(host, strconv.FormatUint(uint64(port), 10)), nil
}

func (c *Client) join(d *net.Dialer, addr string) error {
	const Handshake = 0x00
	addrSrv, err := parseAddress(d.Resolver, addr)
	if err != nil {
		return LoginErr{"resolved address", err}
	}

	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addrSrv)
	if err != nil {
		return LoginErr{"split address", err}
	}
	port, err := strconv.ParseUint(portStr, 0, 16)
	if err != nil {
		return LoginErr{"parse port", err}
	}

	// Dial connection
	c.Conn, err = mcnet.DialMC(addrSrv)
	if err != nil {
		return LoginErr{"connect server", err}
	}

	// Handshake
	err = c.Conn.WritePacket(pk.Marshal(
		Handshake,
		pk.VarInt(ProtocolVersion), // Protocol version
		pk.String(host),            // Host
		pk.UnsignedShort(port),     // Port
		pk.Byte(2),
	))
	if err != nil {
		return LoginErr{"handshake", err}
	}

	// Login Start
	err = c.Conn.WritePacket(pk.Marshal(
		packetid.LoginStart,
		pk.String(c.Auth.Name),
	))
	if err != nil {
		return LoginErr{"login start", err}
	}

	for {
		//Receive Packet
		var p pk.Packet
		if err = c.Conn.ReadPacket(&p); err != nil {
			return LoginErr{"receive packet", err}
		}

		//Handle Packet
		switch p.ID {
		case packetid.Disconnect: //Disconnect
			var reason chat.Message
			err = p.Scan(&reason)
			if err != nil {
				return LoginErr{"disconnect", err}
			}
			return LoginErr{"disconnect", DisconnectErr(reason)}

		case packetid.EncryptionBeginClientbound: //Encryption Request
			if err := handleEncryptionRequest(c, p); err != nil {
				return LoginErr{"encryption", err}
			}

		case packetid.Success: //Login Success
			err := p.Scan(
				(*pk.UUID)(&c.UUID),
				(*pk.String)(&c.Name),
			)
			if err != nil {
				return LoginErr{"login success", err}
			}
			return nil

		case packetid.Compress: //Set Compression
			var threshold pk.VarInt
			if err := p.Scan(&threshold); err != nil {
				return LoginErr{"compression", err}
			}
			c.Conn.SetThreshold(int(threshold))

		case packetid.LoginPluginRequest: //Login Plugin Request
			var (
				msgid   pk.VarInt
				channel pk.Identifier
				data    pk.ByteArray
			)
			if err := p.Scan(&msgid, &channel, &data); err != nil {
				return LoginErr{"Login Plugin", err}
			}
			if handler, ok := c.LoginPlugin[string(channel)]; ok {
				respdata, err := handler([]byte(data))
				if err != nil {
					return LoginErr{"Login Plugin", err}
				}
				if err := c.Conn.WritePacket(pk.Marshal(packetid.LoginPluginResponse, msgid, pk.Boolean(true), pk.ByteArray(respdata))); err != nil {
					return LoginErr{"login Plugin", err}
				}
			} else {
				if err := c.Conn.WritePacket(pk.Marshal(packetid.LoginPluginResponse, msgid, pk.Boolean(false))); err != nil {
					return LoginErr{"login Plugin", err}
				}
			}
		}
	}
}

type LoginErr struct {
	Stage string
	Err   error
}

func (l LoginErr) Error() string {
	return "bot: " + l.Stage + " error: " + l.Err.Error()
}

func (l LoginErr) Unwrap() error {
	return l.Err
}

type DisconnectErr chat.Message

func (d DisconnectErr) Error() string {
	return "disconnect because: " + chat.Message(d).String()
}
