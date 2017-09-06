package flags

import (
	"errors"
	"flag"
)

// Load is a function for reading input flags to the program.
//
// Returns a map[key] = value with the flags or error if a required flag is not provided.
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
