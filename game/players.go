package game

// Player struct represents a user for the server.
type Player struct {
	ID   string
	Nick string
	Addr Address
}

// Players is a group of Player.
type Players []Player
