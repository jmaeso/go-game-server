package game

import "net"

// Address is a wrapper for being able to work with net.UDPAddr and net.TCPAddr
type Address struct {
	IP   net.IP
	Port int
	Zone string // IPv6 scoped addressing zone
}
