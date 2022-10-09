package types

import "time"

const (
	// ModuleName defines the module name
	ModuleName = "checkers"

	// StoreKey defines the primary module store key
	StoreKey = ModuleName

	// RouterKey defines the module's message routing key
	RouterKey = ModuleName

	// MemStoreKey defines the in-memory store key
	MemStoreKey = "mem_checkers"
)

func KeyPrefix(p string) []byte {
	return []byte(p)
}

const (
	SystemInfoKey = "SystemInfo/value/"
)

// Keys for create game event
const (
	GameCreatedEventType      = "new-game-created" // Indicates what event type to listen to
	GameCreatedEventCreator   = "creator"          // Subsidiary information
	GameCreatedEventGameIndex = "game-index"       // What game is relevant
	GameCreatedEventBlack     = "black"            // Is it relevant to me?
	GameCreatedEventRed       = "red"              // Is it relevant to me?
)

// Keys for make move event
const (
	MovePlayedEventType      = "move-played"
	MovePlayedEventCreator   = "creator"
	MovePlayedEventGameIndex = "game-index"
	MovePlayedEventCapturedX = "captured-x"
	MovePlayedEventCapturedY = "captured-y"
	MovePlayedEventWinner    = "winner"
)

// Keys for reject game event
const (
	GameRejectedEventType      = "game-rejected"
	GameRejectedEventCreator   = "creator"
	GameRejectedEventGameIndex = "game-index"
)

const (
	NoFifoIndex = "-1"
)

const (
	MaxTurnDuration = time.Duration(24 * 3_600 * 1000_000_000) // 24 hours
	DeadlineLayout  = "2006-01-02 15:04:05.999999999 +0000 UTC"
)
