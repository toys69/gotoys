package stringx

import (
	"strings"
)

// ContainsInSeparated checks if the substring sub exists in the separator-separated string s.
// For example: ContainsInSeparated("a,b,c", "b", ",") returns true.
func ContainsInSeparated(s, sub, sep string) bool {
	for x := range strings.SplitSeq(s, sep) {
		if x == sub {
			return true
		}
	}

	return false
}

func SplitAndClean(s string, sep string) []string {
	ss := strings.Split(s, sep)
	i := 0
	for _, s := range ss {
		s = strings.TrimSpace(s)
		if len(s) == 0 {
			continue
		}

		ss[i] = s
		i++
	}
	return ss[:i:i]
}
