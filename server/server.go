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
