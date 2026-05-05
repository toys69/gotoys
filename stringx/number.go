package stringx

import (
	"strconv"
	"unsafe"

	"github.com/toys69/gotoys/constraints"
	"github.com/toys69/gotoys/rs"
)

func IToA[I constraints.Integer](i I) string {
	if i > 0 {
		return strconv.FormatUint(uint64(i), 10)
	}
	return strconv.FormatInt(int64(i), 10)
}

func AToI[I constraints.Signed](s string) (ret rs.Result[I]) {
	return ParseInt[I](s, 10)
}

func AToU[I constraints.Unsigned](s string) (ret rs.Result[I]) {
	return ParseUint[I](s, 10)
}

func bitSizeOf[I constraints.Integer]() int {
	return int(unsafe.Sizeof(zero[I]())) << 3
}
func zero[T any]() (ret T) { return ret }

func ParseInt[I constraints.Signed](s string, base int) rs.Result[I] {
	i, err := strconv.ParseInt(s, base, bitSizeOf[I]())
	return rs.Of(I(i), err)
}

func ParseUint[I constraints.Unsigned](s string, base int) rs.Result[I] {
	i, err := strconv.ParseUint(s, base, bitSizeOf[I]())
	return rs.Of(I(i), err)
}

func FormatInt[I constraints.Integer](i I, base int) string {
	if i > 0 {
		return strconv.FormatUint(uint64(i), base)
	}
	return strconv.FormatInt(int64(i), base)
}

func FormatUint[I constraints.Unsigned](i I, base int) string {
	return strconv.FormatUint(uint64(i), base)
}
