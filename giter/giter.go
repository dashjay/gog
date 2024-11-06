//go:build go1.23
// +build go1.23

package giter

import (
	"strings"

	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/optional"
)

func AllFromSeq[T any](seq Seq[T], f func(T) bool) bool {
	for t := range seq {
		if !f(t) {
			return false
		}
	}
	return true
}

func AnyFromSeq[T any](seq Seq[T], f func(T) bool) bool {
	for t := range seq {
		if f(t) {
			return true
		}
	}
	return false
}

func AvgFromSeq[T constraints.Number](seq Seq[T]) float64 {
	var sum T
	count := 0
	for t := range seq {
		sum += t
		count++
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func AvgByFromSeq[V any, T constraints.Number](seq Seq[V], f func(V) T) float64 {
	var sum T
	count := 0
	for v := range seq {
		sum += f(v)
		count++
	}
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func Contains[T comparable](seq Seq[T], in T) bool {
	for v := range seq {
		if in == v {
			return true
		}
	}
	return false
}

func ContainsBy[T any](seq Seq[T], f func(T) bool) bool {
	for v := range seq {
		if f(v) {
			return true
		}
	}
	return false
}

func ContainsAny[T comparable](seq Seq[T], in []T) bool {
	if len(in) == 0 {
		return false
	}
	m := make(map[T]struct{}, len(in))
	for _, v := range in {
		m[v] = struct{}{}
	}
	for v := range seq {
		if _, exists := m[v]; exists {
			return true
		}
	}
	return false
}

func ContainsAll[T comparable](seq Seq[T], in []T) bool {
	if len(in) == 0 {
		return true
	}
	m := make(map[T]struct{}, len(in))
	for _, v := range in {
		m[v] = struct{}{}
	}
	for v := range seq {
		if _, exists := m[v]; exists {
			delete(m, v)
			if len(m) == 0 {
				return true
			}
		}
	}
	return len(m) == 0
}

func Count[T any](seq Seq[T]) int {
	var count int
	for range seq {
		count++
	}
	return count
}

func Find[T any](seq Seq[T], f func(T) bool) (val T, found bool) {
	for v := range seq {
		if f(v) {
			val = v
			found = true
			return
		}
	}
	return
}

func FindO[T any](seq Seq[T], f func(T) bool) optional.O[T] {
	for v := range seq {
		if f(v) {
			return optional.FromValue(v)
		}
	}
	return optional.Empty[T]()
}

func ForEach[T any](seq Seq[T], f func(T) bool) {
	for v := range seq {
		if !f(v) {
			break
		}
	}
}

func ForEachIdx[T any](seq Seq[T], f func(idx int, v T) bool) {
	idx := 0
	for v := range seq {
		if !f(idx, v) {
			break
		}
		idx++
	}
}

func HeadO[T any](seq Seq[T]) optional.O[T] {
	for v := range seq {
		return optional.FromValue(v)
	}
	return optional.Empty[T]()
}

func Head[T any](seq Seq[T]) (v T, hasOne bool) {
	for t := range seq {
		v = t
		hasOne = true
		return
	}
	return
}

func Join[T ~string](seq Seq[T], sep T) T {
	elems := make([]string, 0, 10)
	for v := range seq {
		elems = append(elems, string(v))
	}
	return T(strings.Join(elems, string(sep)))
}

func Max[T constraints.Ordered](seq Seq[T]) T {
	first := true
	var _max T
	for v := range seq {
		if first {
			_max = v
			first = false
		} else if _max < v {
			_max = v
		}
	}
	return _max
}

func Min[T constraints.Ordered](seq Seq[T]) T {
	first := true
	var _min T
	for v := range seq {
		if first {
			_min = v
			first = false
		} else if _min > v {
			_min = v
		}
	}
	return _min
}
