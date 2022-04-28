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
