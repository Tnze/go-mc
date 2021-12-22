package server

import (
	"container/list"
	"context"
	"time"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// keepAliveInterval represents the interval when the server sends keep alive
const keepAliveInterval = time.Second * 15

// keepAliveWaitInterval represents how long does the player expired
const keepAliveWaitInterval = time.Second * 30

type KeepAlive struct {
	join chan *Player
	quit chan *Player
	tick chan *Player

	pingList  *list.List
	waitList  *list.List
	listTimer *time.Timer
	waitTimer *time.Timer
	// The Notchian server uses a system-dependent time in milliseconds to generate the keep alive ID value.
	// We don't do that here for security reason.
	keepAliveID int64

	onPlayerExpire func(p *Player)
}

func NewKeepAlive() (k KeepAlive) {
	k.join = make(chan *Player)
	k.quit = make(chan *Player)
	k.tick = make(chan *Player)
	k.pingList = list.New()
	k.waitList = list.New()
	k.listTimer = time.NewTimer(keepAliveInterval)
	k.waitTimer = time.NewTimer(keepAliveWaitInterval)
	return
}

func (k *KeepAlive) AddPlayer(p *Player) {
	k.join <- p
	p.Add(PacketHandler{
		ID: packetid.ServerboundKeepAlive,
		F: func(packet Packet757) error {
			var KeepAliveID pk.Long
			if err := pk.Packet(packet).Scan(&KeepAliveID); err != nil {
				return err
			}
			k.tick <- p
			return nil
		},
	})
}

func (k *KeepAlive) RemovePlayer(p *Player) {
	k.quit <- p
}

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

func (k KeepAlive) pushPlayer(p *Player) {
	k.pingList.PushBack(keepAliveItem{
		player:   p,
		lastTick: time.Now(),
	})
}

//goland:noinspection GoDeferInLoop
func (k *KeepAlive) removePlayer(p *Player) {
	// find player in pingList
	e := k.pingList.Front()
	for e != nil {
		if e.Value.(keepAliveItem).player.UUID == p.UUID {
			if e.Prev() == nil {
				defer keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
			}
			k.pingList.Remove(e)
			return
		}
		e = e.Next()
	}
	// find player in waitList
	e = k.waitList.Front()
	for e != nil {
		if e.Value.(keepAliveItem).player.UUID == p.UUID {
			if e.Prev() == nil {
				defer keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
			}
			k.waitList.Remove(e)
			return
		}
		e = e.Next()
	}
	// player not found
	panic("keepalive: fail to remove player: " + p.UUID.String() + ": not found")
}

func (k *KeepAlive) pingPlayer(now time.Time) {
	if elem := k.pingList.Front(); elem != nil {
		player := elem.Value.(keepAliveItem).player
		// Send Clientbound KeepAlive packet.
		player.WritePacket(Packet757(pk.Marshal(
			packetid.ClientboundKeepAlive,
			pk.Long(k.keepAliveID),
		)))
		k.keepAliveID++
		// Clientbound KeepAlive packet is sent, move the player to waiting list.
		k.pingList.Remove(elem)
		k.waitList.PushBack(keepAliveItem{
			player:   player,
			lastTick: now,
		})
	}
	// Wait for next earliest player
	keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
}

func (k *KeepAlive) tickPlayer(p *Player) {
	elem := k.waitList.Front()
	for elem != nil {
		if elem.Value.(keepAliveItem).player.UUID == p.UUID {
			k.waitList.Remove(elem)
			break
		}
		elem = elem.Next()
	}

	if elem == nil {
		panic("keepalive: fail to tick player: " + p.UUID.String() + " not found")
	}

	if elem.Prev() == nil {
		if !k.waitTimer.Stop() {
			<-k.waitTimer.C
		}
		defer keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
	}
	k.waitList.Remove(elem)
	k.pingList.PushBack(keepAliveItem{
		player:   p,
		lastTick: time.Now(),
	})
}

func (k *KeepAlive) kickPlayer() {
	if elem := k.waitList.Front(); elem != nil {
		player := elem.Value.(keepAliveItem).player
		k.waitList.Remove(elem)
		k.onPlayerExpire(player)
	}
	keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
}

func keepAliveSetTimer(l *list.List, timer *time.Timer, interval time.Duration) {
	if first := l.Front(); first != nil {
		item := first.Value.(keepAliveItem)
		interval -= time.Since(item.lastTick)
		if interval < 0 {
			interval = 0
		}
	}
	timer.Reset(interval)
	return
}

type keepAliveItem struct {
	player   *Player
	lastTick time.Time
}
