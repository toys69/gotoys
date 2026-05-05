package stringx

import "strings"

// ContainsInSeparated checks if the substring sub exists in the separator-separated string s.
// For example: ContainsInSeparated("a,b,c", "b", ",") returns true.
func ContainsInSeparated(s, sub, sep string) bool {
	i := strings.Index(s, sub)
	if i < 0 {
		return false
	}

	for x := range strings.SplitSeq(s[i:], sep) {
		if x == sub {
			return true
		}
	}

	return false
}
