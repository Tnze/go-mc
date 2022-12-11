package msg

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"fmt"
	"time"

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
			Priority: 64, ID: packetid.ClientboundPlayerChat,
			F: func(packet pk.Packet) error {
				var (
					sender          pk.UUID
					index           pk.VarInt
					signature       pk.Option[sign.Signature, *sign.Signature]
					body            sign.MessageBody
					unsignedContent pk.Option[chat.Message, *chat.Message]
					filter          sign.FilterMask
					chatType        chat.Type
				)
				if err := packet.Scan(&sender, &index, &signature, &body, &unsignedContent, &filter, &chatType); err != nil {
					return err
				}

				var content chat.Message
				if unsignedContent.Has {
					content = unsignedContent.Val
				} else {
					content = chat.Text(body.PlainMsg)
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

// SendMessage send chat message to server.
// Currently only support offline-mode or "Not Secure" chat
func (m *Manager) SendMessage(msg string) error {
	if len(msg) > 256 {
		return errors.New("message length greater than 256")
	}

	var salt int64
	if err := binary.Read(rand.Reader, binary.BigEndian, &salt); err != nil {
		return err
	}

	err := m.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundChat,
		pk.String(msg),
		pk.Long(time.Now().UnixMilli()),
		pk.Long(salt),
		pk.Boolean(false), // signature
		sign.HistoryUpdate{
			Acknowledged: pk.NewFixedBitSet(20),
		},
	))
	return err
}
