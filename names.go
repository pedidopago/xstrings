package xstrings

import (
	"strings"

	"github.com/forPelevin/gomoji"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
)

func FirstName(v string) string {
	// remove special characters
	v = Clean(v)
	// remove emojis
	v = gomoji.RemoveEmojis(v)
	v = strings.Replace(v, ",", " ", -1)
	v = strings.Replace(v, "  ", " ", -1)
	v = strings.Replace(v, "  ", " ", -1)
	vs := strings.Split(v, " ")
	return vs[0]
}

func RemoveEmojis(v string) string {
	return gomoji.RemoveEmojis(v)
}

func ContainsEmoji(v string) bool {
	return gomoji.ContainsEmoji(v)
}

func isMnOrDingbats(r rune) bool {
	if isMn(r) {
		return true
	}
	// dingbats
	if r >= 0x2700 && r <= 0x27BF {
		return true
	}
	switch r {
	case '.', ',', ';', ':', '!', '?', '(', ')', '[', ']', '{', '}', '/', '\\', '"', '\'', '\t', '\n', '\r', '\v', '\f', '\a', '\b', '\000':
		return true
	}
	return false
}

func NormalizeForNameExcludingInvalidChars(v string) string {
	t := transform.Chain(norm.NFKD, runes.Remove(containsRuneFunc(isMnOrDingbats)), norm.NFC)
	// t := transform.Chain(norm.NFKD, transform.RemoveFunc(isMn), norm.NFC)
	v, _, _ = transform.String(t, v)
	return strings.TrimSpace(NormalizeForName(RemoveEmojis(v)))
}
