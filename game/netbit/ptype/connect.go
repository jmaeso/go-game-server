package ptype

type Connect struct {
	Nick string `bit:"10"`
}

func (c *Connect) EncodeNetBit() ([]byte, error) {

	// TODO: Create bit stream.
	return nil, nil
}
func (c *Connect) DecodeNetBit([]byte) error {

	return nil
}
