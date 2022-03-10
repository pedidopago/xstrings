package xstrings

import "strings"

// FormatPhoneNumber removes all non-numeric characters from a string.
//
//		FormatPhoneNumber("+1-234-567-8901") // "+12345678901"
func FormatPhoneNumber(s string) string {
	plus := false
	s = strings.TrimSpace(s)
	if strings.HasPrefix(s, "+") {
		plus = true
	}
	s = strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
	if plus {
		return "+" + s
	}
	return s
}

// FormatNumeric removes all non-numeric characters from a string.
//
//		FormatNumeric("+1-234-567-8901") // "12345678901"
func FormatNumeric(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// Length returns the string length.
func Length(s string) int {
	return len([]rune(s))
}
