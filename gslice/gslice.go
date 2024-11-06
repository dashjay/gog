package gslice

import (
	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/giter"
)

// All returns true if all elements in the slice satisfy the condition provided by f.
// return false if any element in the slice does not satisfy the condition provided by f.
//
//	EXAMPLE:
//
//	assert.True(t, giter.All([]int{1, 2, 3}, func(x int) bool { return x > 0 }))
//	assert.False(t, giter.All([]int{-1, 1, 2, 3}, func(x int) bool { return x > 0 }))
func All[T any](in []T, f func(T) bool) bool {
	return giter.AllFromSeq(giter.FromSlice(in), f)
}

// Any returns true if any element in the slice satisfy the condition provided by f.
// return false if none of  element in the slice satisfy the condition provided by f.
//
//	EXAMPLE:
//
//	assert.True(t, gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == 0 }))
//	assert.False(t, gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == -1 }))
func Any[T any](in []T, f func(T) bool) bool {
	return giter.AnyFromSeq(giter.FromSlice(in), f)
}

// Avg returns the average value of the items in slice.
//
//	EXAMPLE:
//
//	assert.Equal(t, float(2), gslice.Avg([]int{1, 2, 3}))
//	assert.Equal(t, float(0), gslice.Avg([]int{}))
func Avg[T constraints.Number](in []T) float64 {
	return giter.AvgFromSeq(giter.FromSlice(in))
}

// AvgBy returns the averaged of each item's value evaluated by f.
//
//	EXAMPLE:
//
//	assert.Equal(t, float64(2), gslice.AvgBy([]string{"1", "2", "3"}, func(x string) int {
//		i, _ := strconv.Atoi(x)
//		return i
//	}))
func AvgBy[V any, T constraints.Number](in []V, f func(V) T) float64 {
	return giter.AvgByFromSeq(giter.FromSlice(in), f)
}
