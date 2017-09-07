package netbit

const tag = "bit"

type Packet struct {
	Type Packer `bit:"5"`
	Body Packer `bit:"507"` /*Connect or Welcome or etc.*/
}

type Packer interface {
	// Encode() ([]byte, error)
	// Decode() (interface{}, error)
	Encoder
	Decoder
}

// func (p *Packet) Encode() ([]byte, error) {
// 	var err error

// 	typeStream, err := p.Type.Encode()
// 	if err != nil {
// 		return nil, errors.New("Packet_encode: could not encode packet type. err: " + err.Error())
// 	}

// 	bodyStream, err := p.Body.Encode()
// 	if err != nil {
// 		return nil, errors.New("Packet_encode: could not encode body. err: " + err.Error())
// 	}

// 	// TODO: concatenate streams
// 	log.Printf("Packet_encode: encoded type: %v\n", typeStream)
// 	log.Printf("Packet_encode: encoded body: %v\n", bodyStream)

// 	return []byte{}, nil
// }

// func (p *Packet) Decode() (*Packet, error) {
// 	var err error

// 	pt, err := p.Type.Decode()
// 	if err != nil {
// 		return nil, errors.New("Packet_decode: could not decode packet type. err: " + err.Error())
// 	}

// 	body, err := p.Body.Decode()
// 	if err != nil {
// 		return nil, errors.New("Packet_decode: could not decode body. err: " + err.Error())
// 	}

// 	return &Packet{Type: pt, Body: body}, nil
// }
