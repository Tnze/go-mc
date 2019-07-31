package save

import (
	"github.com/Tnze/go-mc/nbt"
	"io"
)

type PlayerData struct {
	Pos    [3]float64
	Motion [3]float64
}

func ReadPlayerData(r io.Reader) (data PlayerData, err error) {
	err = nbt.NewDecoder(r).Decode(&data)
	return
}
