package option

// Map applies the given function to the value of the Option if it is Some, returns None otherwise.
func Map[T1 any, T2 any](o Option[T1], f func(T1) T2) Option[T2] {
	if o.IsSome() {
		return Some(f(o.value))
	}
	return None[T2]()
}

// MapOr applies the given function to the value of the Option if it is Some, returns the given default value otherwise.
func MapOr[T1 any, T2 any](o Option[T1], defaultValue T2, f func(T1) T2) Option[T2] {
	if o.IsSome() {
		return Some(f(o.value))
	}
	return Some(defaultValue)
}

// MapOrElse applies the given function to the value of the Option if it is Some, returns the result of the given function otherwise.
func MapOrElse[T1 any, T2 any](o Option[T1], defaultValue func() T2, f func(T1) T2) Option[T2] {
	if o.IsSome() {
		return Some(f(o.value))
	}
	return Some(defaultValue())
}

// FlatMap applies the given function to the value of the Option if it is Some, returns None otherwise.
func FlatMap[T1 any, T2 any](o Option[T1], f func(T1) Option[T2]) Option[T2] {
	if o.IsSome() {
		return f(o.value)
	}
	return None[T2]()
}
