package giter

import (
	"github.com/dashjay/gog/internal/gassert"
	"github.com/dashjay/gog/optional"
)

// FromSlice received a slice and returned a Seq for this slice.
func FromSlice[T any](in []T) Seq[T] {
	return func(yield func(T) bool) {
		for i := 0; i < len(in); i++ {
			if !yield(in[i]) {
				break
			}
		}
	}
}

// At return the element at index from seq.
func At[T any](seq Seq[T], index int) optional.O[T] {
	gassert.MustBePositive(index)
	elements := PullOut(seq, index+1)
	if index >= len(elements) {
		return optional.Empty[T]()
	}
	return optional.FromValue(elements[index])
}

func FromSliceReverse[T any, Slice ~[]T](in Slice) Seq[T] {
	return func(yield func(T) bool) {
		for i := len(in) - 1; i >= 0; i-- {
			if !yield(in[i]) {
				break
			}
		}
	}
}

// Reverse return a reversed seq.
func Reverse[T any](seq Seq[T]) Seq[T] {
	all := PullOut(seq, -1)
	return func(yield func(T) bool) {
		for i := len(all) - 1; i >= 0; i-- {
			if !yield(all[i]) {
				break
			}
		}
	}
}
