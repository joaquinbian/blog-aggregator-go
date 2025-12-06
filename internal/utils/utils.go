package utils

import (
	"fmt"
)

func GetNameFromArgs(args []string, name string) (string, error) {
	if len(args) == 0 {
		return "", fmt.Errorf("error: no args provided for the %v command", name)
	} else if len(args) != 1 {
		return "", fmt.Errorf("usage: %s <name>", name)
	}

	user := args[0]

	return user, nil
}
