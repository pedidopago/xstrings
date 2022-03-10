package xstrings_test

import (
	"testing"

	"github.com/stretchr/testify/require"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestFormatPhoneNumber(t *testing.T) {
	assert.Equal(t, "+5511900000000", xstrings.FormatPhoneNumber("+55 11 90000-0000"))
	assert.Equal(t, "11900000000", xstrings.FormatPhoneNumber(" (11) 90000-0000🏄"))
}

func TestFormatNumeric(t *testing.T) {
	assert.Equal(t, "11234567890", xstrings.FormatNumeric("+1(12)34567890"))
	assert.Equal(t, "", xstrings.FormatNumeric("👀"))
}

func TestValidCPF(t *testing.T) {
	const validCpf = "939.511.980-28"
	const validCpfOnlyNumbers = "93951198028"
	const invalidCpf = "123.456.789-10"
	v := xstrings.ValidCPF(validCpf)
	require.Equal(t, validCpfOnlyNumbers, v)
	v = xstrings.ValidCPF(invalidCpf)
	require.Len(t, v, 0)
}

func TestValidCNPJ(t *testing.T) {
	const validCnpj = "41.143.201/0001-25"
	const validCnpjOnlyNumbers = "41143201000125"
	const invalidCnpj = "12.345.678/0001-90"
	v := xstrings.ValidCNPJ(validCnpj)
	require.Equal(t, validCnpjOnlyNumbers, v)
	v = xstrings.ValidCNPJ(invalidCnpj)
	require.Len(t, v, 0)
}

func TestRemoveDiacritics(t *testing.T) {
	const withDiacritics = "aáàâäeéèêëiíìîïoóòôöuúùûü"
	const withoutDiacritics = "aaaaaeeeeeiiiiiooooouuuuu"
	v := xstrings.RemoveDiacritics(withDiacritics)
	require.Equal(t, withoutDiacritics, v)
}

func TestBlacklist(t *testing.T) {
	require.Equal(t, "31232211122", xstrings.Blacklist("312.32 2.111-22", " .-"))
}
