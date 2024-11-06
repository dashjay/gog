//go:build go1.23
// +build go1.23

package giter

import "github.com/dashjay/gog/constraints"

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
