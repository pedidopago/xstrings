package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestNormalization(t *testing.T) {
	assert.Equal(t, "Rua 12 Apto 123", xstrings.NormalizeForAddress(" Rua   12    Apto   123    "))
	assert.Equal(t, "Rua 123 Apto 1232", xstrings.NormalizeForAddress(" Rua   123,    Apto   1232    "))
	assert.Equal(t, "Rua 123 Apto 1232", xstrings.NormalizeForAddress(" Rua   123.    Apto   1232    "))
}