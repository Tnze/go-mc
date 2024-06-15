package bot

import (
	"fmt"
	"io"
	"unsafe"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/registry"
)

type ConfigHandler interface {
	GetCookie(key pk.Identifier) []byte
	SetCookie(key pk.Identifier, payload []byte)

	PushResourcePack(res ResourcePack)
	PopResourcePack(id pk.UUID)
	PopAllResourcePack()

	SelectDataPacks(packs []DataPack) []DataPack
}

type ResourcePack struct {
	ID            pk.UUID
	URL           string
	Hash          string
	Forced        bool
	PromptMessage *chat.Message // Optional
}

type ConfigData struct {
	Registries   registry.NetworkCodec
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
		case packetid.ClientboundConfigCookieRequest:
			var key pk.Identifier
			err := p.Scan(&key)
			if err != nil {
				return ConfigErr{"cookie request", err}
			}
			cookieContent := c.ConfigHandler.GetCookie(key)
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigCookieResponse,
				pk.ByteArray(cookieContent),
			))
			if err != nil {
				return ConfigErr{"cookie response", err}
			}

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

		case packetid.ClientboundConfigResetChat:
			// TODO

		case packetid.ClientboundConfigRegistryData:
			// err := p.Scan(pk.NBT(&c.ConfigData.Registries))
			// if err != nil {
			// 	return ConfigErr{"registry data", err}
			// }

		case packetid.ClientboundConfigResourcePackPop: // TODO
			var id pk.Option[pk.UUID, *pk.UUID]
			err := p.Scan(&id)
			if err != nil {
				return ConfigErr{"resource pack pop", err}
			}

		case packetid.ClientboundConfigResourcePackPush:
			var id pk.UUID
			var Url, Hash pk.String
			var Forced pk.Boolean
			var PromptMessage pk.Option[chat.Message, *chat.Message]
			err := p.Scan(
				&id,
				&Url,
				&Hash,
				&Forced,
				&PromptMessage,
			)
			if err != nil {
				return ConfigErr{"resource pack", err}
			}
			res := ResourcePack{
				ID:     id,
				URL:    string(Url),
				Hash:   string(Hash),
				Forced: bool(Forced),
			}
			if PromptMessage.Has {
				res.PromptMessage = &PromptMessage.Val
			}
			c.ConfigHandler.PushResourcePack(res)

		case packetid.ClientboundConfigStoreCookie:
			var key pk.Identifier
			var payload pk.ByteArray
			err := p.Scan(&key, &payload)
			if err != nil {
				return ConfigErr{"store cookie", err}
			}
			c.ConfigHandler.SetCookie(key, []byte(payload))

		case packetid.ClientboundConfigTransfer:
			var host pk.String
			var port pk.VarInt
			err := p.Scan(&host, &port)
			if err != nil {
				return ConfigErr{"store cookie", err}
			}
			// TODO: trnasfer to the server

		case packetid.ClientboundConfigUpdateEnabledFeatures:
			err := p.Scan(pk.Array((*[]pk.Identifier)(unsafe.Pointer(&c.ConfigData.FeatureFlags))))
			if err != nil {
				return ConfigErr{"update enabled features", err}
			}

		case packetid.ClientboundConfigUpdateTags:
			// TODO: Handle Tags
		case packetid.ClientboundConfigSelectKnownPacks:
			packs := []DataPack{}
			err := p.Scan(pk.Array(&packs))
			if err != nil {
				return ConfigErr{"select known packs", err}
			}
			knwonPacks := c.ConfigHandler.SelectDataPacks(packs)
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigSelectKnownPacks,
				pk.Array(knwonPacks),
			))
			if err != nil {
				return ConfigErr{"select known packs", err}
			}

		case packetid.ClientboundConfigCustomReportDetails:
		case packetid.ClientboundConfigServerLinks:
		}
	}
}

type DataPack struct {
	Namespace string
	ID        string
	Version   string
}

func (d DataPack) WriteTo(w io.Writer) (n int64, err error) {
	n, err = pk.String(d.Namespace).WriteTo(w)
	if err != nil {
		return n, err
	}
	n1, err := pk.String(d.ID).WriteTo(w)
	if err != nil {
		return n + n1, err
	}
	n2, err := pk.String(d.Version).WriteTo(w)
	return n + n1 + n2, err
}

func (d *DataPack) ReadFrom(r io.Reader) (n int64, err error) {
	n, err = (*pk.String)(&d.Namespace).ReadFrom(r)
	if err != nil {
		return n, err
	}
	n1, err := (*pk.String)(&d.ID).ReadFrom(r)
	if err != nil {
		return n + n1, err
	}
	n2, err := (*pk.String)(&d.Version).ReadFrom(r)
	return n + n1 + n2, err
}

type DefaultConfigHandler struct {
	cookies       map[pk.Identifier][]byte
	resourcesPack []ResourcePack
}

func NewDefaultConfigHandler() *DefaultConfigHandler {
	return &DefaultConfigHandler{
		cookies:       make(map[pk.String][]byte),
		resourcesPack: make([]ResourcePack, 0),
	}
}

func (d *DefaultConfigHandler) GetCookie(key pk.Identifier) []byte {
	return d.cookies[key]
}

func (d *DefaultConfigHandler) SetCookie(key pk.Identifier, payload []byte) {
	d.cookies[key] = payload
}

func (d *DefaultConfigHandler) PushResourcePack(res ResourcePack) {
	d.resourcesPack = append(d.resourcesPack, res)
}

func (d *DefaultConfigHandler) PopResourcePack(id pk.UUID) {
	for i, v := range d.resourcesPack {
		if id == v.ID {
			d.resourcesPack = append(d.resourcesPack[:i], d.resourcesPack[i+1:]...)
			break
		}
	}
}

func (d *DefaultConfigHandler) PopAllResourcePack() {
	d.resourcesPack = d.resourcesPack[:0]
}

func (d *DefaultConfigHandler) SelectDataPacks(packs []DataPack) []DataPack {
	return []DataPack{}
}
