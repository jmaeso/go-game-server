package netbit

const tag = "bit"

type Packet struct {
	Type PacketType `bit:"5"`
	Body Packer     `bit:"507"` /*Connect or Welcome or etc.*/
}

type Packer interface {
	Encode(*Packet) ([]byte, error)
	Decode([]byte) (*Packet, error)
}
