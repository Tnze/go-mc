package bot

import (
	"unsafe"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/registry"
)

type ConfigData struct {
	Registries   registry.NetworkCodec
	ResourcePack struct {
		URL           string
		Hash          string
		Forced        bool
		PromptMessage *chat.Message // Optional
	}
	FeatureFlags []string
}

type ConfigErr struct {
	Stage string
	Err   error
}

func (l ConfigErr) Error() string {
	return "bot: configuration error: [" + l.Stage + "] " + l.Err.Error()
}

func (l ConfigErr) Unwrap() error {
	return l.Err
}

func (c *Client) joinConfiguration(conn *net.Conn) error {
	receiving := "config custom payload"
	for {
		var p pk.Packet
		if err := conn.ReadPacket(&p); err != nil {
			return ConfigErr{receiving, err}
		}

		switch packetid.ClientboundPacketID(p.ID) {
		case packetid.ClientboundConfigCustomPayload:
			var channel pk.Identifier
			var data pk.PluginMessageData
			err := p.Scan(&channel, &data)
			if err != nil {
				return ConfigErr{"custom payload", err}
			}
			// TODO: Provide configuration custom data handling interface

		case packetid.ClientboundConfigDisconnect:
			var reason chat.Message
			err := p.Scan(&reason)
			if err != nil {
				return ConfigErr{"disconnect", err}
			}
			return ConfigErr{"disconnect", DisconnectErr(reason)}

		case packetid.ClientboundConfigFinishConfiguration:
			err := conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigFinishConfiguration,
			))
			if err != nil {
				return ConfigErr{"finish config", err}
			}
			return nil

		case packetid.ClientboundConfigKeepAlive:
			var keepAliveID pk.Long
			err := p.Scan(&keepAliveID)
			if err != nil {
				return ConfigErr{"keep alive", err}
			}
			// send it back
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigKeepAlive,
				keepAliveID,
			))
			if err != nil {
				return ConfigErr{"keep alive", err}
			}

		case packetid.ClientboundConfigPing:
			var pingID pk.Int
			err := p.Scan(&pingID)
			if err != nil {
				return ConfigErr{"ping", err}
			}
			// send it back
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigPong,
				pingID,
			))
			if err != nil {
				return ConfigErr{"pong", err}
			}

		case packetid.ClientboundConfigRegistryData:
			err := p.Scan(pk.NBT(&c.ConfigData.Registries))
			if err != nil {
				return ConfigErr{"registry data", err}
			}

		case packetid.ClientboundConfigResourcePack:
			var Url, Hash pk.String
			var Forced pk.Boolean
			var PromptMessage pk.Option[chat.Message, *chat.Message]
			err := p.Scan(
				&Url,
				&Hash,
				&Forced,
				&PromptMessage,
			)
			if err != nil {
				return ConfigErr{"resource pack", err}
			}
			c.ConfigData.ResourcePack.URL = string(Url)
			c.ConfigData.ResourcePack.Hash = string(Hash)
			c.ConfigData.ResourcePack.Forced = bool(Forced)
			if PromptMessage.Has {
				c.ConfigData.ResourcePack.PromptMessage = &PromptMessage.Val
			}

		case packetid.ClientboundConfigUpdateEnabledFeatures:
			err := p.Scan(pk.Array((*[]pk.Identifier)(unsafe.Pointer(&c.ConfigData.FeatureFlags))))
			if err != nil {
				return ConfigErr{"update enabled features", err}
			}

		case packetid.ClientboundConfigUpdateTags:
			// TODO: Handle Tags
		}
	}
}
