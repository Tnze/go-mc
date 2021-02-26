package bot

import (
	"fmt"
	"net"
	"strconv"
	"time"

	"github.com/Tnze/go-mc/data/packetid"
	mcnet "github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// PingAndList check server status and list online player.
// Returns a JSON data with server status, and the delay.
//
// For more information for JSON format, see https://wiki.vg/Server_List_Ping#Response
func PingAndList(addr string) ([]byte, time.Duration, error) {
	addrSrv, err := parseAddress(&net.Resolver{}, addr)
	if err != nil {
		return nil, 0, LoginErr{"parse address", err}
	}

	conn, err := mcnet.DialMC(addrSrv)
	if err != nil {
		return nil, 0, LoginErr{"dial connection", err}
	}
	return pingAndList(addr, conn)
}

// PingAndListTimeout PingAndLIstTimeout is the version of PingAndList with max request time.
func PingAndListTimeout(addr string, timeout time.Duration) ([]byte, time.Duration, error) {
	deadLine := time.Now().Add(timeout)

	addrSrv, err := parseAddress(&net.Resolver{}, addr)
	if err != nil {
		return nil, 0, LoginErr{"parse address", err}
	}

	conn, err := mcnet.DialMCTimeout(addrSrv, timeout)
	if err != nil {
		return nil, 0, err
	}

	err = conn.Socket.SetDeadline(deadLine)
	if err != nil {
		return nil, 0, LoginErr{"set deadline", err}
	}

	return pingAndList(addr, conn)
}

func pingAndList(addr string, conn *mcnet.Conn) ([]byte, time.Duration, error) {
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	if err != nil {
		return nil, 0, LoginErr{"split address", err}
	}
	port, err := strconv.Atoi(portStr)
	if err != nil {
		return nil, 0, LoginErr{"parse port", err}
	}

	const Handshake = 0x00
	//握手
	err = conn.WritePacket(pk.Marshal(
		Handshake,                  //Handshake packet ID
		pk.VarInt(ProtocolVersion), //Protocol version
		pk.String(host),            //Server's address
		pk.UnsignedShort(port),
		pk.Byte(1),
	))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send handshake packect fail: %v", err)
	}

	//LIST
	//请求服务器状态
	err = conn.WritePacket(pk.Marshal(
		packetid.PingStart,
	))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send list packect fail: %v", err)
	}

	var p pk.Packet
	//服务器返回状态
	if err := conn.ReadPacket(&p); err != nil {
		return nil, 0, fmt.Errorf("bot: recv list packect fail: %v", err)
	}
	var s pk.String
	err = p.Scan(&s)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: scan list packect fail: %v", err)
	}

	//PING
	startTime := time.Now()
	err = conn.WritePacket(pk.Marshal(
		packetid.PingServerbound,
		pk.Long(startTime.Unix()),
	))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send ping packect fail: %v", err)
	}

	if err = conn.ReadPacket(&p); err != nil {
		return nil, 0, fmt.Errorf("bot: recv pong packect fail: %v", err)
	}
	var t pk.Long
	err = p.Scan(&t)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: scan pong packect fail: %v", err)
	}
	if t != pk.Long(startTime.Unix()) {
		return nil, 0, fmt.Errorf("bot: pong packect no match: %v", err)
	}

	return []byte(s), time.Since(startTime), err
}
