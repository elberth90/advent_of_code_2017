package bank

import (
	"fmt"
)

// Element holds a value and pointers to next and previous elements
type Element struct {
	next  *Element
	prev  *Element
	Value int
}

// Next returns next Element
func (e *Element) Next() *Element {
	return e.next
}

// Prev returns previous Element
func (e *Element) Prev() *Element {
	return e.prev
}

// List holds information about length of the list and pointers to head and tail Elements
type List struct {
	head *Element
	tail *Element
	len  int
}

// New returns empty List
func New() *List {
	return new(List)
}

// NewPredefined returns List filled with given elements
func NewPredefined(values ...int) *List {
	l := new(List)
	for _, v := range values {
		l.PushFront(v)
	}
	return l
}

// Len returns length of the list
func (l *List) Len() int {
	return l.len
}

// Front returns head Element
func (l *List) Front() *Element {
	return l.head
}

// Back return tail Element
func (l *List) Back() *Element {
	return l.tail
}

// PushFront add element at the beginning of the List
func (l *List) PushFront(value int) *List {
	e := Element{Value: value}
	if l.len == 0 {
		l.head = &e
		l.tail = &e
		e.next = &e
		e.prev = &e
	} else {
		l.head.prev = &e
		e.next = l.head
		e.prev = l.tail
		l.head = &e
		l.tail.next = &e
	}
	l.len++
	return l
}

// Max return Element with maximum value
func (l *List) Max() *Element {
	currentMax := l.Back()
	for i, e := 0, l.Back(); i < l.len-1; i++ {
		e = e.Prev()
		if e.Value > currentMax.Value {
			currentMax = e
		}
	}

	return currentMax
}

// Hash get hash of list
func (l *List) Hash() string {
	var input string
	for i, e := 0, l.Back(); i < l.len-1; i++ {
		e = e.Prev()
		input = fmt.Sprintf("%s%d", input, e.Value)
	}

	return input
}
