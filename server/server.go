// Package server provide a minecraft server framework.
// You can build the server you want by combining the various functional modules provided here.
// An example can be found in examples/frameworkServer.
//
// A server is roughly divided into two parts:
//
//	+-----------------------------------------------------------------+
//	|                     Go-MC Server Framework                      |
//	+--------------------------------------+--------------------------+
//	|               Gate                   |         GamePlay         |
//	+--------------------+-----------------+                          |
// 	|    LoginHandler    | ListPingHandler |                          |
// 	+--------------------+--------+--------+-----------+--------------+
// 	| MojangLoginHandler |  Info  |     PlayerList     |  Others....  |
//	+--------------------+--------+--------------------+--------------+
//
// Gate, which is used to respond to the client login request, provide login verification,
// respond to the List Ping Request and providing the online players' information.
//
// Gameplay, which is used to handle all things after a player successfully logs in
// (that is, after the LoginSuccess package is sent),
// and is responsible for functions including player status, chunk management, keep alive, chat, etc.
//
// The implement of Gameplay will provide later.
package server

import "github.com/Tnze/go-mc/net"

const ProtocolVersion = 757

type Server struct {
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
		go s.acceptConn(&conn)
	}
}

func (s *Server) acceptConn(conn *net.Conn) {
	defer conn.Close()
	protocol, intention, err := s.handshake(conn)
	if err != nil {
		return
	}

	switch intention {
	case 1: // list ping
		s.acceptListPing(conn)
	case 2: // login
		name, id, err := s.AcceptLogin(conn, protocol)
		if err != nil {
			return
		}
		s.AcceptPlayer(name, id, protocol, conn)
	}
}
