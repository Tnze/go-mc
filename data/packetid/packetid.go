package packetid

/*
Reference: https://wiki.vg/index.php?title=Protocol&oldid=17753#Packet_format
*/

// Handshake client -> server
const (
	SPacketHandshake       = iota
	SPacketLegacyHandshake = 0xFE
)

// Status server -> client
const (
	CPacketStatusResponse = iota
	CPacketStatusPing
)

// Status client -> server
const (
	SPacketStatusRequest = iota
	SPacketStatusPing
)

// Login server -> client
const (
	CPacketLoginDisconnect = iota
	CPacketEncryptionRequest
	CPacketLoginSuccess
	CPacketSetCompression
	CPacketLoginPluginRequest
)

// Login client -> server
const (
	SPacketLoginStart = iota
	SPacketEncryptionResponse
	SPacketLoginPluginResponse
)

// Play server -> client
const (
	CPacketSpawnEntity = iota
	CPacketSpawnExperienceOrb
	CPacketSpawnPlayer
	CPacketEntityAnimation
	CPacketAwardStats
	CPacketAcknowledgeBlockChange
	CPacketSetBlockDestroyStage
	CPacketBlockEntityData
	CPacketBlockAction
	CPacketBlockUpdate
	CPacketBossBar
	CPacketServerDifficulty
	CPacketChatPreview
	CPacketClearTitles
	CPacketCommandSuggestions
	CPacketCommands
	CPacketCloseContainer
	CPacketSetContainerContent
	CPacketSetContainerProperty
	CPacketSetContainerSlot
	CPacketSetCooldown
	CPacketPluginMessage
	CPacketCustomSoundEffect
	CPacketDisconnect
	CPacketEntityStatus
	CPacketExplosion
	CPacketUnloadChunk
	CPacketChangeGameState
	CPacketOpenHorseWindow
	CPacketInitializeBorder
	CPacketKeepAlive
	CPacketChunkData
	CPacketWorldEvent
	CPacketParticles
	CPacketUpdateLight
	CPacketLogin
	CPacketMapData
	CPacketTradeList
	CPacketEntityPosition
	CPacketEntityPositionRotation
	CPacketEntityRotation
	CPacketVehicleMove
	CPacketOpenBook
	CPacketOpenContainer
	CPacketOpenSignEditor
	CPacketPing
	CPacketGhostRecipe
	CPacketPlayerAbilities
	CPacketChatMessage
	CPacketEndCombat
	CPacketEnterCombat
	CPacketCombatDeath
	CPacketPlayerInfo
	CPacketLookAt
	CPacketSyncPosition
	CPacketUpdateRecipeBook
	CPacketRemoveEntities
	CPacketRemoveEntityEffect
	CPacketResourcePackSend
	CPacketRespawn
	CPacketEntityHeadLook
	CPacketUpdateSectionBlock
	CPacketSelectAdvancementTab
	CPacketServerData
	CPacketSetActionBarText
	CPacketSetBorderCenter
	CPacketSetBorderLerpSize
	CPacketSetBorderSize
	CPacketSetBorderWarningDelay
	CPacketSetBorderWarningDistance
	CPacketSetCamera
	CPacketSetHeldItem
	CPlayerSetCenterChunk
	CPacketSetRenderDistance
	CPacketSetSpawnPosition
	CPacketSetDisplayChat
	CPacketDisplayObjective
	CPacketSetEntityMetadata
	CPacketLinkEntities
	CPacketEntityVelocity
	CPacketSetEquipment
	CPacketSetExperience
	CPacketUpdateHealth
	CPacketScoreboardObjective
	CPacketSetPassengers
	CPacketSetPlayerTeam
	CPacketSetScore
	CPacketSetSimulationDistance
	CPacketSetSubtitleText
	CPacketSetTime
	CPacketSetTitleText
	CPacketSetTitlesAnimation
	CPacketEntitySoundEffect
	CPacketSoundEffect
	CPacketStopSound
	CPacketSystemMessage
	CPacketSetTabList
	CPacketTagQueryResponse
	CPacketPickupItem
	CPacketTeleportEntity
	CPacketUpdateAdvancements
	CPacketUpdateAttributes
	CPacketEntityEffect
	CPacketUpdateRecipes
	CPacketSetTags
)

// Play client -> server
const (
	SPacketTeleportConfirm = iota
	SPacketQueryBlockEntityTag
	SPacketSetDifficulty
	SPacketChatCommand
	SPacketChatMessage
	SPacketChatPreview
	SPacketClientCommand
	SPacketClientSettings
	SPacketCommandSuggestion
	SPacketClickWindowButton
	SPacketClickWindow
	SPacketCloseWindow
	SPacketPluginMessage
	SPacketEditBook
	SPacketQueryEntityTag
	SPacketInteract
	SPacketJigsawGenerate
	SPacketKeepAlive
	SPacketLockDifficulty
	SPacketPlayerPosition
	SPacketPlayerPositionRotation
	SPacketPlayerRotation
	SPacketPlayerOnGround
	SPacketMoveVehicle
	SPacketPadddleBoat
	SPacketPickItem
	SPacketCraftRecipeRequest
	SPacketPlayerAbilities
	SPacketPlayerAction
	SPacketPlayerCommand
	SPacketPlayerInput
	SPacketPongResponse
	SPacketChangeRecipeBookState
	SPacketSetSeenRecipe
	SPacketRenameItem
	SPacketResourcePack
	SPacketSelectTrade
	SPacketSetBeaconEffect
	SPacketSetHeldItem
	SPacketSetCommandBlock
	SPacketSetCommandMinecart
	SPacketSetCreativeModeSlot
	SPacketSetJigsawBlock
	SPacketSetStructureBlock
	SPacketSetSign
	SPacketSwingArm
	SPacketTeleportToEntity
	SPacketUseItemOn
	SPacketUseItem
)
