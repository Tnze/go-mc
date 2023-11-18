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
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/net/queue"
	"github.com/Tnze/go-mc/yggdrasil/user"
)

// ProtocolVersion is the protocol version number of minecraft net protocol
const (
	ProtocolVersion = 764
	DefaultPort     = mcnet.DefaultPort
)

type JoinOptions struct {
	MCDialer mcnet.MCDialer
	Context  context.Context

	// Indicate not to fetch and sending player's PubKey
	NoPublicKey bool

	// Specify the player PubKey to use.
	// If nil, it will be obtained from Mojang when joining
	KeyPair *user.KeyPairResp

	QueueRead  queue.Queue[pk.Packet]
	QueueWrite queue.Queue[pk.Packet]
}

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err error) {
	return c.JoinServerWithOptions(addr, JoinOptions{})
}

// JoinServerWithDialer is similar to JoinServer but using a net.Dialer.
func (c *Client) JoinServerWithDialer(dialer *net.Dialer, addr string) (err error) {
	return c.JoinServerWithOptions(addr, JoinOptions{
		MCDialer: (*mcnet.Dialer)(dialer),
	})
}

func (c *Client) JoinServerWithOptions(addr string, options JoinOptions) (err error) {
	if options.MCDialer == nil {
		options.MCDialer = &mcnet.DefaultDialer
	}
	if options.Context == nil {
		options.Context = context.Background()
	}
	if options.QueueRead == nil {
		options.QueueRead = queue.NewLinkedQueue[pk.Packet]()
	}
	if options.QueueWrite == nil {
		options.QueueWrite = queue.NewLinkedQueue[pk.Packet]()
	}
	return c.join(addr, options)
}

func (c *Client) join(addr string, options JoinOptions) error {
	const Handshake = 0x00

	// Split Host and Port. The DialMCContext will do this once,
	// but we need the result for sending handshake packet here.
	host, portStr, err := net.SplitHostPort(addr)
	var port uint64
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host = addr
			port = 25565
		} else {
			return LoginErr{"split address", err}
		}
	} else {
		port, err = strconv.ParseUint(portStr, 0, 16)
		if err != nil {
			return LoginErr{"parse port", err}
		}
	}

	// Dial connection
	conn, err := options.MCDialer.DialMCContext(options.Context, addr)
	if err != nil {
		return LoginErr{"connect server", err}
	}

	// Handshake
	err = conn.WritePacket(pk.Marshal(
		Handshake,
		pk.VarInt(ProtocolVersion), // Protocol version
		pk.String(host),            // Host
		pk.UnsignedShort(port),     // Port
		pk.VarInt(2),
	))
	if err != nil {
		return LoginErr{"handshake", err}
	}

	// Login Start
	if err := c.joinLogin(conn); err != nil {
		return err
	}

	// Configuration
	if err := c.joinConfiguration(conn); err != nil {
		return err
	}
	c.Conn = warpConn(conn, options.QueueRead, options.QueueWrite)
	return nil
}

type DisconnectErr chat.Message

func (d DisconnectErr) Error() string {
	return "disconnect because: " + chat.Message(d).String()
}
