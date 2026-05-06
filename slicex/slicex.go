package slicex

import (
	stdslices "slices"
)

// S makes a slice from variadic arguments
func S[T comparable](s ...T) []T { return s }

// ContainsAny reports whether any element of a is present in b.
func ContainsAny[S ~[]E, E comparable](a, b S) bool {
	return ContainsFunc(a, func(t E) bool { return Contains(b, t) })
}

// Contains reports whether v is present in s.
func Contains[S ~[]E, E comparable](s S, v E) bool {
	return stdslices.Contains(s, v)
}

// ContainsFunc reports whether at least one
// element e of s satisfies f(e).
func ContainsFunc[S ~[]E, E any](s S, f func(E) bool) bool {
	return stdslices.ContainsFunc(s, f)
}
