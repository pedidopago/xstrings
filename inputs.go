package xstrings

import (
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

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

// NumbersOnly accepts 0-9 range only
//
// Deprecated: use `FormatNumeric` instead
func NumbersOnly(str string) string {
	return FormatNumeric(str)
}

func runeV(r rune) int {
	switch r {
	case '9':
		return 9
	case '8':
		return 8
	case '7':
		return 7
	case '6':
		return 6
	case '5':
		return 5
	case '4':
		return 4
	case '3':
		return 3
	case '2':
		return 2
	case '1':
		return 1
	}
	return 0
}

// Whitelist will only return characters found in "allowed"
func Whitelist(v string, allowed string) string {
	runeray := make(map[rune]bool)
	for _, r := range allowed {
		runeray[r] = true
	}
	runeb := make([]rune, 0, len(v))
	for _, r := range v {
		if runeray[r] {
			runeb = append(runeb, r)
		}
	}
	return string(runeb)
}

// Blacklist will remove certain characters
func Blacklist(v string, cutset string) string {
	runeray := make(map[rune]bool)
	for _, r := range cutset {
		runeray[r] = true
	}
	runeb := make([]rune, 0, len(v))
	for _, r := range v {
		if !runeray[r] {
			runeb = append(runeb, r)
		}
	}
	return string(runeb)
}

type containsRuneFunc func(rune) bool

func (f containsRuneFunc) Contains(r rune) bool {
	return f(r)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// RemoveDiacritics translates é á Í into e a I
func RemoveDiacritics(v string) string {
	t := transform.Chain(norm.NFD, runes.Remove(containsRuneFunc(isMn)), norm.NFC)
	result, _, _ := transform.String(t, v)
	return result
}
