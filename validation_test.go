package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/require"
)

func TestIsValidCPF(t *testing.T) {
	const validCpf = "939.511.980-28"
	const validCpfOnlyNumbers = "93951198028"
	const invalidCpf = "123.456.789-10"
	require.True(t, xstrings.IsValidCPF(validCpf))
	require.True(t, xstrings.IsValidCPF(validCpfOnlyNumbers))
	require.False(t, xstrings.IsValidCPF(invalidCpf))
}

func TestIsValidCNPJ(t *testing.T) {
	const validCnpj = "41.143.201/0001-25"
	const validCnpjOnlyNumbers = "41143201000125"
	const invalidCnpj = "12.345.678/0001-90"
	require.True(t, xstrings.IsValidCNPJ(validCnpj))
	require.True(t, xstrings.IsValidCNPJ(validCnpjOnlyNumbers))
	require.False(t, xstrings.IsValidCNPJ(invalidCnpj))
}

func TestIsValidEmail(t *testing.T) {
	const validEmail = "someone@gmail.com"
	const invalidEmail = "someone@.gmail"
	const invalidEmail2 = "someone@gmail."
	const validEmail2 = "some+one.example@gmail.com"
	require.True(t, xstrings.IsValidEmail(validEmail))
	require.False(t, xstrings.IsValidEmail(invalidEmail))
	require.False(t, xstrings.IsValidEmail(invalidEmail2))
	require.True(t, xstrings.IsValidEmail(validEmail2))
}

func TestIsValidIntlMobilePhoneNumber(t *testing.T) {
	const validMobile = "+5511999999999"
	const validMobile2 = "+5511 999999999"
	const validMobile3 = "5511-999999999"
	const validMobile4 = "1 250 555 0199"
	const invalidMobile = "+5511 12345678"
	const invalidMobile2 = "+5511 1234567800"
	const invalidMobile3 = "+1 2 3 4 5 6"
	require.True(t, xstrings.IsValidIntlMobilePhoneNumber(validMobile))
	require.True(t, xstrings.IsValidIntlMobilePhoneNumber(validMobile2))
	require.True(t, xstrings.IsValidIntlMobilePhoneNumber(validMobile3))
	require.True(t, xstrings.IsValidIntlMobilePhoneNumber(validMobile4))
	require.False(t, xstrings.IsValidIntlMobilePhoneNumber(invalidMobile))
	require.False(t, xstrings.IsValidIntlMobilePhoneNumber(invalidMobile2))
	require.False(t, xstrings.IsValidIntlMobilePhoneNumber(invalidMobile3))
}
