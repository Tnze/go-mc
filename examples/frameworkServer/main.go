package main

import (
	"context"
	_ "embed"
	"flag"
	"github.com/Tnze/go-mc/server/ecs"
	"github.com/Tnze/go-mc/server/world"
	"image"
	_ "image/png"
	"log"
	"os"

	"github.com/Tnze/go-mc/chat"
	"github.com/Tnze/go-mc/server"
	"github.com/Tnze/go-mc/server/command"
	"github.com/Tnze/go-mc/server/player"
)

var motd = chat.Message{Text: "A Minecraft Server ", Extra: []chat.Message{{Text: "Powered by go-mc", Color: "yellow"}}}
var addr = flag.String("Address", "127.0.0.1:25565", "Listening address")
var iconPath = flag.String("ServerIcon", "./server-icon.png", "The path to server icon")
var maxPlayer = flag.Int("MaxPlayer", 16384, "The maximum number of players")
var regionPath = flag.String("Regions", "./save/testdata/region/", "The region files")
var playerdataPath = flag.String("PlayerData", "./save/testdata/playerdata", "The player data files")

func main() {
	flag.Parse()
	logger := Logger{log.Default()}
	playerList := server.NewPlayerList(*maxPlayer)
	serverInfo, err := server.NewPingInfo(playerList, server.ProtocolName, server.ProtocolVersion, motd, readIcon())
	if err != nil {
		logger.Fatalf("Set server info error: %v", err)
	}
	keepAlive := server.NewKeepAlive()
	commands := command.NewGraph()
	handleFunc := func(ctx context.Context, args []command.ParsedData) error {
		logger.Printf("Command: args: %v", args)
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
		logger,
		playerList,
		keepAlive,
		commands,
	)
	ecs.Register[world.Dimension, *ecs.HashMapStorage[world.Dimension]](game.World)
	dimList := world.NewDimensionManager(game)
	dimList.Add(game.CreateEntity(world.NewDimension(
		"minecraft:overworld", *regionPath,
	)), "minecraft:overworld")
	player.SpawnSystem(game, *playerdataPath)
	player.PosAndRotSystem(game)
	go game.Run(context.Background())

	s := server.Server{
		ListPingHandler: serverInfo,
		LoginHandler: &server.MojangLoginHandler{
			OnlineMode:   true,
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

type Logger struct{ *log.Logger }

func (l Logger) Init(g *server.Game) {
	l.Print("Server init")
}

func (l Logger) Run(ctx context.Context) {
	l.Print("Server is running")
}

func (l Logger) ClientJoin(c *server.Client, p *server.Player) {
	l.Printf("Player join [%s]%v from %v", p.Name, p.UUID, c.Socket.RemoteAddr())
}

func (l Logger) ClientLeft(c *server.Client, p *server.Player, reason error) {
	l.Printf("Player left [%s]%v reason: %v", p.Name, p.UUID, reason)
}
