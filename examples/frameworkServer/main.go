package main

import (
	_ "embed"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/server"
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
	defaultDimension.LoadChunk(server.ChunkPos{X: 0, Z: 0}, server.EmptyChunk(16))
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
