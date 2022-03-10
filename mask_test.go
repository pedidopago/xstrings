package xstrings

import (
	"testing"

	"github.com/stretchr/testify/assert"
)

func TestMaskStepped(t *testing.T) {
	assert.Equal(t, "Lui* ** Sil**", maskStepped("Luis da Silva", 3, 3))
	assert.Equal(t, "Kev** *aco* **ger*", maskStepped("Kevin Bacon Rogers", 3, 3))
	assert.Equal(t, "Le*** Je***ns", maskStepped("Leroy Jenkins", 2, 3))
}

func TestMaskDocument(t *testing.T) {
	assert.Equal(t, "***123545**", MaskDocument("32312354564"))
	assert.Equal(t, "***.123.545-**", MaskDocument("323.123.5 4 5 - 64"))
	assert.Equal(t, "***.463.432-**", MaskDocument("7A3.463.432-2Z"))
}
