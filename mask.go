package xstrings

import "strings"

// MaskEmail masks something@gmail.com into s********@gmail.com
func MaskEmail(v string) string {
	if v == "" {
		return ""
	}
	if len([]rune(v)) == 1 {
		return v
	}
	at := strings.Index(v, "@")
	if at == -1 {
		vr := []rune(v)
		return string(vr[0]) + strings.Repeat("*", Length(v)-1)
	}
	vr := []rune(v)
	return string(vr[0]) + strings.Repeat("*", Length(v[:at])-1) + v[at:]
}

// MaskUsername masks username into use*****
func MaskUsername(v string) string {
	if v == "" {
		return ""
	}
	const hintLength = 3
	vr := []rune(v)
	if len(vr) <= hintLength {
		return v
	}
	return string(vr[:hintLength]) + strings.Repeat("*", len(vr)-hintLength)
}

// MaskStepped masks "John Doe da Silva" into "Jo** **e d* ***va"
func MaskStepped(v string) string {
	return maskStepped(v, 2, 4)
}

// MaskDocument is used to maks the first and last digits of a document. It ignores
// punctuation and spaces.
//
// Example:
//
// 	MaskDocument("232.312.223-77") // "***.312.223-**"
func MaskDocument(v string) string {
	if v == "" {
		return ""
	}
	if Length(v) < 6 {
		return Mask(v)
	}
	first := 3
	last := 2
	ln := Length(Blacklist(v, " .-/")) - last
	n := 0
	var result strings.Builder
	for _, r := range v {
		switch r {
		case ' ':
			continue
		case '.', '-', '/':
			result.WriteRune(r)
			continue
		}
		if n < first || n >= ln {
			result.WriteRune('*')
		} else {
			result.WriteRune(r)
		}
		n++
	}
	return result.String()
}

//MaskAddressNumber: mask the first numbers of an address and shows the last number.
//
// Example:
//
//	MaskAddressNumber("12345") into ****5
func MaskAddressNumber(v string) string {
	if v == "" {
		return ""
	}
	return maskFirst(v, 1)
}

func Mask(v string) string {
	if strings.Contains(v, "@") && len(v) > 4 {
		x := strings.IndexRune(v, '@')
		return v[0:1] + maskFirst(v[1:x], 3) + "@" + v[x+1:x+2] + maskFirst(v[x+2:], 3)
	}
	return maskFirst(v, 4)
}

func maskFirst(v string, minshow int) string {
	if len([]rune(v)) <= minshow {
		return strings.Repeat("*", len([]rune(v)))
	}
	cut := len([]rune(v)) - minshow
	return strings.Repeat("*", cut) + v[cut:]
}

func maskStepped(v string, clearStep, maskStep int) string {
	var result strings.Builder
	cc := 0
	for _, r := range v {
		if r == ' ' {
			result.WriteRune(r)
			continue
		}
		cc++
		if cc <= 0 {
			result.WriteRune('*')
			continue
		}
		result.WriteRune(r)
		if cc >= clearStep {
			cc = -maskStep
		}
	}
	return result.String()
}
