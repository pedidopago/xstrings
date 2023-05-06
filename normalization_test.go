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
	assert.Equal(t, "São Paulo", xstrings.NormalizeForAddress("São    PaUlo  "))
	assert.Equal(t, "Rua 23 de Maio", xstrings.NormalizeForAddress(" rua 23 DE maio"))
	assert.Equal(t, "", xstrings.NormalizeForAddress(" "))
	assert.Equal(t, "12", xstrings.NormalizeForAddress(" 12 "))
	assert.Equal(t, "01232092", xstrings.NormalizeForAddress(" 01232-092 "))
	assert.Equal(t, "Lúcio Mauro Araújo", xstrings.NormalizeForName("LúcIo  MaUro Araújo  "))
	assert.Equal(t, "Clau", xstrings.NormalizeForNameExcludingInvalidChars("𝒄𝒍𝒂𝒖 ❀"))
	assert.Equal(t, "", xstrings.NormalizeForNameExcludingInvalidChars(" "))
	assert.Equal(t, "", xstrings.NormalizeForNameExcludingInvalidChars(","))
	assert.Equal(t, "Andressa Carvalho", xstrings.NormalizeForNameExcludingInvalidChars("ᴀɴᴅʀᴇꜱꜱᴀ ᴄᴀʀᴠᴀʟʜᴏ"))
}
