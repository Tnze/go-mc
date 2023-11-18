package packetid

// This file might be remove in the future go-mc!!!

// Login Clientbound
const (
	LoginDisconnect = iota
	LoginEncryptionRequest
	LoginSuccess
	LoginCompression
	LoginPluginRequest
)

// Login Serverbound
const (
	LoginStart = iota
	LoginEncryptionResponse
	LoginPluginResponse
)

// Status Clientbound
const (
	StatusResponse = iota
	StatusPongResponse
)

// Status Serverbound
const (
	StatusRequest = iota
	StatusPingRequest
)
