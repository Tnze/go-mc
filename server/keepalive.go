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
	p.handlers[packetid.ServerboundKeepAlive] = append(
		p.handlers[packetid.ServerboundKeepAlive],
		func(packet Packet757) error {
			var KeepAliveID pk.Long
			if err := pk.Packet(packet).Scan(&KeepAliveID); err != nil {
				return err
			}
			k.tick <- p
			return nil
		},
	)
}

func (k *KeepAlive) RemovePlayer(p *Player) {
	k.quit <- p
}

func (k *KeepAlive) Run(ctx context.Context) {
	for {
	Select:
		select {
		case <-ctx.Done():
			return

		case p := <-k.join:
			k.pingList.PushBack(keepAliveItem{
				player:   p,
				lastTick: time.Now(),
			})

		case p := <-k.quit:
			// find player in pingList
			e := k.pingList.Front()
			for e != nil {
				if e.Value.(keepAliveItem).player.UUID == p.UUID {
					isFirst := e.Prev() == nil
					k.pingList.Remove(e)
					if isFirst {
						keepAliveSetTimer(k.pingList, k.listTimer, keepAliveInterval)
					}
					break Select
				}
				e = e.Next()
			}
			// find player in waitList
			e = k.waitList.Front()
			for e != nil {
				if e.Value.(keepAliveItem).player.UUID == p.UUID {
					isFirst := e.Prev() == nil
					k.waitList.Remove(e)
					if isFirst {
						keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
					}
					break Select
				}
				e = e.Next()
			}
			// player not found
			panic("keepalive: fail to remove player: " + p.UUID.String() + ": not found")

		case now := <-k.listTimer.C:
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

		case <-k.waitTimer.C:
			if elem := k.waitList.Front(); elem != nil {
				player := elem.Value.(keepAliveItem).player
				k.waitList.Remove(elem)
				k.onPlayerExpire(player)
			}
			keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)

		case p := <-k.tick:
			elem := k.waitList.Front()
			for elem != nil {
				if elem.Value.(keepAliveItem).player.UUID == p.UUID {
					k.waitList.Remove(elem)
					goto Success
				}
				elem = elem.Next()
			}
			panic("keepalive: fail to tick player: " + p.UUID.String() + " not found")
		Success:
			isFirst := elem.Prev() == nil
			k.waitList.Remove(elem)
			if isFirst {
				if !k.waitTimer.Stop() {
					<-k.waitTimer.C
				}
				keepAliveSetTimer(k.waitList, k.waitTimer, keepAliveWaitInterval)
			}
			k.pingList.PushBack(keepAliveItem{
				player:   p,
				lastTick: time.Now(),
			})
		}
	}
}

func keepAliveSetTimer(l *list.List, timer *time.Timer, interval time.Duration) {
	if first := l.Front(); first != nil {
		item := first.Value.(keepAliveItem)
		interval -= time.Since(item.lastTick)
	}
	timer.Reset(interval)
	return
}

type keepAliveItem struct {
	player   *Player
	lastTick time.Time
}
