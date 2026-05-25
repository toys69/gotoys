package xconv

import (
	"errors"
	"testing"

	"github.com/stretchr/testify/require"
	"github.com/toys69/gotoys/rs"
)

// ---- 测试辅助类型 ----

type Greeter interface {
	Greet() string
}

// Person 值接收者实现 Greeter
type Person struct{ Name string }

func (p Person) Greet() string { return "Hello, " + p.Name }

// Robot 指针接收者实现 Greeter（值类型不实现）
type Robot struct{ ID string }

func (r *Robot) Greet() string { return "Beep, " + r.ID }

// Alien 不实现 Greeter
type Alien struct{ Planet string }

// ---- 测试用例 ----

func TestGet_NilObj(t *testing.T) {
	getter := func(p Person) string { return p.Name }
	result := Extract[Person, string](nil, getter)
	require.False(t, result.OK())
	require.True(t, errors.Is(result.Err(), rs.ErrUnset))
}

func TestGet_ConcreteTypeMatch(t *testing.T) {
	p := Person{Name: "Alice"}
	getter := func(pp Person) string { return pp.Name }
	result := Extract[Person, string](p, getter)
	require.True(t, result.OK())
	require.Equal(t, "Alice", result.Val())
}

func TestGet_InterfaceMatch_ValueReceiver(t *testing.T) {
	// Person 值接收者实现 Greeter，直接走快路径
	p := Person{Name: "Bob"}
	getter := func(g Greeter) string { return g.Greet() }
	result := Extract[Greeter, string](p, getter)
	require.True(t, result.OK())
	require.Equal(t, "Hello, Bob", result.Val())
}

func TestGet_InterfaceMatch_PointerReceiver(t *testing.T) {
	// Robot 值类型不实现 Greeter，但 *Robot 实现了，走反射路径
	r := Robot{ID: "R2D2"}
	getter := func(g Greeter) string { return g.Greet() }
	result := Extract[Greeter, string](r, getter)
	require.True(t, result.OK())
	require.Equal(t, "Beep, R2D2", result.Val())
}

func TestGet_InterfaceMatch_AlreadyPointer(t *testing.T) {
	// 传入已经是 *Robot，直接走快路径
	r := &Robot{ID: "C3PO"}
	getter := func(g Greeter) string { return g.Greet() }
	result := Extract[Greeter, string](r, getter)
	require.True(t, result.OK())
	require.Equal(t, "Beep, C3PO", result.Val())
}

func TestGet_NoMatch(t *testing.T) {
	// Alien 既不实现 Greeter，*Alien 也不实现
	a := Alien{Planet: "Mars"}
	getter := func(g Greeter) string { return g.Greet() }
	result := Extract[Greeter, string](a, getter)
	require.False(t, result.OK())
	require.True(t, errors.Is(result.Err(), rs.ErrUnset))
}

func TestGet_NilInterface(t *testing.T) {
	// nil 接口值传入 any 后 != nil，但其动态类型为空，反射无法处理 → None
	var g Greeter
	getter := func(gg Greeter) string { return gg.Greet() }
	result := Extract[Greeter, string](g, getter)
	require.False(t, result.OK())
	require.True(t, errors.Is(result.Err(), rs.ErrUnset))
}

func TestGet_InterfaceMatch_PointerReceiver_TableDriven(t *testing.T) {
	tests := []struct {
		name   string
		obj    any
		getter func(Greeter) string
		wantOK bool
		want   string
	}{
		{
			name:   "nil",
			obj:    nil,
			getter: func(g Greeter) string { return g.Greet() },
			wantOK: false,
			want:   "",
		},
		{
			name:   "Person value receiver",
			obj:    Person{Name: "Dave"},
			getter: func(g Greeter) string { return g.Greet() },
			wantOK: true,
			want:   "Hello, Dave",
		},
		{
			name:   "Robot pointer receiver - value passed",
			obj:    Robot{ID: "X99"},
			getter: func(g Greeter) string { return g.Greet() },
			wantOK: true,
			want:   "Beep, X99",
		},
		{
			name:   "Robot pointer receiver - pointer passed",
			obj:    &Robot{ID: "Y77"},
			getter: func(g Greeter) string { return g.Greet() },
			wantOK: true,
			want:   "Beep, Y77",
		},
		{
			name:   "Alien no match",
			obj:    Alien{Planet: "Venus"},
			getter: func(g Greeter) string { return g.Greet() },
			wantOK: false,
			want:   "",
		},
	}
	for _, tt := range tests {
		t.Run(tt.name, func(t *testing.T) {
			result := Extract[Greeter, string](tt.obj, tt.getter)
			require.Equal(t, tt.wantOK, result.OK())
			if tt.wantOK {
				require.Equal(t, tt.want, result.Val())
			}
		})
	}
}

func TestGet_PreservesOriginalValue(t *testing.T) {
	// 验证反射路径不会修改原始值
	r := Robot{ID: "Original"}
	getter := func(g Greeter) string { return g.Greet() }
	_ = Extract[Greeter, string](r, getter)
	require.Equal(t, "Original", r.ID)
}
