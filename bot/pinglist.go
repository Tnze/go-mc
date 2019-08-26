package bot

import (
	"fmt"
	"time"

	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// PingAndList check server status and list online player.
// Returns a JSON data with server status, and the delay.
//
// For more information for JSON format, see https://wiki.vg/Server_List_Ping#Response
func PingAndList(addr string, port int) ([]byte, time.Duration, error) {
	conn, err := net.DialMC(fmt.Sprintf("%s:%d", addr, port))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: dial fail: %v", err)
	}
	return pingAndList(addr, port, conn)
}

// PingAndListTimeout PingAndLIstTimeout is the version of PingAndList with max request time.
func PingAndListTimeout(addr string, port int, timeout time.Duration) ([]byte, time.Duration, error) {
	deadLine := time.Now().Add(timeout)

	conn, err := net.DialMCTimeout(fmt.Sprintf("%s:%d", addr, port), timeout)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: dial fail: %v", err)
	}

	err = conn.Socket.SetDeadline(deadLine)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: set deadline fail: %v", err)
	}

	return pingAndList(addr, port, conn)
}

func pingAndList(addr string, port int, conn *net.Conn) ([]byte, time.Duration, error) {
	//握手
	err := conn.WritePacket(
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
