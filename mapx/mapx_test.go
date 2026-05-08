package mapx

import (
	"testing"
)

func TestMakeIfNil(t *testing.T) {
	type M1 map[int]int

	t.Logf("%#v", MakeIfNil[M1](nil))
	t.Logf("%#v", MakeIfNil[M1](nil, 10))
	t.Logf("%#v", MakeIfNil(M1{}))
	t.Logf("%#v", MakeIfNil(M1{1: 1}))
	t.Logf("%#v", CloneOrMake[M1](nil))
}
