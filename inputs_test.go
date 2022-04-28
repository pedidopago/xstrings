package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestFormatPhoneNumber(t *testing.T) {
	assert.Equal(t, "+5511900000000", xstrings.FormatPhoneNumber("+55 11 90000-0000"))
	assert.Equal(t, "11900000000", xstrings.FormatPhoneNumber(" (11) 90000-0000🏄"))
}

func TestFormatNumeric(t *testing.T) {
	assert.Equal(t, "11234567890", xstrings.FormatNumeric("+1(12)34567890"))
	assert.Equal(t, "", xstrings.FormatNumeric("👀"))
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
