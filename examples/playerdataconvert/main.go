// playerdataconvert is a program to convert player data form offline server to online server.
//
// When a player with official account login connect to a offline-mode server,
// the server store the player data with their "offline UUID". While you open
// the online-mode switch, the player data loose.
//
// By using this tool, you can convert the offline data into online data.
// The players will keep everything they got, yay!
package main

import (
	"encoding/json"
	"flag"
	"fmt"
	"os"
	"path/filepath"

	"github.com/google/uuid"
)

var (
	savePath            = flag.String("save", ".", "The save folder with \"usercache.json\" file inside")
	convertPlayerData   = flag.Bool("cplayerdata", true, "Whether convert files at /world/playerdata/*.dat")
	convertEntities     = flag.Bool("centities", true, "Whether convert pets' Owner at /world/entities/*")
	convertAdvancements = flag.Bool("cadvancements", true, "Whether convert advancements at /world/advancements/*")
)

func main() {
	flag.Parse()

	usercaches, err := readUsercache(filepath.Join(*savePath, "usercache.json"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse usercache file: %v\n", err)
		return
	}
	fmt.Printf("Successfully reading usercache\n")
	m := mappingUsers(usercaches)

	if *convertPlayerData {
		readPlayerdata(filepath.Join(*savePath, "world", "playerdata"), m)
	}

	if *convertEntities {
		readEntities(filepath.Join(*savePath, "world", "entities"), m)
	}

	if *convertAdvancements {
		readAdvancements(filepath.Join(*savePath, "world", "advancements"), m)
	}
}

type UserCache struct {
	Name string    `json:"name"`
	UUID uuid.UUID `json:"uuid"`
}

func readUsercache(path string) ([]UserCache, error) {
	data, err := os.ReadFile(path)
	if err != nil {
		return nil, err
	}

	var usercache []UserCache
	err = json.Unmarshal(data, &usercache)
	if err != nil {
		return nil, err
	}
	return usercache, nil
}

func mappingUsers(users []UserCache) map[uuid.UUID]UserCache {
	m := make(map[uuid.UUID]UserCache)
	for _, user := range users {
		name := user.Name
		// // You can add your maps here
		// if v, ok := offlineOnlineMaps[name]; ok {
		// 	name = v
		// }

		name, id, err := usernameToUUID(name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to fetch username for %s from Mojang server: %v\n", name, err)
			continue
		}

		fmt.Printf("[%s] %v -> %v\n", name, user.UUID, id)
		m[user.UUID] = UserCache{name, id}
	}
	return m
}
