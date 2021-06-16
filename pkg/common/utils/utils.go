package utils

import (
	"strings"
)

func IsValidString(s string) bool {
	return len(strings.TrimSpace(s)) > 0
}

func IsValidStringPointer(s *string) bool {

	if s == nil {
		return false
	}

	return IsValidString(*s)
}

func StrToPr(s string) *string {

	return &s

}
