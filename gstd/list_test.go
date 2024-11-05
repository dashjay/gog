package gstd_test

import (
	"strconv"
	"testing"

	"github.com/dashjay/gog/gstd"
	"github.com/stretchr/testify/assert"
)

func TestListWrapper(t *testing.T) {
	t.Run("simple push back", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}

		ele := l.Front()
		for i := 0; i < 100; i++ {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}

		ele = l.Back()
		for i := 99; i >= 0; i-- {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Prev()
		}
	})

	t.Run("simple push front", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 99; i >= 0; i-- {
			l.PushFront(strconv.Itoa(i))
		}

		ele := l.Front()
		for i := 0; i < 100; i++ {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}

		ele = l.Back()
		for i := 99; i >= 0; i-- {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Prev()
		}
	})

	t.Run("init", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}

		ele := l.Front()
		for i := 0; i < 100; i++ {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}

		l.Init()
		assert.Equal(t, 0, l.Len())
	})

	t.Run("insert after & before", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}
		ele := l.Front()
		for ele.Value() != "50" {
			ele = ele.Next()
		}
		assert.Equal(t, "50", ele.Value())
		l.InsertAfter("after-50", ele)
		l.InsertBefore("before-50", ele)

		assert.Equal(t, "after-50", ele.Next().Value())
		assert.Equal(t, "before-50", ele.Prev().Value())
	})

	t.Run("move after & before", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}
		ele := l.Front()
		for ele.Value() != "50" {
			ele = ele.Next()
		}

		l.MoveBefore(ele, l.Front())
		assert.Equal(t, "50", ele.Value())
		assert.Equal(t, l.Front(), ele)

		l.MoveAfter(ele, l.Back())
		assert.Equal(t, "50", ele.Value())
		assert.Equal(t, l.Back(), ele)

	})

	t.Run("move to & back", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}
		ele := l.Front()
		for ele.Value() != "50" {
			ele = ele.Next()
		}

		l.MoveToFront(ele)
		assert.Equal(t, "50", ele.Value())
		assert.Equal(t, l.Front(), ele)

		l.MoveToBack(ele)
		assert.Equal(t, "50", ele.Value())
		assert.Equal(t, l.Back(), ele)
	})

	t.Run("remove", func(t *testing.T) {
		l := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			l.PushBack(strconv.Itoa(i))
		}
		ele := l.Front()
		for ele.Value() != "50" {
			ele = ele.Next()
		}
		val := l.Remove(ele)
		assert.Equal(t, "50", val)

		ele = l.Front()
		for i := 0; i < 100; i++ {
			if i == 50 {
				continue
			}
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}
	})

	t.Run("push back list", func(t *testing.T) {
		left := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			left.PushBack(strconv.Itoa(i))
		}
		right := gstd.NewList[string]()
		for i := 100; i < 200; i++ {
			right.PushBack(strconv.Itoa(i))
		}

		left.PushBackList(right)

		ele := left.Front()
		for i := 0; i < 200; i++ {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}
	})

	t.Run("push front list", func(t *testing.T) {
		left := gstd.NewList[string]()
		for i := 0; i < 100; i++ {
			left.PushBack(strconv.Itoa(i))
		}
		right := gstd.NewList[string]()
		for i := 100; i < 200; i++ {
			right.PushBack(strconv.Itoa(i))
		}

		right.PushFrontList(left)

		ele := right.Front()
		for i := 0; i < 200; i++ {
			assert.Equal(t, strconv.Itoa(i), ele.Value())
			ele = ele.Next()
		}
	})
}

// test copy from src/container/list/list_test.go
func checkListLen(t *testing.T, l *gstd.List[int], len int) bool {
	if n := l.Len(); n != len {
		t.Errorf("l.Len() = %d, want %d", n, len)
		return false
	}
	return true
}

func checkListPointers(t *testing.T, l *gstd.List[int], es []*gstd.Element[int]) {
	//root := l

	if !checkListLen(t, l, len(es)) {
		return
	}

	// zero length lists must be the zero value or properly initialized (sentinel circle)
	//if len(es) == 0 {
	//	if l.root.next != nil && l.root.next != root || l.root.prev != nil && l.root.prev != root {
	//		t.Errorf("l.root.next = %p, l.root.prev = %p; both should both be nil or %p", l.root.next, l.root.prev, root)
	//	}
	//	return
	//}
	// len(es) > 0

	// check internal and external prev/next connections
	for i, e := range es {
		prev := (*gstd.Element[int])(nil)
		Prev := (*gstd.Element[int])(nil)
		if i > 0 {
			prev = es[i-1]
			Prev = prev
		}
		//if p := e.prev; p != prev {
		//	t.Errorf("elt[%d](%p).prev = %p, want %p", i, e, p, prev)
		//}
		if p := e.Prev(); p != Prev {
			t.Errorf("elt[%d](%p).Prev() = %p, want %p", i, e, p, Prev)
		}

		next := (*gstd.Element[int])(nil)
		Next := (*gstd.Element[int])(nil)
		if i < len(es)-1 {
			next = es[i+1]
			Next = next
		}
		//if n := e.next; n != next {
		//	t.Errorf("elt[%d](%p).next = %p, want %p", i, e, n, next)
		//}
		if n := e.Next(); n != Next {
			t.Errorf("elt[%d](%p).Next() = %p, want %p", i, e, n, Next)
		}
	}
}

func TestList(t *testing.T) {
	l := gstd.NewList[int]()
	checkListPointers(t, l, []*gstd.Element[int]{})

	// Single element list
	e := l.PushFront(1)
	checkListPointers(t, l, []*gstd.Element[int]{e})
	l.MoveToFront(e)
	checkListPointers(t, l, []*gstd.Element[int]{e})
	l.MoveToBack(e)
	checkListPointers(t, l, []*gstd.Element[int]{e})
	l.Remove(e)
	checkListPointers(t, l, []*gstd.Element[int]{})

	// Bigger list
	e2 := l.PushFront(2)
	e1 := l.PushFront(1)
	e3 := l.PushBack(3)
	e4 := l.PushBack(999)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e3, e4})

	l.Remove(e2)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e3, e4})

	l.MoveToFront(e3) // move from middle
	checkListPointers(t, l, []*gstd.Element[int]{e3, e1, e4})

	l.MoveToFront(e1)
	l.MoveToBack(e3) // move from middle
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e3})

	l.MoveToFront(e3) // move from back
	checkListPointers(t, l, []*gstd.Element[int]{e3, e1, e4})
	l.MoveToFront(e3) // should be no-op
	checkListPointers(t, l, []*gstd.Element[int]{e3, e1, e4})

	l.MoveToBack(e3) // move from front
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e3})
	l.MoveToBack(e3) // should be no-op
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e3})

	e2 = l.InsertBefore(2, e1) // insert before front
	checkListPointers(t, l, []*gstd.Element[int]{e2, e1, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e4) // insert before middle
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertBefore(2, e3) // insert before back
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e2, e3})
	l.Remove(e2)

	e2 = l.InsertAfter(2, e1) // insert after front
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e4, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e4) // insert after middle
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e2, e3})
	l.Remove(e2)
	e2 = l.InsertAfter(2, e3) // insert after back
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e3, e2})
	l.Remove(e2)

	// Check standard iteration.
	sum := 0
	for e := l.Front(); e != nil; e = e.Next() {
		sum += e.Value()
	}
	if sum != 4 {
		t.Errorf("sum over l = %d, want 4", sum)
	}

	// Clear all elements by iterating
	var next *gstd.Element[int]
	for e := l.Front(); e != nil; e = next {
		next = e.Next()
		l.Remove(e)
	}
	checkListPointers(t, l, []*gstd.Element[int]{})
}

func checkList(t *testing.T, l *gstd.List[int], es []any) {
	if !checkListLen(t, l, len(es)) {
		return
	}

	i := 0
	for e := l.Front(); e != nil; e = e.Next() {
		le := e.Value()
		if le != es[i] {
			t.Errorf("elt[%d].Value = %v, want %v", i, le, es[i])
		}
		i++
	}
}

func TestExtending(t *testing.T) {
	l1 := gstd.NewList[int]()
	l2 := gstd.NewList[int]()

	l1.PushBack(1)
	l1.PushBack(2)
	l1.PushBack(3)

	l2.PushBack(4)
	l2.PushBack(5)

	l3 := gstd.NewList[int]()
	l3.PushBackList(l1)
	checkList(t, l3, []any{1, 2, 3})
	l3.PushBackList(l2)
	checkList(t, l3, []any{1, 2, 3, 4, 5})

	l3 = gstd.NewList[int]()
	l3.PushFrontList(l2)
	checkList(t, l3, []any{4, 5})
	l3.PushFrontList(l1)
	checkList(t, l3, []any{1, 2, 3, 4, 5})

	checkList(t, l1, []any{1, 2, 3})
	checkList(t, l2, []any{4, 5})

	l3 = gstd.NewList[int]()
	l3.PushBackList(l1)
	checkList(t, l3, []any{1, 2, 3})
	l3.PushBackList(l3)
	checkList(t, l3, []any{1, 2, 3, 1, 2, 3})

	l3 = gstd.NewList[int]()
	l3.PushFrontList(l1)
	checkList(t, l3, []any{1, 2, 3})
	l3.PushFrontList(l3)
	checkList(t, l3, []any{1, 2, 3, 1, 2, 3})

	l3 = gstd.NewList[int]()
	l1.PushBackList(l3)
	checkList(t, l1, []any{1, 2, 3})
	l1.PushFrontList(l3)
	checkList(t, l1, []any{1, 2, 3})
}

func TestRemove(t *testing.T) {
	l := gstd.NewList[int]()
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2})
	e := l.Front()
	l.Remove(e)
	checkListPointers(t, l, []*gstd.Element[int]{e2})
	l.Remove(e)
	checkListPointers(t, l, []*gstd.Element[int]{e2})
}

func TestIssue4103(t *testing.T) {
	l1 := gstd.NewList[int]()
	l1.PushBack(1)
	l1.PushBack(2)

	l2 := gstd.NewList[int]()
	l2.PushBack(3)
	l2.PushBack(4)

	e := l1.Front()
	l2.Remove(e) // l2 should not change because e is not an element of l2
	if n := l2.Len(); n != 2 {
		t.Errorf("l2.Len() = %d, want 2", n)
	}

	l1.InsertBefore(8, e)
	if n := l1.Len(); n != 3 {
		t.Errorf("l1.Len() = %d, want 3", n)
	}
}

func TestIssue6349(t *testing.T) {
	l := gstd.NewList[int]()
	l.PushBack(1)
	l.PushBack(2)

	e := l.Front()
	l.Remove(e)
	if e.Value() != 1 {
		t.Errorf("e.value = %d, want 1", e.Value())
	}
	if e.Next() != nil {
		t.Errorf("e.Next() != nil")
	}
	if e.Prev() != nil {
		t.Errorf("e.Prev() != nil")
	}
}

func TestMove(t *testing.T) {
	l := gstd.NewList[int]()
	e1 := l.PushBack(1)
	e2 := l.PushBack(2)
	e3 := l.PushBack(3)
	e4 := l.PushBack(4)

	l.MoveAfter(e3, e3)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e3, e4})
	l.MoveBefore(e2, e2)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e3, e4})

	l.MoveAfter(e3, e2)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e3, e4})
	l.MoveBefore(e2, e3)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e2, e3, e4})

	l.MoveBefore(e2, e4)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e3, e2, e4})
	e2, e3 = e3, e2

	l.MoveBefore(e4, e1)
	checkListPointers(t, l, []*gstd.Element[int]{e4, e1, e2, e3})
	e1, e2, e3, e4 = e4, e1, e2, e3

	l.MoveAfter(e4, e1)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e4, e2, e3})
	e2, e3, e4 = e4, e2, e3

	l.MoveAfter(e2, e3)
	checkListPointers(t, l, []*gstd.Element[int]{e1, e3, e2, e4})
}

// Test PushFront, PushBack, PushFrontList, PushBackList with uninitialized List
func TestZeroList(t *testing.T) {
	var l1 = new(gstd.List[int])
	l1.PushFront(1)
	checkList(t, l1, []any{1})

	var l2 = new(gstd.List[int])
	l2.PushBack(1)
	checkList(t, l2, []any{1})

	var l3 = new(gstd.List[int])
	l3.PushFrontList(l1)
	checkList(t, l3, []any{1})

	var l4 = new(gstd.List[int])
	l4.PushBackList(l2)
	checkList(t, l4, []any{1})
}

// Test that a list l is not modified when calling InsertBefore with a mark that is not an element of l.
func TestInsertBeforeUnknownMark(t *testing.T) {
	var l gstd.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.InsertBefore(1, gstd.NewElement[int](0))
	checkList(t, &l, []any{1, 2, 3})
}

// Test that a list l is not modified when calling InsertAfter with a mark that is not an element of l.
func TestInsertAfterUnknownMark(t *testing.T) {
	var l gstd.List[int]
	l.PushBack(1)
	l.PushBack(2)
	l.PushBack(3)
	l.InsertAfter(1, gstd.NewElement[int](0))
	checkList(t, &l, []any{1, 2, 3})
}

// Test that a list l is not modified when calling MoveAfter or MoveBefore with a mark that is not an element of l.
func TestMoveUnknownMark(t *testing.T) {
	var l1 gstd.List[int]
	e1 := l1.PushBack(1)

	var l2 gstd.List[int]
	e2 := l2.PushBack(2)

	l1.MoveAfter(e1, e2)
	checkList(t, &l1, []any{1})
	checkList(t, &l2, []any{2})

	l1.MoveBefore(e1, e2)
	checkList(t, &l1, []any{1})
	checkList(t, &l2, []any{2})
}
