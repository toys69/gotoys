package slicex

import (
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsAny(t *testing.T) {
	require.True(t, ContainsAny(S(1, 2, 3), S(2, 3, 4)))
	require.False(t, ContainsAny(S(1, 2, 3), S(4, 5, 6)))
	require.False(t, ContainsAny(S(1, 2, 3), nil))
}
