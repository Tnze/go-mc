package data

// Clientbound packet IDs
const (
	SpawnObject int32 = iota //0x00
	SpawnExperienceOrb
	SpawnLivingEntity
	SpawnPainting
	SpawnPlayer
	EntityAnimationClientbound
	Statistics
	AcknowledgePlayerDigging
	BlockBreakAnimation
	BlockEntityData
	BlockAction
	BlockChange
	BossBar
	ServerDifficulty
	ChatMessageClientbound
	TabComplete

	DeclareCommands //0x10
	WindowConfirmationClientbound
	CloseWindowClientbound
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
	KeepAliveClientbound

	ChunkData //0x20
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
	CraftRecipeResponse

	PlayerAbilitiesClientbound //0x30
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
	MultiBlockChange
	SelectAdvancementTab
	WorldBorder
	Camera
	HeldItemChangeClientbound

	UpdateViewPosition //0x40
	UpdateViewDistance
	SpawnPosition
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
	GenerateStructure

	KeepAliveServerbound //0x10
	LockDifficulty
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

	ResourcePackStatus //0x20
	AdvancementTab
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
	UseItem //0x2E
)
