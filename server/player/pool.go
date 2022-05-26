package player

import (
	"compress/gzip"
	"os"
	"path/filepath"

	"github.com/google/uuid"

	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
)

type storage struct {
	playerdataDir string
}

func (s *storage) GetPlayer(id uuid.UUID) (data save.PlayerData, err error) {
	filename := id.String() + ".dat"

	f, err := os.Open(filepath.Join(s.playerdataDir, filename))
	if err != nil {
		return save.PlayerData{}, err
	}
	defer f.Close()

	r, err := gzip.NewReader(f)
	if err != nil {
		return save.PlayerData{}, err
	}

	return save.ReadPlayerData(r)
}

func (s *storage) PutPlayer(pos level.ChunkPos, c *level.Chunk) (err error) {
	return nil
}
