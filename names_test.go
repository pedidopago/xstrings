package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestNormalizeForNameExcludingInvalidChars(t *testing.T) {
	tests := []struct {
		name     string
		expected string
	}{
		{"ɑɳɗʀɛiɑ", "Andreia"},
		{"Lu@na", "Luana"},
		{"@anna_sibylla", "Anna Sibylla"},
	}

	for i, test := range tests {
		assert.Equal(t, xstrings.NormalizeForNameExcludingInvalidChars(test.name), test.expected, i)
	}
}
