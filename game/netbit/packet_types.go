package netbit

// PacketType defines the type of message.
// Maximum 5 bits = [0-31] = 32 different PacketTypes.
type PacketType int

const (
	// CONNECT is the packet to request a connection
	CONNECT PacketType = iota

	// WELCOME is the packet for telling client is connected.
	WELCOME

	// FULL is the packet for telling clients the server is full.
	FULL

	// DISCONNECT is the packet that tells a client request disconnection.
	DISCONNECT
)

func (pt PacketType) EncodeNetBit() ([]byte, error) {

	return nil, nil
}

func (pt PacketType) DecodeNetBit([]byte) error {

	return nil
}
