package utils

import (
	"fmt"
	"regexp"
	"time"
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

func ParsePubDate(s string) (time.Time, error) {
	// Try a few common RSS/Atom formats
	layouts := []string{
		time.RFC1123Z,                   // "Mon, 02 Jan 2006 15:04:05 -0700"
		time.RFC1123,                    // "Mon, 02 Jan 2006 15:04:05 MST"
		time.RFC822Z,                    // "02 Jan 06 15:04 -0700"
		time.RFC822,                     // "02 Jan 06 15:04 MST"
		time.RFC3339,                    // "2006-01-02T15:04:05Z07:00"
		"Mon, 02 Jan 2006 15:04:05 MST", // some feeds use weird variants
	}

	var lastErr error
	for _, layout := range layouts {
		t, err := time.Parse(layout, s)
		if err == nil {
			return t, nil
		}
		lastErr = err
	}
	return time.Time{}, lastErr
}
