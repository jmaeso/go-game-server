package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/jmaeso/go-game-server/game"
	"github.com/jmaeso/go-game-server/tools/flags"
	"github.com/jmaeso/go-game-server/tools/yaml"
)

func main() {
	flags, err := flags.Load()
	if err != nil {
		fmt.Printf("could not start program. err: %s\n", err.Error())
		os.Exit(1)
	}

	var settings game.Settings
	if err = yaml.Load("config/"+flags["environment"]+".yml", &settings); err != nil {
		fmt.Printf("could not load config. err: %s\n", err.Error())
	}

	var server game.Server
	if err := server.Init(settings.Server); err != nil {
		fmt.Printf("could not init server. err: %s\n", err.Error())
	}

	if err = server.Handle(); err != nil {
		fmt.Printf("could not handle message. err: %s\n", err.Error())
	}
	// for {
	// 	testhandleClient(server.Conn)
	// }

}

func testhandleClient(conn *net.UDPConn) {
	var buff [512]byte
	_, addr, err := conn.ReadFromUDP(buff[0:])
	if err != nil {
		return
	}

	daytime := time.Now().String()
	conn.WriteToUDP([]byte(daytime), addr)
}
