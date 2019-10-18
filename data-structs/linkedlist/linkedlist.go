package main

import "fmt"

type Element struct {
	next *Element
	prev *Element

	list *LinkedList

	Value interface{}
}

func (e *Element) Next() *Element {
	if p := e.next; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

func (e *Element) Prev() *Element {
	if p := e.prev; e.list != nil && p != &e.list.root {
		return p
	}
	return nil
}

type LinkedList struct {
	root Element
	len  int
}

func (l *LinkedList) Init() *LinkedList {
	l.root.next = &l.root
	l.root.prev = &l.root
	l.len = 0
	return l
}

func New() *LinkedList { return new(LinkedList).Init() }

func (l LinkedList) Len() int { return l.len }

func (list *LinkedList) Append(e Element) {
	c := &list.root
	for c.next != nil {
		c = c.next
	}
	c.next = &e
}

func NewLinkedList() *LinkedList {
	return new(LinkedList)
}

func main() {
	fmt.Printf("%#v\n", NewLinkedList())
}
