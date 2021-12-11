package server

import (
	"encoding/json"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

type ListPingHandler interface {
	Name() string
	Protocol() int
	MaxPlayer() int
	OnlinePlayer() int
	PlayerSamples() []PlayerSample
	Description() *chat.Message
}

type PlayerSample struct {
	Name string    `json:"name"`
	ID   uuid.UUID `json:"id"`
}

func (s *Server) acceptListPing(conn *net.Conn) {
	var p pk.Packet
	for i := 0; i < 2; i++ { // Ping or List. Only allow check twice
		err := conn.ReadPacket(&p)
		if err != nil {
			return
		}

		switch p.ID {
		case packetid.StatusResponse: //List
			var resp []byte
			resp, err = s.listResp()
			if err != nil {
				break
			}
			err = conn.WritePacket(pk.Marshal(0x00, pk.String(resp)))
		case packetid.StatusPongResponse: //Ping
			err = conn.WritePacket(p)
		}
		if err != nil {
			return
		}
	}
}

func (s *Server) listResp() ([]byte, error) {
	var list struct {
		Version struct {
			Name     string `json:"name"`
			Protocol int    `json:"protocol"`
		} `json:"version"`
		Players struct {
			Max    int            `json:"max"`
			Online int            `json:"online"`
			Sample []PlayerSample `json:"sample"`
		} `json:"players"`
		Description *chat.Message `json:"description"`
		FavIcon     string        `json:"favicon,omitempty"`
	}

	list.Version.Name = s.ListPingHandler.Name()
	list.Version.Protocol = s.ListPingHandler.Protocol()
	list.Players.Max = s.ListPingHandler.MaxPlayer()
	list.Players.Online = s.ListPingHandler.OnlinePlayer()
	list.Players.Sample = s.ListPingHandler.PlayerSamples()
	list.Description = s.ListPingHandler.Description()

	return json.Marshal(list)
}
