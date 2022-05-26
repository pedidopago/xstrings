package xstrings

import (
	"regexp"
	"strings"

	"github.com/avito-tech/normalize"
)

var addressLowercases = map[string]bool{
	"da":  true,
	"das": true,
	"do":  true,
	"dos": true,
	"de":  true,
}

func NormalizeForAddress(s string) string {
	v := normalize.Normalize(s, withRemoveSpecialChars(),
		normalize.WithFixRareCyrillicChars(),
		normalize.WithCyrillicToLatinLookAlike(),
		normalize.WithUmlautToLatinLookAlike())
	v = leadClosingWhitespacePattern.ReplaceAllString(v, "")
	v = insideWhitespacePattern.ReplaceAllString(v, " ")
	vsplit := strings.Split(v, " ")
	vb := new(strings.Builder)
	for i, s := range vsplit {
		if i > 0 {
			vb.WriteString(" ")
		}
		if len(s) > 1 {
			if addressLowercases[strings.ToLower(s)] {
				vb.WriteString(strings.ToLower(s))
			} else {
				srunes := []rune(s)
				vb.WriteString(strings.ToUpper(string(srunes[0])))
				vb.WriteString(strings.ToLower(string(srunes[1:])))
			}
		} else {
			vb.WriteString(strings.ToLower(s))
		}
	}
	return vb.String()
}

var (
	specialCharsPattern          = regexp.MustCompile(`(?i:[^äãõéáíóñöüa-zа-яё0-9\s])`)
	leadClosingWhitespacePattern = regexp.MustCompile(`^[\s\p{Zs}]+|[\s\p{Zs}]+$`)
	insideWhitespacePattern      = regexp.MustCompile(`[\s\p{Zs}]{2,}`)
)

// withRemoveSpecialChars any char except latin/cyrillic letters, German umlauts (`ä`, `ö`, `ü`) and digits are removed
func withRemoveSpecialChars() normalize.Option {
	return func(str string) string {
		return specialCharsPattern.ReplaceAllString(str, "")
	}
}
