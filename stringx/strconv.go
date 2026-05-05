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
	i, err := strconv.ParseInt(s, 10, bitSizeOf[I]())
	return rs.Of(I(i), err)
}

func AToU[I constraints.Unsigned](s string) (ret rs.Result[I]) {
	i, err := strconv.ParseUint(s, 10, bitSizeOf[I]())
	return rs.Of(I(i), err)
}

func bitSizeOf[I constraints.Integer]() int {
	return int(unsafe.Sizeof(zero[I]())) << 3
}
func zero[T any]() (ret T) { return ret }
