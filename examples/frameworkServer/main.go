package main

import (
	"context"
	_ "embed"
	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
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
	defaultDimension := server.NewSimpleDim(256)
	chunk00 := level.ChunkFromSave(readChunk00(), 256)
	defaultDimension.LoadChunk(level.ChunkPos{X: 0, Z: 0}, chunk00)

	game := server.NewGame(defaultDimension, playerList)
	game.Run(context.Background())

	s := server.Server{
		ListPingHandler: serverInfo,
		LoginHandler: &server.MojangLoginHandler{
			OnlineMode: false,
			Threshold:  256,
		},
		GamePlay: game,
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

func readChunk00() *save.Chunk {
	r, err := region.Open("./save/testdata/region/r.0.0.mca")
	if err != nil {
		panic(err)
	}
	defer r.Close()

	var c save.Chunk
	data, err := r.ReadSector(0, 0)
	if err != nil {
		panic(err)
	}
	err = c.Load(data)
	if err != nil {
		panic(err)
	}
	return &c
}
