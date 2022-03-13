package server

import (
	"context"
	"io"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type PlayerInfo struct {
	updateDelay chan playerInfo
	join        chan *Player
	quit        chan *Player
	ticker      *time.Ticker
}

type playerInfo struct {
	player *Player
	delay  time.Duration
}

func (p *playerInfo) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.UUID(p.player.UUID),
		pk.String(p.player.Name),
		pk.Array([]pk.FieldEncoder{}),
		pk.VarInt(p.player.Gamemode),
		pk.VarInt(p.delay),
		pk.Boolean(false),
	}.WriteTo(w)
}

type playerInfoList struct {
	list map[uuid.UUID]playerInfo
}

func (p *playerInfoList) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.VarInt(len(p.list)).WriteTo(w)
	if err != nil {
		return
	}
	var n1 int64
	for _, p := range p.list {
		n1, err = p.WriteTo(w)
		n += n1
		if err != nil {
			return
		}
	}
	return
}

type playerDelayUpdate playerInfo

func (p playerDelayUpdate) WriteTo(w io.Writer) (n int64, err error) {
	return pk.Tuple{
		pk.UUID(p.player.UUID),
		pk.VarInt(p.delay.Milliseconds()),
	}.WriteTo(w)
}

const (
	actionAddPlayer = iota
	actionUpdateGamemode
	actionUpdateLatency
	actionUpdateDisplayName
	actionRemovePlayer
)

type PlayerDelaySource interface {
	AddPlayerDelayUpdateHandler(f func(p *Player, delay time.Duration))
}

func NewPlayerInfo(updateFreq time.Duration, delaySource PlayerDelaySource) *PlayerInfo {
	p := &PlayerInfo{
		updateDelay: make(chan playerInfo),
		join:        make(chan *Player),
		quit:        make(chan *Player),
		ticker:      time.NewTicker(updateFreq),
	}
	if delaySource != nil {
		delaySource.AddPlayerDelayUpdateHandler(p.onPlayerDelayUpdate)
	}
	return p
}

func (p *PlayerInfo) Init(*Game) {}

func (p *PlayerInfo) Run(ctx context.Context) {
	players := &playerInfoList{list: make(map[uuid.UUID]playerInfo)}
	var delayBuffer []playerDelayUpdate
	for {
		select {
		case player := <-p.join:
			info := playerInfo{player: player, delay: 0}
			pack := Packet758(pk.Marshal(
				packetid.ClientboundPlayerInfo,
				pk.VarInt(actionAddPlayer),
				pk.VarInt(1),
				&info,
			))
			players.list[player.UUID] = info
			for _, p := range players.list {
				p.player.WritePacket(pack)
			}
			player.WritePacket(Packet758(pk.Marshal(
				packetid.ClientboundPlayerInfo,
				pk.VarInt(actionAddPlayer),
				players,
			)))
		case player := <-p.quit:
			delete(players.list, player.UUID)
			pack := Packet758(pk.Marshal(
				packetid.ClientboundPlayerInfo,
				pk.VarInt(actionRemovePlayer),
				pk.VarInt(1),
				pk.UUID(player.UUID),
			))
			for _, p := range players.list {
				p.player.WritePacket(pack)
			}
		case change := <-p.updateDelay:
			delayBuffer = append(delayBuffer, playerDelayUpdate(change))
		case <-p.ticker.C:
			pack := Packet758(pk.Marshal(
				packetid.ClientboundPlayerInfo,
				pk.VarInt(actionUpdateLatency),
				pk.Array(&delayBuffer),
			))
			for _, p := range players.list {
				p.player.WritePacket(pack)
			}
			delayBuffer = delayBuffer[:0]
		case <-ctx.Done():
			break
		}
	}
}

func (p *PlayerInfo) AddPlayer(player *Player)    { p.join <- player }
func (p *PlayerInfo) RemovePlayer(player *Player) { p.quit <- player }
func (p *PlayerInfo) onPlayerDelayUpdate(player *Player, delay time.Duration) {
	p.updateDelay <- playerInfo{player: player, delay: delay}
}
