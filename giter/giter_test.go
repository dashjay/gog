package giter_test

import (
	"strconv"
	"testing"

	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/giter"
	"github.com/stretchr/testify/assert"
)

func _range(a, b int) []int {
	var res []int
	for i := a; i < b; i++ {
		res = append(res, i)
	}
	return res
}

func avg[T constraints.Number](in []T) float64 {
	if len(in) == 0 {
		return 0
	}
	var sum T
	for i := 0; i < len(in); i++ {
		sum += in[i]
	}
	return float64(sum) / float64(len(in))
}

func TestSlices(t *testing.T) {
	t.Run("test all", func(t *testing.T) {
		assert.True(t, giter.AllFromSeq(giter.FromSlice([]int{1, 2, 3}), func(x int) bool { return x > 0 }))
		assert.False(t, giter.AllFromSeq(giter.FromSlice([]int{-1, 1, 2, 3}), func(x int) bool { return x > 0 }))
		assert.True(t, giter.AllFromSeq(giter.FromSlice(_range(1, 9999)), func(x int) bool { return x > 0 }))
		assert.False(t, giter.AllFromSeq(giter.FromSlice(_range(0, 9999)), func(x int) bool { return x > 0 }))
	})

	t.Run("test any", func(t *testing.T) {
		assert.True(t, giter.AnyFromSeq(giter.FromSlice([]int{0, 1, 2, 3}), func(x int) bool { return x == 0 }))
		assert.False(t, giter.AnyFromSeq(giter.FromSlice([]int{0, 1, 2, 3}), func(x int) bool { return x == -1 }))
		assert.True(t, giter.AnyFromSeq(giter.FromSlice(_range(1, 9999)), func(x int) bool { return x == 5000 }))
		assert.False(t, giter.AnyFromSeq(giter.FromSlice(_range(0, 9999)), func(x int) bool { return x < 0 }))
	})

	t.Run("test avg & avg by", func(t *testing.T) {
		assert.Equal(t, avg(_range(1, 101)), giter.AvgFromSeq(giter.FromSlice(_range(1, 101))))
		assert.Equal(t, float64(0), giter.AvgFromSeq(giter.FromSlice([]int{})))
		assert.Equal(t, float64(0), giter.AvgFromSeq(giter.FromSlice(_range(-50, 51))))

		assert.Equal(t, float64(2), giter.AvgByFromSeq(giter.FromSlice([]string{"1", "2", "3"}), func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
		assert.Equal(t, float64(0), giter.AvgByFromSeq(giter.FromSlice([]string{"0"}), func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
		assert.Equal(t, float64(0), giter.AvgByFromSeq(giter.FromSlice([]string{}), func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
	})

	t.Run("test contains", func(t *testing.T) {
		// contains
		assert.True(t, giter.Contains(giter.FromSlice([]int{1, 2, 3}), 1))
		assert.False(t, giter.Contains(giter.FromSlice([]int{-1, 2, 3}), 1))

		// contains by
		assert.True(t, giter.ContainsBy(giter.FromSlice([]string{"1", "2", "3"}), func(x string) bool {
			i, _ := strconv.Atoi(x)
			return i == 1
		}))
		assert.False(t, giter.ContainsBy(giter.FromSlice([]string{"1", "2", "3"}), func(x string) bool {
			i, _ := strconv.Atoi(x)
			return i == -1
		}))

		// contains any
		assert.True(t, giter.ContainsAny(giter.FromSlice([]string{"1", "2", "3"}), []string{"1", "99", "1000"}))
		assert.False(t, giter.ContainsAny(giter.FromSlice([]string{"1", "2", "3"}), []string{"-1"}))
		assert.False(t, giter.ContainsAny(giter.FromSlice([]string{"1", "2", "3"}), []string{}))

		// contains all
		assert.True(t, giter.ContainsAll(giter.FromSlice([]string{"1", "2", "3"}), []string{"1", "2", "3"}))
		assert.False(t, giter.ContainsAll(giter.FromSlice([]string{"1", "2", "3"}), []string{"1", "99", "1000"}))
		assert.True(t, giter.ContainsAll(giter.FromSlice([]string{"1", "2", "3"}), []string{}))
	})
}
