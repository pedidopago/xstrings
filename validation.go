package xstrings

import (
	"net/mail"
)

// IsValidCNPJ returns true if the input is a valid CNPJ.
func IsValidCNPJ(cnpj string) bool {
	return validateCNPJ(cnpj) != ""
}

// IsValidCPF returns true if the input is a valid cpf
func IsValidCPF(cpf string) bool {
	return validateCPF(cpf) != ""
}

// IsValidEmail returns true if the email appears to be valid. Uses net/mail
// internally.
func IsValidEmail(email string) bool {
	_, err := mail.ParseAddress(email)
	return err == nil
}

// IsValidIntlMobilePhoneNumber returns true if the input is a valid
// international mobile phone number.
func IsValidIntlMobilePhoneNumber(phone string) bool {
	return isValidInternationalPhoneNumber(phone, true)
}

// IsValidIntlPhoneNumber returns true if the input is a valid international
// phone number.
func IsValidIntlPhoneNumber(phone string) bool {
	return isValidInternationalPhoneNumber(phone, false)
}

// ValidCPF returns a string containing the valid cpf with numbers only if the input is valid and an empty string otherwise
//
// Deprecated: use IsValidCPF instead
func ValidCPF(cpf string) (validCpf string) {
	return validateCPF(cpf)
}

// ValidCNPJ returns a string containing the valid cnpj with numbers only if the input is valid and an empty string otherwise
//
// Deprecated: use IsValidCNPJ instead
func ValidCNPJ(cnpj string) (validCnpj string) {
	return validateCNPJ(cnpj)
}

func isValidInternationalPhoneNumber(phone string, mustBeMobile bool) bool {
	fullnum := FormatNumeric(phone)
	if len(fullnum) < 7 || len(fullnum) > 15 {
		// https://en.wikipedia.org/wiki/E.164
		return false
	}
	if fullnum[:2] == "55" {
		if mustBeMobile {
			return len(fullnum) == 13
		}
		return len(fullnum) >= 12 && len(fullnum) <= 13
	}
	return true
}

func validateCPF(cpf string) (validCpf string) {
	v := FormatNumeric(cpf)
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

func validateCNPJ(cnpj string) (validCnpj string) {
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
