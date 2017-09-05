package flags

import (
	"errors"
	"flag"
)

func Load() (map[string]string, error) {
	flags := make(map[string]string)

	environmentFlag := flag.String("environment", "", "environment of the server.")
	flag.Parse()

	if *environmentFlag == "" {
		return nil, errors.New("environment flag is required")
	}

	flags["environment"] = *environmentFlag

	return flags, nil
}
