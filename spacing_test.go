package xstrings

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestTrimLeadingSpaces(t *testing.T) {
	require.Equal(t, "a b c   ", TrimLeadingSpaces("   a b c   "))
	require.Equal(t, "a b c  ", TrimLeadingSpaces("a b c  "))
	require.Equal(t, "a b c", TrimLeadingSpaces("    a b c"))
	require.Equal(t, "a b c", TrimLeadingSpaces("a b c"))
}

func TestTrimTrailingSpaces(t *testing.T) {
	require.Equal(t, "   a b c", TrimTrailingSpaces("   a b c   "))
	require.Equal(t, "a b c", TrimTrailingSpaces("a b c  "))
	require.Equal(t, "    a b c", TrimTrailingSpaces("    a b c"))
	require.Equal(t, "a b c", TrimTrailingSpaces("a b c"))
}

func TestTrimLeadingAndTrailingSpaces(t *testing.T) {
	require.Equal(t, "a b c", TrimLeadingAndTrailingSpaces("   a b c   "))
	require.Equal(t, "a b c", TrimLeadingAndTrailingSpaces("a b c  "))
	require.Equal(t, "a b c", TrimLeadingAndTrailingSpaces("    a b c"))
	require.Equal(t, "a b c", TrimLeadingAndTrailingSpaces("a b c"))
}
