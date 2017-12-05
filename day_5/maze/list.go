package maze

// Element holds a value and pointers to next and previous elements
type Element struct {
	prev  *Element
	next  *Element
	Value int
}

// Next returns next Element
func (e *Element) Next() *Element {
	if n := e.next; n != nil {
		return n
	}

	return nil
}

// Prev returns previous Element
func (e *Element) Prev() *Element {
	if n := e.prev; n != nil {
		return n
	}

	return nil
}

// JumpBy jumps to next or previous element depending on given value
func (e *Element) JumpBy(jumps int) *Element {
	if jumps == 0 {
		return e
	}

	if jumps < 0 {
		if e.prev == nil {
			return nil
		}
		jumps += 1
		return e.prev.JumpBy(jumps)
	}

	if e.next == nil {
		return nil
	}

	jumps -= 1
	return e.next.JumpBy(jumps)
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
		l.Push(v)
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

// Push add element at the end of the List
func (l *List) Push(value int) *List {
	e := Element{Value: value}
	if l.len == 0 {
		l.head = &e
		l.tail = &e
	} else {
		e.prev = l.tail
		l.tail.next = &e
		l.tail = &e
	}
	l.len += 1
	return l
}
