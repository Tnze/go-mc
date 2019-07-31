package save

import (
	"github.com/Tnze/go-mc/nbt"
	"io"
)

type Level struct {
	Data struct {
		DataVersion int32
		NBTVersion  int32 `nbt:"version"`
		Version     struct {
			ID       int32 `nbt:"Id"`
			Name     string
			Snapshot byte
		}
		GameType         int32
		Difficulty       byte
		DifficultyLocked byte
		HardCore         byte `nbt:"hardcore"`
		Initialized      byte `nbt:"initialized"`
		AllowCommands    byte `nbt:"allowCommands"`

		MapFeatures      byte
		LevelName        string
		GeneratorName    string `nbt:"generatorName"`
		GeneratorVersion int32  `nbt:"generatorVersion"`
		RandomSeed       int64

		SpawnX, SpawnY, SpawnZ int32

		BorderCenterX, BorderCenterZ float64
		BorderDamagePerBlock         float64
		BorderSafeZone               float64
		BorderSize                   float64
		BorderSizeLerpTarget         float64
		BorderSizeLerpTime           int64
		BorderWarningBlocks          float64
		BorderWarningTime            float64

		GameRules map[string]string
		DataPacks struct {
			Enabled, Disabled []string
		}
		DimensionData struct {
			TheEnd struct {
				DragonFight struct {
					Gateways         []int32
					DragonKilled     byte
					PreviouslyKilled byte
				}
			} `nbt:"1"`
		}

		Raining          byte  `nbt:"raining"`
		Thundering       byte  `nbt:"thundering"`
		RainTime         int32 `nbt:"rainTime"`
		ThunderTime      int32 `nbt:"thunderTime"`
		ClearWeatherTime int32 `nbt:"clearWeatherTime"`

		Time       int64
		DayTime    int64
		LastPlayed int64
	}
}

func ReadLevel(r io.Reader) (data Level, err error) {
	err = nbt.NewDecoder(r).Decode(&data)
	return
}
