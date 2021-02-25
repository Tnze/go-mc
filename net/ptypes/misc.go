package ptypes

import (
	"io"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

// SoundEffect is a client-bound packet used to play a specific sound ID
// on the client.
type SoundEffect struct {
	Sound         pk.VarInt
	Category      pk.VarInt
	X, Y, Z       pk.Int
	Volume, Pitch pk.Float
}

func (p *SoundEffect) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.Sound, &p.Category, &p.X, &p.Y, &p.Z, &p.Volume, &p.Pitch)
}

// NamedSoundEffect is a client-bound packet used to play a sound with the
// specified name on the client.
type NamedSoundEffect struct {
	Sound         pk.String
	Category      pk.VarInt
	X, Y, Z       pk.Int
	Volume, Pitch pk.Float
}

func (p *NamedSoundEffect) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.Sound, &p.Category, &p.X, &p.Y, &p.Z, &p.Volume, &p.Pitch)
}

// ChatMessageClientbound represents a chat message forwarded by the server.
type ChatMessageClientbound struct {
	S      chat.Message
	Pos    pk.Byte
	Sender pk.UUID
}

func (p *ChatMessageClientbound) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.S, &p.Pos, &p.Sender)
}

// UpdateHealth encodes player health/food information from the server.
type UpdateHealth struct {
	Health         pk.Float
	Food           pk.VarInt
	FoodSaturation pk.Float
}

func (p *UpdateHealth) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.Health, &p.Food, &p.FoodSaturation)
}

// PluginData encodes the custom data encoded in a plugin message.
type PluginData []byte

func (p PluginData) Encode() []byte {
	return []byte(p)
}

func (p *PluginData) Decode(r pk.DecodeReader) error {
	d, err := io.ReadAll(r)
	if err != nil {
		return err
	}
	*p = d
	return nil
}

// PluginMessage represents a packet with a customized payload.
type PluginMessage struct {
	Channel pk.Identifier
	Data    PluginData
}

func (p *PluginMessage) Decode(pkt pk.Packet) error {
	return pkt.Scan(&p.Channel, &p.Data)
}

func (p *PluginMessage) Encode() pk.Packet {
	return pk.Marshal(
		packetid.CustomPayloadServerbound,
		p.Channel,
		p.Data,
	)
}
