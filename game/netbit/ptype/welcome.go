package ptype

import "github.com/jmaeso/go-game-server/game/netbit"

type Welcome struct {
	ID         string `bit:"10"`
	NumPlayers int    `bit:"3"` //max 8 players
}

func (w *Welcome) Encode(*netbit.Packet) ([]byte, error) {

	return nil, nil
}

func (w *Welcome) Decode([]byte) (*netbit.Packet, error) {

	return nil, nil
}
