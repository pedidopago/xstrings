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

// var runeMaps = map[rune]rune{
// 	'ᴀ': 'A', // 0x1d00
// 	'ᴁ': 'A', // 0x1d01
// 	'ᴂ': 'A', // 0x1d02
// 	'ᴃ': 'B', // 0x1d03
// 	'ᴄ': 'C', // 0x1d04
// }

type replRange struct {
	start   rune
	end     rune
	replace func(rune) rune
}

func (rr replRange) contains(r rune) bool {
	return r >= rr.start && r <= rr.end
}

var replacements = []replRange{
	{
		start: 0x0041,
		end:   0x005A,
		replace: func(r rune) rune {
			return r
		},
	},
	{
		start: 0x0061,
		end:   0x007A,
		replace: func(r rune) rune {
			return r
		},
	},
	{
		start: 0x00C0,
		end:   0x00D6,
		replace: func(r rune) rune {
			return r
		},
	},
	{
		start: 0x00D9,
		end:   0x00DD,
		replace: func(r rune) rune {
			return r
		},
	},
	{
		start: 0x00E0,
		end:   0x00FF,
		replace: func(r rune) rune {
			return r
		},
	},
	//TODO: https://www.compart.com/en/unicode/block/U+0100
	{
		// ᴀ ᴁ ᴂ
		start: 0x1d00,
		end:   0x1d02,
		replace: func(r rune) rune {
			return 'A'
		},
	},
	{
		// ᴃ ᴄ ᴅ
		start: 0x1d03,
		end:   0x1d05,
		replace: func(r rune) rune {
			return r - 0x1d03 + 'B'
		},
	},
}

func mapRunesForName(v string) string {
	return strings.Map(func(r rune) rune {
		for _, rr := range replacements {
			if rr.contains(r) {
				return rr.replace(r)
			}
		}
		return ' '
	}, v)
}

func NormalizeForNameExcludingInvalidChars(v string) string {
	t := transform.Chain(norm.NFKD, runes.Remove(containsRuneFunc(isMnOrDingbats)), norm.NFC)
	// t := transform.Chain(norm.NFKD, transform.RemoveFunc(isMn), norm.NFC)
	v, _, _ = transform.String(t, v)
	vf := strings.TrimSpace(NormalizeForName(RemoveEmojis(v)))
	vf = mapRunesForName(vf)
	vf = strings.Join(strings.Fields(vf), " ")
	return vf
}
