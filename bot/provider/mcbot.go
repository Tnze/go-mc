// Package bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at examples/ .
package provider

import (
	"context"
	"encoding/base64"
	"encoding/pem"
	"errors"
	"fmt"
	"github.com/Tnze/go-mc/yggdrasil/user"
	"io"
	"net"
	"strconv"

	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ProtocolVersion is the protocol version number of minecraft net protocol
const ProtocolVersion = 759
const DefaultPort = mcnet.DefaultPort

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err basic.Error) {
	return c.join(context.Background(), &mcnet.DefaultDialer, addr)
}

// JoinServerWithDialer is similar to JoinServer but using a Dialer.
func (c *Client) JoinServerWithDialer(d *net.Dialer, addr string) (err error) {
	dialer := (*mcnet.Dialer)(d)
	return c.join(context.Background(), dialer, addr)
}

func (c *Client) join(ctx context.Context, d *mcnet.Dialer, addr string) basic.Error {
	const Handshake = 0x00
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	var port uint64
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host = addr
			port = 25565
		} else {
			return basic.Error{Err: basic.NullValue, Info: fmt.Errorf("split host and port: %w", err)}
		}
	} else {
		port, err = strconv.ParseUint(portStr, 0, 16)
		if err != nil {
			return basic.Error{Err: basic.NullValue, Info: fmt.Errorf("parse port: %w", err)}
		}
	}

	// Dial connection
	if c.Conn, err = d.DialMCContext(ctx, addr); err != nil {
		return basic.Error{Err: basic.DialError, Info: fmt.Errorf("dial connection: %w", err)}
	}

	// Handshake
	if err := c.Conn.WritePacket(pk.Marshal(
		Handshake,
		pk.VarInt(ProtocolVersion),
		pk.String(host),
		pk.UnsignedShort(port),
		pk.VarInt(2),
	)); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("handshake: %w", err)}
	}
	// Login Start
	if err := c.Conn.WritePacket(pk.Marshal(
		packetid.SPacketLoginStart,
		pk.String(c.Auth.Name),
		pk.Boolean(false), // TODO: Support online mode
		//keyPair(c.Auth.KeyPair),
	)); !err.Is(basic.NoError) {
		return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("login start: %w", err)}
	}
	for {
		//Receive Packet
		var p pk.Packet
		if err := c.Conn.ReadPacket(&p); !err.Is(basic.NoError) {
			return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("read packet: %w", err)}
		}

		//Handle Packet
		switch p.ID {
		case packetid.CPacketLoginDisconnect:
			var reason chat.Message
			if err := p.Scan(&reason); err != nil {
				return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("login disconnect: %w", err)}
			}

		case packetid.CPacketEncryptionRequest:
			if err := handleEncryptionRequest(c, p); err != nil {
				return basic.Error{Err: basic.DialError, Info: fmt.Errorf("encryption request: %w", err)}
			}

		case packetid.CPacketLoginSuccess:
			if err := p.Scan(
				(*pk.UUID)(&c.Player.UUID),
				(*pk.String)(&c.Player.DisplayName),
			); err != nil {
				return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("login success: %w", err)}
			} else {
				return basic.Error{Err: basic.NoError, Info: nil}
			}
		case packetid.CPacketSetCompression:
			var threshold pk.VarInt
			if err := p.Scan(&threshold); err != nil {
				return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("set compression: %w", err)}
			}
			c.Conn.SetThreshold(int(threshold))

		case packetid.CPacketPluginMessage:
			var (
				msgid   pk.VarInt
				channel pk.Identifier
				data    pk.PluginMessageData
			)
			if err := p.Scan(&msgid, &channel, &data); err != nil {
				return basic.Error{Err: basic.ReaderError, Info: fmt.Errorf("plugin message: %w", err)}
			}

			handler, ok := c.LoginPlugin[string(channel)]
			if ok {
				data, err = handler(data)
				if err != nil {
					return basic.Error{Err: basic.DialError, Info: fmt.Errorf("plugin message: %w", err)}
				}
			}

			if err := c.Conn.WritePacket(pk.Marshal(
				packetid.CPacketPluginMessage,
				msgid, pk.Boolean(ok),
				pk.Opt{If: ok, Value: data},
			)); !err.Is(basic.NoError) {
				return basic.Error{Err: basic.WriterError, Info: fmt.Errorf("plugin message: %w", err)}
			}
		}
	}
}

type keyPair user.KeyPairResp

func (k keyPair) WriteTo(w io.Writer) (int64, error) {
	block, _ := pem.Decode([]byte(k.KeyPair.PublicKey))
	if block == nil {
		return 0, errors.New("pem decode error: no data is found")
	}
	signature, err := base64.StdEncoding.DecodeString(k.PublicKeySignatureV2)
	if err != nil {
		return 0, err
	}
	return pk.Tuple{
		pk.Long(k.ExpiresAt.UnixMilli()),
		pk.ByteArray(block.Bytes),
		pk.ByteArray(signature),
	}.WriteTo(w)
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
