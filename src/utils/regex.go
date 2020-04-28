package utils

import (
	"regexp"
)

// IsCountryCode - check if the value may be a country code
func IsCountryCode(arg string) bool {
	var re = regexp.MustCompile("[a-zA-Z]{2,}")

	match := re.FindString(arg)
	if len(arg) != len(match) {
		return false
	}
	return true
}