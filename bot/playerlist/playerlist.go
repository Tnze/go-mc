// Package playerlist contains a PlayerList struct that used to manage player information.
//
// The [PlayerList] contains a list of [PlayerInfo] which is received from server when client join.
// The playerlist contains every players' information of name, display name, uuid, gamemode, latency, public key, etc.
// And can be used to render the "TAB List". Other packages may also require playerlist to work,
// for example, the bot/msg package.
package playerlist

import (
	"bytes"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/chat/sign"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/yggdrasil/user"
)

type PlayerList struct {
	PlayerInfos map[uuid.UUID]*PlayerInfo
}

func New(c *bot.Client) *PlayerList {
	pl := PlayerList{
		PlayerInfos: make(map[uuid.UUID]*PlayerInfo),
	}
	c.Events.AddListener(
		bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerInfoUpdate,
			F: pl.handlePlayerInfoUpdatePacket,
		},
		bot.PacketHandler{
			Priority: 64, ID: packetid.ClientboundPlayerInfoRemove,
			F: pl.handlePlayerInfoRemovePacket,
		},
	)
	return &pl
}

func (pl *PlayerList) handlePlayerInfoUpdatePacket(p pk.Packet) error {
	r := bytes.NewReader(p.Data)

	action := pk.NewFixedBitSet(6)
	if _, err := action.ReadFrom(r); err != nil {
		return err
	}

	var length pk.VarInt
	if _, err := length.ReadFrom(r); err != nil {
		return err
	}

	for i := 0; i < int(length); i++ {
		var id pk.UUID
		if _, err := id.ReadFrom(r); err != nil {
			return err
		}

		player, ok := pl.PlayerInfos[uuid.UUID(id)]
		if !ok { // create new player info if not exist
			player = new(PlayerInfo)
			pl.PlayerInfos[uuid.UUID(id)] = player
		}

		// add player
		if action.Get(0) {
			var name pk.String
			var properties []user.Property
			if _, err := (pk.Tuple{&name, pk.Array(&properties)}).ReadFrom(r); err != nil {
				return err
			}
			player.GameProfile = GameProfile{
				ID:         uuid.UUID(id),
				Name:       string(name),
				Properties: properties,
			}
		}
		// initialize chat
		if action.Get(1) {
			var chatSession pk.Option[sign.Session, *sign.Session]
			if _, err := chatSession.ReadFrom(r); err != nil {
				return err
			}
			if chatSession.Has {
				player.ChatSession = chatSession.Pointer()
				player.ChatSession.InitValidate()
			} else {
				player.ChatSession = nil
			}
		}
		// update gamemode
		if action.Get(2) {
			var gamemode pk.VarInt
			if _, err := gamemode.ReadFrom(r); err != nil {
				return err
			}
			player.Gamemode = int32(gamemode)
		}
		// update listed
		if action.Get(3) {
			var listed pk.Boolean
			if _, err := listed.ReadFrom(r); err != nil {
				return err
			}
			player.Listed = bool(listed)
		}
		// update latency
		if action.Get(4) {
			var latency pk.VarInt
			if _, err := latency.ReadFrom(r); err != nil {
				return err
			}
			player.Latency = int32(latency)
		}
		// display name
		if action.Get(5) {
			var displayName pk.Option[chat.Message, *chat.Message]
			if _, err := displayName.ReadFrom(r); err != nil {
				return err
			}
			if displayName.Has {
				player.DisplayName = &displayName.Val
			} else {
				player.DisplayName = nil
			}
		}
	}

	return nil
}

func (pl *PlayerList) handlePlayerInfoRemovePacket(p pk.Packet) error {
	return nil
}

type PlayerInfo struct {
	GameProfile
	ChatSession *sign.Session
	Gamemode    int32
	Latency     int32
	Listed      bool
	DisplayName *chat.Message
}

type GameProfile struct {
	ID         uuid.UUID
	Name       string
	Properties []user.Property
}
