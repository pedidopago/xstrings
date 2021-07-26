package xstrings_test

import (
	"testing"

	"github.com/pedidopago/xstrings"
	"github.com/stretchr/testify/assert"
)

func TestClean(t *testing.T) {
	assert.Equal(t, "hello!", xstrings.Clean("hel\vlo!"))
	assert.Equal(t, "hello! 2 \nz", xstrings.Clean("hel\vlo! 2\t\nz"))
	assert.Equal(t, "", xstrings.Clean("\x00"))
	assert.Equal(t, "", xstrings.Clean(""))
}
