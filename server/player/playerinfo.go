package player

//import (
//	"context"
//	"io"
//	"time"
//
//	"github.com/Tnze/go-mc/data/packetid"
//	pk "github.com/Tnze/go-mc/net/packet"
//	"github.com/Tnze/go-mc/server"
//	"github.com/Tnze/go-mc/server/ecs"
//)
//
//type PlayerInfo struct {
//	updateDelay chan playerDelayUpdate
//	quit        chan clientAndPlayer
//}
//type clientAndPlayer struct {
//	*server.Client
//	*server.Player
//}
//
//type playerInfoList struct {
//	players ecs.MaskedStorage[server.Player]
//	delays  ecs.MaskedStorage[server.ClientDelay]
//}
//
//func (p playerInfoList) WriteTo(w io.Writer) (n int64, err error) {
//	n, err = pk.VarInt(p.players.Len).WriteTo(w)
//	if err != nil {
//		return
//	}
//	var n1 int64
//	p.players.And(p.delays.BitSetLike).Range(func(eid ecs.Index) {
//		p := playerDelayUpdate{
//			player: p.players.Get(eid),
//			delay:  p.delays.Get(eid).Delay,
//		}
//		n1, err = p.WriteTo(w)
//		n += n1
//		if err != nil {
//			return
//		}
//	})
//	return
//}
//
//type playerDelayUpdate struct {
//	player *server.Player
//	delay  time.Duration
//}
//
//func (p playerDelayUpdate) WriteTo(w io.Writer) (n int64, err error) {
//	return pk.Tuple{
//		pk.UUID(p.player.UUID),
//		pk.VarInt(p.delay.Milliseconds()),
//	}.WriteTo(w)
//}
//
//const (
//	actionAddPlayer = iota
//	actionUpdateGamemode
//	actionUpdateLatency
//	actionUpdateDisplayName
//	actionRemovePlayer
//)
//
//type DelaySource interface {
//	AddPlayerDelayUpdateHandler(f func(c *server.Client, p *server.Player, delay time.Duration))
//}
//
//func NewPlayerInfo(delaySource DelaySource) *PlayerInfo {
//	updateChan := make(chan playerDelayUpdate)
//	p := &PlayerInfo{
//		updateDelay: updateChan,
//		quit:        make(chan clientAndPlayer),
//	}
//	if delaySource != nil {
//		delaySource.AddPlayerDelayUpdateHandler(func(client *server.Client, player *server.Player, delay time.Duration) {
//			updateChan <- playerDelayUpdate{player: player, delay: delay}
//		})
//	}
//	return p
//}
//
//type playerInfoSystemJoin struct{}
//
//func (p playerInfoSystemJoin) Update(w *ecs.World) {
//	clients := ecs.GetComponent[server.Client](w)
//	players := ecs.GetComponent[server.Player](w)
//	delays := ecs.GetComponent[server.ClientDelay](w)
//}
//
//func (p *PlayerInfo) Init(g *server.Game) {
//	var delayBuffer []playerDelayUpdate
//	clients := ecs.GetComponent[server.Client](g.World)
//	players := ecs.GetComponent[server.Player](g.World)
//	delays := ecs.GetComponent[server.ClientDelay](g.World)
//	g.Dispatcher.Add(ecs.FuncSystem(func(client *server.Client, player *server.Player, delay server.ClientDelay) {
//		info := server.ClientDelay{Delay: 0}
//		pack := server.Packet758(pk.Marshal(
//			packetid.ClientboundPlayerInfo,
//			pk.VarInt(actionAddPlayer),
//			pk.VarInt(1),
//
//			pk.UUID(player.UUID),
//			pk.String(player.Name),
//			pk.Array([]pk.FieldEncoder{}),
//			pk.VarInt(profile.Gamemode),
//			pk.VarInt(0),
//			pk.Boolean(false),
//		))
//		delays.Set(client.Index, info)
//		clients.Range(func(eid ecs.Index) {
//			clients.Get(eid).WritePacket(pack)
//		})
//		client.WritePacket(server.Packet758(pk.Marshal(
//			packetid.ClientboundPlayerInfo,
//			pk.VarInt(actionAddPlayer),
//			playerInfoList{players: players, delays: delays},
//		)))
//	}), "PlayerInfoSystem:Join", nil)
//	g.Dispatcher.Add(ecs.FuncSystem(func() {
//		for {
//			select {
//			case cp := <-p.quit:
//				pack := server.Packet758(pk.Marshal(
//					packetid.ClientboundPlayerInfo,
//					pk.VarInt(actionRemovePlayer),
//					pk.VarInt(1),
//					pk.UUID(cp.UUID),
//				))
//				for _, p := range players.list {
//					cp.WritePacket(pack)
//				}
//			case change := <-p.updateDelay:
//				delayBuffer = append(delayBuffer, change)
//			default:
//				if len(delayBuffer) > 0 {
//					pack := server.Packet758(pk.Marshal(
//						packetid.ClientboundPlayerInfo,
//						pk.VarInt(actionUpdateLatency),
//						pk.Array(&delayBuffer),
//					))
//					players.Range(func(eid ecs.Index) {
//						players.Get(eid).(*server.Client).WritePacket(pack)
//					})
//					delayBuffer = delayBuffer[:0]
//				}
//				return
//			}
//		}
//	}), "PlayerInfoSystem", nil)
//}
//
//func (p *PlayerInfo) Run(context.Context)                                     {}
//func (p *PlayerInfo) ClientJoin(client *server.Client, player *server.Player) {}
//func (p *PlayerInfo) ClientLeft(client *server.Client, player *server.Player) {
//	p.quit <- clientAndPlayer{
//		Client: client,
//		Player: player,
//	}
//}
