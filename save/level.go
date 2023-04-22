package save

import (
	"io"

	"github.com/Tnze/go-mc/nbt"
)

type Level struct {
	Data LevelData
}

type LevelData struct {
	AllowCommands                byte `nbt:"allowCommands"`
	BorderCenterX, BorderCenterZ float64
	BorderDamagePerBlock         float64
	BorderSafeZone               float64
	BorderSize                   float64
	BorderSizeLerpTarget         float64
	BorderSizeLerpTime           int64
	BorderWarningBlocks          float64
	BorderWarningTime            float64
	ClearWeatherTime             int32 `nbt:"clearWeatherTime"`
	CustomBossEvents             map[string]CustomBossEvent
	DataPacks                    struct {
		Enabled, Disabled []string
	}
	DataVersion      int32
	DayTime          int64
	Difficulty       byte
	DifficultyLocked bool
	DimensionData    struct {
		TheEnd struct {
			DragonFight struct {
				Gateways         []int32
				DragonKilled     byte
				PreviouslyKilled byte
			}
		} `nbt:"1"`
	}
	DragonFight struct {
		Gateways           []int32
		DragonKilled       bool
		NeedsStateScanning bool
		PreviouslyKilled   bool
	}
	GameRules              map[string]string
	WorldGenSettings       WorldGenSettings
	GameType               int32
	HardCore               bool `nbt:"hardcore"`
	Initialized            bool `nbt:"initialized"`
	LastPlayed             int64
	LevelName              string
	MapFeatures            bool
	Player                 map[string]any
	Raining                bool  `nbt:"raining"`
	RainTime               int32 `nbt:"rainTime"`
	RandomSeed             int64
	ScheduledEvents        []nbt.RawMessage
	ServerBrands           []string
	SizeOnDisk             int64
	SpawnAngle             float32
	SpawnX, SpawnY, SpawnZ int32
	Thundering             bool  `nbt:"thundering"`
	ThunderTime            int32 `nbt:"thunderTime"`
	Time                   int64
	Version                struct {
		ID       int32 `nbt:"Id"`
		Name     string
		Series   string
		Snapshot byte
	}
	StorageVersion             int32 `nbt:"version"`
	WanderingTraderId          []int32
	WanderingTraderSpawnChance int32
	WanderingTraderSpawnDelay  int32
	WasModded                  bool
}

type CustomBossEvent struct {
	Players        [][]int32
	Color          string
	CreateWorldFog bool
	DarkenScreen   bool
	Max            int32
	Value          int32
	Name           string
	Overlay        string
	PlayBossMusic  bool
	Visible        bool
}

func ReadLevel(r io.Reader) (data Level, err error) {
	decoder := nbt.NewDecoder(r)
	decoder.DisallowUnknownFields()
	_, err = decoder.Decode(&data)
	return
}
