package stringx

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsInSeparated(t *testing.T) {
	require.True(t, ContainsInSeparated("a", "a", ","))
	require.True(t, ContainsInSeparated("a,b,c", "b", ","))

	require.False(t, ContainsInSeparated("a", "aaaa", ","))
	require.False(t, ContainsInSeparated("a,b,c", "d", ","))
}
