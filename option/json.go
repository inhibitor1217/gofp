package option

import (
	"bytes"
	"encoding/json"
)

var (
	_ json.Marshaler   = Option[any]{}
	_ json.Unmarshaler = (*Option[any])(nil)
)

// MarshalJSON implements the json.Marshaler interface.
func (o Option[T]) MarshalJSON() ([]byte, error) {
	if o.IsSome() {
		return json.Marshal(o.value)
	}
	return json.Marshal(nil)
}

// UnmarshalJSON implements the json.Unmarshaler interface.
func (o *Option[T]) UnmarshalJSON(data []byte) error {
	if bytes.Equal(data, []byte("null")) {
		*o = None[T]()
		return nil
	}

	var value T
	if err := json.Unmarshal(data, &value); err != nil {
		return err
	}
	*o = Some(value)
	return nil
}
