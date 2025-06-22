package option

// As casts the value into Option of a given type with best effort.
func As[T any](o *Option[T], v any) {
	switch v := v.(type) {
	case T:
		*o = Some(v)
	case *T:
		if v != nil {
			*o = Some(*v)
		} else {
			*o = None[T]()
		}
	case Option[T]:
		*o = v
	case *Option[T]:
		if v != nil {
			*o = *v
		} else {
			*o = None[T]()
		}
	default:
		*o = None[T]()
	}
}
