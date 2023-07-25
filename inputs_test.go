package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
	"github.com/stretchr/testify/require"
)

func TestNormalizeIntlPhoneNumberStr(t *testing.T) {
	assert.Equal(t, "+5511900000000", xstrings.NormalizeIntlPhoneNumberStr("+55 11 90000-0000", true))
	assert.Equal(t, "5511900000000", xstrings.NormalizeIntlPhoneNumberStr("+55 11 90000-0000", false))
	assert.Equal(t, "11900000000", xstrings.NormalizeIntlPhoneNumberStr(" (11) 90000-0000🏄", false))
}

func TestNormalizeNumericStr(t *testing.T) {
	assert.Equal(t, "11234567890", xstrings.NormalizeNumericStr("+1(12)34567890"))
	assert.Equal(t, "", xstrings.NormalizeNumericStr("👀"))
	assert.Equal(t, "123", xstrings.NormalizeNumericStr("1a2s3d😎"))
}

func TestNormalizeAplhaNumericStr(t *testing.T) {
	assert.Equal(t, "112A34567890", xstrings.NormalizeAphanumericStr("+1(12A)34567890"))
	assert.Equal(t, "XYZ", xstrings.NormalizeAphanumericStr("👀xyz"))
	assert.Equal(t, "1A2S3D", xstrings.NormalizeAphanumericStr("1a2s3d😎"))
}

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

func TestWhitelist(t *testing.T) {
	require.Equal(t, "12", xstrings.Whitelist("123", "12"))
}

func TestMatchNumericPrefix(t *testing.T) {
	require.True(t, xstrings.MatchNumericPrefix("5511999999999", 50, 55))
	require.False(t, xstrings.MatchNumericPrefix("5511999999999", 49, 54))
}

func TestSplitIntlPhoneNumber(t *testing.T) {
	const brphone = "+5511989999999"
	cc, ac, num := xstrings.SplitIntlPhoneNumber(brphone)
	require.Equal(t, "55", cc)
	require.Equal(t, "11", ac)
	require.Equal(t, "989999999", num)
	const usphone = "+1 313 423 5234"
	cc, ac, num = xstrings.SplitIntlPhoneNumber(usphone)
	require.Equal(t, "1", cc)
	require.Equal(t, "313", ac)
	require.Equal(t, "4235234", num)
}
