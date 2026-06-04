package rs

import (
	"bytes"
	"strconv"
)

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

func (b Bool) IsTrue() bool {
	return b.IsSet() && *b.v
}

func (b Bool) IsFalse() bool {
	return b.IsSet() && !*b.v
}

func (b Bool) MarshalText() ([]byte, error) {
	return b.MarshalJSON()
}
func (b Bool) MarshalJSON() ([]byte, error) {
	if b.IsNil() {
		return []byte("null"), nil
	}
	return []byte(strconv.FormatBool(*b.v)), nil
}

func (b *Bool) UnmarshalText(text []byte) error {
	text = bytes.TrimSpace(text)
	if len(text) == 0 || bytes.Equal(text, []byte("null")) {
		b.v = nil
		return nil
	}

	v, err := strconv.ParseBool(string(text))
	if err != nil {
		return err
	}
	b.v = &v
	return nil
}

func (b *Bool) UnmarshalJSON(bs []byte) error {
	return b.UnmarshalText(bs)
}
