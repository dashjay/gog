package gstl_test

import (
	"testing"

	"github.com/dashjay/gog/gstl"
	"github.com/stretchr/testify/assert"
)

func TestStack(t *testing.T) {
	stack := gstl.NewStack[int]()
	stack = gstl.NewStackWithCap[int](100)
	t.Run("simple test", func(t *testing.T) {
		for i := 0; i < 100; i++ {
			stack.Push(i)
		}

		for i := 99; i >= 0; i-- {
			tr := stack.TopRef()
			assert.Equal(t, i, *tr)
			assert.Equal(t, i, stack.Top())
			assert.Equal(t, i, stack.Pop())
			assert.Equal(t, i, stack.Len())
		}

		assert.Panics(t, func() {
			stack.Top()
		})
		assert.Panics(t, func() {
			stack.Pop()
		})
		assert.Panics(t, func() {
			stack.TopRef()
		})
		assert.True(t, stack.Empty())
	})
}
