package utils

import (
	"fmt"
	"regexp"
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

// ValidateUrl checks if the provided URL is valid using regex.
func ValidateUrl(url string) bool {

	if url == "" {
		return false
	}

	// Regex pattern for validating URLs
	urlPattern := `^https?:\/\/(www\.)?[-a-zA-Z0-9@:%._\+~#=]{1,256}\.[a-zA-Z0-9()]{1,6}\b([-a-zA-Z0-9()@:%_\+.~#?&//=]*)$`

	match, err := regexp.MatchString(urlPattern, url)
	if err != nil {
		return false
	}

	if !match {
		return false
	}

	return true
}
