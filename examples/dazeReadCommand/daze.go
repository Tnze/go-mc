// Daze could join an offline-mode server as client.
// Just standing there and do nothing. Automatically reborn after five seconds of death.
package main

import (
	"flag"
	"io"
	"log"
	"strconv"

	//"github.com/mattn/go-colorable"

	"github.com/Tnze/go-mc/bot"
	"github.com/Tnze/go-mc/bot/basic"
	"github.com/Tnze/go-mc/bot/screen"
	_ "github.com/Tnze/go-mc/data/lang/zh-cn"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
)

var address = flag.String("address", "127.0.0.1", "The server address")
var client *bot.Client
var player *basic.Player
var screenManager *screen.Manager

func main() {
	flag.Parse()
	//log.SetOutput(colorable.NewColorableStdout())
	client = bot.NewClient()
	client.Auth.Name = "Daze"
	player = basic.NewPlayer(client, basic.DefaultSettings, basic.EventsListener{})
	client.Events.AddListener(bot.PacketHandler{
		ID:       packetid.ClientboundCommands,
		Priority: 50,
		F:        onCommands,
	})

	//Login
	err := client.JoinServer(*address)
	if err != nil {
		log.Fatal(err)
	}
	log.Println("Login success")

	//JoinGame
	for {
		if err = client.HandleGame(); err == nil {
			panic("HandleGame never return nil")
		}
		log.Fatal(err)
	}
}

func onCommands(p pk.Packet) error {
	var nodes []Node
	var root pk.VarInt
	err := p.Scan(pk.Array(&nodes), &root)
	if err != nil {
		return err
	}
	log.Printf("Root index: %d", root)
	return nil
}

type Node struct {
}

func (n Node) ReadFrom(r io.Reader) (int64, error) {
	var Flags pk.Byte
	var Children []pk.VarInt
	var Redirect pk.VarInt
	var Name pk.String
	var Parser pk.Identifier
	var Properties = Prop{Type: &Parser}
	var SuggestionsType pk.Identifier
	m, err := pk.Tuple{
		&Flags,
		pk.Array(&Children),
		pk.Opt{
			Has:   func() bool { return Flags&0x08 != 0 },
			Field: &Redirect,
		},
		pk.Opt{
			Has:   func() bool { return Flags&0x03 == 2 || Flags&0x03 == 1 },
			Field: &Name,
		},
		pk.Opt{
			Has:   func() bool { return Flags&0x03 == 2 },
			Field: pk.Tuple{&Parser, &Properties},
		},
		pk.Opt{
			Has:   func() bool { return Flags&0x10 != 0 },
			Field: &SuggestionsType,
		},
	}.ReadFrom(r)
	if err != nil {
		return m, err
	}
	var redirect string
	if Flags&0x08 != 0 {
		redirect = "Redirect: " + strconv.Itoa(int(Redirect))
	}
	var parser string
	if Flags&0x03 == 2 {
		redirect = string("Parser: " + Parser)
	}
	log.Printf("Type: %2d\tName: %s\tChildren: %v\t%v\t%v", Flags&0x03, Name, Children, redirect, parser)

	return m, nil
}

type Prop struct {
	Type *pk.Identifier
}

func (p Prop) ReadFrom(r io.Reader) (int64, error) {
	var Flags pk.Byte
	switch *p.Type {
	case "brigadier:double":
		var Min, Max pk.Double
		return pk.Tuple{
			&Flags,
			pk.Opt{Has: func() bool { return Flags&0x01 != 0 }, Field: &Min},
			pk.Opt{Has: func() bool { return Flags&0x02 != 0 }, Field: &Max},
		}.ReadFrom(r)
	case "brigadier:float":
		var Min, Max pk.Float
		return pk.Tuple{
			&Flags,
			pk.Opt{Has: func() bool { return Flags&0x01 != 0 }, Field: &Min},
			pk.Opt{Has: func() bool { return Flags&0x02 != 0 }, Field: &Max},
		}.ReadFrom(r)
	case "brigadier:integer":
		var Min, Max pk.Int
		return pk.Tuple{
			&Flags,
			pk.Opt{Has: func() bool { return Flags&0x01 != 0 }, Field: &Min},
			pk.Opt{Has: func() bool { return Flags&0x02 != 0 }, Field: &Max},
		}.ReadFrom(r)
	case "brigadier:long":
		var Min, Max pk.Long
		return pk.Tuple{
			&Flags,
			pk.Opt{Has: func() bool { return Flags&0x01 != 0 }, Field: &Min},
			pk.Opt{Has: func() bool { return Flags&0x02 != 0 }, Field: &Max},
		}.ReadFrom(r)
	case "brigadier:string":
		return new(pk.VarInt).ReadFrom(r)
	case "minecraft:entity":
		return new(pk.Byte).ReadFrom(r)
	case "minecraft:score_holder":
		return new(pk.Byte).ReadFrom(r)
	case "minecraft:range":
		return new(pk.Boolean).ReadFrom(r)
	default:
		return 0, nil
	}
}
