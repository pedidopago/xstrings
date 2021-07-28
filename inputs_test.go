package xstrings_test

import (
	"testing"

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
