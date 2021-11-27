package server

import (
	"github.com/Tnze/go-mc/net"
	"github.com/google/uuid"
)

type GamePlay interface {
	AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn)
}
