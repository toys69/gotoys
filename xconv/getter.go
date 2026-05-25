package xconv

import (
	"reflect"

	"github.com/toys69/gotoys/rs"
)

// Extract 尝试将 obj 断言为类型 S，然后调用 getter 提取结果。
// 当 obj 的值类型 T 未实现接口 S，但 *T 实现了 S 时，会自动取地址重试。
func Extract[S any, V any](obj any, getter func(S) V) rs.Result[V] {
	if obj == nil {
		return rs.None[V]()
	}

	// 快速路径：obj 直接满足 S
	if s, ok := obj.(S); ok {
		return rs.OK(getter(s))
	}

	// 慢路径：检查 *T 是否实现了接口 S（值类型 T 未实现，指针类型 *T 实现了）
	i := reflect.TypeFor[S]()
	if i.Kind() == reflect.Interface {
		typ := reflect.TypeOf(obj)
		if typ != nil && reflect.PointerTo(typ).Implements(i) {
			v := reflect.New(typ)
			v.Elem().Set(reflect.ValueOf(obj))
			// 已验证 *T 实现 S，直接断言，避免递归调用
			return rs.OK(getter(v.Interface().(S)))
		}
	}

	return rs.None[V]()
}
