package server

import (
	"container/heap"
	"errors"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"sync"
	"time"
)

// keepAliveInterval represents the interval when the server sends keep alive
const keepAliveInterval = time.Second * 15

// keepAliveWaitInterval represents how long does the player expired
const keepAliveWaitInterval = time.Second * 30

type KeepAlive struct {
	list     keepAliveHeap
	listLock sync.Mutex

	waitList  keepAliveHeap
	waitItem  keepAliveItem
	waitLock  sync.Mutex
	waitTimer *time.Timer

	onPlayerExpire func(p *Player)
}

func NewKeepAlive() (k KeepAlive) {
	heap.Init(&k.list)
	heap.Init(&k.waitList)
	return
}

func (k *KeepAlive) AddPlayer(p *Player) {
	k.listLock.Lock()
	defer k.listLock.Unlock()

	now := time.Now()
	heap.Push(&k.list, keepAliveItem{
		player:   p,
		lastTick: now,
	})
}

func (k *KeepAlive) RemovePlayer(p *Player) {
	k.listLock.Lock()
	defer k.listLock.Unlock()

	for i, v := range k.list {
		if v.player.UUID == p.UUID {
			heap.Remove(&k.list, i)
			return
		}
	}
	panic("KeepAlive: Fail to remove player: " + p.UUID.String() + ": not found")
}

func (k *KeepAlive) Run() {
	listTimer := time.NewTimer(keepAliveInterval)
	k.waitTimer = time.NewTimer(keepAliveWaitInterval)
	var listItem keepAliveItem
	// The Notchian server uses a system-dependent time in milliseconds to generate the keep alive ID value.
	// We don't do that here for security reason.
	var keepAliveID int64

	for {
		select {
		case now := <-listTimer.C:
			if listItem.player == nil {
				listItem = getEarliest(&k.list, &k.listLock, listTimer, keepAliveInterval)
				break
			}

			// Send Clientbound KeepAlive packet.
			listItem.player.WritePacket(Packet757(pk.Marshal(
				packetid.ClientboundKeepAlive,
				pk.Long(keepAliveID),
			)))
			keepAliveID++

			// Clientbound KeepAlive packet is sent, add the player to waiting list.
			k.waitLock.Lock()
			heap.Push(&k.waitList, keepAliveItem{
				player:   listItem.player,
				lastTick: now,
			})
			k.waitLock.Unlock()

			// Wait for next earliest player
			listItem = getEarliest(&k.list, &k.listLock, listTimer, keepAliveInterval)

		case <-k.waitTimer.C:
			if k.waitItem.player == nil {
				k.waitItem = getEarliest(&k.waitList, &k.waitLock, k.waitTimer, keepAliveWaitInterval)
				break
			}
			k.onPlayerExpire(k.waitItem.player)
			k.waitItem = getEarliest(&k.waitList, &k.waitLock, k.waitTimer, keepAliveWaitInterval)
		}
	}
}

func (k *KeepAlive) TickPlayer(player *Player, packet Packet757) error {
	var KeepAliveID pk.Long
	if err := pk.Packet(packet).Scan(&KeepAliveID); err != nil {
		return err
	}
	k.waitLock.Lock()
	defer k.waitLock.Unlock()

	if k.waitItem.player.UUID == player.UUID {
		if k.waitTimer.Stop() {
			k.waitItem = getEarliest(&k.waitList, &k.waitLock, k.waitTimer, keepAliveWaitInterval)
		} else {
			<-k.waitTimer.C
			return errors.New("keepalive: player " + player.UUID.String() + " is already expired")
		}
	} else {
		for i, v := range k.waitList {
			if v.player.UUID == player.UUID {
				heap.Remove(&k.waitList, i)
				goto Success
			}
		}
		return errors.New("keepalive: player " + player.UUID.String() + " not found")
	}
Success:
	heap.Push(&k.list, keepAliveItem{
		player:   player,
		lastTick: time.Now(),
	})
	return nil
}

func getEarliest(list *keepAliveHeap, listLock *sync.Mutex, timer *time.Timer, interval time.Duration) (item keepAliveItem) {
	listLock.Lock()
	defer listLock.Unlock()

	if list.Len() == 0 {
		timer.Reset(interval)
	} else {
		item = heap.Pop(list).(keepAliveItem)
		timer.Reset(time.Until(item.lastTick.Add(interval)))
	}
	return
}

type keepAliveItem struct {
	player   *Player
	lastTick time.Time
}

type keepAliveHeap []keepAliveItem

func (k *keepAliveHeap) Len() int           { return len(*k) }
func (k *keepAliveHeap) Less(i, j int) bool { return (*k)[i].lastTick.Before((*k)[j].lastTick) }
func (k *keepAliveHeap) Swap(i, j int)      { (*k)[i], (*k)[j] = (*k)[j], (*k)[i] }
func (k *keepAliveHeap) Push(x interface{}) { *k = append(*k, x.(keepAliveItem)) }
func (k *keepAliveHeap) Pop() (x interface{}) {
	i := len(*k) - 1
	x = (*k)[i]
	*k = (*k)[:i]
	return
}
