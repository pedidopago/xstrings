package xstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskStepped(t *testing.T) {
	assert.Equal(t, "Lui* ** Sil**", maskStepped("Luis da Silva", 3, 3))
	assert.Equal(t, "", maskStepped("", 3, 3))
	assert.Equal(t, "Kev** *aco* **ger*", maskStepped("Kevin Bacon Rogers", 3, 3))
	assert.Equal(t, "Le*** Je***ns", maskStepped("Leroy Jenkins", 2, 3))
}

func TestMaskDocument(t *testing.T) {
	assert.Equal(t, "***123545**", MaskDocument("32312354564"))
	assert.Equal(t, "***.123.545-**", MaskDocument("323.123.5 4 5 - 64"))
	assert.Equal(t, "***.463.432-**", MaskDocument("7A3.463.432-2Z"))
	assert.Equal(t, "**", MaskDocument("12"))
	assert.Equal(t, "", MaskDocument(""))
}

func TestMaskAddressNumber(t *testing.T) {
	assert.Equal(t, "***4", MaskAddressNumber("1234"))
	assert.Equal(t, "**3", MaskAddressNumber("123"))
	assert.Equal(t, "", MaskAddressNumber(""))
}

func TestMaskPrefix(t *testing.T) {
	assert.Equal(t, "***", MaskPrefix("123", 0))
	assert.Equal(t, "", MaskPrefix("", 0))
	assert.Equal(t, "*****ade", MaskPrefix("Lemonade", 3))
	assert.Equal(t, "*****chet", MaskPrefix("Trebuchet", 4))
}

func TestMaskSuffix(t *testing.T) {
	assert.Equal(t, "***", MaskSuffix("123", 0))
	assert.Equal(t, "", MaskSuffix("", 0))
	assert.Equal(t, "Lem*****", MaskSuffix("Lemonade", 3))
	assert.Equal(t, "Treb*****", MaskSuffix("Trebuchet", 4))
}
