package server

import (
	"container/list"
	"sync"
)

// PlayerList is a player list based on linked-list.
// This struct should not be copied after used.
type PlayerList struct {
	maxPlayer int
	players   *list.List
	// Only the linked-list is protected by this Mutex.
	// Because others field never change after created.
	playersLock sync.Mutex
}

// NewPlayerList create a PlayerList which implement ListPingHandler.
func NewPlayerList(maxPlayers int) *PlayerList {
	return &PlayerList{
		maxPlayer: maxPlayers,
		players:   list.New(),
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
