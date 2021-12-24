package server

import (
	"context"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
)

type GlobalChat struct {
	msg  chan chatItem
	join chan *Player
	quit chan *Player

	players map[uuid.UUID]*Player
}

type chatItem struct {
	p    *Player
	text string
}

func NewGlobalChat() *GlobalChat {
	return &GlobalChat{
		msg:     make(chan chatItem),
		join:    make(chan *Player),
		quit:    make(chan *Player),
		players: make(map[uuid.UUID]*Player),
	}
}

func (g *GlobalChat) Init(game *Game) {
	game.AddHandler(&PacketHandler{
		ID: packetid.ServerboundChat,
		F: func(player *Player, packet Packet757) error {
			var msg pk.String
			if err := pk.Packet(packet).Scan(&msg); err != nil {
				return err
			}
			text, _ := chat.TransCtrlSeq(string(msg), false)
			g.msg <- chatItem{p: player, text: text}
			return nil
		},
	})
}

func (g *GlobalChat) Run(ctx context.Context) {
	for {
		select {
		case <-ctx.Done():
			return
		case item := <-g.msg:
			packet := Packet757(pk.Marshal(
				packetid.ClientboundChat,
				item.toMessage(),
				pk.Byte(0),
				pk.UUID(item.p.UUID),
			))
			for _, p := range g.players {
				err := p.WritePacket(packet)
				if err != nil {
					p.PutErr(err)
				}
			}
		case p := <-g.join:
			g.players[p.UUID] = p
		case p := <-g.quit:
			delete(g.players, p.UUID)
		}
	}
}

func (g *GlobalChat) AddPlayer(player *Player) { g.join <- player }

func (g *GlobalChat) RemovePlayer(p *Player) { g.quit <- p }

func (c chatItem) toMessage() chat.Message {
	return chat.TranslateMsg(
		"chat.type.text",
		chat.Message{
			Text:       c.p.Name,
			ClickEvent: chat.SuggestCommand("/msg " + c.p.Name),
			HoverEvent: chat.ShowEntity(playerToSNBT(c.p)),
		},
		chat.Text(c.text),
	)
}

func playerToSNBT(p *Player) string {
	var s nbt.StringifiedMessage
	entity := struct {
		ID   string `nbt:"id"`
		Name string `nbt:"name"`
	}{
		ID:   p.UUID.String(),
		Name: p.Name,
	}

	data, err := nbt.Marshal(entity)
	if err != nil {
		panic(err)
	}

	err = nbt.Unmarshal(data, &s)
	if err != nil {
		panic(err)
	}

	return string(s)
}
