package server

import (
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/registry"
	"github.com/Tnze/go-mc/yggdrasil/user"
	"github.com/google/uuid"
)

type ConfigHandler interface {
	AcceptConfig(name string, id uuid.UUID, profilePubKey *user.PublicKey, properties []user.Property, protocol int32, conn *net.Conn) error
}

type Configurations struct {
	Registries registry.NetworkCodec
}

func (c *Configurations) AcceptConfig(name string, id uuid.UUID, profilePubKey *user.PublicKey, properties []user.Property, protocol int32, conn *net.Conn) error {
	err := conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigRegistryData,
		pk.NBT(c.Registries),
	))
	if err != nil {
		return err
	}
	err = conn.WritePacket(pk.Marshal(
		packetid.ClientboundConfigFinishConfiguration,
	))
	if err != nil {
		return err
	}
	// receive ack
	var p pk.Packet
	for {
		err = conn.ReadPacket(&p)
		if err != nil {
			return err
		}
		if packetid.ServerboundPacketID(p.ID) == packetid.ServerboundConfigFinishConfiguration {
			return nil
		}
	}
}

type ConfigFailErr struct {
	reason chat.Message
}

func (c ConfigFailErr) Error() string {
	return "config error: " + c.reason.ClearString()
}
