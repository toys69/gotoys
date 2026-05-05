package stringx

import (
	"math"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/toys69/gotoys/rs"
)

func TestIToA(t *testing.T) {
	require.Equal(t, "123", IToA(123))
	require.Equal(t, "18446744073709551615", IToA(uint64(math.MaxUint64)))
	require.Equal(t, "9223372036854775807", IToA(math.MaxInt64))
	require.Equal(t, "-9223372036854775808", IToA(math.MinInt64))
}

func TestAToI(t *testing.T) {
	require.Equal(t, rs.OK(int64(123)), AToI[int64]("123"))
	require.Equal(t, rs.OK(int64(9223372036854775807)), AToI[int64]("9223372036854775807"))
	require.Equal(t, rs.OK(int64(-9223372036854775808)), AToI[int64]("-9223372036854775808"))
	require.Equal(t, rs.OK(uint64(18446744073709551615)), AToU[uint64]("18446744073709551615"))

	require.Error(t, AToI[int8]("65535").Err())
	require.Error(t, AToU[uint8]("65535").Err())
	require.Error(t, AToU[uint8]("-65535").Err())

	require.Error(t, AToI[int64]("123.456").Err())
	require.Error(t, AToU[uint64]("123.456").Err())
}
