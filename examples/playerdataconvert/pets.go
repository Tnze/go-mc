package main

import (
	"bytes"
	"compress/gzip"
	"compress/zlib"
	"errors"
	"fmt"
	"io"
	"os"
	"path/filepath"

	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/nbt/dynbt"
	"github.com/Tnze/go-mc/save/region"
	"github.com/google/uuid"
)

func readEntities(dir string, m map[uuid.UUID]UserCache) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to read entities dir: %v\n", err)
		return
	}
	for i := range entries {
		readEntityMcaFile(filepath.Join(dir, entries[i].Name()), m)
	}
}

func readEntityMcaFile(path string, m map[uuid.UUID]UserCache) {
	r, err := region.Open(path)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open entities region file %s: %v\n", path, err)
		return
	}
	defer r.Close()

	for i := 0; i < 32; i++ {
		for j := 0; j < 32; j++ {
			if !r.ExistSector(i, j) {
				continue
			}

			data, err := r.ReadSector(i, j)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to read entities region sector: %v\n", err)
				continue
			}

			newdata, err := readEntityMcaSector(data, m)
			if err != nil {
				fmt.Fprintf(os.Stderr, "Failed to parse entities region sector: %v\n", err)
				continue
			}

			if newdata != nil {
				err = r.WriteSector(i, j, newdata)
				if err != nil {
					fmt.Fprintf(os.Stderr, "Failed to update region sector data: %v\n", err)
					continue
				}
			}
		}
	}
}

func readEntityMcaSector(data []byte, m map[uuid.UUID]UserCache) ([]byte, error) {
	var r io.Reader = bytes.NewReader(data[1:])
	var err error
	switch data[0] {
	default:
		return nil, errors.New("unknown compression")
	case 1:
		r, err = gzip.NewReader(r)
	case 2:
		r, err = zlib.NewReader(r)
	case 3:
	}
	if err != nil {
		return nil, err
	}

	var nbtdata dynbt.Value
	_, err = nbt.NewDecoder(r).Decode(&nbtdata)
	if err != nil {
		return nil, err
	}

	updated := false

	entities := nbtdata.Get("Entities")
	if entities == nil {
		return nil, fmt.Errorf("no Entities field in nbt, what happen?")
	}
	entities2 := entities.List()
	for _, entity := range entities2 {
		id := entity.Get("id").String()
		if owner := entity.Get("Owner"); owner != nil {
			owner, _ := intArrayToUUID(owner.IntArray())
			fmt.Printf("Found %s: owner=%s\n", id, owner)

			if owner.Version() != 3 {
				continue
			}

			if newOwner, ok := m[owner]; ok {
				ownerInts := uuidToIntArray(newOwner.UUID)
				entity.Set("Owner", dynbt.NewIntArray(ownerInts[:]))
				updated = true
			}
		}
	}

	if updated {
		var w bytes.Buffer
		w.WriteByte(1)
		gw := gzip.NewWriter(&w)

		err := nbt.NewEncoder(gw).Encode(&nbtdata, "")
		if err != nil {
			gw.Close()
			return nil, err
		}

		err = gw.Close()
		return w.Bytes(), err
	}
	return nil, nil
}
