package main

import (
	"container/list"
	_ "embed"
	"log"
	"sync"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/net"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
	"github.com/google/uuid"
)

type MyServer struct {
	PlayerList     *list.List
	PlayerListLock sync.Mutex
	server.MojangLoginHandler
	Settings struct {
		Name      string
		MaxPlayer int
		MOTD      chat.Message
	}
}

func main() {
	var ms MyServer
	ms.PlayerList = list.New()
	ms.MojangLoginHandler.OnlineMode = true
	ms.MojangLoginHandler.Threshold = 256

	ms.Settings.Name = "MyServer"
	ms.Settings.MaxPlayer = 1
	ms.Settings.MOTD = chat.Message{Text: "A Minecraft Server ", Extra: []chat.Message{{Text: "Powered by go-mc", Color: "yellow"}}}

	s := server.Server{
		ListPingHandler: &ms,
		LoginHandler:    &ms,
		GamePlay:        &ms,
	}
	if err := s.Listen(":25565"); err != nil {
		log.Fatalf("Listen error: %v", err)
	}
}

func (m *MyServer) Name() string {
	return m.Settings.Name
}

func (m *MyServer) Protocol() int {
	return server.ProtocolVersion
}

func (m *MyServer) MaxPlayer() int {
	return m.Settings.MaxPlayer
}

func (m *MyServer) OnlinePlayer() int {
	m.PlayerListLock.Lock()
	defer m.PlayerListLock.Unlock()
	return m.PlayerList.Len()
}

func (m *MyServer) PlayerSamples() (sample []server.PlayerSample) {
	m.PlayerListLock.Lock()
	defer m.PlayerListLock.Unlock()
	// get first 10 players
	sample = make([]server.PlayerSample, 0, 10)
	v := m.PlayerList.Front()
	for i := 0; i < 10; i++ {
		if v == nil {
			break
		}
		sample = append(sample, v.Value.(server.PlayerSample))
		v = v.Next()
	}
	return
}

func (m *MyServer) Description() chat.Message {
	return m.Settings.MOTD
}

func (m *MyServer) AcceptPlayer(name string, id uuid.UUID, protocol int32, conn *net.Conn) {
	// Add player into PlayerList
	m.PlayerListLock.Lock()
	if m.PlayerList.Len() >= m.Settings.MaxPlayer {
		err := conn.WritePacket(pk.Marshal(packetid.KickDisconnect,
			chat.TranslateMsg("multiplayer.disconnect.server_full"),
		))
		if err != nil {
			log.Printf("Write packet fail: %v", err)
		}
		m.PlayerListLock.Unlock()
		return
	}
	elem := m.PlayerList.PushBack(server.PlayerSample{
		Name: name,
		ID:   id,
	})
	m.PlayerListLock.Unlock()
	defer func() {
		m.PlayerListLock.Lock()
		defer m.PlayerListLock.Unlock()
		// remove player in PlayerList
		m.PlayerList.Remove(elem)
	}()

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

//go:embed DimensionCodec.snbt
var dimensionCodecSNBT string

//go:embed Dimension.snbt
var dimensionSNBT string

func (m *MyServer) joinGame(conn *net.Conn) error {
	return conn.WritePacket(pk.Marshal(packetid.Login,
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
		pk.VarInt(m.Settings.MaxPlayer),                    // Max Players
		pk.VarInt(15),                                      // View Distance
		pk.VarInt(15),                                      // Simulation Distance
		pk.Boolean(false),                                  // Reduced Debug Info
		pk.Boolean(true),                                   // Enable respawn screen
		pk.Boolean(false),                                  // Is Debug
		pk.Boolean(true),                                   // Is Flat
	))
}

func (m *MyServer) playerPositionAndLook(conn *net.Conn) error {
	return conn.WritePacket(pk.Marshal(packetid.PositionClientbound,
		// https://wiki.vg/index.php?title=Protocol&oldid=16067#Player_Position_And_Look_.28clientbound.29
		pk.Double(0), pk.Double(0), pk.Double(0), // XYZ
		pk.Float(0), pk.Float(0), // Yaw Pitch
		pk.Byte(0),        // flag
		pk.VarInt(0),      // TP ID
		pk.Boolean(false), // Dismount vehicle
	))
}
