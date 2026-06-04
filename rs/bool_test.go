package rs

import (
	"encoding/json"
	"testing"

	"github.com/stretchr/testify/require"
)

func TestBool_UnmarshalText(t *testing.T) {

	type T1 struct {
		A, B, C Bool
	}

	var v1, v2 T1
	v1.A.Store(true)
	v1.B.Store(false)
	bs1, err1 := json.Marshal(v1)
	require.NoError(t, err1)
	t.Log(string(bs1))
	require.Equal(t, `{"A":true,"B":false,"C":null}`, string(bs1))

	err2 := json.Unmarshal(bs1, &v2)
	require.NoError(t, err2)

	require.Equal(t, v1, v2)

}
