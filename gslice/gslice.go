package gslice

import (
	"math/rand"

	"github.com/dashjay/gog/giter"
	"github.com/dashjay/gog/internal/constraints"
	"github.com/dashjay/gog/optional"
)

// All returns true if all elements in the slice satisfy the condition provided by f.
// return false if any element in the slice does not satisfy the condition provided by f.
//
// EXAMPLE:
//
//	gslice.All([]int{1, 2, 3}, func(x int) bool { return x > 0 }) 👉 true
//	gslice.All([]int{-1, 1, 2, 3}, func(x int) bool { return x > 0 }) 👉 false
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

// Map returns a new slice with the results of applying the given function to every element in this slice.
//
// EXAMPLE:
//
//	gslice.Map([]int{1, 2, 3}, func(x int) int { return x * 2 }) 👉 [2, 4, 6]
//	gslice.Map([]int{1, 2, 3}, strconv.Itoa) 👉 ["1", "2", "3"]
func Map[T any, U any](in []T, f func(T) U) []U {
	out := make([]U, len(in))
	for i := range in {
		out[i] = f(in[i])
	}
	return out
}

// Clone returns a copy of the slice.
//
// EXAMPLE:
//
//	gslice.Clone([]int{1, 2, 3}) 👉 [1, 2, 3]
func Clone[T any](in []T) []T {
	if in == nil {
		return nil
	}
	return giter.ToSlice(giter.FromSlice(in))
}

// CloneBy returns a copy of the slice with the results of applying the given function to every element in this slice.
//
// EXAMPLE:
//
//	gslice.CloneBy([]int{1, 2, 3}, func(x int) int { return x * 2 }) 👉 [2, 4, 6]
//	gslice.CloneBy([]int{1, 2, 3}, strconv.Itoa) 👉 ["1", "2", "3"]
func CloneBy[T any, U any](in []T, f func(T) U) []U {
	if in == nil {
		return nil
	}
	return Map(in, f)
}

// Concat concatenates the slices.
//
// EXAMPLE:
//
//	gslice.Concat([]int{1, 2, 3}, []int{4, 5, 6}) 👉 [1, 2, 3, 4, 5, 6]
//	gslice.Concat([]int{1, 2, 3}, []int{}) 👉 [1, 2, 3]
func Concat[T any](vs ...[]T) []T {
	var seqs = make([]giter.Seq[T], 0, len(vs))
	for _, v := range vs {
		seqs = append(seqs, giter.FromSlice(v))
	}
	return giter.ToSlice(giter.Concat(seqs...))
}

// Subset returns a subset slice from the slice.
// if start < -1 means that we take subset from right-to-left
//
// EXAMPLE:
//
//	gslice.Subset([]int{1, 2, 3}, 0, 2) 👉 [1, 2]
//	gslice.Subset([]int{1, 2, 3}, -1, 2) 👉 [2, 3]
func Subset[T any, Slice ~[]T](in Slice, start, count int) Slice {
	if count < 0 {
		count = 0
	}
	if start >= len(in) || -start > len(in) {
		return nil
	}
	if start >= 0 {
		return giter.ToSlice(giter.Limit(giter.Skip(giter.FromSlice(in), start), count))
	} else {
		return giter.ToSlice(giter.Limit(giter.Skip(giter.FromSlice(in), len(in)+start), count))
	}
}

// SubsetInPlace returns a subset slice copied from the slice.
// if start < -1 means that we take subset from right-to-left
// EXAMPLE:
//
//	gslice.SubsetInPlace([]int{1, 2, 3}, 0, 2) 👉 [1, 2]
//	gslice.SubsetInPlace([]int{1, 2, 3}, -1, 2) 👉 [2, 3]
func SubsetInPlace[T any, Slice ~[]T](in Slice, start int, count uint) Slice {
	size := len(in)

	if start < 0 {
		start = size + start
		if start < 0 {
			return Slice{}
		}
	}
	if start > size {
		return Slice{}
	}

	if count > uint(size)-uint(start) {
		count = uint(size - start)
	}
	return in[start : start+int(count)]
}

// Replace replaces the count elements in the slice from 'from' to 'to'.
//
// EXAMPLE:
//
//	gslice.Replace([]int{1, 2, 3}, 2, 4, 1) 👉 [1, 4, 3]
//	gslice.Replace([]int{1, 2, 2}, 2, 4, -1) 👉 [1, 4, 4]
func Replace[T comparable, Slice ~[]T](in Slice, from, to T, count int) []T {
	return giter.ToSlice(giter.Replace(giter.FromSlice(in), from, to, count))
}

// ReplaceAll replaces all elements in the slice from 'from' to 'to'.
//
// EXAMPLE:
//
//	gslice.ReplaceAll([]int{1, 2, 3}, 2, 4) 👉 [1, 4, 3]
//	gslice.ReplaceAll([]int{1, 2, 2}, 2, 4) 👉 [1, 4, 4]
func ReplaceAll[T comparable, Slice ~[]T](in Slice, from, to T) []T {
	return Replace(in, from, to, -1)
}

// ReverseClone reverses the slice.
//
// EXAMPLE:
//
//	gslice.ReverseClone([]int{1, 2, 3}) 👉 [3, 2, 1]
//	gslice.ReverseClone([]int{}) 👉 []int{}
//	gslice.ReverseClone([]int{3, 2, 1}) 👉 [1, 2, 3]
func ReverseClone[T any, Slice ~[]T](in Slice) Slice {
	// why we do not use slices.Reverse() directly ?
	// because lower version golang may has not package "slices"
	return giter.ToSlice(giter.FromSliceReverse(in))
}

// Reverse reverses the slice.
//
// EXAMPLE:
//
//	gslice.Reverse([]int{1, 2, 3}) 👉 [3, 2, 1]
//	gslice.Reverse([]int{}) 👉 []int{}
func Reverse[T any, Slice ~[]T](in Slice) {
	for i, j := 0, len(in)-1; i < j; i, j = i+1, j-1 {
		in[i], in[j] = in[j], in[i]
	}
}

// Repeat returns a new slice with the elements repeated 'count' times.
//
// EXAMPLE:
//
//	gslice.Repeat([]int{1, 2, 3}, 3) 👉 [1, 2, 3, 1, 2, 3, 1, 2, 3]
//	gslice.Repeat([]int{1, 2, 3}, 0) 👉 []int{}
func Repeat[T any, Slice ~[]T](in Slice, count int) Slice {
	return giter.ToSlice(giter.Repeat(giter.FromSlice(in), count))
}

// RepeatBy returns a new slice with the elements return by f repeated 'count' times.
//
// EXAMPLE:
//
//	gslice.RepeatBy(3, func(i int) int { return i }) 👉 [0, 1, 2]
//	gslice.RepeatBy(3, func(i int) string { return strconv.Itoa(i) }) 👉 []string{"1", "2", "3"}
func RepeatBy[T any](n int, f func(i int) T) []T {
	out := make([]T, 0, n)
	for i := 0; i < n; i++ {
		out = append(out, f(i))
	}
	return out
}

// Shuffle shuffles the slice.
//
// EXAMPLE:
//
//	gslice.Shuffle([]int{1, 2, 3}) 👉 [2, 1, 3] (random)
//	gslice.Shuffle([]int{}) 👉 []int{}
func Shuffle[T any, Slice ~[]T](in Slice) Slice {
	return giter.ToSlice(giter.FromSliceShuffle(in))
}

// ShuffleInPlace shuffles the slice.
//
// EXAMPLE:
//
//	array := []int{1, 2, 3}
//	gslice.ShuffleInPlace(array) 👉 [2, 1, 3] (random)
func ShuffleInPlace[T any, Slice ~[]T](in Slice) {
	// why we do not use slices.Shuffle() directly?
	// because lower version golang may has not package "slices"
	rand.Shuffle(len(in), func(i, j int) {
		in[i], in[j] = in[j], in[i]
	})
}
