//go:build !go1.23
// +build !go1.23

package giter

import (
	"strings"

	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/optional"
)

func AllFromSeq[T any](seq Seq[T], f func(T) bool) bool {
	res := true
	seq(func(v T) bool {
		if !f(v) {
			res = false
			return false
		}
		return true
	})
	return res
}

func AnyFromSeq[T any](seq Seq[T], f func(T) bool) bool {
	res := false
	seq(func(v T) bool {
		if f(v) {
			res = true
			return false
		}
		return true
	})
	return res
}

func AvgFromSeq[T constraints.Number](seq Seq[T]) float64 {
	var sum T
	count := 0

	seq(func(t T) bool {
		sum += t
		count++
		return true
	})
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func AvgByFromSeq[V any, T constraints.Number](seq Seq[V], f func(V) T) float64 {
	var sum T
	count := 0

	seq(func(v V) bool {
		sum += f(v)
		count++
		return true
	})
	if count == 0 {
		return 0
	}
	return float64(sum) / float64(count)
}

func Contains[T comparable](seq Seq[T], in T) bool {
	res := false
	seq(func(t T) bool {
		if in == t {
			res = true
			return false
		}
		return true
	})
	return res
}

func ContainsBy[T any](seq Seq[T], f func(T) bool) bool {
	res := false
	seq(func(t T) bool {
		if f(t) {
			res = true
			return false
		}
		return true
	})
	return res
}

func ContainsAny[T comparable](seq Seq[T], in []T) bool {
	if len(in) == 0 {
		return false
	}
	m := make(map[T]struct{}, len(in))
	for _, v := range in {
		m[v] = struct{}{}
	}
	res := false
	seq(func(t T) bool {
		if _, exists := m[t]; exists {
			res = true
			return false
		}
		return true
	})
	return res
}

func ContainsAll[T comparable](seq Seq[T], in []T) bool {
	if len(in) == 0 {
		return true
	}
	m := make(map[T]struct{}, len(in))
	for _, v := range in {
		m[v] = struct{}{}
	}
	seq(func(t T) bool {
		if _, exists := m[t]; exists {
			delete(m, t)
			if len(m) == 0 {
				return false
			}
		}
		return true
	})
	return len(m) == 0
}

func Count[T any](seq Seq[T]) int {
	var count int
	seq(func(t T) bool {
		count++
		return true
	})
	return count
}

func Find[T any](seq Seq[T], f func(T) bool) (val T, found bool) {
	seq(func(t T) bool {
		found = f(t)
		if found {
			val = t
			return false
		}
		return true
	})
	return
}

func FindO[T any](seq Seq[T], f func(T) bool) optional.O[T] {
	var res = optional.Empty[T]()
	seq(func(t T) bool {
		if f(t) {
			res = optional.FromValue(t)
			return false
		}
		return true
	})
	return res
}

func ForEach[T any](seq Seq[T], f func(T) bool) {
	seq(func(t T) bool {
		return f(t)
	})
}

func ForEachIdx[T any](seq Seq[T], f func(idx int, v T) bool) {
	i := 0
	seq(func(t T) bool {
		c := f(i, t)
		i++
		return c
	})
}

func HeadO[T any](seq Seq[T]) optional.O[T] {
	res := optional.Empty[T]()
	seq(func(t T) bool {
		res = optional.FromValue(t)
		return false
	})
	return res
}

func Head[T any](seq Seq[T]) (v T, hasOne bool) {
	seq(func(t T) bool {
		v = t
		hasOne = true
		return false
	})
	return
}

func Join[T ~string](seq Seq[T], sep T) T {
	//var out T
	//first := false
	//seq(func(t T) bool {
	//	if first {
	//		first = true
	//	} else {
	//		out += sep
	//	}
	//	out += t
	//	return true
	//})
	//return out

	elems := make([]string, 0, 10)
	seq(func(t T) bool {
		elems = append(elems, string(t))
		return true
	})
	return T(strings.Join(elems, string(sep)))
}

func Max[T constraints.Ordered](seq Seq[T]) (r optional.O[T]) {
	first := true
	var _max T
	seq(func(v T) bool {
		if first {
			_max = v
			first = false
		} else if _max < v {
			_max = v
		}
		return true
	})
	if first {
		return
	}
	return optional.FromValue(_max)
}

func MaxBy[T constraints.Ordered](seq Seq[T], less func(T, T) bool) (r optional.O[T]) {
	first := true
	var _max T
	seq(func(v T) bool {
		if first {
			_max = v
			first = false
		} else if less(_max, v) {
			_max = v
		}
		return true
	})
	if first {
		return
	}
	return optional.FromValue(_max)
}

func Min[T constraints.Ordered](seq Seq[T]) (r optional.O[T]) {
	first := true
	var _min T
	seq(func(v T) bool {
		if first {
			_min = v
			first = false
		} else if _min > v {
			_min = v
		}
		return true
	})
	if first {
		return
	}
	return optional.FromValue(_min)
}

func MinBy[T constraints.Ordered](seq Seq[T], less func(T, T) bool) (r optional.O[T]) {
	first := true
	var _min T
	seq(func(v T) bool {
		if first {
			_min = v
			first = false
		} else if less(v, _min) {
			_min = v
		}
		return true
	})
	if first {
		return
	}
	return optional.FromValue(_min)
}
