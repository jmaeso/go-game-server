package app

import (
	"errors"
	"fmt"
	"net"
)

type (
	Server struct {
		Conn *net.UDPConn
	}
)

func (s *Server) Init(address string) error {
	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return errors.New("server_init: could not resolve address. err: " + err.Error())
	}

	s.Conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		fmt.Printf("server_init: could not listen address. err: " + err.Error())
	}

	return nil
}
