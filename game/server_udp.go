package game

import (
	"errors"
	"log"
	"net"
	"time"

	"github.com/jmaeso/go-game-server/game/netbit"
)

// UDPServer is the representation for a UDP network connection.
type UDPServer struct {
	Conn *net.UDPConn
}

// Init is a required function previous to use the server.
//
// settings: Has to contain Port attr. Url is optional (localhost by default).
//
// Returns an error if the resultant address is bad or connection can't be established.
func (s *UDPServer) Init(settings *ServerSettings) error {
	url := settings.URL
	port := settings.Port
	if port == "" {
		return errors.New("udp_server_init: port required")
	}

	address := url + ":" + port

	udpAddr, err := net.ResolveUDPAddr("udp4", address)
	if err != nil {
		return errors.New("udp_server_init: could not resolve address. err: " + err.Error())
	}

	s.Conn, err = net.ListenUDP("udp", udpAddr)
	if err != nil {
		return errors.New("udp_server_init: could not start connection. err: " + err.Error())
	}

	return nil
}

// Handle starts to listen and process client's messages.
//
// Returns an error if a message is unprocesable.
func (s *UDPServer) Handle() error {
	var err error

	for err == nil {
		var buff [512]byte
		_, addr, err := s.Conn.ReadFrom(buff[0:])
		if err != nil {
			err = errors.New("udp_server_handle: could not read message. err: " + err.Error())
		}

		if err = s.process(buff, addr); err != nil {
			err = errors.New("udp_server_handle: could not process message. err: " + err.Error())
		}
	}

	return err
}

func (s *UDPServer) process(content [512]byte, emiter net.Addr) error {

	var userMsg netbit.Packet

	if err := netbit.Decode(content, &userMsg); err != nil {
		return err
	}

	log.Printf("MSG content: %v\n", userMsg)
	log.Printf("From: %v\n\n\n", emiter)
	s.debugResponse(emiter)
	return nil
}

func (s *UDPServer) debugResponse(dest net.Addr) {
	daytime := time.Now().String()
	s.Conn.WriteTo([]byte(daytime), dest)
}
