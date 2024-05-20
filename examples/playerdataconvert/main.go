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
	"compress/gzip"
	"encoding/binary"
	"encoding/json"
	"flag"
	"fmt"
	"io/fs"
	"net/http"
	"os"
	"path/filepath"

	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/nbt/dynbt"
	"github.com/google/uuid"
)

var savePath = flag.String("save", "The save folder with \"usercache.json\" file inside", "")

func main() {
	flag.Parse()
	save, err := os.ReadDir(*savePath)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open dir: %v", err)
		return
	}

	var usercache fs.DirEntry
	for i := range save {
		name := save[i].Name()
		if name == "usercache.json" && !save[i].IsDir() {
			usercache = save[i]
		}
	}
	if usercache == nil {
		fmt.Fprintf(os.Stderr, "usercache.json not found")
		return
	}
	usercaches := readUsercache(filepath.Join(*savePath, usercache.Name()))
	fmt.Printf("Successfully reading usercache\n")
	readPlayerdata(filepath.Join(*savePath, "world", "playerdata"), usercaches)
}

type UserCache struct {
	Name      string `json:"name"`
	UUID      string `json:"uuid"`
	ExpiresOn string `json:"expiresOn"`
}

func readUsercache(path string) []UserCache {
	data, err := os.ReadFile(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read usercache file: %v\n", err)
		return nil
	}

	var usercache []UserCache
	err = json.Unmarshal(data, &usercache)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse usercache file: %v\n", err)
		return nil
	}
	return usercache
}

func readPlayerdata(dir string, users []UserCache) {
	for _, user := range users {
		nbtdata, err := readNbtData(dir, &user)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read %s's nbt data\n", user.Name)
			continue
		}

		// Get old UUID
		uuidInts := nbtdata.Get("UUID").IntArray()
		uuidBytes, err := intArrayToUUID(uuidInts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read %s's UUID\n", user.Name)
			continue
		}

		if ver := uuidBytes.Version(); ver != 3 { // v3 is for offline players
			fmt.Printf("Ignoring UUID: %v version: %d\n", uuidBytes, ver)
			continue
		}

		// Get new UUID
		name, id, err := usernameToUUID(user.Name)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to fetch username for %s from Mojang server: %v\n", user.Name, err)
			continue
		}

		fmt.Printf("[%s] %v -> %v\n", name, uuidBytes, id)

		// Update UUID
		ints := uuidToIntArray(id)
		nbtdata.Set("UUID", dynbt.NewIntArray(ints[:]))

		// Create new .dat file
		err = writeNbtData(dir, id.String(), &nbtdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write %s's .dat file: %v\n", name, err)
			continue
		}
	}
}

func readNbtData(dir string, user *UserCache) (dynbt.Value, error) {
	file, err := os.Open(filepath.Join(dir, user.UUID+".dat"))
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read %s's userdata: %v\n", user.Name, err)
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to decompress %s's userdata: %v\n", user.Name, err)
	}

	var nbtdata dynbt.Value
	_, err = nbt.NewDecoder(r).Decode(&nbtdata)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to parse %s's userdata: %v\n", user.Name, err)
	}
	return nbtdata, nil
}

func writeNbtData(dir string, id string, nbtdata *dynbt.Value) error {
	newDatFilePath := filepath.Join(dir, id+".dat")
	file, err := os.Create(newDatFilePath)
	if err != nil {
		return err
	}

	w := gzip.NewWriter(file)
	err = nbt.NewEncoder(w).Encode(&nbtdata, "")
	if err != nil {
		return err
	}

	err = w.Close()
	if err != nil {
		return err
	}

	err = file.Close()
	if err != nil {
		return err
	}
	return nil
}

func usernameToUUID(name string) (string, uuid.UUID, error) {
	var id uuid.UUID
	resp, err := http.Get("https://api.mojang.com/users/profiles/minecraft/" + name)
	if err != nil {
		return "", id, err
	}

	var body struct {
		Name string `json:"name"`
		ID   string `json:"id"`
	}
	err = json.NewDecoder(resp.Body).Decode(&body)
	if err != nil {
		return "", id, err
	}

	id, err = uuid.Parse(body.ID)
	return body.Name, id, err
}

func intArrayToUUID(uuidInts []int32) (id uuid.UUID, err error) {
	if uuidLen := len(uuidInts); uuidLen != 4 {
		err = fmt.Errorf("invalid UUID len: %d * int32", uuidLen)
		return
	}
	for i, v := range uuidInts {
		binary.BigEndian.PutUint32(id[i*4:], uint32(v))
	}
	return
}

func uuidToIntArray(id uuid.UUID) (ints [4]int32) {
	for i := range ints {
		ints[i] = int32(binary.BigEndian.Uint32(id[i*4:]))
	}
	return
}
