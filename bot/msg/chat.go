package msg

import (
	"crypto/rand"
	"encoding/binary"
	"errors"
	"time"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/playerlist"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/chat/sign"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

type Manager struct {
	c      *bot.Client
	p      *basic.Player
	pl     *playerlist.PlayerList
	events EventsHandler

	sign.SignatureCache
}

func New(c *bot.Client, p *basic.Player, pl *playerlist.PlayerList, events EventsHandler) *Manager {
	m := &Manager{
		c:              c,
		p:              p,
		pl:             pl,
		events:         events,
		SignatureCache: sign.NewSignatureCache(),
	}
	c.Events.AddListener(
		bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerChat,
			F: m.handlePacket,
		},
	)
	return m
}

func (m *Manager) handlePacket(packet pk.Packet) error {
	var (
		sender          pk.UUID
		index           pk.VarInt
		signature       pk.Option[sign.Signature, *sign.Signature]
		body            sign.PackedMessageBody
		unsignedContent pk.Option[chat.Message, *chat.Message]
		filter          sign.FilterMask
		chatType        chat.Type
	)
	if err := packet.Scan(&sender, &index, &signature, &body, &unsignedContent, &filter, &chatType); err != nil {
		return err
	}

	unpackedMsg, err := body.Unpack(&m.SignatureCache)
	if err != nil {
		return InvalidChatPacket
	}
	senderInfo, ok := m.pl.PlayerInfos[uuid.UUID(sender)]
	if !ok {
		return InvalidChatPacket
	}
	ct := m.p.WorldInfo.RegistryCodec.ChatType.FindByID(chatType.ID)
	if ct == nil {
		return InvalidChatPacket
	}

	var message sign.Message
	if senderInfo.ChatSession != nil {
		message.Prev = sign.Prev{
			Index:   int(index),
			Sender:  uuid.UUID(sender),
			Session: senderInfo.ChatSession.SessionID,
		}
	} else {
		message.Prev = sign.Prev{
			Index:   0,
			Sender:  uuid.UUID(sender),
			Session: uuid.Nil,
		}
	}
	message.Signature = signature.Pointer()
	message.MessageBody = unpackedMsg
	message.Unsigned = unsignedContent.Pointer()
	message.FilterMask = filter

	var validated bool
	if senderInfo.ChatSession != nil {
		if !senderInfo.ChatSession.VerifyAndUpdate(&message) {
			return ValidationFailed
		}
		validated = true
		// store signature into signatureCache
		m.PopOrInsert(signature.Pointer(), message.LastSeen)
	}

	var content chat.Message
	if unsignedContent.Has {
		content = unsignedContent.Val
	} else {
		content = chat.Text(body.PlainMsg)
	}

	if m.events.PlayerChatMessage == nil {
		return nil
	}
	msg := chatType.Decorate(content, &ct.Chat)
	return m.events.PlayerChatMessage(msg, validated)
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

var (
	InvalidChatPacket       = errors.New("invalid chat packet")
	ValidationFailed  error = bot.DisconnectErr(chat.TranslateMsg("multiplayer.disconnect.chat_validation_failed"))
)
