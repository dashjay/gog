package gslice

import (
	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/giter"
	"github.com/dashjay/gog/optional"
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

// AvgN returns the average value of the items
//
//	EXAMPLE:
//
//	assert.Equal(t, float(2), gslice.AvgN(1, 2, 3)
//	assert.Equal(t, float(0), gslice.AvgN()
func AvgN[T constraints.Number](inputs ...T) float64 {
	return giter.AvgFromSeq(giter.FromSlice(inputs))
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

// Contains returns true if the slice contains the value v.
//
//	EXAMPLE:
//	assert.True(t, gslice.Contains([]int{1, 2, 3}, 1))
//	assert.False(t, gslice.Contains([]int{-1, 2, 3}, 1))
func Contains[T comparable](in []T, v T) bool {
	return giter.Contains(giter.FromSlice(in), v)
}

// ContainsBy returns true if the slice contains the value v evaluated by f.
//
//	EXAMPLE:
//
//	assert.True(t, gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
//		i, _ := strconv.Atoi(x)
//		return i == 1
//	}))
//	assert.False(t, gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
//		i, _ := strconv.Atoi(x)
//		return i == -1
//	}))
func ContainsBy[T any](in []T, f func(T) bool) bool {
	return giter.ContainsBy(giter.FromSlice(in), f)
}

// ContainsAny returns true if the slice contains any value in v.
//
//	EXAMPLE:
//
//	assert.True(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{"1", "99", "1000"}))
//	assert.False(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{"-1"}))
//	assert.False(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{}))
func ContainsAny[T comparable](in []T, v []T) bool {
	return giter.ContainsAny(giter.FromSlice(in), v)
}

// ContainsAll returns true if the slice contains all values in v.
//
//	EXAMPLE:
//
//	assert.True(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
//	assert.False(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "99", "1000"}))
//	assert.True(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{}))
func ContainsAll[T comparable](in []T, v []T) bool {
	return giter.ContainsAll(giter.FromSlice(in), v)
}

// Count returns the number of items in the slice.
//
//	EXAMPLE:
//
//	assert.Equal(t, 3, gslice.Count([]int{1, 2, 3}))
//	assert.Equal(t, 0, gslice.Count([]int{}))
func Count[T any](in []T) int {
	return giter.Count(giter.FromSlice(in))
}

// Find returns the first item in the slice that satisfies the condition provided by f.
//
//	EXAMPLE:
//
//	val, found := gslice.Find([]int{1, 2, 3}, func(x int) bool { return x == 1 })
//	assert.True(t, found)
//	assert.Equal(t, 1, val)
//	val, found = gslice.Find([]int{1, 2, 3}, func(x int) bool { return x == -1 })
//	assert.False(t, found)
func Find[T any](in []T, f func(T) bool) (val T, found bool) {
	return giter.Find(giter.FromSlice(in), f)
}

// FindO returns the first item in the slice that satisfies the condition provided by f.
//
//	EXAMPLE:
//	assert.Equal(t, 1,
//		giter.FindO(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == 1 }).Must())
//	assert.False(t,
//		giter.FindO(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == -1 }).Ok())
func FindO[T any](in []T, f func(T) bool) optional.O[T] {
	return giter.FindO(giter.FromSlice(in), f)
}

// ForEach iterates over each item in the slice, stop if f returns false.
//
//	EXAMPLE:
//	ForEach([]int{1, 2, 3}, func(x int) bool {
//		fmt.Println(x)
//		return true
//	}
//	Output:
//	1
//	2
//	3
func ForEach[T any](in []T, f func(T) bool) {
	giter.ForEach(giter.FromSlice(in), f)
}

// ForEachIdx iterates over each item in the slice, stop if f returns false.
//
//	EXAMPLE:
//	ForEach([]int{1, 2, 3}, func(idx, x int) bool {
//		fmt.Println(idx, x)
//		return true
//	}
//	Output:
//	0 1
//	1 2
//	2 3
func ForEachIdx[T any](in []T, f func(idx int, v T) bool) {
	giter.ForEachIdx(giter.FromSlice(in), f)
}

// HeadO returns the first item in the slice.
//
//	EXAMPLE:
//
//	assert.Equal(t, 0, gslice.HeadO(_range(0, 10)).Must())
//	assert.False(t, gslice.HeadO(_range(0, 0)).Ok())
func HeadO[T any](in []T) optional.O[T] {
	return giter.HeadO(giter.FromSlice(in))
}

// Head returns the first item in the slice.
//
//	EXAMPLE:
//	assert.Equal(t, 0, optional.FromValue2(gslice.Head(_range(0, 10))).Must())
//	assert.False(t, optional.FromValue2(gslice.Head(_range(0, 0))).Ok())
func Head[T any](in []T) (v T, hasOne bool) {
	return giter.Head(giter.FromSlice(in))
}

// Join joins the slice with sep.
//
//	EXAMPLE:
//	assert.Equal(t, "1.2.3", gslice.Join([]string{"1", "2", "3"}, "."))
//	assert.Equal(t, "", gslice.Join([]string{}, "."))
func Join[T ~string](in []T, sep T) T {
	return giter.Join(giter.FromSlice(in), sep)
}

// Min returns the minimum value in the slice.
//
//	EXAMPLE:
//	assert.Equal(t, 1, gslice.Min([]int{1, 2, 3}))
//	assert.Equal(t, 0, gslice.Min([]int{}))
func Min[T constraints.Ordered](in []T) T {
	return giter.Min(giter.FromSlice(in))
}

// MinN returns the minimum value in the slice.
//
//	EXAMPLE:
//	assert.Equal(t, 1, gslice.MinN(1, 2, 3))
func MinN[T constraints.Ordered](in ...T) T {
	return Min(in)
}

// Max returns the maximum value in the slice.
//
//	EXAMPLE:
//	assert.Equal(t, 3, gslice.Max([]int{1, 2, 3}))
//	assert.Equal(t, 0, gslice.Max([]int{}))
func Max[T constraints.Ordered](in []T) T {
	return giter.Max(giter.FromSlice(in))
}

// MaxN returns the maximum value in the slice.
//
//	EXAMPLE:
//	assert.Equal(t, 3, gslice.MaxN(1, 2, 3))
func MaxN[T constraints.Ordered](in ...T) T {
	return Max(in)
}
