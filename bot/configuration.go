package bot

import (
	"bytes"
	"errors"
	"fmt"
	"io"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
)

type ConfigHandler interface {
	EnableFeature(features []pk.Identifier)

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
	for {
		var p pk.Packet
		if err := conn.ReadPacket(&p); err != nil {
			return ConfigErr{"config custom payload", err}
		}

		switch packetid.ClientboundPacketID(p.ID) {
		case packetid.ClientboundConfigCookieRequest:
			var key pk.Identifier
			err := p.Scan(&key)
			if err != nil {
				return ConfigErr{"cookie request", err}
			}
			cookieContent := c.Cookies[string(key)]
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigCookieResponse,
				key, pk.OptionEncoder[pk.ByteArray]{
					Has: cookieContent != nil,
					Val: pk.ByteArray(cookieContent),
				},
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
			//
			// There are two types of Custom packet.
			// One for Login stage, the other for config and play stage.
			// The first one called "Custom Query", and the second one called "Custom Payload".
			// We can know the different by their name, the "query" is one request to one response, paired.
			// But the second one can be sent in any order.
			//
			// And the custome payload packet seems to be same in config stage and play stage.
			// How do we provide API for that?

		case packetid.ClientboundConfigDisconnect:
			const ErrStage = "disconnect"
			var reason chat.Message
			err := p.Scan(&reason)
			if err != nil {
				return ConfigErr{ErrStage, err}
			}
			return ConfigErr{ErrStage, DisconnectErr(reason)}

		case packetid.ClientboundConfigFinishConfiguration:
			err := conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigFinishConfiguration,
			))
			if err != nil {
				return ConfigErr{"finish config", err}
			}
			return nil

		case packetid.ClientboundConfigKeepAlive:
			const ErrStage = "keep alive"
			var keepAliveID pk.Long
			err := p.Scan(&keepAliveID)
			if err != nil {
				return ConfigErr{ErrStage, err}
			}
			// send it back
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigKeepAlive,
				keepAliveID,
			))
			if err != nil {
				return ConfigErr{ErrStage, err}
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
			const ErrStage = "registry"
			var registryID pk.Identifier

			r := bytes.NewReader(p.Data)
			_, err := registryID.ReadFrom(r)
			if err != nil {
				return ConfigErr{ErrStage, err}
			}

			registry := c.Registries.Registry(string(registryID))
			if registry == nil {
				return ConfigErr{ErrStage, errors.New("unknown registry: " + string(registryID))}
			}

			_, err = registry.ReadFrom(r)
			if err != nil {
				return ConfigErr{ErrStage, fmt.Errorf("failed to read registry %s: %w", registryID, err)}
			}

		case packetid.ClientboundConfigResourcePackPop:
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
			c.Cookies[string(key)] = []byte(payload)

		case packetid.ClientboundConfigTransfer:
			var host pk.String
			var port pk.VarInt
			err := p.Scan(&host, &port)
			if err != nil {
				return ConfigErr{"transfer", err}
			}
			// TODO: trnasfer to the specific server
			// How does it work? Just connect the new server, and re-start at handshake?

		case packetid.ClientboundConfigUpdateEnabledFeatures:
			features := []pk.Identifier{}
			err := p.Scan(pk.Array(&features))
			if err != nil {
				return ConfigErr{"update enabled features", err}
			}
			c.ConfigHandler.EnableFeature(features)

		case packetid.ClientboundConfigUpdateTags:
			const ErrStage = "update tags"
			r := bytes.NewReader(p.Data)

			var length pk.VarInt
			_, err := length.ReadFrom(r)
			if err != nil {
				return ConfigErr{ErrStage, err}
			}

			var registryID pk.Identifier
			for i := 0; i < int(length); i++ {
				_, err = registryID.ReadFrom(r)
				if err != nil {
					return ConfigErr{ErrStage, err}
				}

				registry := c.Registries.Registry(string(registryID))
				if registry == nil {
					// TODO: Sice our registry system is incompelted, ignore all tags bind to non-exist registry
					_, err = idleTagsDecoder{}.ReadFrom(r)
					if err != nil {
						return ConfigErr{ErrStage, err}
					}
					continue
					// return ConfigErr{ErrStage, errors.New("unknown registry: " + string(registryID))}
				}

				_, err = registry.ReadTagsFrom(r)
				if err != nil {
					return ConfigErr{ErrStage, err}
				}
			}

		case packetid.ClientboundConfigSelectKnownPacks:
			const ErrStage = "select known packs"
			packs := []DataPack{}
			err := p.Scan(pk.Array(&packs))
			if err != nil {
				return ConfigErr{ErrStage, err}
			}
			knwonPacks := c.ConfigHandler.SelectDataPacks(packs)
			err = conn.WritePacket(pk.Marshal(
				packetid.ServerboundConfigSelectKnownPacks,
				pk.Array(knwonPacks),
			))
			if err != nil {
				return ConfigErr{ErrStage, err}
			}

		case packetid.ClientboundConfigCustomReportDetails:
			const ErrStage = "custom report details"
			var length pk.VarInt
			var title, description pk.String
			r := bytes.NewReader(p.Data)
			_, err := length.ReadFrom(r)
			if err != nil {
				return ConfigErr{ErrStage, err}
			}
			for i := 0; i < int(length); i++ {
				_, err = title.ReadFrom(r)
				if err != nil {
					return ConfigErr{ErrStage, err}
				}
				_, err = description.ReadFrom(r)
				if err != nil {
					return ConfigErr{ErrStage, err}
				}
				c.CustomReportDetails[string(title)] = string(description)
			}

		case packetid.ClientboundConfigServerLinks:
			// TODO
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
	resourcesPack []ResourcePack
}

func NewDefaultConfigHandler() *DefaultConfigHandler {
	return &DefaultConfigHandler{
		resourcesPack: make([]ResourcePack, 0),
	}
}

func (d *DefaultConfigHandler) EnableFeature(features []pk.Identifier) {}

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

type idleTagsDecoder struct{}

func (idleTagsDecoder) ReadFrom(r io.Reader) (int64, error) {
	var count pk.VarInt
	var tag pk.Identifier
	var length pk.VarInt
	n, err := count.ReadFrom(r)
	if err != nil {
		return n, err
	}
	for i := 0; i < int(count); i++ {
		var n1, n2, n3 int64
		n1, err = tag.ReadFrom(r)
		if err != nil {
			return n + n1, err
		}
		n2, err = length.ReadFrom(r)
		if err != nil {
			return n + n1 + n2, err
		}
		n += n1 + n2

		var id pk.VarInt
		for i := 0; i < int(length); i++ {
			n3, err = id.ReadFrom(r)
			if err != nil {
				return n + n3, err
			}
			n += n3
		}
	}
	return n, nil
}
