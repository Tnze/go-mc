package packetid

// This file might be remove in the future go-mc!!!

// Login Clientbound
//
// Deprecated: Legacy name, might be removed in future version
const (
	LoginDisconnect        = ClientboundLoginLoginDisconnect
	LoginEncryptionRequest = ClientboundLoginHello
	LoginSuccess           = ClientboundLoginGameProfile
	LoginCompression       = ClientboundLoginLoginCompression
	LoginPluginRequest     = ClientboundLoginCustomQuery

	ClientboundLoginDisconnect        = ClientboundLoginLoginDisconnect
	ClientboundLoginEncryptionRequest = ClientboundLoginHello
	ClientboundLoginSuccess           = ClientboundLoginGameProfile
	ClientboundLoginCompression       = ClientboundLoginLoginCompression
	ClientboundLoginPluginRequest     = ClientboundLoginCustomQuery
)

// Login Serverbound
//
// Deprecated: Legacy name, might be removed in future version
const (
	LoginStart              = ServerboundLoginHello
	LoginEncryptionResponse = ServerboundLoginKey
	LoginPluginResponse     = ServerboundLoginCustomQueryAnswer

	ServerboundLoginStart              = ServerboundLoginHello
	ServerboundLoginEncryptionResponse = ServerboundLoginKey
	ServerboundLoginPluginResponse     = ServerboundLoginCustomQueryAnswer
	ServerboundLoginAcknowledged       = ServerboundLoginLoginAcknowledged
)

// Status Clientbound
//
// Deprecated: Legacy name, might be removed in future version
const (
	StatusResponse     = ClientboundStatusStatusResponse
	StatusPongResponse = ClientboundStatusPongResponse

	ClientboundStatusResponse = ClientboundStatusStatusResponse
)

// Status Serverbound
//
// Deprecated: Legacy name, might be removed in future version
const (
	StatusRequest     = ServerboundStatusStatusRequest
	StatusPingRequest = ServerboundStatusPingRequest

	ServerboundStatusRequest = ServerboundStatusStatusRequest
)
