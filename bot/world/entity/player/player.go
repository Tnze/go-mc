package player

import "github.com/Tnze/go-mc/bot/world/entity"

// Player includes the player's status.
type Player struct {
	entity.Entity
	UUID [2]int64 //128bit UUID

	X, Y, Z    float64
	Yaw, Pitch float32
	OnGround   bool

	HeldItem  int //拿着的物品栏位

	Health         float32 //血量
	Food           int32   //饱食度
	FoodSaturation float32 //食物饱和度
}
