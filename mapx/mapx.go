package mapx

import (
	"cmp"
	"maps"
)

func CloneOrMake[M ~map[K]V, K comparable, V any](m M) M {
	if m == nil {
		return make(M)
	}

	return maps.Clone(m)
}

func MakeIfNil[M ~map[K]V, K comparable, V any](m M, capacity ...int) M {
	if m == nil {
		return make(M, cmp.Or(capacity...))
	}
	return m
}
