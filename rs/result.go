package rs

import (
	"cmp"
	stderrors "errors"
	"fmt"
	"os"
	"slices"
	"strings"

	"github.com/pkg/errors"
)

var Dev = slices.ContainsFunc(os.Args, func(s string) bool {
	return strings.Contains(s, "test.v")
})

var ErrUnset = stderrors.New("value is unset")

var WrapErr = func(err error) error {
	if Dev {
		err = errors.WithStack(err)
	}
	return err
}

// errOK    = stderrors.New("ok")

type okFlag struct{}

func (okFlag) Error() string { return "OK" }

type Result[T any] struct {
	val T
	err error
}

func Of[T any](val T, err error) Result[T] {
	if err == nil {
		return OK(val)
	}
	return Err[T](err)
}

func Opt[T any](val T, ok bool) Result[T] {
	if ok {
		return OK(val)
	}
	return None[T]()
}

func OK[T any](val T) Result[T] { return Result[T]{val: val, err: okFlag{}} }
func None[T any]() Result[T]    { return Err[T](ErrUnset) }
func Err[T any](err error) Result[T] {
	if WrapErr != nil {
		err = WrapErr(err)
	}
	return Result[T]{err: err}
}

func (t Result[T]) Unwrap() (T, error) { return t.Val(), t.Err() }
func (t Result[T]) Simple() any        { return cmp.Or[any](t.Err(), t.val) }

func (t Result[T]) Val() T { return t.val }
func (t Result[T]) Or(def T) T {
	if t.OK() {
		return t.val
	}
	return def
}

func (t Result[T]) Err() error {
	switch t.err.(type) {
	case okFlag:
		return nil
	case nil:
		return ErrUnset
	default:
		return t.err
	}
}

func (t Result[T]) OK() bool {
	return isInstanceOf[okFlag](t.err)
}

func (t Result[T]) String() string {
	if err := t.Err(); err != nil {
		return "Error:" + err.Error()
	}
	return fmt.Sprint(t.val)
}

func (t Result[T]) WithNone() Result[T]           { return None[T]() }
func (t Result[T]) WithValue(val T) Result[T]     { return OK(val) }
func (t Result[T]) WithError(err error) Result[T] { return Err[T](err) }

func isInstanceOf[T any](x any) bool {
	_, ok := x.(T)
	return ok
}
