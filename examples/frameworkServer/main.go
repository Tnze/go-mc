package main

import (
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
	"github.com/google/uuid"
)

type MyServer struct {
	playerList *server.PlayerList
}

const MaxPlayer = 20
const IconPath = "./server-icon.png"

var motd = chat.Message{Text: "A Minecraft Server ", Extra: []chat.Message{{Text: "Powered by go-mc", Color: "yellow"}}}

func main() {
	playerList := server.NewPlayerList(MaxPlayer)
	serverInfo, err := server.NewPingInfo(playerList, server.ProtocolName, server.ProtocolVersion, motd, readIcon())
	if err != nil {
		log.Fatalf("Set server info error: %v", err)
	}
	ms := MyServer{
		playerList: playerList,
	}

	s := server.Server{
		ListPingHandler: serverInfo,
		LoginHandler: &server.MojangLoginHandler{
			OnlineMode: true,
			Threshold:  256,
		},
		GamePlay: &ms,
	}
	if err := s.Listen(":25565"); err != nil {
		log.Fatalf("Listen error: %v", err)
	}
}

func (m *MyServer) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	// Add player into PlayerList
	remove := m.playerList.TryInsert(server.PlayerSample{
		Name: name,
		ID:   id,
	})
	if remove == nil {
		err := conn.WritePacket(pk.Marshal(packetid.ClientboundDisconnect,
			chat.TranslateMsg("multiplayer.disconnect.server_full"),
		))
		if err != nil {
			log.Printf("Write packet fail: %v", err)
		}
		return
	}
	defer remove()

	if err := m.joinGame(conn); err != nil {
		log.Printf("Write packet fail: %v", err)
		return
	}
	if err := m.playerPositionAndLook(conn); err != nil {
		log.Printf("Write packet fail: %v", err)
		return
	}

	var p pk.Packet
	for {
		err := conn.ReadPacket(&p)
		if err != nil {
			log.Printf("Read packet fail: %v", err)
			break
		}

		log.Printf("Read packet: %#X", p.ID)
	}
}

func readIcon() image.Image {
	f, err := os.Open(IconPath)
	// if the file doesn't exist, return nil
	if os.IsNotExist(err) {
		return nil
	} else if err != nil {
		log.Fatalf("Open icon file error: %v", err)
	}
	defer f.Close()

	icon, _, err := image.Decode(f)
	if err != nil {
		log.Fatalf("Decode image error: %v", err)
	}
	return icon
}

//go:embed DimensionCodec.snbt
var dimensionCodecSNBT string

//go:embed Dimension.snbt
var dimensionSNBT string

func (m *MyServer) joinGame(conn *net.Conn) error {
	return conn.WritePacket(pk.Marshal(packetid.ClientboundLogin,
		pk.Int(0),          // EntityID
		pk.Boolean(false),  // Is hardcore
		pk.UnsignedByte(1), // Gamemode
		pk.Byte(1),         // Previous Gamemode
		pk.VarInt(1),       // World Count
		pk.Ary{Len: 1, Ary: []pk.Identifier{"world"}},      // World Names
		pk.NBT(nbt.StringifiedMessage(dimensionCodecSNBT)), // Dimension codec
		pk.NBT(nbt.StringifiedMessage(dimensionSNBT)),      // Dimension
		pk.Identifier("world"),                             // World Name
		pk.Long(0),                                         // Hashed Seed
		pk.VarInt(MaxPlayer),                               // Max Players
		pk.VarInt(15),                                      // View Distance
		pk.VarInt(15),                                      // Simulation Distance
		pk.Boolean(false),                                  // Reduced Debug Info
		pk.Boolean(true),                                   // Enable respawn screen
		pk.Boolean(false),                                  // Is Debug
		pk.Boolean(true),                                   // Is Flat
	))
}

func (m *MyServer) playerPositionAndLook(conn *net.Conn) error {
	return conn.WritePacket(pk.Marshal(packetid.ClientboundPlayerPosition,
		pk.Double(0), pk.Double(0), pk.Double(0), // XYZ
		pk.Float(0), pk.Float(0), // Yaw Pitch
		pk.Byte(0),        // flag
		pk.VarInt(0),      // TP ID
		pk.Boolean(false), // Dismount vehicle
	))
}
