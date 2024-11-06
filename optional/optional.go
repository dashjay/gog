package optional

import "fmt"

type O[T any] struct {
	value T
	ok    bool
}

func FromValue[T any](v T) O[T] {
	return O[T]{value: v, ok: true}
}

func FromValue2[T any](v T, ok bool) O[T] {
	return O[T]{value: v, ok: ok}
}

func Empty[T any]() O[T] {
	return O[T]{ok: false}
}

func (o O[T]) Ptr() *T {
	if o.ok {
		return &o.value
	}
	return nil
}

func (o O[T]) Ok() bool {
	return o.ok
}

func (o O[T]) Must() T {
	if o.ok {
		return o.value
	}
	panic(fmt.Sprintf("Optional.O[%T] has no valid value", o.value))
}

func (o O[T]) ValueOr(dft T) T {
	if o.ok {
		return o.value
	}
	return dft
}

func (o O[T]) ValueOrZero() T {
	if o.ok {
		return o.value
	}
	var empty T
	return empty
}
