package option_test

import (
	"testing"

	"github.com/inhibitor1217/gofp/option"
	"github.com/stretchr/testify/assert"
)

func TestAs(t *testing.T) {
	var o option.Option[string]

	option.As(&o, "foo")
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap())

	option.As(&o, 42)
	assert.True(t, o.IsNone())

	foo := "foo"
	option.As(&o, &foo)
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap())

	var v *string
	option.As(&o, v)
	assert.True(t, o.IsNone())

	option.As(&o, nil)
	assert.True(t, o.IsNone())

	option.As(&o, option.None[string]())
	assert.True(t, o.IsNone())

	option.As(&o, option.Some("foo"))
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap())

	fooOpt := option.Some("foo")
	option.As(&o, &fooOpt)
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap())

	var v2 *option.Option[string]
	option.As(&o, v2)
	assert.True(t, o.IsNone())
}

func TestAsInterface(t *testing.T) {
	var o option.Option[A]

	option.As(&o, a{foo: "foo"})
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap().Foo())

	option.As(&o, b{bar: "bar"})
	assert.True(t, o.IsNone())

	aStruct := a{foo: "foo"}
	option.As(&o, &aStruct)
	assert.True(t, o.IsSome())
	assert.Equal(t, "foo", o.Unwrap().Foo())

	var v *A
	option.As(&o, v)
	assert.True(t, o.IsNone())

	option.As(&o, nil)
	assert.True(t, o.IsNone())

	option.As(&o, option.None[A]())
	assert.True(t, o.IsNone())
}
