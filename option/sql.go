package option

import (
	"database/sql"
	"fmt"
)

var (
	_ sql.Scanner = (*Option[any])(nil)
)

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
