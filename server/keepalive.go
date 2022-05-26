package server

import (
	"container/list"
	"context"
	"errors"
	"time"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server/ecs"
)

// keepAliveInterval represents the interval when the server sends keep alive
const keepAliveInterval = time.Second * 15

// keepAliveWaitInterval represents how long does the player expired
const keepAliveWaitInterval = time.Second * 30

type ClientDelay struct {
	Delay time.Duration
}

type KeepAlive struct {
	join chan *Client
	quit chan *Client
	tick chan *Client

	pingList  *list.List
	waitList  *list.List
	listIndex map[*Client]*list.Element
	listTimer *time.Timer
	waitTimer *time.Timer
	// The Notchian server uses a system-dependent time in milliseconds to generate the keep alive ID value.
	// We don't do that here for security reason.
	keepAliveID int64

	updatePlayerDelay []func(p *Client, delay time.Duration)
}

func NewKeepAlive() (k *KeepAlive) {
	return &KeepAlive{
		join:        make(chan *Client),
		quit:        make(chan *Client),
		tick:        make(chan *Client),
		pingList:    list.New(),
		waitList:    list.New(),
		listIndex:   make(map[*Client]*list.Element),
		listTimer:   time.NewTimer(keepAliveInterval),
		waitTimer:   time.NewTimer(keepAliveWaitInterval),
		keepAliveID: 0,
	}
}

func (k *KeepAlive) AddPlayerDelayUpdateHandler(f func(p *Client, delay time.Duration)) {
	k.updatePlayerDelay = append(k.updatePlayerDelay, f)
}

// Init implement Component for KeepAlive
func (k *KeepAlive) Init(g *Game) {
	ecs.Register[ClientDelay, *ecs.HashMapStorage[ClientDelay]](g.World)
	k.AddPlayerDelayUpdateHandler(func(p *Client, delay time.Duration) {
		c := ClientDelay{Delay: delay}
		ecs.GetComponent[ClientDelay](g.World).SetValue(p.Index, c)
	})
	g.AddHandler(&PacketHandler{
		ID: packetid.ServerboundKeepAlive,
		F: func(client *Client, player *Player, packet Packet758) error {
			var KeepAliveID pk.Long
			if err := pk.Packet(packet).Scan(&KeepAliveID); err != nil {
				return err
			}
			k.tick <- client
			return nil
		},
	})
}

// Run implement Component for KeepAlive
func (k *KeepAlive) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case p := <-k.join:
			k.pushPlayer(p)
		case p := <-k.quit:
			k.removePlayer(p)
		case now := <-k.listTimer.C:
			k.pingPlayer(now)
		case <-k.waitTimer.C:
			k.kickPlayer()
		case p := <-k.tick:
			k.tickPlayer(p)
		}
	}
}

// ClientJoin implement Component for KeepAlive
func (k *KeepAlive) ClientJoin(client *Client, _ *Player) { k.join <- client }

// ClientLeft implement Component for KeepAlive
func (k *KeepAlive) ClientLeft(client *Client, _ *Player) { k.quit <- client }

func (k KeepAlive) pushPlayer(p *Client) {
	k.listIndex[p] = k.pingList.PushBack(
		keepAliveItem{player: p, t: time.Now()},
	)
}

func (k *KeepAlive) removePlayer(p *Client) {
	elem := k.listIndex[p]
	delete(k.listIndex, p)
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
		p := k.pingList.Remove(elem).(keepAliveItem).player
		// Send Clientbound KeepAlive packet.
		p.WritePacket(Packet758(pk.Marshal(
			packetid.ClientboundKeepAlive,
			pk.Long(k.keepAliveID),
		)))
		k.keepAliveID++
		// Clientbound KeepAlive packet is sent, move the player to waiting list.
		k.listIndex[p] = k.waitList.PushBack(
			keepAliveItem{player: p, t: now},
		)
	}
	// Wait for next earliest player
	keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
}

func (k *KeepAlive) tickPlayer(p *Client) {
	elem, ok := k.listIndex[p]
	if !ok {
		p.PutErr(errors.New("keepalive: fail to tick player: client not found"))
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
		f(p, delay)
	}
	// move the player to ping list
	k.listIndex[p] = k.pingList.PushBack(
		keepAliveItem{player: p, t: now},
	)
}

func (k *KeepAlive) kickPlayer() {
	if elem := k.waitList.Front(); elem != nil {
		player := k.waitList.Remove(elem).(keepAliveItem).player
		k.waitList.Remove(elem)
		player.PutErr(errors.New("keepalive: client did not response"))
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
	player *Client
	t      time.Time
}
