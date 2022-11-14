package xstrings

import (
	"strings"

	"github.com/forPelevin/gomoji"
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
