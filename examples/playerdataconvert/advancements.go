package main

import (
	"fmt"
	"os"
	"path/filepath"
	"strings"

	"github.com/google/uuid"
)

func readAdvancements(dir string, m map[uuid.UUID]UserCache) {
	entries, err := os.ReadDir(dir)
	if err != nil {
		fmt.Fprintf(os.Stderr, "Failed to open playerdata folder: %v", err)
		return
	}
	for _, files := range entries {
		filename := files.Name()
		if ext := filepath.Ext(filename); ext != ".json" {
			fmt.Fprintf(os.Stderr, "Unkown file type: %s\n", ext)
			continue
		}

		// Parse old UUID from filename
		oldID, err := uuid.Parse(strings.TrimSuffix(filename, ".json"))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Unable to parse filename as uuid: %v\n", err)
			continue
		}

		if ver := oldID.Version(); ver != 3 { // v3 is for offline players
			fmt.Printf("Ignoring UUID: %v version: %d\n", oldID, ver)
			continue
		}

		newUser, ok := m[oldID]
		if !ok {
			fmt.Printf("Skip user: %v\n", oldID)
			continue
		}

		content, err := os.ReadFile(filepath.Join(dir, filename))
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to read json file: %v\n", err)
			continue
		}

		newFile := newUser.UUID.String() + ".json"
		err = os.WriteFile(filepath.Join(dir, newFile), content, 0o666)
		if err != nil {
			fmt.Fprintf(os.Stderr, "Failed to write json file: %v\n", err)
			continue
		}

		fmt.Printf("Converted advancement file: %s\n", newFile)
	}
}
