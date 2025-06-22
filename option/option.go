package option

// Option is a container for an optional value of type T.
type Option[T any] struct {
	some  bool
	value T
}

// Some returns an Option with the given value.
func Some[T any](value T) Option[T] {
	return Option[T]{
		some:  true,
		value: value,
	}
}

// None returns an Option with no value.
func None[T any]() Option[T] {
	return Option[T]{
		some: false,
	}
}

// FromPtr returns an Option with the given value if the pointer is not nil, None otherwise.
func FromPtr[T any](ptr *T) Option[T] {
	if ptr == nil {
		return None[T]()
	}
	return Some(*ptr)
}

// FromResult returns an Option with the given value if the result is Ok, None otherwise.
func FromResult[T any](value T, ok bool) Option[T] {
	if !ok {
		return None[T]()
	}
	return Some(value)
}

// FromTry returns an Option with the given value if the function returns no error, None otherwise.
func FromTry[T any](value T, err error) Option[T] {
	if err != nil {
		return None[T]()
	}
	return Some(value)
}

// FromPredicate returns an Option with the given value if the predicate returns true, None otherwise.
func FromPredicate[T any](value T, predicate func(T) bool) Option[T] {
	if predicate(value) {
		return Some(value)
	}
	return None[T]()
}

// IsSome returns true if the Option is Some, false otherwise.
func (o Option[T]) IsSome() bool {
	return o.some
}

// IsNone returns true if the Option is None, false otherwise.
func (o Option[T]) IsNone() bool {
	return !o.some
}

// Unwrap returns the value of the Option if it is Some, panics otherwise.
func (o Option[T]) Unwrap() T {
	if o.IsNone() {
		panic("tried to unwrap None")
	}
	return o.value
}

// UnwrapOr returns the value of the Option if it is Some, the given default value otherwise.
func (o Option[T]) UnwrapOr(defaultValue T) T {
	if o.IsSome() {
		return o.value
	}
	return defaultValue
}

// UnwrapOrElse returns the value of the Option if it is Some, the result of the given function otherwise.
func (o Option[T]) UnwrapOrElse(defaultValue func() T) T {
	if o.IsSome() {
		return o.value
	}
	return defaultValue()
}

// Ptr returns a pointer to the value of the Option if it is Some, nil otherwise.
func (o Option[T]) Ptr() *T {
	if o.IsSome() {
		return &o.value
	}
	return nil
}
