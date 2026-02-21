package somers

type someType = any

type Option[T someType] struct {
	value T
	none  bool
}

func Some[T someType](value T) Option[T] {
	return Option[T]{
		value: value,
		none:  false,
	}
}

func None[T someType]() Option[T] {
	return Option[T]{
		none: true,
	}
}

func (s *Option[T]) IsNone() bool {
	return s.none
}

func (s *Option[T]) IsSome() bool {
	return !s.IsNone()
}

func (s *Option[T]) Unwrap() T {
	if s.IsNone() {
		panic("called `Option.Unwrap()` on a `None` value")
	}
	return s.value
}

func (s *Option[T]) UnwrapOr(defaultValue T) T {
	if s.IsNone() {
		return defaultValue
	}
	return s.value
}

func (s *Option[T]) UnwarpOrError(err error) (T, error) {
	if s.IsNone() {
		var zero T
		return zero, err
	}
	return s.value, nil
}
