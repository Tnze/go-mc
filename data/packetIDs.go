package data

// Clientbound packet IDs
const (
	SpawnObject int32 = iota //0x00
	SpawnExperienceOrb
	SpawnGlobalEntity
	SpawnMob
	SpawnPainting
	SpawnPlayer
	AnimationClientbound
	Statistics
	AcknowledgePlayerDigging
	BlockBreakAnimation
	UpdateBlockEntity
	BlockAction
	BlockChange
	BossBar
	ServerDifficulty
	ChatMessageClientbound

	MultiBlockChange //0x10
	TabComplete
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

	OpenHorseWindow //0x20
	KeepAliveClientbound
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

	OpenSignEditor //0x30
	CraftRecipeResponse
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

	HeldItemChangeClientbound //0x40
	UpdateViewPosition
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

	Title //0x50
	EntitySoundEffect
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
	Tags //0x5C
)

// Serverbound packet IDs
const (
	TeleportConfirm int32 = iota //0x00
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
