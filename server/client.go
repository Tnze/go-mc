package server

import (
	"strconv"

	"github.com/Tnze/go-mc/internal/queue"
	pk "github.com/Tnze/go-mc/net/packet"
)

// Packet758 is a packet in protocol 757.
// We are using type system to force programmers to update packets.
type (
	Packet758 pk.Packet
	Packet757 pk.Packet
)

type WritePacketError struct {
	Err error
	ID  int32
}

func (s WritePacketError) Error() string {
	return "server: send packet " + strconv.FormatInt(int64(s.ID), 16) + " error: " + s.Err.Error()
}

func (s WritePacketError) Unwrap() error {
	return s.Err
}

type PacketQueue = queue.Queue[pk.Packet]
