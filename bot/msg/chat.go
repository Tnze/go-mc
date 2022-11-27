package msg

import (
	"fmt"
	"io"
	"time"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/google/uuid"
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
	c.Events.AddListener(bot.PacketHandler{
		Priority: 64, ID: packetid.ClientboundPlayerChat,
		F: func(packet pk.Packet) error {
			var message PlayerMessage
			if err := packet.Scan(&message); err != nil {
				return err
			}

			var content chat.Message
			if message.content.formatted != nil {
				content = *message.content.formatted
			} else {
				content = chat.Text(message.content.plainMsg)
			}

			ct := p.WorldInfo.RegistryCodec.ChatType.FindByID(message.chatType.ID)
			if ct == nil {
				return fmt.Errorf("chat type %d not found", message.chatType.ID)
			}

			msg := (*chat.Type)(&message.chatType).Decorate(content, &ct.Chat)
			return handler(msg)
		},
	})
}

type PlayerMessage struct {
	// SignedMessageHeader
	signature []byte
	sender    uuid.UUID
	// MessageSignature
	msgSignature []byte
	// SignedMessageBody
	content      msgContent
	timestamp    time.Time
	salt         int64
	prevMessages []prevMsg
	// Optional<Component>
	unsignedContent *chat.Message
	// FilterMask
	filterType int32
	filterSet  pk.BitSet
	// ChatType
	chatType chatType
}

func (p *PlayerMessage) String() string {
	return p.content.plainMsg
}

func (p *PlayerMessage) ReadFrom(r io.Reader) (n int64, err error) {
	var hasMsgSign, hasUnsignedContent pk.Boolean
	var timestamp pk.Long
	var unsignedContent chat.Message
	n, err = pk.Tuple{
		&hasMsgSign,
		pk.Opt{
			Has:   &hasMsgSign,
			Field: (*pk.ByteArray)(&p.signature),
		},
		(*pk.UUID)(&p.sender),
		(*pk.ByteArray)(&p.msgSignature),
		&p.content,
		&timestamp,
		(*pk.Long)(&p.salt),
		pk.Array(&p.prevMessages),
		&hasUnsignedContent,
		pk.Opt{
			Has:   &hasUnsignedContent,
			Field: &unsignedContent,
		},
		(*pk.VarInt)(&p.filterType),
		pk.Opt{
			Has:   func() bool { return p.filterType == 2 },
			Field: &p.filterSet,
		},
		&p.chatType,
	}.ReadFrom(r)
	if err != nil {
		return
	}
	p.timestamp = time.UnixMilli(int64(timestamp))
	if hasUnsignedContent {
		p.unsignedContent = &unsignedContent
	}
	return n, err
}

type msgContent struct {
	plainMsg  string
	formatted *chat.Message
}

func (m *msgContent) ReadFrom(r io.Reader) (n int64, err error) {
	var hasFormatted pk.Boolean
	n1, err := (*pk.String)(&m.plainMsg).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := hasFormatted.ReadFrom(r)
	if err != nil {
		return n1 + n2, err
	}
	if hasFormatted {
		m.formatted = new(chat.Message)
		n3, err := m.formatted.ReadFrom(r)
		return n1 + n2 + n3, err
	}
	return n1 + n2, err
}

type prevMsg struct {
	sender    uuid.UUID
	signature []byte
}

func (p *prevMsg) ReadFrom(r io.Reader) (n int64, err error) {
	return pk.Tuple{
		(*pk.UUID)(&p.sender),
		(*pk.ByteArray)(&p.signature),
	}.ReadFrom(r)
}

type chatType chat.Type

func (c *chatType) ReadFrom(r io.Reader) (n int64, err error) {
	var hasTargetName pk.Boolean
	n1, err := (*pk.VarInt)(&c.ID).ReadFrom(r)
	if err != nil {
		return n1, err
	}
	n2, err := c.SenderName.ReadFrom(r)
	if err != nil {
		return n1 + n2, err
	}
	n3, err := hasTargetName.ReadFrom(r)
	if err != nil {
		return n1 + n2 + n3, err
	}
	if hasTargetName {
		c.TargetName = new(chat.Message)
		n4, err := c.TargetName.ReadFrom(r)
		return n1 + n2 + n3 + n4, err
	}
	return n1 + n2 + n3, nil
}
