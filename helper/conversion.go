package helper

import (
	"regexp"
	"strings"
)

func FormatPhoneNumber(phoneNumber string) string {
	phoneNumber = regexp.MustCompile(`\D`).ReplaceAllString(phoneNumber, "")

	if strings.HasPrefix(phoneNumber, "0") {
		return strings.Replace(phoneNumber, "0", "+62", 1)
	}

	return phoneNumber
}
