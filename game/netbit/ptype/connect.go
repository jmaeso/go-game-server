package ptype

import "github.com/jmaeso/go-game-server/game/netbit"

type Connect struct {
	Nick string `bit:"10"`
}

func (c *Connect) Encode(*netbit.Packet) ([]byte, error) {

	return nil, nil
}
func (c *Connect) Decode([]byte) (*netbit.Packet, error) {

	return nil, nil
}
