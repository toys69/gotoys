package stringx

import (
	"strconv"

	"github.com/toys69/gotoys/rs"
)

func ParseBool(s string) rs.Result[bool] {
	return rs.Of(strconv.ParseBool(s))
}

func FormatBool(b bool) string {
	return strconv.FormatBool(b)
}

func IsTrue(s string) bool {
	v, err := strconv.ParseBool(s)
	return v && err == nil
}

func IsFalse(s string) bool {
	v, err := strconv.ParseBool(s)
	return !v && err == nil
}
