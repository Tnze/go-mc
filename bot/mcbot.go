// Package bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at cmd/ .
package bot

import (
	"context"
	"fmt"
	"net"
	"strconv"

	"github.com/Tnze/go-mc/data"
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ProtocolVersion , the protocol version number of minecraft net protocol
const ProtocolVersion = 754
const DefaultPort = 25565

// JoinServer connect a Minecraft server for playing the game.
// Using roughly the same way to parse address as minecraft.
func (c *Client) JoinServer(addr string) (err error) {
	return c.JoinServerWithDialer(&net.Dialer{}, addr)
}

func parseAddress(r *net.Resolver, addr string) (string, error) {
	const missingPort = "missing port in address"
	var port uint16
	host, portStr, err := net.SplitHostPort(addr)
	if addrErr, ok := err.(*net.AddrError); ok && addrErr.Err == missingPort {
		host, port = addr, DefaultPort
	} else if err != nil {
		return "", err
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

// JoinServerWithDialer is similar to JoinServer but using a Dialer.
func (c *Client) JoinServerWithDialer(d *net.Dialer, addr string) (err error) {
	addr, err = parseAddress(d.Resolver, addr)
	if err != nil {
		return fmt.Errorf("parse address error: %w", err)
	}
	return c.join(d, addr)
}

func (c *Client) join(d *net.Dialer, addr string) (err error) {
	conn, err := d.Dial("tcp", addr)
	if err != nil {
		err = fmt.Errorf("bot: connect server fail: %v", err)
		return err
	}
	//Set Conn
	c.conn = mcnet.WrapConn(conn)

	//Get Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		err = fmt.Errorf("bot: connect server fail: %v", err)
		return err
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		err = fmt.Errorf("bot: connect server fail: %v", err)
		return err
	}

	//Handshake
	err = c.conn.WritePacket(
		//Handshake Packet
		pk.Marshal(
			0x00,                       //Handshake packet ID
			pk.VarInt(ProtocolVersion), //Protocol version
			pk.String(host),            //Server's address
			pk.UnsignedShort(port),
			pk.Byte(2),
		))
	if err != nil {
		err = fmt.Errorf("bot: send handshake packect fail: %v", err)
		return
	}

	//Login
	err = c.conn.WritePacket(
		//LoginStart Packet
		pk.Marshal(0, pk.String(c.Name)))
	if err != nil {
		err = fmt.Errorf("bot: send login start packect fail: %v", err)
		return
	}

	for {
		//Recive Packet
		var pack pk.Packet
		pack, err = c.conn.ReadPacket()
		if err != nil {
			err = fmt.Errorf("bot: recv packet for Login fail: %v", err)
			return
		}

		//Handle Packet
		switch pack.ID {
		case 0x00: //Disconnect
			var reason pk.String
			err = pack.Scan(&reason)
			if err != nil {
				err = fmt.Errorf("bot: read Disconnect message fail: %v", err)
			} else {
				err = fmt.Errorf("bot: connect disconnected by server: %s", reason)
			}
			return
		case 0x01: //Encryption Request
			if err := handleEncryptionRequest(c, pack); err != nil {
				return fmt.Errorf("bot: encryption fail: %v", err)
			}
		case 0x02: //Login Success
			// uuid, l := pk.UnpackString(pack.Data)
			// name, _ := unpackString(pack.Data[l:])
			return nil
		case 0x03: //Set Compression
			var threshold pk.VarInt
			if err := pack.Scan(&threshold); err != nil {
				return fmt.Errorf("bot: set compression fail: %v", err)
			}
			c.conn.SetThreshold(int(threshold))
		case 0x04: //Login Plugin Request
			if err := handlePluginPacket(c, pack); err != nil {
				return fmt.Errorf("bot: handle plugin packet fail: %v", err)
			}
		}
	}
}

// Conn return the MCConn of the Client.
// Only used when you want to handle the packets by yourself
func (c *Client) Conn() *mcnet.Conn {
	return c.conn
}

// SendMessage sends a chat message.
func (c *Client) SendMessage(msg string) error {
	return c.conn.WritePacket(
		pk.Marshal(
			data.ChatServerbound,
			pk.String(msg),
		),
	)
}
