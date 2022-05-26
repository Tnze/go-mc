package player

import (
	"fmt"
	"github.com/Tnze/go-mc/data/packetid"
	pk "github.com/Tnze/go-mc/net/packet"
	"github.com/Tnze/go-mc/server"
	"github.com/Tnze/go-mc/server/ecs"
	"github.com/Tnze/go-mc/server/world"
	"log"
)

type PlayerProfile struct {
	Dim ecs.Index
}

type playerSpawnSystem struct {
	storage
}

func (p playerSpawnSystem) Update(w *ecs.World) {
	clients := ecs.GetComponent[server.Client](w)
	players := ecs.GetComponent[server.Player](w)
	profiles := ecs.GetComponent[PlayerProfile](w)
	dimensionRes := ecs.GetResource[world.DimensionList](w)
	players.AndNot(profiles.BitSetLike).Range(func(eid ecs.Index) {
		player := players.GetValue(eid)
		client := clients.GetValue(eid)
		profile, err := p.GetPlayer(player.UUID)
		if err != nil {
			client.PutErr(fmt.Errorf("read player data fail: %w", err))
			return
		}
		log.Println("load player info successes", profile)
		dim, ok := dimensionRes.Find(profile.Dimension)
		if !ok {
			panic("dimension " + profile.Dimension + " not found")
		}
		profiles.SetValue(eid, PlayerProfile{Dim: dim})
		client.WritePacket(server.Packet758(pk.Marshal(
			packetid.ClientboundLogin,
			pk.Int(eid),                     // Entity ID
			pk.Boolean(false),               // Is hardcore
			pk.Byte(profile.PlayerGameType), // Gamemode
			pk.Byte(-1),                     // Prev Gamemode
			dimensionRes,
			pk.NBT(dimensionRes.DimCodecSNBT),
			pk.NBT(dimensionRes.DimSNBT),
			pk.Identifier(profile.Dimension), // World Name
			pk.Long(1234567),                 // Hashed seed
			pk.VarInt(0),                     // Max Players (Ignored by client)
			pk.VarInt(15),                    // View Distance
			pk.VarInt(15),                    // Simulation Distance
			pk.Boolean(false),                // Reduced Debug Info
			pk.Boolean(true),                 // Enable respawn screen
			pk.Boolean(false),                // Is Debug
			pk.Boolean(true),                 // Is Flat
		)))
	})
}

func SpawnSystem(g *server.Game, playerdataPath string) {
	ecs.Register[PlayerProfile, *ecs.HashMapStorage[PlayerProfile]](g.World)
	g.Dispatcher.Add(playerSpawnSystem{storage: storage{playerdataPath}}, "go-mc:player:SpawnSystem", nil)
}

// PosAndRotSystem add a system to g.Dispatcher that
// receive player movement packets and update Pos and Rot component
// Require component Pos and Rot to be registered before.
func PosAndRotSystem(g *server.Game) {
	type posUpdate struct {
		ecs.Index
		server.Pos
	}
	updateChan := make(chan posUpdate)
	ecs.Register[server.Pos, *ecs.HashMapStorage[server.Pos]](g.World)
	ecs.Register[server.Rot, *ecs.HashMapStorage[server.Rot]](g.World)
	g.Dispatcher.Add(ecs.FuncSystem(func() {
		posStorage := ecs.GetComponent[server.Pos](g.World)
		for {
			select {
			case event := <-updateChan:
				if v := posStorage.GetValue(event.Index); v != nil {
					*v = event.Pos
				}
			default:
				return
			}
		}
	}), "go-mc:player:PosAndRotSystem", nil)

	g.AddHandler(&server.PacketHandler{
		ID: packetid.ServerboundMovePlayerPos,
		F: func(client *server.Client, player *server.Player, packet server.Packet758) error {
			var X, FeetY, Z pk.Double
			var OnGround pk.Boolean
			err := pk.Packet(packet).Scan(&X, &FeetY, &Z, &OnGround)
			if err != nil {
				return err
			}
			updateChan <- posUpdate{
				Index: client.Index,
				Pos: server.Pos{
					X: float64(X),
					Y: float64(FeetY),
					Z: float64(Z),
				},
			}
			return nil
		},
	})
}
