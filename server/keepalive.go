package server

import (
	"container/list"
	"context"
	"errors"
	"github.com/Tnze/go-mc/chat"
	"time"
)

// keepAliveInterval represents the interval when the server sends keep alive
const keepAliveInterval = time.Second * 15

// keepAliveWaitInterval represents how long does the player expired
const keepAliveWaitInterval = time.Second * 30

type KeepAliveClient interface {
	SendKeepAlive(id int64)
	SendDisconnect(reason chat.Message)
}

type KeepAlive struct {
	join chan KeepAliveClient
	quit chan KeepAliveClient
	tick chan KeepAliveClient

	pingList  *list.List
	waitList  *list.List
	listIndex map[KeepAliveClient]*list.Element
	listTimer *time.Timer
	waitTimer *time.Timer
	// The Notchian server uses a system-dependent time in milliseconds to generate the keep alive ID value.
	// We don't do that here for security reason.
	keepAliveID int64

	updatePlayerDelay []func(p KeepAliveClient, delay time.Duration)
}

func NewKeepAlive[C KeepAliveClient]() (k *KeepAlive) {
	return &KeepAlive{
		join:        make(chan KeepAliveClient),
		quit:        make(chan KeepAliveClient),
		tick:        make(chan KeepAliveClient),
		pingList:    list.New(),
		waitList:    list.New(),
		listIndex:   make(map[KeepAliveClient]*list.Element),
		listTimer:   time.NewTimer(keepAliveInterval),
		waitTimer:   time.NewTimer(keepAliveWaitInterval),
		keepAliveID: 0,
	}
}

func (k *KeepAlive) AddPlayerDelayUpdateHandler(f func(c KeepAliveClient, delay time.Duration)) {
	k.updatePlayerDelay = append(k.updatePlayerDelay, f)
}

func (k *KeepAlive) ClientJoin(client KeepAliveClient) { k.join <- client }
func (k *KeepAlive) ClientTick(client KeepAliveClient) { k.tick <- client }
func (k *KeepAlive) ClientLeft(client KeepAliveClient) { k.quit <- client }

// Run implement Component for KeepAlive
func (k *KeepAlive) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case c := <-k.join:
			k.pushPlayer(c)
		case c := <-k.quit:
			k.removePlayer(c)
		case now := <-k.listTimer.C:
			k.pingPlayer(now)
		case <-k.waitTimer.C:
			k.kickPlayer()
		case c := <-k.tick:
			k.tickPlayer(c)
		}
	}
}

func (k KeepAlive) pushPlayer(c KeepAliveClient) {
	k.listIndex[c] = k.pingList.PushBack(
		keepAliveItem{player: c, t: time.Now()},
	)
}

func (k *KeepAlive) removePlayer(c KeepAliveClient) {
	elem := k.listIndex[c]
	delete(k.listIndex, c)
	if elem.Prev() == nil {
		// At present, it is difficult to distinguish
		// which linked list the player is in,
		// so both timers will be reset
		defer keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
		defer keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
	}
	k.pingList.Remove(elem)
	k.waitList.Remove(elem)
}

func (k *KeepAlive) pingPlayer(now time.Time) {
	if elem := k.pingList.Front(); elem != nil {
		c := k.pingList.Remove(elem).(keepAliveItem).player
		// Send Clientbound KeepAlive packet.
		c.SendKeepAlive(k.keepAliveID)
		k.keepAliveID++
		// Clientbound KeepAlive packet is sent, move the player to waiting list.
		k.listIndex[c] = k.waitList.PushBack(
			keepAliveItem{player: c, t: now},
		)
	}
	// Wait for next earliest player
	keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
}

func (k *KeepAlive) tickPlayer(c KeepAliveClient) {
	elem, ok := k.listIndex[c]
	if !ok {
		panic(errors.New("keepalive: fail to tick player: client not found"))
		return
	}
	if elem.Prev() == nil {
		if !k.waitTimer.Stop() {
			<-k.waitTimer.C
		}
		defer keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
	}
	// update delay of player
	now := time.Now()
	delay := now.Sub(k.waitList.Remove(elem).(keepAliveItem).t)
	for _, f := range k.updatePlayerDelay {
		f(c, delay)
	}
	// move the player to ping list
	k.listIndex[c] = k.pingList.PushBack(
		keepAliveItem{player: c, t: now},
	)
}

func (k *KeepAlive) kickPlayer() {
	if elem := k.waitList.Front(); elem != nil {
		c := k.waitList.Remove(elem).(keepAliveItem).player
		k.waitList.Remove(elem)
		c.SendDisconnect(chat.TranslateMsg("disconnect.timeout"))
	}
	keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
}

func keepAliveSetTimer(l *list.List, timer *time.Timer, interval time.Duration) {
	if first := l.Front(); first != nil {
		item := first.Value.(keepAliveItem)
		interval -= time.Since(item.t)
		if interval < 0 {
			interval = 0
		}
	}
	timer.Reset(interval)
	return
}

type keepAliveItem struct {
	player KeepAliveClient
	t      time.Time
}
