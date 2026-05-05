package xconv

import (
	"github.com/toys69/gotoys/constraints"
	"github.com/toys69/gotoys/rs"
	"github.com/toys69/gotoys/stringx"
)

func IToA[I constraints.Integer](i I) string {
	return stringx.IToA(i)
}

func AToI[I constraints.Signed](s string) rs.Result[I] {
	return stringx.AToI[I](s)
}

func AToU[I constraints.Unsigned](s string) rs.Result[I] {
	return stringx.AToU[I](s)
}
