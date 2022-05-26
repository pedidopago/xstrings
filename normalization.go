package xstrings

import (
	"regexp"

	"github.com/avito-tech/normalize"
)

func NormalizeForAddress(s string) string {
	v := normalize.Normalize(s, withRemoveSpecialChars(),
		normalize.WithFixRareCyrillicChars(),
		normalize.WithCyrillicToLatinLookAlike(),
		normalize.WithUmlautToLatinLookAlike())
	v = leadClosingWhitespacePattern.ReplaceAllString(v, "")
	v = insideWhitespacePattern.ReplaceAllString(v, " ")
	return v
}

var (
	specialCharsPattern          = regexp.MustCompile(`(?i:[^äöüa-zа-яё0-9\s])`)
	leadClosingWhitespacePattern = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	insideWhitespacePattern      = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

// withRemoveSpecialChars any char except latin/cyrillic letters, German umlauts (`ä`, `ö`, `ü`) and digits are removed
func withRemoveSpecialChars() normalize.Option {
	return func(str string) string {
		return specialCharsPattern.ReplaceAllString(str, "")
	}
}
