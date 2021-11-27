package server

import (
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (s *Server) handshake(conn *net.Conn) (protocol int32, intention int32, err error) {
	var (
		Protocol, Intention pk.VarInt
		ServerAddress       pk.String        // ignored
		ServerPort          pk.UnsignedShort // ignored
	)
	// receive handshake packet
	var p pk.Packet
	err = conn.ReadPacket(&p)
	if err != nil {
		return 0, 0, err
	}
	err = p.Scan(&Protocol, &ServerAddress, &ServerPort, &Intention)
	return int32(Protocol), int32(Intention), err
}
