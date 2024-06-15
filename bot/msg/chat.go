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

// The Manager is used to receive and send chat messages.
type Manager struct {
	c      *bot.Client
	p      *basic.Player
	pl     *playerlist.PlayerList
	events EventsHandler

	sign.SignatureCache
}

// New returns a new chat manager.
func New(c *bot.Client, p *basic.Player, pl *playerlist.PlayerList, events EventsHandler) *Manager {
	m := &Manager{
		c:              c,
		p:              p,
		pl:             pl,
		events:         events,
		SignatureCache: sign.NewSignatureCache(),
	}
	if events.SystemChat != nil {
		c.Events.AddListener(bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundSystemChat,
			F: m.handleSystemChat,
		})
	}
	if events.PlayerChatMessage != nil {
		c.Events.AddListener(bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerChat,
			F: m.handlePlayerChat,
		})
	}
	if events.DisguisedChat != nil {
		c.Events.AddListener(bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundDisguisedChat,
			F: m.handleDisguisedChat,
		})
	}
	return m
}

func (m *Manager) handleSystemChat(p pk.Packet) error {
	var msg chat.Message
	var overlay pk.Boolean
	if err := p.Scan(&msg, &overlay); err != nil {
		return err
	}
	return m.events.SystemChat(msg, bool(overlay))
}

func (m *Manager) handlePlayerChat(packet pk.Packet) error {
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
		return InvalidChatPacket{err}
	}
	senderInfo, ok := m.pl.PlayerInfos[uuid.UUID(sender)]
	if !ok {
		return InvalidChatPacket{ErrUnknownPlayer}
	}
	ct := m.c.Registries.ChatType.FindByID(chatType.ID)
	if ct == nil {
		return InvalidChatPacket{ErrUnknwonChatType}
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
			return ErrValidationFailed
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
	msg := chatType.Decorate(content, &ct.Chat)
	return m.events.PlayerChatMessage(msg, validated)
}

func (m *Manager) handleDisguisedChat(packet pk.Packet) error {
	var (
		message  chat.Message
		chatType chat.Type
	)
	if err := packet.Scan(&message, &chatType); err != nil {
		return err
	}

	ct := m.c.Registries.ChatType.FindByID(chatType.ID)
	if ct == nil {
		return InvalidChatPacket{ErrUnknwonChatType}
	}
	msg := chatType.Decorate(message, &ct.Chat)

	return m.events.DisguisedChat(msg)
}

// SendMessage send chat message to server.
// Doesn't support sending message with signature currently.
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

// SendMessage send chat message to server.
// Doesn't support sending message with signature currently.
func (m *Manager) SendCommand(command string) error {
	if len(command) > 256 {
		return errors.New("message length greater than 256")
	}
	var salt int64
	if err := binary.Read(rand.Reader, binary.BigEndian, &salt); err != nil {
		return err
	}

	err := m.c.Conn.WritePacket(pk.Marshal(
		packetid.ServerboundChatCommand,
		pk.String(command),
		pk.Long(time.Now().UnixMilli()),
		pk.Long(salt),
		pk.Ary[pk.VarInt]{Ary: []pk.Tuple{}},
		sign.HistoryUpdate{
			Acknowledged: pk.NewFixedBitSet(20),
		},
	))
	return err
}

type InvalidChatPacket struct {
	err error
}

func (i InvalidChatPacket) Error() string {
	if i.err == nil {
		return "invalid chat packet"
	}
	return "invalid chat packet: " + i.err.Error()
}

func (i InvalidChatPacket) Unwrap() error {
	return i.err
}

var (
	ErrUnknownPlayer          = errors.New("unknown player")
	ErrUnknwonChatType        = errors.New("unknown chat type")
	ErrValidationFailed error = bot.DisconnectErr(chat.TranslateMsg("multiplayer.disconnect.chat_validation_failed"))
)
