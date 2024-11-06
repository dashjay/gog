package gslice_test

import (
	"strconv"
	"testing"

	"github.com/dashjay/gog/constraints"
	"github.com/dashjay/gog/gslice"
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
		assert.True(t, gslice.All([]int{1, 2, 3}, func(x int) bool { return x > 0 }))
		assert.False(t, gslice.All([]int{-1, 1, 2, 3}, func(x int) bool { return x > 0 }))
		assert.True(t, gslice.All(_range(1, 9999), func(x int) bool { return x > 0 }))
		assert.False(t, gslice.All(_range(0, 9999), func(x int) bool { return x > 0 }))
	})

	t.Run("test any", func(t *testing.T) {
		assert.True(t, gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == 0 }))
		assert.False(t, gslice.Any([]int{0, 1, 2, 3}, func(x int) bool { return x == -1 }))
		assert.True(t, gslice.Any(_range(1, 9999), func(x int) bool { return x == 5000 }))
		assert.False(t, gslice.Any(_range(0, 9999), func(x int) bool { return x < 0 }))
	})

	t.Run("test avg", func(t *testing.T) {
		assert.Equal(t, avg(_range(1, 101)), gslice.Avg(_range(1, 101)))
		assert.Equal(t, float64(0), gslice.Avg([]int{}))
		assert.Equal(t, float64(0), gslice.Avg(_range(-50, 51)))
	})

	t.Run("test avg by", func(t *testing.T) {
		assert.Equal(t, float64(2), gslice.AvgBy([]string{"1", "2", "3"}, func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
	})
}
