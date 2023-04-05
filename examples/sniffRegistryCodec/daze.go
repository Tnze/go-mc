// sniffRegistryCodec is an example that acts as a client,
// connects to the server and saves its RegistryCodec to a .nbt file.
package main

import (
	"flag"
	"log"
	"os"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
	"github.com/Tnze/go-mc/data/packetid"
	"github.com/Tnze/go-mc/nbt"
	pk "github.com/Tnze/go-mc/net/packet"
)

var (
	address = flag.String("address", "127.0.0.1", "The server address")
	client  *bot.Client
	player  *basic.Player
)

func main() {
	flag.Parse()
	// log.SetOutput(colorable.NewColorableStdout())
	client = bot.NewClient()
	client.Auth.Name = "Daze"
	player = basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{})

	// To receive the raw NBT data, create a new packet handler
	// instead of just reading player.RegistryCodec in GameStart event.
	client.Events.AddListener(bot.PacketHandler{
		ID:       packetid.ClientboundLogin,
		Priority: 50,
		F:        onLogin,
	})

	// Login
	err := client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	// JoinGame
	for {
		if err = client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}
		log.Fatal(err)
	}
}

func onLogin(p pk.Packet) error {
	var DimensionNames []pk.Identifier
	var RegistryCodec nbt.RawMessage
	err := p.Scan(
		new(pk.Int),               // Entity ID
		new(pk.Boolean),           // Is hardcore
		new(pk.Byte),              // Gamemode
		new(pk.Byte),              // Previous Gamemode
		pk.Array(&DimensionNames), // Dimension Names
		pk.NBT(&RegistryCodec),    // Registry Codec (Only care about this)
		// ...Ignored
	)
	if err != nil {
		return err
	}
	err = saveToFile("RegistryCodec.nbt", RegistryCodec)
	if err != nil {
		return err
	}
	log.Print("Successfully written RegistryCodec.nbt")
	return nil
}

func saveToFile(filename string, value any) (err error) {
	f, err := os.Create(filename)
	if err != nil {
		return nil
	}
	defer func(f *os.File) {
		if err2 := f.Close(); err == nil {
			err = err2
		}
	}(f)
	return nbt.NewEncoder(f).Encode(value, "")
}
