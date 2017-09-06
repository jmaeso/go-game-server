package game

import (
	"errors"
	"log"
	"net"
	"time"
)

type Server struct {
	Conn *net.UDPConn
}

func (s *Server) Init(settings *ServerSettings) error {
	url := settings.URL
	port := settings.Port
	if port == "" {
		return errors.New("server_init: port required")
	}

	address := url + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return errors.New("server_init: could not resolve address. err: " + err.Error())
	}

	s.Conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return errors.New("server_init: could not start connection. err: " + err.Error())
	}

	return nil
}

func (s *Server) Handle() error {
	var err error

	for err == nil {
		var buff [512]byte
		_, addr, err := s.Conn.ReadFrom(buff[0:])
		if err != nil {
			err = errors.New("server_handle: could not read message. err: " + err.Error())
		}

		if err = s.process(buff, addr); err != nil {
			err = errors.New("server_handle: could not process message. err: " + err.Error())
		}
	}

	return err
}

func (s *Server) process(content [512]byte, emiter net.Addr) error {

	log.Printf("MSG content: %v\n", content)
	log.Printf("From: %v\n\n\n", emiter)
	s.debugResponse(emiter)
	return nil
}

func (s *Server) debugResponse(dest net.Addr) {
	daytime := time.Now().String()
	s.Conn.WriteTo([]byte(daytime), dest)
}
