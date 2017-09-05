package main

import (
	"fmt"
	"net"
	"os"
	"time"

	"github.com/jmaeso/go-game-server/app"
	"github.com/jmaeso/go-game-server/tools"
)

func main() {
	flags, err := tools.LoadFlags()
	if err != nil {
		fmt.Printf("could not start program. err: %s\n", err.Error())
		os.Exit(1)
	}

	// TODO: Load config depending on flag
	fmt.Println(flags)

	var server app.Server
	if err := server.Init(":1200"); err != nil {
		fmt.Printf("could not init server. err: %s\n", err.Error())
	}

	for {
		testhandleClient(server.Conn)
	}

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
