package main

import (
	"compress/gzip"
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/Tnze/go-mc/nbt"
	"github.com/Tnze/go-mc/nbt/dynbt"
	"github.com/google/uuid"
)

func readPlayerdata(dir string, m map[uuid.UUID]UserCache) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open playerdata folder: %v", err)
		return
	}
	for _, files := range entries {
		filename := files.Name()
		if ext := filepath.Ext(filename); ext != ".dat" {
			fmt.Fprintf(os.Stderr, "Unkown file type: %s\n", ext)
			continue
		}

		// Parse old UUID from filename
		oldID, err := uuid.Parse(strings.TrimSuffix(filename, ".dat"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse filename as uuid: %v\n", err)
			continue
		}

		nbtdata, err := readNbtData(filepath.Join(dir, filename))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read %s nbt data\n", filename)
			continue
		}

		// Read old UUID from nbt
		uuidInts := nbtdata.Get("UUID").IntArray()
		uuidBytes, err := intArrayToUUID(uuidInts)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read %s UUID\n", filename)
			continue
		}

		// Does they matche?
		if oldID != uuidBytes {
			fmt.Fprintf(os.Stderr, "UUID in filename and nbt data don't match, what happend?\n")
		}

		if ver := uuidBytes.Version(); ver != 3 { // v3 is for offline players
			fmt.Printf("Ignoring UUID: %v version: %d\n", uuidBytes, ver)
			continue
		}

		newUser, ok := m[oldID]
		if !ok {
			fmt.Printf("Skip user: %v\n", oldID)
			continue
		}

		// Update UUID
		ints := uuidToIntArray(newUser.UUID)
		nbtdata.Set("UUID", dynbt.NewIntArray(ints[:]))

		// Create new .dat file
		err = writeNbtData(dir, newUser.UUID.String(), &nbtdata)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to write %s's .dat file: %v\n", newUser.Name, err)
			continue
		}
	}
}

func readNbtData(filepath string) (nbtdata dynbt.Value, err error) {
	file, err := os.Open(filepath)
	if err != nil {
		return nbtdata, fmt.Errorf("failed to open userdata: %w", err)
	}
	defer file.Close()

	r, err := gzip.NewReader(file)
	if err != nil {
		return nbtdata, fmt.Errorf("failed to decompress userdata: %w", err)
	}

	_, err = nbt.NewDecoder(r).Decode(&nbtdata)
	if err != nil {
		return nbtdata, fmt.Errorf("failed to parse userdata: %w", err)
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
