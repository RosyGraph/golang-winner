package main

import "testing"

func TestNewLinkedList(t *testing.T) {
	t.Run("empty linked list", func(t *testing.T) {
		got := NewLinkedList()
		want := *LinkedList{}

		if got != want {
			t.Errorf("got %v want %v", got, want)
		}
	})
}

func TestAppend(t *testing.T) {
	list := NewLinkedList()
	e := Element{val: 1, next: nil}
	list.Append(e)
	got := list.head.val
	want := 1

	if got != want {
		t.Errorf("got %d want %d", got, want)
	}
}
