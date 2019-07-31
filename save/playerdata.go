package save

import (
	"encoding/binary"
	"github.com/Tnze/go-mc/nbt"
	"github.com/google/uuid"
	"io"
)

type PlayerData struct {
	DataVersion int32

	Dimension    int32
	Pos          [3]float64
	Motion       [3]float64
	Rotation     [2]float32
	FallDistance float32
	FallFlying   byte
	OnGround     byte

	UUID                uuid.UUID `nbt:"-"`
	UUIDLeast, UUIDMost int64

	PlayerGameType  int32 `nbt:"playerGameType"`
	Air             int16
	DeathTime       int16
	Fire            int16
	HurtTime        int16
	Health          float32
	HurtByTimestamp int32

	Invulnerable     byte
	SeenCredits      byte `nbt:"seenCredits"`
	SelectedItemSlot int32
	Score            int32
	AbsorptionAmount float32

	XpLevel int32
	XpP     float32
	XpTotal int32
	XpSeed  int32

	FoodExhaustionLevel float32 `nbt:"foodExhaustionLevel"`
	FoodLevel           int32   `nbt:"foodLevel"`
	FoodSaturationLevel float32 `nbt:"foodSaturationLevel"`
	FoodTickTimer       int32   `nbt:"foodTickTimer"`

	Attributes []struct {
		Base float64
		Name string
	}
	Abilities struct {
		FlySpeed     float32 `nbt:"flySpeed"`
		WalkSpeed    float32 `nbt:"warkSpeed"`
		Flying       byte    `nbt:"flying"`
		InstantBuild byte    `nbt:"instabuild"`
		Invulnerable byte    `nbt:"invulnerable"`
		MayBuild     byte    `nbt:"mayBuild"`
		MayFly       byte    `nbt:"mayFly"`
	} `nbt:"abilities"`
}

func ReadPlayerData(r io.Reader) (data PlayerData, err error) {
	err = nbt.NewDecoder(r).Decode(&data)
	//parse UUID from two int64s
	binary.BigEndian.PutUint64(data.UUID[:], uint64(data.UUIDMost))
	binary.BigEndian.PutUint64(data.UUID[8:], uint64(data.UUIDLeast))
	return
}
