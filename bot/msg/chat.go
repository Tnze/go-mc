package msg

import (
	"encoding/hex"
	"fmt"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/chat/sign"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Manager struct {
	c *bot.Client
	p *basic.Player
}

func New(c *bot.Client, p *basic.Player, events EventsHandler) *Manager {
	attachPlayerMsg(c, p, events.PlayerChatMessage)
	return &Manager{c, p}
}

func attachPlayerMsg(c *bot.Client, p *basic.Player, handler func(msg chat.Message) error) {
	c.Events.AddListener(
		bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerChatHeader,
			F: func(packet pk.Packet) error {
				fmt.Println(packetid.ClientboundPacketID(packet.ID))
				fmt.Println(hex.Dump(packet.Data))
				return nil
			},
		},
		bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerChat,
			F: func(packet pk.Packet) error {
				var message sign.PlayerMessage
				var chatType chat.Type
				fmt.Println(packetid.ClientboundPacketID(packet.ID))
				fmt.Println(hex.Dump(packet.Data))
				if err := packet.Scan(&message, &chatType); err != nil {
					return err
				}

				var content chat.Message
				if message.MessageBody.DecoratedMsg != nil {
					data, _ := message.MessageBody.DecoratedMsg.MarshalJSON()
					if err := content.UnmarshalJSON(data); err != nil {
						return err
					}
				} else {
					content = chat.Text(message.MessageBody.PlainMsg)
				}

				ct := p.WorldInfo.RegistryCodec.ChatType.FindByID(chatType.ID)
				if ct == nil {
					return fmt.Errorf("chat type %d not found", chatType.ID)
				}

				msg := chatType.Decorate(content, &ct.Chat)
				return handler(msg)
			},
		})
}
