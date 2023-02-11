package msg

import "github.com/Tnze/go-mc/chat"

// EventsHandler is a collection of event handlers.
// Fill the fields with your handler functions and pass this struct to [New] to create the msg manager.
// The handler functions will be called when the corresponding event is triggered.
// Leave the fields as nil if you don't want to handle the event.
type EventsHandler struct {
	// SystemChat handles messages sent by gaming system.
	//
	// In vanilla client:
	// If overlay is false, the message will be displayed in the chat box.
	// If overlay is true, the message will be displayed on the top of the hot-bar.
	SystemChat func(msg chat.Message, overlay bool) error

	// PlayerChatMessage handles messages sent by players.
	//
	// Message signing system is added in 1.19. The message and its context could be signed by the player's private key.
	// The manager tries to verify the message signature through the player's public key,
	// and return the result as validated boolean.
	PlayerChatMessage func(msg chat.Message, validated bool) error

	// DisguisedChat handles DisguisedChat message.
	//
	// DisguisedChat message used to send system chat.
	// Now it is used to send messages from "/say" command from server console.
	DisguisedChat func(msg chat.Message) error
}
