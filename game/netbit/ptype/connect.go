package ptype

type Connect struct {
	Nick string `bit:"10"`
}

func (c Connect) EncodeNetBit() ([]byte, error) {

	// TODO: Create bit stream.
	return nil, nil
}

// Fullfy Nick with data
func (c Connect) DecodeNetBit(data []byte) error {
	c.Nick = string(data[:])

	return nil
}
