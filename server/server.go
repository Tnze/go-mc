// Package server provide a minecraft server framework.
// You can build the server you want by combining the various functional modules provided here.
// An example can be found in examples/frameworkServer.
//
// # This package is under rapid development, and any API may be subject to break changes
//
// A server is roughly divided into two parts: Gate and GamePlay
//
//	+------------------------------------------------------------------------------+
//	|                             Go-MC Server Framework                           |
//	|--------------------------------------+---------------------------------------|
//	|               Gate                   |                GamePlay               |
//	|--------------------+-----------------+---------------+-----------------------|
//	|    LoginHandler    |         ListPingHandler         |        Others..       |
//	|--------------------|------------+----+---------------|-----------------------+
//	| MojangLoginHandler |  PingInfo  |     PlayerList     |  [go-mc/server], etc. |
//	+--------------------+------------+--------------------+-----------------------+
//
// Gate, which is used to respond to the client login request, provide login verification,
// respond to the List Ping Request and providing the online players' information.
//
// Gameplay, which is used to handle all things after a player successfully logs in
// (that is, after the LoginSuccess package is sent),
// and is responsible for functions including player status, chunk management, keep alive, chat, etc.
//
// The implement of Gameplay is provided at [go-mc/server]. You can also write your version.
//
// [go-mc/server]: https://github.com/go-mc/server
package server

import (
	"errors"
	"log"

	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

const (
	ProtocolName    = "1.19.3"
	ProtocolVersion = 761
)

type Server struct {
	*log.Logger
	ListPingHandler
	LoginHandler
	GamePlay
}

func (s *Server) Listen(addr string) error {
	listener, err := net.ListenMC(addr)
	if err != nil {
		return err
	}

	for {
		conn, err := listener.Accept()
		if err != nil {
			return err
		}
		go s.AcceptConn(&conn)
	}
}

func (s *Server) AcceptConn(conn *net.Conn) {
	defer conn.Close()
	protocol, intention, err := s.handshake(conn)
	if err != nil {
		return
	}

	switch intention {
	case 1: // list ping
		s.acceptListPing(conn)
	case 2: // login
		name, id, profilePubKey, properties, err := s.AcceptLogin(conn, protocol)
		if err != nil {
			var loginErr LoginFailErr
			if errors.As(err, &loginErr) {
				_ = conn.WritePacket(pk.Marshal(
					packetid.LoginDisconnect,
					loginErr.reason,
				))
			}
			if s.Logger != nil {
				s.Logger.Printf("client %v login error: %v", conn.Socket.RemoteAddr(), err)
			}
			return
		}
		s.AcceptPlayer(name, id, profilePubKey, properties, protocol, conn)
	}
}
