package gstd

import "container/list"

// Element wraps container/list.Element.
type Element[T any] struct {
	ele *list.Element
}

func NewElement[T any](v int) *Element[T] {
	ele := new(list.Element)
	ele.Value = v
	return &Element[T]{ele: ele}
}

func (e *Element[T]) Value() T {
	return e.ele.Value.(T)
}

func (e *Element[T]) Next() *Element[T] {
	if e.ele == nil || e.ele.Next() == nil {
		return nil
	}
	return &Element[T]{e.ele.Next()}
}

func (e *Element[T]) Prev() *Element[T] {
	if e.ele == nil || e.ele.Prev() == nil {
		return nil
	}
	return &Element[T]{e.ele.Prev()}
}

// List wraps container/list.List.
type List[T any] struct {
	_list *list.List
}

func NewList[T any]() *List[T] {
	return &List[T]{_list: list.New()}
}

func (l *List[T]) Back() *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.Back()}
}

func (l *List[T]) Front() *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.Front()}
}

func (l *List[T]) Init() *List[T] {
	l._list.Init()
	return l
}

func (l *List[T]) InsertAfter(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.InsertAfter(v, mark.ele)}
}

func (l *List[T]) InsertBefore(v T, mark *Element[T]) *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.InsertBefore(v, mark.ele)}
}

func (l *List[T]) Len() int {
	l.lazyInit()
	return l._list.Len()
}

func (l *List[T]) MoveAfter(e, mark *Element[T]) {
	l.lazyInit()
	l._list.MoveAfter(e.ele, mark.ele)
}

func (l *List[T]) MoveBefore(e, mark *Element[T]) {
	l.lazyInit()
	l._list.MoveBefore(e.ele, mark.ele)
}

func (l *List[T]) MoveToBack(e *Element[T]) {
	l.lazyInit()
	l._list.MoveToBack(e.ele)
}

func (l *List[T]) MoveToFront(e *Element[T]) {
	l.lazyInit()
	l._list.MoveToFront(e.ele)

}

func (l *List[T]) PushBack(v T) *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.PushBack(v)}
}

func (l *List[T]) PushBackList(other *List[T]) {
	l.lazyInit()
	l._list.PushBackList(other._list)
}

func (l *List[T]) PushFront(v T) *Element[T] {
	l.lazyInit()
	return &Element[T]{l._list.PushFront(v)}
}

func (l *List[T]) PushFrontList(other *List[T]) {
	l.lazyInit()
	l._list.PushFrontList(other._list)

}
func (l *List[T]) Remove(e *Element[T]) T {
	l.lazyInit()
	l._list.Remove(e.ele)
	return e.Value()
}

func (l *List[T]) lazyInit() {
	if l._list == nil {
		l._list = list.New()
	}
}
