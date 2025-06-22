package option_test

import (
	"encoding/json"
	"testing"

	"github.com/inhibitor1217/gofp/option"
	"github.com/stretchr/testify/assert"
)

type inner struct {
	A string `json:"a"`
	B int    `json:"b"`
}

type test struct {
	A string                `json:"a"`
	B option.Option[string] `json:"b"`
	C option.Option[string] `json:"c"`
	D option.Option[int]    `json:"d"`
	E option.Option[int]    `json:"e"`
	F option.Option[inner]  `json:"f"`
	G option.Option[inner]  `json:"g"`
}

func TestMarshalJSON(t *testing.T) {
	out, err := json.Marshal(test{
		A: "a",
		B: option.Some("b"),
		C: option.None[string](),
		D: option.Some(1),
		E: option.None[int](),
		F: option.Some(inner{
			A: "a",
			B: 1,
		}),
		G: option.None[inner](),
	})
	assert.NoError(t, err)
	assert.Equal(t, `{"a":"a","b":"b","c":null,"d":1,"e":null,"f":{"a":"a","b":1},"g":null}`, string(out))
}

func TestUnmarshalJSON(t *testing.T) {
	t.Run("null", func(t *testing.T) {
		var out test
		err := json.Unmarshal([]byte(`{"a":"a","b":"b","c":null,"d":1,"e":null,"f":{"a":"a","b":1},"g":null}`), &out)
		assert.NoError(t, err)
		assert.Equal(t, test{
			A: "a",
			B: option.Some("b"),
			C: option.None[string](),
			D: option.Some(1),
			E: option.None[int](),
			F: option.Some(inner{
				A: "a",
				B: 1,
			}),
			G: option.None[inner](),
		}, out)
	})

	t.Run("undefined", func(t *testing.T) {
		var out test
		err := json.Unmarshal([]byte(`{"a":"a","b":"B","d":1,"f":{"a":"a","b":1}}`), &out)
		assert.NoError(t, err)
		assert.Equal(t, test{
			A: "a",
			B: option.Some("B"),
			C: option.None[string](),
			D: option.Some(1),
			E: option.None[int](),
			F: option.Some(inner{
				A: "a",
				B: 1,
			}),
			G: option.None[inner](),
		}, out)
	})
}
