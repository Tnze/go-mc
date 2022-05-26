package world

import (
	"fmt"
	"path/filepath"

	"github.com/Tnze/go-mc/level"
	"github.com/Tnze/go-mc/save"
	"github.com/Tnze/go-mc/save/region"
)

// TODO: Cache regions and chunks
type storage struct {
	regionDir string
}

func (s *storage) GetChunk(pos level.ChunkPos) (lc *level.Chunk, err error) {
	filename := fmt.Sprintf("r.%d.%d.mca", pos.X>>5, pos.Z>>5)

	var r *region.Region
	r, err = region.Open(filepath.Join(s.regionDir, filename))
	if err != nil {
		return nil, err
	}
	defer func() {
		err2 := r.Close()
		if err == nil && err2 != nil {
			err = err2
		}
	}()

	sector, err := r.ReadSector(region.In(pos.X, pos.Z))
	if err != nil {
		return nil, err
	}

	var sc save.Chunk
	err = sc.Load(sector)
	if err != nil {
		return nil, err
	}

	return level.ChunkFromSave(&sc)
}

func (s *storage) PutChunk(pos level.ChunkPos, c *level.Chunk) (err error) {
	var sc save.Chunk
	err = level.ChunkToSave(c, &sc)
	if err != nil {
		return
	}

	var data []byte
	data, err = sc.Data(1)
	if err != nil {
		return
	}

	filename := fmt.Sprintf("r.%d.%d.mca", pos.X>>5, pos.Z>>5)
	var r *region.Region
	r, err = region.Open(filepath.Join(s.regionDir, filename))
	if err != nil {
		return err
	}
	defer func() {
		err2 := r.Close()
		if err == nil && err2 != nil {
			err = err2
		}
	}()
	x, z := region.In(pos.X, pos.Z)
	err = r.WriteSector(x, z, data)
	return
}
