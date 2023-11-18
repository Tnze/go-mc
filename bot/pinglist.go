package bot

import (
	"context"
	"errors"
	"fmt"
	"net"
	"os"
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
	conn, err := mcnet.DialMC(addr)
	if err != nil {
		return nil, 0, LoginErr{"dial connection", err}
	}
	return pingAndList(context.Background(), addr, conn)
}

// PingAndListTimeout is the version of PingAndList with max request time.
func PingAndListTimeout(addr string, timeout time.Duration) ([]byte, time.Duration, error) {
	ctx, cancel := context.WithTimeout(context.Background(), timeout)
	defer cancel()
	return PingAndListContext(ctx, addr)
}

func PingAndListContext(ctx context.Context, addr string) ([]byte, time.Duration, error) {
	conn, err := mcnet.DefaultDialer.DialMCContext(ctx, addr)
	if err != nil {
		return nil, 0, err
	}
	return pingAndList(ctx, addr, conn)
}

func pingAndList(ctx context.Context, addr string, conn *mcnet.Conn) (data []byte, delay time.Duration, err error) {
	if deadline, hasDeadline := ctx.Deadline(); hasDeadline {
		if err := conn.Socket.SetDeadline(deadline); err != nil {
			return nil, 0, err
		}
		defer func() {
			// Reset deadline
			if err2 := conn.Socket.SetDeadline(time.Time{}); err2 != nil {
				if err2 == nil {
					err = err2
				}
				return
			}
			// Map error type
			if errors.Is(err, os.ErrDeadlineExceeded) {
				err = context.DeadlineExceeded
			}
		}()
	}
	// Split Host and Port
	host, portStr, err := net.SplitHostPort(addr)
	var port uint64
	if err != nil {
		var addrErr *net.AddrError
		const missingPort = "missing port in address"
		if errors.As(err, &addrErr) && addrErr.Err == missingPort {
			host, port, err = addr, DefaultPort, nil
		} else {
			return nil, 0, LoginErr{"split address", err}
		}
	} else {
		port, err = strconv.ParseUint(portStr, 0, 16)
		if err != nil {
			return nil, 0, LoginErr{"parse port", err}
		}
	}

	const Handshake = 0x00
	// 握手
	err = conn.WritePacket(pk.Marshal(
		Handshake,                  // Handshake packet ID
		pk.VarInt(ProtocolVersion), // Protocol version
		pk.String(host),            // Server's address
		pk.UnsignedShort(port),
		pk.Byte(1),
	))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send handshake packect fail: %v", err)
	}

	// LIST
	// 请求服务器状态
	err = conn.WritePacket(pk.Marshal(
		packetid.ServerboundStatusRequest,
	))
	if err != nil {
		return nil, 0, fmt.Errorf("bot: send list packect fail: %v", err)
	}

	var p pk.Packet
	// 服务器返回状态
	if err := conn.ReadPacket(&p); err != nil {
		return nil, 0, fmt.Errorf("bot: recv list packect fail: %v", err)
	}
	var s pk.String
	err = p.Scan(&s)
	if err != nil {
		return nil, 0, fmt.Errorf("bot: scan list packect fail: %v", err)
	}

	// PING
	startTime := time.Now()
	err = conn.WritePacket(pk.Marshal(
		packetid.ServerboundStatusPingRequest,
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
