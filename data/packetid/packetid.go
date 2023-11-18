package packetid

//go:generate stringer -type ClientboundPacketID
//go:generate stringer -type ServerboundPacketID
type (
	ClientboundPacketID int32
	ServerboundPacketID int32
)

// Login Clientbound
const (
	ClientboundLoginDisconnect        ClientboundPacketID = iota // LoginDisconnect
	ClientboundLoginEncryptionRequest                            // Hello
	ClientboundLoginSuccess                                      // GameProfile
	ClientboundLoginCompression                                  // LoginCompression
	ClientboundLoginPluginRequest                                // CustomQuery
)

// Login Serverbound
const (
	ServerboundLoginStart              ServerboundPacketID = iota // Hello
	ServerboundLoginEncryptionResponse                            // Key
	ServerboundLoginPluginResponse                                // CustomQueryAnswer
	ServerboundLoginAcknowledged                                  // LoginAcknowledged
)

// Status Clientbound
const (
	ClientboundStatusResponse ClientboundPacketID = iota
	ClientboundStatusPongResponse
)

// Status Serverbound
const (
	ServerboundStatusRequest ServerboundPacketID = iota
	ServerboundStatusPingRequest
)

// Configuration Clientbound
const (
	ClientboundConfigCustomPayload ClientboundPacketID = iota
	ClientboundConfigDisconnect
	ClientboundConfigFinishConfiguration
	ClientboundConfigKeepAlive
	ClientboundConfigPing
	ClientboundConfigRegistryData
	ClientboundConfigResourcePack
	ClientboundConfigUpdateEnabledFeatures
	ClientboundConfigUpdateTags
)

// Configuration Serverbound
const (
	ServerboundConfigClientInformation ServerboundPacketID = iota
	ServerboundConfigCustomPayload
	ServerboundConfigFinishConfiguration
	ServerboundConfigKeepAlive
	ServerboundConfigPong
	ServerboundConfigResourcePack
)

// Game Clientbound
const (
	BundleDelimiter ClientboundPacketID = iota
	ClientboundAddEntity
	ClientboundAddExperienceOrb
	ClientboundAnimate
	ClientboundAwardStats
	ClientboundBlockChangedAck
	ClientboundBlockDestruction
	ClientboundBlockEntityData
	ClientboundBlockEvent
	ClientboundBlockUpdate
	ClientboundBossEvent
	ClientboundChangeDifficulty
	ClientboundChunkBatchFinished
	ClientboundChunkBatchStart
	ClientboundChunksBiomes
	ClientboundClearTitles
	ClientboundCommandSuggestions
	ClientboundCommands
	ClientboundContainerClose
	ClientboundContainerSetContent
	ClientboundContainerSetData
	ClientboundContainerSetSlot
	ClientboundCooldown
	ClientboundCustomChatCompletions
	ClientboundCustomPayload
	ClientboundDamageEvent
	ClientboundDeleteChat
	ClientboundDisconnect
	ClientboundDisguisedChat
	ClientboundEntityEvent
	ClientboundExplode
	ClientboundForgetLevelChunk
	ClientboundGameEvent
	ClientboundHorseScreenOpen
	ClientboundHurtAnimation
	ClientboundInitializeBorder
	ClientboundKeepAlive
	ClientboundLevelChunkWithLight
	ClientboundLevelEvent
	ClientboundLevelParticles
	ClientboundLightUpdate
	ClientboundLogin
	ClientboundMapItemData
	ClientboundMerchantOffers
	ClientboundMoveEntityPos
	ClientboundMoveEntityPosRot
	ClientboundMoveEntityRot
	ClientboundMoveVehicle
	ClientboundOpenBook
	ClientboundOpenScreen
	ClientboundOpenSignEditor
	ClientboundPing
	ClientboundPongResponse
	ClientboundPlaceGhostRecipe
	ClientboundPlayerAbilities
	ClientboundPlayerChat
	ClientboundPlayerCombatEnd
	ClientboundPlayerCombatEnter
	ClientboundPlayerCombatKill
	ClientboundPlayerInfoRemove
	ClientboundPlayerInfoUpdate
	ClientboundPlayerLookAt
	ClientboundPlayerPosition
	ClientboundRecipe
	ClientboundRemoveEntities
	ClientboundRemoveMobEffect
	ClientboundResourcePack
	ClientboundRespawn
	ClientboundRotateHead
	ClientboundSectionBlocksUpdate
	ClientboundSelectAdvancementsTab
	ClientboundServerData
	ClientboundSetActionBarText
	ClientboundSetBorderCenter
	ClientboundSetBorderLerpSize
	ClientboundSetBorderSize
	ClientboundSetBorderWarningDelay
	ClientboundSetBorderWarningDistance
	ClientboundSetCamera
	ClientboundSetCarriedItem
	ClientboundSetChunkCacheCenter
	ClientboundSetChunkCacheRadius
	ClientboundSetDefaultSpawnPosition
	ClientboundSetDisplayObjective
	ClientboundSetEntityData
	ClientboundSetEntityLink
	ClientboundSetEntityMotion
	ClientboundSetEquipment
	ClientboundSetExperience
	ClientboundSetHealth
	ClientboundSetObjective
	ClientboundSetPassengers
	ClientboundSetPlayerTeam
	ClientboundSetScore
	ClientboundSetSimulationDistance
	ClientboundSetSubtitleText
	ClientboundSetTime
	ClientboundSetTitleText
	ClientboundSetTitlesAnimation
	ClientboundSoundEntity
	ClientboundSound
	ClientboundStartConfiguration
	ClientboundStopSound
	ClientboundSystemChat
	ClientboundTabList
	ClientboundTagQuery
	ClientboundTakeItemEntity
	ClientboundTeleportEntity
	ClientboundUpdateAdvancements
	ClientboundUpdateAttributes
	ClientboundUpdateMobEffect
	ClientboundUpdateRecipes
	ClientboundUpdateTags
	ClientboundPacketIDGuard
)

// Game Serverbound
const (
	ServerboundAcceptTeleportation ServerboundPacketID = iota
	ServerboundBlockEntityTagQuery
	ServerboundChangeDifficulty
	ServerboundChatAck
	ServerboundChatCommand
	ServerboundChat
	ServerboundChatSessionUpdate
	ServerboundChunkBatchReceived
	ServerboundClientCommand
	ServerboundClientInformation
	ServerboundCommandSuggestion
	ServerboundConfigurationAcknowledged
	ServerboundContainerButtonClick
	ServerboundContainerClick
	ServerboundContainerClose
	ServerboundCustomPayload
	ServerboundEditBook
	ServerboundEntityTagQuery
	ServerboundInteract
	ServerboundJigsawGenerate
	ServerboundKeepAlive
	ServerboundLockDifficulty
	ServerboundMovePlayerPos
	ServerboundMovePlayerPosRot
	ServerboundMovePlayerRot
	ServerboundMovePlayerStatusOnly
	ServerboundMoveVehicle
	ServerboundPaddleBoat
	ServerboundPickItem
	ServerboundPingRequest
	ServerboundPlaceRecipe
	ServerboundPlayerAbilities
	ServerboundPlayerAction
	ServerboundPlayerCommand
	ServerboundPlayerInput
	ServerboundPong
	ServerboundRecipeBookChangeSettings
	ServerboundRecipeBookSeenRecipe
	ServerboundRenameItem
	ServerboundResourcePack
	ServerboundSeenAdvancements
	ServerboundSelectTrade
	ServerboundSetBeacon
	ServerboundSetCarriedItem
	ServerboundSetCommandBlock
	ServerboundSetCommandMinecart
	ServerboundSetCreativeModeSlot
	ServerboundSetJigsawBlock
	ServerboundSetStructureBlock
	ServerboundSignUpdate
	ServerboundSwing
	ServerboundTeleportToEntity
	ServerboundUseItemOn
	ServerboundUseItem
	ServerboundPacketIDGuard
)
