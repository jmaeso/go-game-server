package netbit

import (
	"fmt"
	"reflect"
	"strconv"
)

type Decoder interface {
	DecodeNetBit([]byte) error
}

func Decode(data []byte, p *Packet) error {

	typeNumBits := reflect.TypeOf(*p).Field(0).Tag.Get("bit")
	fmt.Printf("tag_value: %v\n", typeNumBits)
	numBits, err := strconv.ParseUint(typeNumBits, 10, 32)
	if err != nil {
		return err
	}
	//p.Type = data[0] & ((1 << (numBits)) - 1)
	fmt.Printf("binary stream: %v\n", data[0]&((1<<(numBits))-1))
	dt := int(data[0] & ((1 << (numBits)) - 1))
	switch dt {
	case int(CONNECT):
		fmt.Printf("Packet type = CONNECT \n")
	case int(WELCOME):
		fmt.Printf("Packet type = WELCOME \n")
	case int(FULL):
		fmt.Printf("Packet type = FULL \n")
	case int(DISCONNECT):
		fmt.Printf("Packet type = DISCONNECT \n")
	default:
		fmt.Printf("int value of dt: %d\n", dt)
	}

	p.Type = PacketType(dt)

	if err := p.DecodeNetBit(data[:numBits]); err != nil {
		return err
	}

	return nil
}
