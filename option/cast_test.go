package option_test

import (
	"testing"

	"github.com/inhibitor1217/gofp/option"
	"github.com/stretchr/testify/assert"
)

func TestFromCast(t *testing.T) {
	v := "foo"
	n := 42

	vAny := any(v)
	nAny := any(n)

	vOpt1 := option.FromCast[string](vAny)
	vOpt2 := option.FromCast[int](vAny)
	nOpt1 := option.FromCast[int](nAny)
	nOpt2 := option.FromCast[string](nAny)

	assert.True(t, vOpt1.IsSome())
	assert.Equal(t, v, vOpt1.UnwrapOrZero())

	assert.True(t, vOpt2.IsNone())

	assert.True(t, nOpt1.IsSome())
	assert.Equal(t, n, nOpt1.UnwrapOrZero())

	assert.True(t, nOpt2.IsNone())
}

type A interface {
	Foo() string
}

type a struct {
	foo string
}

func (a a) Foo() string {
	return a.foo
}

type b struct {
	bar string
}

func TestFromCastInterface(t *testing.T) {
	a := a{foo: "foo"}
	aAny := any(a)
	aOpt := option.FromCast[A](aAny)
	assert.True(t, aOpt.IsSome())
	assert.Equal(t, a, aOpt.Unwrap())

	b := b{bar: "bar"}
	bAny := any(b)
	bOpt := option.FromCast[A](bAny)
	assert.True(t, bOpt.IsNone())
}
