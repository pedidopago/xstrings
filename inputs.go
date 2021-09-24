package xstrings

import (
	"bytes"
	"golang.org/x/text/runes"
	"golang.org/x/text/transform"
	"golang.org/x/text/unicode/norm"
	"strings"
	"unicode"
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

// NumbersOnly accepts 0-9 range only
func NumbersOnly(str string) string {
	var b bytes.Buffer
	for _, r := range str {
		switch r {
		case '0', '1', '2', '3', '4', '5', '6', '7', '8', '9':
			b.WriteRune(r)
		}
	}
	return b.String()
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

// ValidCPF returns a string containing the valid cpf with numbers only if the input is valid and an empty string otherwise
func ValidCPF(cpf string) (validCpf string) {
	v := NumbersOnly(cpf)
	if len(v) != 11 {
		return
	}
	same := true
	for i := 0; i < len(v)-1; i++ {
		if v[i] != v[i+1] {
			same = false
			break
		}
	}
	if same {
		return
	}
	sum := 0
	numeros := v[:9]
	digitos := v[9:]
	for i := 10; i > 1; i-- {
		sum += runeV(rune(numeros[10-i])) * i
	}
	var result int
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	if result != runeV(rune(digitos[0])) {
		return
	}
	numeros = v[:10]
	sum = 0
	for i := 11; i > 1; i-- {
		sum += runeV(rune(numeros[11-i])) * i
	}
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	if result != runeV(rune(digitos[1])) {
		return
	}
	return v
}

var cnpjDigs = []int{6, 5, 4, 3, 2, 9, 8, 7, 6, 5, 4, 3, 2}

// ValidCNPJ returns a string containing the valid cnpj with numbers only if the input is valid and an empty string otherwise
func ValidCNPJ(cnpj string) (validCnpj string) {
	v := NumbersOnly(cnpj)
	if len(v) != 14 {
		return
	}
	same := true
	for i := 0; i < len(v)-1; i++ {
		if v[i] != v[i+1] {
			same = false
			break
		}
	}
	if same {
		return
	}
	sum := 0
	numeros := v[:12]
	digitos := v[12:]
	for i, numero := range numeros {
		sum += runeV(numero) * cnpjDigs[i+1]
	}
	var result int
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	if result != runeV(rune(digitos[0])) {
		return
	}
	numeros = v[:13]
	sum = 0
	for i, numero := range numeros {
		sum += runeV(numero) * cnpjDigs[i]
	}
	if sum%11 < 2 {
		result = 0
	} else {
		result = 11 - sum%11
	}
	if result != runeV(rune(digitos[1])) {
		return
	}
	return v
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
