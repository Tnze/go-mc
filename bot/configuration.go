package bot

import (
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

func (c *Client) joinConfiguration(conn *net.Conn) error {
	receiving := "config custom payload"
	for {
		var p pk.Packet
		if err := conn.ReadPacket(&p); err != nil {
			return LoginErr{receiving, err}
		}

		switch packetid.ClientboundPacketID(p.ID) {
		case packetid.ClientboundConfigCustomPayload:
			var channel pk.Identifier
			var data pk.PluginMessageData
			err := p.Scan(&channel, &data)
			if err != nil {
				return LoginErr{"custom payload", err}
			}
			// TODO: Provide configuration custom data handling interface

		case packetid.ClientboundConfigDisconnect:
		case packetid.ClientboundConfigFinishConfiguration:
			err := conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigFinishConfiguration,
			))
			if err != nil {
				return LoginErr{"finish config", err}
			}
			return nil

		case packetid.ClientboundConfigKeepAlive:
			var keepAliveID pk.Long
			err := p.Scan(&keepAliveID)
			if err != nil {
				return LoginErr{"keep alive", err}
			}
			// send it back
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigKeepAlive,
				keepAliveID,
			))
			if err != nil {
				return LoginErr{"keep alive", err}
			}

		case packetid.ClientboundConfigPing:
		case packetid.ClientboundConfigRegistryData:
			var registryCodec nbt.RawMessage
			err := p.Scan(pk.NBT(&registryCodec))
			if err != nil {
				return LoginErr{"registry data", err}
			}
			// TODO: Handle registries

		case packetid.ClientboundConfigResourcePack:
		case packetid.ClientboundConfigUpdateEnabledFeatures:
		case packetid.ClientboundConfigUpdateTags:
		}
	}
}
