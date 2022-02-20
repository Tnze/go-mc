package main

import (
	"context"
	_ "embed"
	"flag"
	"fmt"
	"image"
	_ "image/png"
	"log"
	"os"
	"path/filepath"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
	"github.com/Tnze/go-mc/server"
	"github.com/Tnze/go-mc/server/command"
)

var motd = chat.Message{Text: "A Minecraft Server ", Extra: []chat.Message{{Text: "Powered by go-mc", Color: "yellow"}}}
var addr = flag.String("Address", ":25565", "Listening address")
var iconPath = flag.String("ServerIcon", "./server-icon.png", "The path to server icon")
var maxPlayer = flag.Int("MaxPlayer", 16384, "The maximum number of players")
var regionPath = flag.String("Regions", "./save/testdata/region/", "The region files")

func main() {
	flag.Parse()
	playerList := server.NewPlayerList(*maxPlayer)
	serverInfo, err := server.NewPingInfo(playerList, server.ProtocolName, server.ProtocolVersion, motd, readIcon())
	if err != nil {
		log.Fatalf("Set server info error: %v", err)
	}

	defaultDimension, err := loadAllRegions(*regionPath)
	if err != nil {
		log.Fatalf("Load chunks fail: %v", err)
	}

	commands := command.NewGraph()
	handleFunc := func(ctx context.Context, args []command.ParsedData) error {
		log.Printf("Command: args: %v", args)
		return nil
	}
	commands.AppendLiteral(commands.Literal("me").
		AppendArgument(commands.Argument("action", command.StringParser(2)).
			HandleFunc(handleFunc)).
		Unhandle(),
	).AppendLiteral(commands.Literal("help").
		AppendArgument(commands.Argument("command", command.StringParser(0)).
			HandleFunc(handleFunc)).
		HandleFunc(handleFunc),
	).AppendLiteral(commands.Literal("list").
		AppendLiteral(commands.Literal("uuids").
			HandleFunc(handleFunc)).
		HandleFunc(handleFunc),
	)

	game := server.NewGame(
		defaultDimension,
		playerList,
		server.NewKeepAlive(),
		server.NewGlobalChat(),
		commands,
	)
	go game.Run(context.Background())

	s := server.Server{
		ListPingHandler: serverInfo,
		LoginHandler: &server.MojangLoginHandler{
			OnlineMode:   false,
			Threshold:    256,
			LoginChecker: playerList,
		},
		GamePlay: game,
	}
	if err := s.Listen(*addr); err != nil {
		log.Fatalf("Listen error: %v", err)
	}
}

func readIcon() image.Image {
	f, err := os.Open(*iconPath)
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

func loadAllRegions(dir string) (*server.SimpleDim, error) {
	mcafiles, err := filepath.Glob(filepath.Join(dir, "r.*.*.mca"))
	if err != nil {
		return nil, err
	}
	dim := server.NewSimpleDim(256)
	for _, file := range mcafiles {
		var rx, rz int
		_, err := fmt.Sscanf(filepath.Base(file), "r.%d.%d.mca", &rx, &rz)
		if err != nil {
			return nil, err
		}
		err = loadAllChunks(dim, file, rx, rz)
		if err != nil {
			return nil, err
		}
	}
	return dim, nil
}

func loadAllChunks(dim *server.SimpleDim, file string, rx, rz int) error {
	r, err := region.Open(file)
	if err != nil {
		return err
	}
	defer r.Close()
	var c save.Chunk
	for x := 0; x < 32; x++ {
		for z := 0; z < 32; z++ {
			if !r.ExistSector(x, z) {
				continue
			}
			data, err := r.ReadSector(x, z)
			if err != nil {
				return err
			}
			if err := c.Load(data); err != nil {
				return err
			}
			chunk := level.ChunkFromSave(&c, 256)
			dim.LoadChunk(level.ChunkPos{X: rx<<5 + x, Z: rz<<5 + z}, chunk)
		}
	}
	return nil
}
