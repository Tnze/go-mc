package data

//Clientbound packet IDs
const (
	SpawnObject byte = iota //0x00
	SpawnExperienceOrb
	SpawnGlobalEntity
	SpawnMob
	SpawnPainting
	SpawnPlayer
	AnimationClientbound
	Statistics
	BlockBreakAnimation
	UpdateBlockEntity
	BlockAction
	BlockChange
	BossBar
	ServerDifficulty
	ChatMessageClientbound
	MultiBlockChange

	TabComplete //0x10
	DeclareCommands
	ConfirmTransaction
	CloseWindow
	WindowItems
	WindowProperty
	SetSlot
	SetCooldown
	PluginMessageClientbound
	NamedSoundEffect
	DisconnectPlay
	EntityStatus
	Explosion
	UnloadChunk
	ChangeGameState
	OpenHorseWindow

	KeepAliveClientbound //0x20
	ChunkData
	Effect
	Particle
	UpdateLight
	JoinGame
	MapData
	TradeList
	EntityRelativeMove
	EntityLookAndRelativeMove
	EntityLook
	Entity
	VehicleMoveClientbound
	OpenBook
	OpenWindow
	OpenSignEditor

	CraftRecipeResponse //0x30
	PlayerAbilitiesClientbound
	CombatEvent
	PlayerInfo
	FacePlayer
	PlayerPositionAndLookClientbound
	UnlockRecipes
	DestroyEntities
	RemoveEntityEffect
	ResourcePackSend
	Respawn
	EntityHeadLook
	SelectAdvancementTab
	WorldBorder
	Camera
	HeldItemChangeClientbound

	UpdateViewPosition //0x40
	UpdateViewDistance
	DisplayScoreboard
	EntityMetadata
	AttachEntity
	EntityVelocity
	EntityEquipment
	SetExperience
	UpdateHealth
	ScoreboardObjective
	SetPassengers
	Teams
	UpdateScore
	SpawnPosition
	TimeUpdate
	Title

	EntitySoundEffect //0x50
	SoundEffect
	StopSound
	PlayerListHeaderAndFooter
	NBTQueryResponse
	CollectItem
	EntityTeleport
	Advancements
	EntityProperties
	EntityEffect
	DeclareRecipes
	Tags //0x5B
)

//Serverbound packet IDs
const (
	TeleportConfirm byte = iota //0x00
	QueryBlockNBT
	SetDifficulty
	ChatMessageServerbound
	ClientStatus
	ClientSettings
	TabCompleteServerbound
	ConfirmTransactionServerbound
	ClickWindowButton
	ClickWindow
	CloseWindowServerbound
	PluginMessageServerbound
	EditBook
	QueryEntityNBT
	UseEntity
	KeepAliveServerbound

	LockDifficulty //0x10
	PlayerPosition
	PlayerPositionAndLookServerbound
	PlayerLook
	Player
	VehicleMoveServerbound
	SteerBoat
	PickItem
	CraftRecipeRequest
	PlayerAbilitiesServerbound
	PlayerDigging
	EntityAction
	SteerVehicle
	RecipeBookData
	NameItem
	ResourcePackStatus

	AdvancementTab //0x20
	SelectTrade
	SetBeaconEffect
	HeldItemChangeServerbound
	UpdateCommandBlock
	UpdateCommandBlockMinecart
	CreativeInventoryAction
	UpdateJigsawBlock
	UpdateStructureBlock
	UpdateSign
	AnimationServerbound
	Spectate
	PlayerBlockPlacement
	UseItem //0x2D
)
