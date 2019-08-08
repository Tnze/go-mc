// Package bot implements a simple Minecraft client that can join a server
// or just ping it for getting information.
//
// Runnable example could be found at cmd/ .
package bot

import (
	"fmt"
	"time"

	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ProtocolVersion , the protocol version number of minecraft net protocol
const ProtocolVersion = 498

//Minecraft Version
const MCVersion = "1.14.4"

// PingAndList check server status and list online player.
// Returns a JSON data with server status, and the delay.
//
// For more information for JSON format, see https://wiki.vg/Server_List_Ping#Response
func PingAndList(addr string, port int) ([]byte, time.Duration, error) {
	conn, err := net.DialMC(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, 0, err
	}

	//握手
	err = conn.WritePacket(
		//Handshake Packet
		pk.Marshal(
			0x00,                       //Handshake packet ID
			pk.VarInt(ProtocolVersion), //Protocol version
			pk.String(addr),            //Server's address
			pk.UnsignedShort(port),
			pk.Byte(1),
		))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send handshake packect fail: %v", err)
	}

	//LIST
	//请求服务器状态
	err = conn.WritePacket(pk.Marshal(0))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send list packect fail: %v", err)
	}

	//服务器返回状态
	recv, err := conn.ReadPacket()
	if err != nil {
		return nil, 0, fmt.Errorf("bot: recv list packect fail: %v", err)
	}
	var s pk.String
	err = recv.Scan(&s)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: scan list packect fail: %v", err)
	}

	//PING
	startTime := time.Now()
	err = conn.WritePacket(pk.Marshal(0x01, pk.Long(startTime.Unix())))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send ping packect fail: %v", err)
	}

	recv, err = conn.ReadPacket()
	if err != nil {
		return nil, 0, fmt.Errorf("bot: recv pong packect fail: %v", err)
	}
	var t pk.Long
	err = recv.Scan(&t)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: scan pong packect fail: %v", err)
	}
	if t != pk.Long(startTime.Unix()) {
		return nil, 0, fmt.Errorf("bot: pong packect no match: %v", err)
	}

	return []byte(s), time.Since(startTime), err
}

// JoinServer connect a Minecraft server for playing the game.
func (c *Client) JoinServer(addr string, port int) (err error) {
	//Connect
	c.conn, err = net.DialMC(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		err = fmt.Errorf("bot: connect server fail: %v", err)
		return
	}

	//Handshake
	err = c.conn.WritePacket(
		//Handshake Packet
		pk.Marshal(
			0x00,                       //Handshake packet ID
			pk.VarInt(ProtocolVersion), //Protocol version
			pk.String(addr),            //Server's address
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
			return //switches the connection state to PLAY.
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
func (c *Client) Conn() *net.Conn {
	return c.conn
}
