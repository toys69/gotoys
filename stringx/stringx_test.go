package stringx

import (
	"cmp"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestContainsInSeparated(t *testing.T) {
	require.True(t, ContainsInSeparated("a", "a", ","))
	require.True(t, ContainsInSeparated("a,b,c", "b", ","))

	require.False(t, ContainsInSeparated("a", "aaaa", ","))
	require.False(t, ContainsInSeparated("a,b,c", "d", ","))
}

func TestSplitAndClean(t *testing.T) {
	tests := []struct {
		name string
		str  string
		sep  string
		want []string
	}{
		{name: "empty", str: "", want: make([]string, 0)},
		{name: "single", str: "a", want: []string{"a"}},
		{name: "multi", str: "a,b,c", want: []string{"a", "b", "c"}},
		{name: "multi-sep", str: "a,,b,c", want: []string{"a", "b", "c"}},
		{name: "multi-sep2", str: "a,,b,c,,d", want: []string{"a", "b", "c", "d"}},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			tt.sep = cmp.Or(tt.sep, ",")
			got := SplitAndClean(tt.str, tt.sep)
			require.EqualValues(t, tt.want, got)
		})
	}
}
