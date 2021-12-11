package server

import (
	"container/list"
	"sync"

	"github.com/Tnze/go-mc/chat"
)

// PlayerList is an implement of ListPingHandler based on linked-list.
// This struct should not be copied after used.
type PlayerList struct {
	name        string
	protocol    int
	maxPlayer   int
	description *chat.Message
	players     *list.List
	// Only the linked-list is protected by this Mutex.
	// Because others field never change after created.
	playersLock sync.Mutex
}

// NewPlayerList create a PlayerList which implement ListPingHandler.
func NewPlayerList(name string, protocol, maxPlayers int, motd *chat.Message) *PlayerList {
	return &PlayerList{
		name:        name,
		protocol:    protocol,
		maxPlayer:   maxPlayers,
		description: motd,
		players:     list.New(),
	}
}

// TryInsert trying to insert player into PlayerList.
// Return nil if the server is full (length of list larger than maxPlayers),
// otherwise return a function which is used to remove the player from PlayerList
func (p *PlayerList) TryInsert(player PlayerSample) (remove func()) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()

	if p.players.Len() >= p.maxPlayer {
		return nil
	}

	elem := p.players.PushBack(player)
	return func() {
		p.playersLock.Lock()
		p.players.Remove(elem)
		p.playersLock.Unlock()
	}
}

func (p *PlayerList) Name() string {
	return p.name
}

func (p *PlayerList) Protocol() int {
	return p.protocol
}

func (p *PlayerList) MaxPlayer() int {
	return p.maxPlayer
}

func (p *PlayerList) OnlinePlayer() int {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	return p.players.Len()
}

func (p *PlayerList) PlayerSamples() (sample []PlayerSample) {
	p.playersLock.Lock()
	defer p.playersLock.Unlock()
	// Up to 10 players can be returned
	length := p.players.Len()
	if length > 10 {
		length = 10
	}
	sample = make([]PlayerSample, length)
	v := p.players.Front()
	for i := 0; i < length; i++ {
		sample[i] = v.Value.(PlayerSample)
		v = v.Next()
	}
	return
}

func (p *PlayerList) Description() *chat.Message {
	return p.description
}
