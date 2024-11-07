package gslice_test

import (
	"strconv"
	"testing"

	"github.com/dashjay/gog/giter"
	"github.com/dashjay/gog/gslice"
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

		assert.Equal(t, avg(_range(1, 101)), gslice.AvgN(_range(1, 101)...))
		assert.Equal(t, float64(0), gslice.AvgN([]int{}...))
		assert.Equal(t, float64(0), gslice.AvgN(_range(-50, 51)...))
	})

	t.Run("test avg by", func(t *testing.T) {
		assert.Equal(t, float64(2), gslice.AvgBy([]string{"1", "2", "3"}, func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
		assert.Equal(t, float64(0), gslice.AvgBy([]string{"0"}, func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
		assert.Equal(t, float64(0), gslice.AvgBy([]string{}, func(x string) int {
			i, _ := strconv.Atoi(x)
			return i
		}))
	})

	t.Run("test contains", func(t *testing.T) {
		// contains
		assert.True(t, gslice.Contains([]int{1, 2, 3}, 1))
		assert.False(t, gslice.Contains([]int{-1, 2, 3}, 1))

		// contains by
		assert.True(t, gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
			i, _ := strconv.Atoi(x)
			return i == 1
		}))
		assert.False(t, gslice.ContainsBy([]string{"1", "2", "3"}, func(x string) bool {
			i, _ := strconv.Atoi(x)
			return i == -1
		}))

		// contains any
		assert.True(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{"1", "99", "1000"}))
		assert.False(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{"-1"}))
		assert.False(t, gslice.ContainsAny([]string{"1", "2", "3"}, []string{}))

		// contains all
		assert.True(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "2", "3"}))
		assert.False(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{"1", "99", "1000"}))
		assert.True(t, gslice.ContainsAll([]string{"1", "2", "3"}, []string{}))
	})

	t.Run("test count", func(t *testing.T) {
		assert.Equal(t, 3, gslice.Count([]int{1, 2, 3}))
		assert.Equal(t, 0, gslice.Count([]int{}))
		assert.Equal(t, 10000, gslice.Count(_range(0, 10000)))
	})

	t.Run("test find", func(t *testing.T) {
		assert.Equal(t, 1,
			optional.FromValue2(gslice.Find(_range(0, 10), func(x int) bool { return x == 1 })).Must())
		assert.False(t, optional.FromValue2(gslice.Find(_range(0, 10), func(x int) bool { return x == -1 })).Ok())

		assert.Equal(t, 1,
			gslice.FindO(_range(0, 10), func(x int) bool { return x == 1 }).Must())
		assert.False(t,
			gslice.FindO(_range(0, 10), func(x int) bool { return x == -1 }).Ok())
	})

	t.Run("test foreach", func(t *testing.T) {
		var res []int
		gslice.ForEach(_range(0, 10), func(i int) bool {
			if i == 5 {
				return false
			}
			res = append(res, i)
			return true
		})
		assert.Equal(t, _range(0, 5), res)

		var idxs []int
		var res2 []int

		gslice.ForEachIdx(_range(0, 10), func(idx int, v int) bool {
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
			optional.FromValue2(gslice.Head(_range(0, 10))).Must())
		assert.False(t,
			optional.FromValue2(gslice.Head(_range(0, 0))).Ok())

		assert.Equal(t, 0,
			gslice.HeadO(_range(0, 10)).Must())
		assert.False(t,
			gslice.HeadO(_range(0, 0)).Ok())
	})

	t.Run("test join", func(t *testing.T) {
		assert.Equal(t, "1.2.3", gslice.Join([]string{"1", "2", "3"}, "."))
		assert.Equal(t, "", gslice.Join([]string{}, "."))
	})

	t.Run("min max", func(t *testing.T) {
		assert.Equal(t, 1, gslice.Min([]int{1, 2, 3}).Must())
		assert.Equal(t, 1, gslice.MinN([]int{1, 2, 3}...).Must())
		assert.Equal(t, 3, gslice.Max([]int{1, 2, 3}).Must())
		assert.Equal(t, 3, gslice.MaxN([]int{1, 2, 3}...).Must())

		assert.False(t, gslice.Min([]int{}).Ok())
		assert.False(t, gslice.Max([]int{}).Ok())

		assert.Equal(t, 3,
			gslice.MinBy([]int{3, 2, 1} /*less = */, func(a, b int) bool { return a > b }).Must())
		assert.Equal(t, 1,
			gslice.MaxBy([]int{1, 2, 3} /*less = */, func(a, b int) bool { return a > b }).Must())
	})

	t.Run("clone", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3}, gslice.Clone([]int{1, 2, 3}))
		assert.Len(t, gslice.Clone([]int{}), 0)
		assert.Len(t, gslice.Clone([]int(nil)), 0)
		assert.Equal(t, []string{"1", "2", "3"}, gslice.CloneBy([]int{1, 2, 3}, strconv.Itoa))
		assert.Len(t, gslice.CloneBy([]int{}, strconv.Itoa), 0)
		assert.Len(t, gslice.CloneBy([]int(nil), strconv.Itoa), 0)
	})

	t.Run("concat", func(t *testing.T) {
		assert.Equal(t, []int{1, 2, 3, 4, 5}, gslice.Concat([]int{1, 2, 3}, []int{4, 5}))
		assert.Equal(t, []int{1, 2, 3}, gslice.Concat([]int{1, 2, 3}))
		assert.Len(t, gslice.Concat([]int{}, []int{}, []int{}), 0)
	})

	t.Run("subset", func(t *testing.T) {
		assert.Len(t, gslice.Subset([]int{1, 2, 3}, 0, -1), 0)
		assert.Len(t, gslice.Subset([]int{1, 2, 3}, 3, 0), 0)
		assert.Len(t, gslice.Subset([]int{1, 2, 3}, -3, 0), 0)
		assert.Len(t, gslice.Subset([]int{1, 2, 3}, 0, 0), 0)
		assert.Equal(t, []int{1}, gslice.Subset([]int{1, 2, 3}, 0, 1))
		assert.Equal(t, []int{1, 2}, gslice.Subset([]int{1, 2, 3}, 0, 2))
		assert.Equal(t, []int{1, 2, 3}, gslice.Subset([]int{1, 2, 3}, 0, 3))

		for i := 0; i < 100; i++ {
			assert.Equal(t, _range(i, i+10), gslice.Subset(_range(0, 200), i, 10))
			assert.Subset(t, _range(0, 200), gslice.Subset(_range(0, 200), i, 10))
		}

		assert.Equal(t, []int{3}, gslice.SubsetInPlace([]int{1, 2, 3}, -1, 1))
		assert.Equal(t, []int{3}, gslice.SubsetInPlace([]int{1, 2, 3}, -1, 2))
		assert.Equal(t, []int{3}, gslice.SubsetInPlace([]int{1, 2, 3}, -1, 3))
		assert.Len(t, gslice.SubsetInPlace([]int{1, 2, 3}, -999, 3), 0)
		assert.Len(t, gslice.SubsetInPlace([]int{1, 2, 3}, 999, 3), 0)

		for i := 1; i < 100; i++ {
			assert.Equal(t, _range(200-i, gslice.MinN(200-i+10, 200).Must()), gslice.Subset(_range(0, 200), -i, 10))
		}

		original := []int{1, 2, 3}
		clonedSubset := gslice.Subset(original, 0, 2)
		clonedSubset[0] = 100
		assert.Equal(t, []int{1, 2, 3}, original)
		assert.Equal(t, []int{100, 2}, clonedSubset)
	})

	t.Run("test replace", func(t *testing.T) {
		assert.Equal(t, append([]int{10}, _range(1, 10)...), gslice.ReplaceAll(_range(0, 10), 0, 10))
		assert.Equal(t, append([]int{10}, _range(1, 10)...), gslice.Replace(_range(0, 10), 0, 10, 1))
		assert.Equal(t, append([]int{10}, _range(1, 10)...), gslice.Replace(_range(0, 10), 0, 10, 5))

		// replace nothing
		assert.Equal(t, _range(0, 10), giter.ToSlice(giter.Replace(giter.FromSlice(_range(0, 10)), 0, 100, 0)))
	})

	t.Run("reverse", func(t *testing.T) {
		assert.Equal(t, []int{3, 2, 1}, gslice.ReverseClone([]int{1, 2, 3}))
		assert.Equal(t, []int{5, 4, 3, 2, 1}, gslice.ReverseClone([]int{1, 2, 3, 4, 5}))

		arr := []int{1, 2, 3}
		gslice.Reverse(arr)
		assert.Equal(t, []int{3, 2, 1}, arr)
	})
}
