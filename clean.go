package xstrings

import (
	"strings"
	"unicode"
)

// Clean removes control characters and other non printable characters
// from a string (except \n).
// It also replaces \t with a single space.
func Clean(s string) string {
	if s == "" {
		return ""
	}
	return strings.Map(func(r rune) rune {
		if unicode.IsGraphic(r) {
			return r
		}
		if r == '\n' {
			return r
		}
		if r == '\t' {
			return ' '
		}
		return -1
	}, s)
}

// CleanPtr is the same as Clean but accepts a pointer to a string.
// If s is nil, it doesn't do anything.
func CleanPtr(s *string) *string {
	if s == nil {
		return nil
	}
	v := Clean(*s)
	return &v
}
