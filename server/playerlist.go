package server

import (
	"context"
	"errors"
	"sync"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// PlayerList is a player list based on linked-list.
// This struct should not be copied after used.
type PlayerList struct {
	maxPlayer int
	players   map[uuid.UUID]*Player
	// Only the linked-list is protected by this Mutex.
	// Because others field never change after created.
	playersLock sync.Mutex
}

// NewPlayerList create a PlayerList which implement ListPingHandler.
func NewPlayerList(maxPlayers int) *PlayerList {
	return &PlayerList{
		maxPlayer: maxPlayers,
		players:   make(map[uuid.UUID]*Player),
	}
}

func (p *PlayerList) Run(context.Context) {}

func (p *PlayerList) AddPlayer(player *Player) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()

	if len(p.players) >= p.maxPlayer {
		err := player.WritePacket(Packet757(pk.Marshal(
			packetid.ClientboundDisconnect,
			chat.TranslateMsg("multiplayer.disconnect.server_full"),
		)))
		if err != nil {
			player.PutErr(err)
		} else {
			player.PutErr(errors.New("playerlist: server full"))
		}
		return
	}

	p.players[player.UUID] = player
}

func (p *PlayerList) RemovePlayer(player *Player) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	delete(p.players, player.UUID)
}

func (p *PlayerList) MaxPlayer() int {
	return p.maxPlayer
}

func (p *PlayerList) OnlinePlayer() int {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	return len(p.players)
}

func (p *PlayerList) PlayerSamples() (sample []PlayerSample) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	// Up to 10 players can be returned
	sample = make([]PlayerSample, len(p.players))
	var i int
	for _, v := range p.players {
		sample[i] = PlayerSample{
			Name: v.Name,
			ID:   v.UUID,
		}
		i++
		if i >= len(p.players) {
			break
		}
	}
	return
}
