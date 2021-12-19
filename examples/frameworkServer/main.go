package main

import (
	_ "embed"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/server"
	"image"
	_ "image/png"
	"log"
	"os"
)

const MaxPlayer = 20
const IconPath = "./server-icon.png"

var motd = chat.Message{Text: "A Minecraft Server ", Extra: []chat.Message{{Text: "Powered by go-mc", Color: "yellow"}}}

func main() {
	playerList := server.NewPlayerList(MaxPlayer)
	serverInfo, err := server.NewPingInfo(playerList, server.ProtocolName, server.ProtocolVersion, motd, readIcon())
	if err != nil {
		log.Fatalf("Set server info error: %v", err)
	}
	defaultDimension := server.NewSimpleDim(16)
	chunk00 := server.EmptyChunk(16)
	for s := 0; s < 16; s++ {
		for i := 0; i < 16*16; i++ {
			chunk00.Sections[s].SetBlock(i, 1)
		}
	}
	defaultDimension.LoadChunk(server.ChunkPos{X: 0, Z: 0}, chunk00)
	s := server.Server{
		ListPingHandler: serverInfo,
		LoginHandler: &server.MojangLoginHandler{
			OnlineMode: true,
			Threshold:  256,
		},
		GamePlay: &server.Game{
			Dim:        defaultDimension,
			PlayerList: playerList,
		},
	}
	if err := s.Listen(":25565"); err != nil {
		log.Fatalf("Listen error: %v", err)
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
