package netbit

import (
	"fmt"
	"reflect"
	"strconv"
)

type Decoder interface {
	DecodeNetBit([512]byte) error
}

func Decode(data [512]byte, p *Packet) error {

	typeNumBits := reflect.TypeOf(*p).Field(0).Tag.Get("bit")
	fmt.Printf("tag_value: %v\n", typeNumBits)
	numBits, err := strconv.ParseUint(typeNumBits, 10, 32)
	if err != nil {
		return err
	}
	//p.Type = data[0] & ((1 << (numBits)) - 1)
	fmt.Printf("binary stream: %v\n", data[0]&((1<<(numBits))-1))
	return nil
}
