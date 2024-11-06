//go:build !go1.23
// +build !go1.23

package giter

type Seq[V any] func(yield func(V) bool)
