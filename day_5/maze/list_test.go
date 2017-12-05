package maze

import (
	"testing"
)

func TestElement(t *testing.T) {
	next := Element{}
	prev := Element{}
	e := Element{prev: &prev, next: &next}

	if e.Next() != &next {
		t.Fatalf("Expected next was `%#v`, got `%#v`", next, e.Next())
	}

	if e.Prev() != &prev {
		t.Fatalf("Expected prev was `%#v`, got `%#v`", prev, e.Prev())
	}

	if next.Next() != nil {
		t.Fatalf("Expected next was nil, got `%#v`", next.Next())
	}

	if next.Prev() != nil {
		t.Fatalf("Expected prev was nil, got `%#v`", next.Prev())
	}
}

func TestNewList(t *testing.T) {
	var expectedLength int
	var length int

	l := New()

	length = l.Len()
	expectedLength = 0
	if length != expectedLength {
		t.Fatalf("Expected length was `%#v`, got `%#v`", expectedLength, length)
	}

	if l.Front() != nil {
		t.Fatal("Front element was not expected")
	}

	if l.Back() != nil {
		t.Fatal("Back element was not expected")
	}

	l.Push(7)
	length = l.Len()
	expectedLength = 1
	if length != expectedLength {
		t.Fatalf("Expected length was `%#v`, got `%#v`", expectedLength, length)
	}

	if l.Front() == nil {
		t.Fatal("Front element was expected")
	}

	if l.Back() == nil {
		t.Fatal("Back element was expected")
	}

	firstExpectedValue := 8
	secondExpectedValue := 1
	l.Push(firstExpectedValue).Push(secondExpectedValue)

	afterJump := l.Front().JumpBy(2)
	if afterJump.Value != secondExpectedValue {
		t.Fatalf("Expected value was `%d`, got `%d`", secondExpectedValue, afterJump.Value)
	}

	afterJump = afterJump.JumpBy(-1)
	if afterJump.Value != firstExpectedValue {
		t.Fatalf("Expected value was `%d`, got `%d`", firstExpectedValue, afterJump.Value)
	}
}

func TestNewPredefinedList(t *testing.T) {
	var expectedLength int
	var length int

	l := NewPredefined(1, 2, 3)

	length = l.Len()
	expectedLength = 3
	if length != expectedLength {
		t.Fatalf("Expected length was `%#v`, got `%#v`", expectedLength, length)
	}

	if l.Front() == nil {
		t.Fatal("Front element was expected")
	}

	if l.Back() == nil {
		t.Fatal("Back element was expected")
	}

	l.Push(7)
	length = l.Len()
	expectedLength = 4
	if length != expectedLength {
		t.Fatalf("Expected length was `%#v`, got `%#v`", expectedLength, length)
	}

	firstExpectedValue := 8
	secondExpectedValue := 1
	l.Push(firstExpectedValue).Push(secondExpectedValue)

	afterJump := l.Front().JumpBy(5)
	if afterJump.Value != secondExpectedValue {
		t.Fatalf("Expected value was `%d`, got `%d`", secondExpectedValue, afterJump.Value)
	}

	afterJump = afterJump.JumpBy(-1)
	if afterJump.Value != firstExpectedValue {
		t.Fatalf("Expected value was `%d`, got `%d`", firstExpectedValue, afterJump.Value)
	}
}
