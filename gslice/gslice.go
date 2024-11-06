package gslice

import (
	"github.com/dashjay/gog/giter"
	"github.com/dashjay/gog/internal/constraints"
	"github.com/dashjay/gog/optional"
)

// All returns true if all elements in the slice satisfy the condition provided by f.
// return false if any element in the slice does not satisfy the condition provided by f.
//
// EXAMPLE:
//
//	giter.All([]int{1, 2, 3}, func(x int) bool { return x > 0 }) 👉 true
//	giter.All([]int{-1, 1, 2, 3}, func(x int) bool { return x > 0 }) 👉 false
func All[T any](in []T, f func(T) bool) bool {
	return giter.AllFromSeq(giter.FromSlice(in), f)
}

// Any returns true if any element in the slice satisfy the condition provided by f.
// return false if none of  element in the slice satisfy the condition provided by f.
//
// EXAMPLE:
//
//	gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == 0 }) 👉 true
//	gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == -1 }) 👉 false
func Any[T any](in []T, f func(T) bool) bool {
	return giter.AnyFromSeq(giter.FromSlice(in), f)
}

// Avg returns the average value of the items in slice (float64).
//
// EXAMPLE:
//
//	gslice.Avg([]int{1, 2, 3}) 👉 float(2)
//	gslice.Avg([]int{}) 👉 float(0)
func Avg[T constraints.Number](in []T) float64 {
	return giter.AvgFromSeq(giter.FromSlice(in))
}

// AvgN returns the average value of the items
//
// EXAMPLE:
//
//	gslice.AvgN(1, 2, 3) 👉 float(2)
//	gslice.AvgN() 👉 float(0)
func AvgN[T constraints.Number](inputs ...T) float64 {
	return giter.AvgFromSeq(giter.FromSlice(inputs))
}

// AvgBy returns the averaged of each item's value evaluated by f.
//
// EXAMPLE:
//
//	gslice.AvgBy([]string{"1", "2", "3"}, func(x string) int {
//		i, _ := strconv.Atoi(x)
//		return i
//	}) 👉 float(2)
func AvgBy[V any, T constraints.Number](in []V, f func(V) T) float64 {
	return giter.AvgByFromSeq(giter.FromSlice(in), f)
}

// Contains returns true if the slice contains the value v.
//
// EXAMPLE:
//
//	gslice.Contains([]int{1, 2, 3}, 1) 👉 true
//	gslice.Contains([]int{-1, 2, 3}, 1) 👉 false
func Contains[T comparable](in []T, v T) bool {
	return giter.Contains(giter.FromSlice(in), v)
}

// ContainsBy returns true if the slice contains the value v evaluated by f.
//
// EXAMPLE:
//
//	gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
//		i, _ := strconv.Atoi(x)
//		return i == 1
//	}) 👉 true
//
//	gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
//		i, _ := strconv.Atoi(x)
//		return i == -1
//	}) 👉 false
func ContainsBy[T any](in []T, f func(T) bool) bool {
	return giter.ContainsBy(giter.FromSlice(in), f)
}

// ContainsAny returns true if the slice contains any value in v.
//
// EXAMPLE:
//
//	gslice.ContainsAny([]string{"1", "2", "3"}, []string{"1", "99", "1000"}) 👉 true
//	gslice.ContainsAny([]string{"1", "2", "3"}, []string{"-1"}) 👉 false
//	gslice.ContainsAny([]string{"1", "2", "3"}, []string{}) 👉 false
func ContainsAny[T comparable](in []T, v []T) bool {
	return giter.ContainsAny(giter.FromSlice(in), v)
}

// ContainsAll returns true if the slice contains all values in v.
//
// EXAMPLE:
//
//	gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "2", "3"})  👉 true
//	gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "99", "1000"}) 👉 false
//	gslice.ContainsAll([]string{"1", "2", "3"}, []string{}) 👉 true
func ContainsAll[T comparable](in []T, v []T) bool {
	return giter.ContainsAll(giter.FromSlice(in), v)
}

// Count returns the number of items in the slice.
//
// EXAMPLE:
//
//	gslice.Count([]int{1, 2, 3}) 👉 3
//	gslice.Count([]int{}) 👉 0
func Count[T any](in []T) int {
	return giter.Count(giter.FromSlice(in))
}

// Find returns the first item in the slice that satisfies the condition provided by f.
//
// EXAMPLE:
//
//	gslice.Find([]int{1, 2, 3}, func(x int) bool { return x == 1 })  👉 1, true
//	gslice.Find([]int{1, 2, 3}, func(x int) bool { return x == -1 }) 👉 0, false
func Find[T any](in []T, f func(T) bool) (val T, found bool) {
	return giter.Find(giter.FromSlice(in), f)
}

// FindO returns the first item in the slice that satisfies the condition provided by f.
//
// EXAMPLE:
//
//	gslice.FindO(_range(0, 10), func(x int) bool { return x == 1 }).Must() 👉 1
//	gslice.FindO(_range(0, 10), func(x int) bool { return x == -1 }).Ok() 👉 false
func FindO[T any](in []T, f func(T) bool) optional.O[T] {
	return giter.FindO(giter.FromSlice(in), f)
}

// ForEach iterates over each item in the slice, stop if f returns false.
//
// EXAMPLE:
//
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
// EXAMPLE:
//
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
// EXAMPLE:
//
//	gslice.HeadO(_range(0, 10)).Must() 👉 0
//	gslice.HeadO(_range(0, 0)).Ok() 👉 false
func HeadO[T any](in []T) optional.O[T] {
	return giter.HeadO(giter.FromSlice(in))
}

// Head returns the first item in the slice.
//
// EXAMPLE:
//
//	optional.FromValue2(gslice.Head(_range(0, 10))).Must() 👉 0
//	optional.FromValue2(gslice.Head(_range(0, 0))).Ok() 👉 false
func Head[T any](in []T) (v T, hasOne bool) {
	return giter.Head(giter.FromSlice(in))
}

// Join joins the slice with sep.
//
// EXAMPLE:
//
//	gslice.Join([]string{"1", "2", "3"}, ".") 👉 "1.2.3"
//	gslice.Join([]string{}, ".") 👉 ""
func Join[T ~string](in []T, sep T) T {
	return giter.Join(giter.FromSlice(in), sep)
}

// Min returns the minimum value in the slice.
//
// EXAMPLE:
//
//	gslice.Min([]int{1, 2, 3}) 👉 1
//	gslice.Min([]int{}) 👉 0
func Min[T constraints.Ordered](in []T) optional.O[T] {
	return giter.Min(giter.FromSlice(in))
}

// MinN returns the minimum value in the slice.
//
// EXAMPLE:
//
//	gslice.MinN(1, 2, 3) 👉 1
func MinN[T constraints.Ordered](in ...T) optional.O[T] {
	return Min(in)
}

// MinBy returns the minimum value evaluated by f in the slice.
//
// EXAMPLE:
//
//	gslice.MinBy([]int{3, 2, 1} /*less = */, func(a, b int) bool { return a > b }).Must() 👉 3
func MinBy[T constraints.Ordered](in []T, f func(T, T) bool) optional.O[T] {
	return giter.MinBy(giter.FromSlice(in), f)
}

// Max returns the maximum value in the slice.
//
// EXAMPLE:
//
//	gslice.Max([]int{1, 2, 3}) 👉 3
//	gslice.Max([]int{}) 👉 0
func Max[T constraints.Ordered](in []T) optional.O[T] {
	return giter.Max(giter.FromSlice(in))
}

// MaxN returns the maximum value in the slice.
//
// EXAMPLE:
//
//	gslice.MaxN(1, 2, 3) 👉 3
func MaxN[T constraints.Ordered](in ...T) optional.O[T] {
	return Max(in)
}

// MaxBy returns the maximum value evaluated by f in the slice.
//
// EXAMPLE:
//
//	gslice.MaxBy([]int{1, 2, 3} /*less = */, func(a, b int) bool { return a > b }).Must() 👉 1
func MaxBy[T constraints.Ordered](in []T, f func(T, T) bool) optional.O[T] {
	return giter.MaxBy(giter.FromSlice(in), f)
}
