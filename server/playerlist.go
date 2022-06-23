package server

import (
	"sync"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/chat"
)

type PlayerListClient interface {
	SendDisconnect(reason chat.Message)
}

// PlayerList is a player list based on linked-list.
// This struct should not be copied after used.
type PlayerList struct {
	maxPlayer int
	players   map[PlayerListClient]PlayerSample
	// Only the field players is protected by this Mutex.
	// Because others field never change after created.
	playersLock sync.Mutex
}

// NewPlayerList create a PlayerList which implement ListPingHandler.
func NewPlayerList(maxPlayers int) *PlayerList {
	return &PlayerList{
		maxPlayer: maxPlayers,
		players:   make(map[PlayerListClient]PlayerSample),
	}
}

func (p *PlayerList) ClientJoin(client PlayerListClient, player PlayerSample) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()

	if len(p.players) >= p.maxPlayer {
		client.SendDisconnect(chat.TranslateMsg("multiplayer.disconnect.server_full"))
		return
	}

	p.players[client] = player
}

func (p *PlayerList) ClientLeft(client PlayerListClient) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	delete(p.players, client)
}

// CheckPlayer implements LoginChecker for PlayerList
func (p *PlayerList) CheckPlayer(string, uuid.UUID, int32) (ok bool, reason chat.Message) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	if len(p.players) >= p.maxPlayer {
		return false, chat.TranslateMsg("multiplayer.disconnect.server_full")
	}
	return true, chat.Message{}
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
	length := len(p.players)
	if length > 10 {
		length = 10
	}
	sample = make([]PlayerSample, length)
	var i int
	for _, player := range p.players {
		sample[i] = player
		i++
		if i >= length {
			break
		}
	}
	return
}

func (p *PlayerList) Len() int {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	return len(p.players)
}

func (p *PlayerList) Range(f func(PlayerListClient, PlayerSample)) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	for client, player := range p.players {
		f(client, player)
	}
}
