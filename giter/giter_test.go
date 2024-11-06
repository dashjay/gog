package giter_test

import (
	"strconv"
	"testing"

	"github.com/dashjay/gog/giter"
	"github.com/dashjay/gog/internal/constraints"
	"github.com/dashjay/gog/optional"
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

	t.Run("test count", func(t *testing.T) {
		assert.Equal(t, len(_range(0, 10)), giter.Count(giter.FromSlice(_range(0, 10))))
	})

	t.Run("test find", func(t *testing.T) {
		assert.Equal(t, 1,
			optional.FromValue2(giter.Find(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == 1 })).Must())
		assert.False(t, optional.FromValue2(giter.Find(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == -1 })).Ok())

		assert.Equal(t, 1,
			giter.FindO(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == 1 }).Must())
		assert.False(t,
			giter.FindO(giter.FromSlice(_range(0, 10)), func(x int) bool { return x == -1 }).Ok())
	})

	t.Run("test foreach", func(t *testing.T) {
		var res []int
		giter.ForEach(giter.FromSlice(_range(0, 10)), func(i int) bool {
			if i == 5 {
				return false
			}
			res = append(res, i)
			return true
		})
		assert.Equal(t, _range(0, 5), res)

		var idxs []int
		var res2 []int

		giter.ForEachIdx(giter.FromSlice(_range(0, 10)), func(idx int, v int) bool {
			if idx == 5 {
				return false
			}
			idxs = append(idxs, idx)
			res2 = append(res2, v)
			return true
		})
		assert.Equal(t, _range(0, 5), idxs)
		assert.Equal(t, _range(0, 5), res2)
	})

	t.Run("test head", func(t *testing.T) {
		assert.Equal(t, 0,
			optional.FromValue2(giter.Head(giter.FromSlice(_range(0, 10)))).Must())
		assert.False(t,
			optional.FromValue2(giter.Head(giter.FromSlice(_range(0, 0)))).Ok())

		assert.Equal(t, 0,
			giter.HeadO(giter.FromSlice(_range(0, 10))).Must())
		assert.False(t,
			giter.HeadO(giter.FromSlice(_range(0, 0))).Ok())
	})

	t.Run("test join", func(t *testing.T) {
		assert.Equal(t, "1.2.3", giter.Join(giter.FromSlice([]string{"1", "2", "3"}), "."))
		assert.Equal(t, "", giter.Join(giter.FromSlice([]string{}), "."))
	})

	t.Run("min max", func(t *testing.T) {
		assert.Equal(t, 1, giter.Min(giter.FromSlice([]int{3, 2, 1})).Must())
		assert.Equal(t, 3, giter.Max(giter.FromSlice([]int{1, 2, 3})).Must())

		assert.False(t, giter.Min(giter.FromSlice([]int{})).Ok())
		assert.False(t, giter.Max(giter.FromSlice([]int{})).Ok())

		assert.Equal(t, 3,
			giter.MinBy(giter.FromSlice([]int{1, 3, 2}) /*less = */, func(a, b int) bool { return a > b }).Must())
		assert.Equal(t, 1,
			giter.MaxBy(giter.FromSlice([]int{3, 1, 2}) /*less = */, func(a, b int) bool { return a > b }).Must())

		assert.False(t, giter.MinBy(giter.FromSlice([]int{}) /*less = */, func(a, b int) bool { return a > b }).Ok())
		assert.False(t, giter.MaxBy(giter.FromSlice([]int{}) /*less = */, func(a, b int) bool { return a > b }).Ok())
	})
}
