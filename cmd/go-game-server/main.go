package main

import (
	"fmt"
	"os"

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

	var server game.UDPServer
	if err := server.Init(settings.Server); err != nil {
		fmt.Printf("could not init server. err: %s\n", err.Error())
	}

	if err = server.Handle(); err != nil {
		fmt.Printf("could not handle message. err: %s\n", err.Error())
	}
}
