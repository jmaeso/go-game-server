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
	numBits, err := strconv.ParseUint(typeNumBits, 10, 64)
	if err != nil {
		return err
	}
	//p.Type = data[0] & ((1 << (numBits)) - 1)
	//Following code goes through right to left
	// data[0] & ((1 << (numBits)) - 1)
	//fmt.Printf("first byte stream: %v\n", binary.BigEndian(data[0]))
	// for _, b := range data {
	// 	for i := uint(len(data)) * 8; i >= 0; i-- {
	// 		fmt.Printf(b & (1 << i) >> i)
	// 	}
	// }

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
