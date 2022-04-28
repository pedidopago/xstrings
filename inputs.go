package xstrings

import (
	"strconv"
	"strings"
	"unicode"

	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

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

// Length returns the string length.
func Length(s string) int {
	return len([]rune(s))
}

// NormalizeIntlPhoneNumberStr removes all non-numeric characters from a string.
//
//		NormalizeIntlPhoneNumberStr("+1-234-567-8901", true) // "+12345678901"
//		NormalizeIntlPhoneNumberStr("+1-234-567-8901", false) // "12345678901"
//		NormalizeIntlPhoneNumberStr("1-234-567-8901", true) // "+12345678901"
//		NormalizeIntlPhoneNumberStr("1-234-567-8901", false) // "12345678901"
func NormalizeIntlPhoneNumberStr(s string, plusSign bool) string {
	num := NormalizeNumericStr(s)
	if plusSign {
		return "+" + num
	}
	return num
}

// NormalizeNumericStr removes all non-numeric characters from a string.
//
//		NormalizeNumericStr("+1-234-567-8901") // "12345678901"
func NormalizeNumericStr(s string) string {
	return strings.Map(func(r rune) rune {
		if r >= '0' && r <= '9' {
			return r
		}
		return -1
	}, s)
}

// RemoveDiacritics translates é á Í into e a I
func RemoveDiacritics(v string) string {
	t := transform.Chain(norm.NFD, runes.Remove(containsRuneFunc(isMn)), norm.NFC)
	result, _, _ := transform.String(t, v)
	return result
}

// SplitIntlPhoneNumber tries to split a E.164 international phone number into
// country code, area code and number.
//
// If it fails, the returned country and area codes will be empty.
//
// Works well with (+)55___________ (Brasil).
func SplitIntlPhoneNumber(s string) (countryCode, areaCode, number string) {
	intlnum := NormalizeNumericStr(s)
	if len(intlnum) < 7 || len(intlnum) > 15 {
		// https://en.wikipedia.org/wiki/E.164
		return "", "", intlnum
	}
	if intlnum[:2] == "55" {
		// Brasil
		return phoneSplit(intlnum, 2, 2)
	}
	if intlnum[0] == '1' {
		// USA, Canada etc (North American Numbering Plan countries and territories)
		return phoneSplit(intlnum, 1, 3)
	}
	if intlnum[:2] == "20" {
		// Egypt
		return phoneSplit(intlnum, 2, 2)
	}
	if intlnum[:3] == "211" {
		// South Sudan
		return phoneSplit(intlnum, 3, 2)
	}
	if intlnum[:3] == "212" {
		// Morocco + Western Sahara
		return phoneSplit(intlnum, 3, 2)
	}
	if intlnum[:3] == "213" {
		// Algeria
		return phoneSplit(intlnum, 3, 2)
	}

	// ...
	//TODO: fill more popular codes from https://en.wikipedia.org/wiki/List_of_country_calling_codes
	//

	// Zones 3–4: Europe

	if intlnum[:2] == "30" {
		// Greece
		return phoneSplit(intlnum, 2, 3)
	}
	if intlnum[:2] == "31" {
		// Netherlands
		//TODO: 3 or 4 area code digits depending on the region starting digits
		// https://en.wikipedia.org/wiki/Telephone_numbers_in_the_Netherlands
		return phoneSplit(intlnum, 2, 3)
	}
	if intlnum[:2] == "32" {
		// Belgium
		//TODO: belgium area codes
		return phoneSplit(intlnum, 2, 2)
	}
	if intlnum[:2] == "33" {
		// France
		//TODO: France area codes
		return phoneSplit(intlnum, 2, 2)
	}
	if intlnum[:2] == "34" {
		// Spain
		//TODO: Spain area codes
		return phoneSplit(intlnum, 2, 2)
	}
	if MatchNumericPrefix(intlnum, 350, 359) {
		// Gibraltar
		// Portugal +
		// Luxembourg
		// ...
		//TODO: all these area codes
		return phoneSplit(intlnum, 3, 2)
	}
	if MatchNumericPrefix(intlnum, 370, 379) {
		return phoneSplit(intlnum, 3, 2)
	}
	if MatchNumericPrefix(intlnum, 36, 37) {
		return phoneSplit(intlnum, 2, 2)
	}
	if MatchNumericPrefix(intlnum, 380, 389) {
		return phoneSplit(intlnum, 3, 2)
	}
	if MatchNumericPrefix(intlnum, 38, 39) {
		return phoneSplit(intlnum, 2, 2)
	}
	if MatchNumericPrefix(intlnum, 420, 429) {
		return phoneSplit(intlnum, 3, 2)
	}
	if MatchNumericPrefix(intlnum, 42, 49) {
		return phoneSplit(intlnum, 2, 2)
	}

	// FIXME: more

	return phoneSplit(intlnum, 2, 2)
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

// FormatPhoneNumber removes all non-numeric characters from a string.
//
//		FormatPhoneNumber("+1-234-567-8901") // "+12345678901"
//
// Deprecated: use `NormalizeIntlPhoneNumberStr` instead
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
//
// Deprecated: use `NormalizeNumericStr` instead
func FormatNumeric(s string) string {
	return NormalizeNumericStr(s)
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

type containsRuneFunc func(rune) bool

func (f containsRuneFunc) Contains(r rune) bool {
	return f(r)
}

func isMn(r rune) bool {
	return unicode.Is(unicode.Mn, r) // Mn: nonspacing marks
}

// MatchNumericPrefix returns true if the string starts with the given prefix range.
//
// MatchNumericPrefix("123456789", 10, 19) // true for "103456789" ...
func MatchNumericPrefix(pnum string, start, end int) bool {
	for i := start; i <= end; i++ {
		pfx := strconv.Itoa(i)
		if strings.HasPrefix(pnum, pfx) {
			return true
		}
	}
	return false
}

func phoneSplit(pnum string, ccd, acd int) (countryCode, areaCode, number string) {
	if len(pnum) < ccd+acd {
		return "", "", pnum
	}
	return pnum[:ccd], pnum[ccd : ccd+acd], pnum[ccd+acd:]
}
