package server

import (
	"container/list"
	"context"
	"errors"
	"github.com/google/uuid"
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
	listIndex map[uuid.UUID]*list.Element
	listTimer *time.Timer
	waitTimer *time.Timer
	// The Notchian server uses a system-dependent time in milliseconds to generate the keep alive ID value.
	// We don't do that here for security reason.
	keepAliveID int64

	updatePlayerDelay func(p *Player, delay time.Duration)
}

func NewKeepAlive() (k *KeepAlive) {
	return &KeepAlive{
		join:        make(chan *Player),
		quit:        make(chan *Player),
		tick:        make(chan *Player),
		pingList:    list.New(),
		waitList:    list.New(),
		listIndex:   make(map[uuid.UUID]*list.Element),
		listTimer:   time.NewTimer(keepAliveInterval),
		waitTimer:   time.NewTimer(keepAliveWaitInterval),
		keepAliveID: 0,
	}
}

func (k *KeepAlive) AddPlayerDelayUpdateHandler(f func(p *Player, delay time.Duration)) *KeepAlive {
	if k.updatePlayerDelay != nil {
		panic("add player update handler twice")
	}
	k.updatePlayerDelay = f
	return k
}

// Init implement Component for KeepAlive
func (k *KeepAlive) Init(g *Game) {
	g.AddHandler(&PacketHandler{
		ID: packetid.ServerboundKeepAlive,
		F: func(player *Player, packet Packet757) error {
			var KeepAliveID pk.Long
			if err := pk.Packet(packet).Scan(&KeepAliveID); err != nil {
				return err
			}
			k.tick <- player
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

// AddPlayer implement Component for KeepAlive
func (k *KeepAlive) AddPlayer(player *Player) { k.join <- player }

// RemovePlayer implement Component for KeepAlive
func (k *KeepAlive) RemovePlayer(p *Player) { k.quit <- p }

func (k KeepAlive) pushPlayer(p *Player) {
	k.listIndex[p.UUID] = k.pingList.PushBack(
		keepAliveItem{player: p, t: time.Now()},
	)
}

func (k *KeepAlive) removePlayer(p *Player) {
	elem := k.listIndex[p.UUID]
	delete(k.listIndex, p.UUID)
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
		err := p.WritePacket(Packet757(pk.Marshal(
			packetid.ClientboundKeepAlive,
			pk.Long(k.keepAliveID),
		)))
		if err != nil {
			p.PutErr(err)
			return
		}
		k.keepAliveID++
		// Clientbound KeepAlive packet is sent, move the player to waiting list.
		k.listIndex[p.UUID] = k.waitList.PushBack(
			keepAliveItem{player: p, t: now},
		)
	}
	// Wait for next earliest player
	keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
}

func (k *KeepAlive) tickPlayer(p *Player) {
	elem, ok := k.listIndex[p.UUID]
	if !ok {
		p.PutErr(errors.New("keepalive: fail to tick player: " + p.UUID.String() + " not found"))
		return
	}
	if elem.Prev() == nil {
		if !k.waitTimer.Stop() {
			<-k.waitTimer.C
		}
		defer keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
	}
	// update delay of player
	t := k.waitList.Remove(elem).(keepAliveItem).t
	now := time.Now()
	if k.updatePlayerDelay != nil {
		k.updatePlayerDelay(p, now.Sub(t))
	}
	// move the player to ping list
	k.listIndex[p.UUID] = k.pingList.PushBack(
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
	player *Player
	t      time.Time
}
