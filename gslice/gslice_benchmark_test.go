//go:build go1.23
// +build go1.23

package gslice_test

import (
	"slices"
	"testing"

	"github.com/dashjay/gog/gslice"
)

func BenchmarkSlice(b *testing.B) {
	const length = 1_000_000
	b.Run("benchmark all", func(b *testing.B) {
		seq := _range(1, length)
		fn := func(i int) bool {
			return i != length-1
		}

		b.Run("baseline", func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				for i := 0; i < len(seq); i++ {
					if !fn(seq[i]) {
						break
					}
				}
			}
		})
		b.Run("gslice", func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				_ = gslice.All(seq, fn)
			}
		})
	})

	b.Run("benchmark any", func(b *testing.B) {
		seq := _range(1, length)
		fn := func(i int) bool {
			return i == length-1
		}

		b.Run("baseline", func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				for i := 0; i < len(seq); i++ {
					_ = fn(seq[i])
				}
			}
		})

		b.Run("gslice", func(b *testing.B) {
			for j := 0; j < b.N; j++ {
				_ = gslice.Any(seq, fn)
			}
		})
	})

	b.Run("benchmark avg", func(b *testing.B) {
		seq := _range(1, length)
		b.Run("baseline", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = avg(seq)
			}

		})
		b.Run("gslice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				_ = gslice.Avg(seq)
			}
		})
	})

	b.Run("benchmark contain", func(b *testing.B) {
		seq := _range(1, length)
		b.Run("baseline", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				slices.Contains(seq, length/2)
			}
		})
		b.Run("gslice", func(b *testing.B) {
			for i := 0; i < b.N; i++ {
				gslice.Contains(seq, length/2)
			}
		})
	})
}
