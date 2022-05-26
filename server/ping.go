package server

import (
	"encoding/base64"
	"encoding/json"
	"errors"
	"image"
	"image/png"
	"strings"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

// ListPingHandler collect server running status info
// which is used to handle client ping and list progress.
type ListPingHandler interface {
	// Name of the server version
	Name() string
	// Protocol number
	Protocol() int
	MaxPlayer() int
	OnlinePlayer() int
	// PlayerSamples is a short list of some player in the server
	PlayerSamples() []PlayerSample

	Description() *chat.Message
	// FavIcon should be a PNG image that is Base64 encoded
	// (without newlines: \n, new lines no longer work since 1.13)
	// and prepended with "data:image/png;base64,".
	//
	// This method can return empty string if no icon is set.
	FavIcon() string
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

	list.Version.Name = s.Name()
	list.Version.Protocol = s.Protocol()
	list.Players.Max = s.MaxPlayer()
	list.Players.Online = s.OnlinePlayer()
	list.Players.Sample = s.PlayerSamples()
	list.Description = s.Description()
	list.FavIcon = s.FavIcon()

	return json.Marshal(list)
}

// PingInfo implement ListPingHandler.
type PingInfo struct {
	name     string
	protocol int
	*PlayerList
	description chat.Message
	favicon     string
}

// NewPingInfo crate a new PingInfo, the icon can be nil.
func NewPingInfo(list *PlayerList, name string, protocol int, motd chat.Message, icon image.Image) (p *PingInfo, err error) {
	var favIcon string
	if icon != nil {
		if !icon.Bounds().Size().Eq(image.Point{X: 64, Y: 64}) {
			return nil, errors.New("icon size is not 64x64")
		}
		// Encode icon into string "data:image/png;base64,......" format
		var sb strings.Builder
		sb.WriteString("data:image/png;base64,")
		w := base64.NewEncoder(base64.StdEncoding, &sb)
		err = png.Encode(w, icon)
		if err != nil {
			return nil, err
		}
		err = w.Close()
		if err != nil {
			return nil, err
		}
		favIcon = sb.String()
	}
	p = &PingInfo{
		name:        name,
		protocol:    protocol,
		PlayerList:  list,
		description: motd,
		favicon:     favIcon,
	}
	return
}

func (p *PingInfo) Name() string {
	return p.name
}

func (p *PingInfo) Protocol() int {
	return p.protocol
}

func (p *PingInfo) FavIcon() string {
	return p.favicon
}

func (p *PingInfo) Description() *chat.Message {
	return &p.description
}
