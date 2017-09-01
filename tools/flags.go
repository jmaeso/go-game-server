package tools

import (
	"errors"
	"flag"
)

func LoadFlags() (map[string]string, error) {
	//var flags map[string]string
	flags := make(map[string]string)

	environmentFlag := flag.String("environment", "", "environment of the server.")
	flag.Parse()

	if *environmentFlag == "" {
		return nil, errors.New("environment flag is required")
	}

	flags["environment"] = *environmentFlag

	return flags, nil
}
