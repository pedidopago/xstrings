package faker_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/pedidopago/xstrings/faker"
	"github.com/stretchr/testify/require"
)

func TestNewRandomCPF(t *testing.T) {
	for i := 0; i < 100; i++ {
		cpf := faker.NewRandomCPF()
		require.True(t, xstrings.IsValidCPF(cpf))
	}
}
