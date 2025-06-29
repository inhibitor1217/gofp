package option

import (
	"database/sql"
	"database/sql/driver"
	"fmt"
)

var (
	_ driver.Valuer = Option[any]{}
	_ sql.Scanner   = (*Option[any])(nil)
)

// Value implements the driver.Valuer interface.
func (o Option[T]) Value() (driver.Value, error) {
	if o.IsSome() {
		return driver.DefaultParameterConverter.ConvertValue(o.value)
	}
	return nil, nil
}

// Scan implements the sql.Scanner interface.
func (o *Option[T]) Scan(src any) error {
	if src == nil {
		*o = None[T]()
		return nil
	}

	var value T
	if scanner, ok := any(value).(sql.Scanner); ok {
		if err := scanner.Scan(src); err != nil {
			return err
		}
	} else {
		return fmt.Errorf("unsupported type %T", value)
	}

	*o = Some(value)
	return nil
}
