package rs

type Bool struct {
	v *bool
}

func (b *Bool) Store(v bool) {
	b.v = &v
}

func (b *Bool) IsNil() bool {
	return b == nil || b.v == nil
}
func (b *Bool) IsSet() bool { return !b.IsNil() }

func (b *Bool) IsTrue() bool {
	return b.IsSet() && *b.v
}

func (b *Bool) IsFalse() bool {
	return b.IsSet() && !*b.v
}
