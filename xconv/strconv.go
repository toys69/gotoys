package xconv

import (
	"github.com/toys69/gotoys/constraints"
	"github.com/toys69/gotoys/rs"
	"github.com/toys69/gotoys/stringx"
)

// IToA ...
// Deprecated: use stringx.IToA
//
//go:fix inline
func IToA[I constraints.Integer](i I) string {
	return stringx.IToA(i)
}

// AToI ...
// Deprecated: use stringx.AToI
//
//go:fix inline
func AToI[I constraints.Signed](s string) rs.Result[I] {
	return stringx.AToI[I](s)
}

// AToU ...
// Deprecated: use stringx.AToU
//
//go:fix inline
func AToU[I constraints.Unsigned](s string) rs.Result[I] {
	return stringx.AToU[I](s)
}
