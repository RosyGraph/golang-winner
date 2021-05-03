package main

import (
	"errors"
	"fmt"
	"strconv"
	"strings"
)

type BinaryHeap []int32

// Add an element to the heap.
func (h *BinaryHeap) Insert(n int32) {
	if len(*h) == 0 {
		*h = BinaryHeap{n}
		return
	}
	*h = append(*h, n)
	(*h).percolateUp()
}

// Return the root of the heap.
func (h *BinaryHeap) Peek() int32 {
	return (*h)[0]
}

// Delete and return the root of the heap.
func (h *BinaryHeap) Extract() (int32, error) {
	if len(*h) == 0 {
		return 0, errors.New("cannot extract from empty heap")
	}
	top := (*h)[0]
	(*h)[0] = (*h)[len(*h)-1]
	*h = (*h)[0 : len(*h)-1]
	(*h).percolateDown(0)
	return top, nil
}

// Return whether n is present in the heap.
func (h *BinaryHeap) Contains(n int32) bool {
	for _, v := range *h {
		if n == v {
			return true
		}
	}
	return false
}

// Delete the first occurrence of n from the heap.
// Return true if the deletion was successful.
func (h *BinaryHeap) Delete(n int32) bool {
	for i, v := range *h {
		if v == n {
			(*h).swap(i, len(*h)-1)
			*h = (*h)[0 : len(*h)-1]
			(*h).percolateDown(i)
			return true
		}
	}
	return false
}

// Swap the elements at index n with the element at index m.
func (h *BinaryHeap) swap(n, m int) {
	tmp := (*h)[n]
	(*h)[n] = (*h)[m]
	(*h)[m] = tmp
}

// Restore the structure of the heap by cascading up the heap.
func (h *BinaryHeap) percolateUp() {
	var i, j int
	i = len(*h) - 1
	j = parent(i)
	for j >= 0 && (*h)[j] < (*h)[i] {
		h.swap(i, j)
		i = j
		j = parent(i)
	}
}

// Restore the structure of the heap by cascading down the heap.
func (h *BinaryHeap) percolateDown(p int) {
	l, r := children(p)
	if len(*h)-1 < l {
		return
	}
	if len(*h)-1 < r {
		if (*h)[p] < (*h)[l] {
			(*h).swap(p, l)
		}
		return
	}

	var mc int
	if (*h)[r] > (*h)[l] {
		mc = r
	} else {
		mc = l
	}
	if (*h)[p] < (*h)[mc] {
		(*h).swap(p, mc)
		(*h).percolateDown(mc)
	}
}

func children(p int) (int, int) {
	return 2*p + 1, 2*p + 2
}

func parent(i int) int {
	if i%2 == 0 {
		return (i - 2) / 2
	} else {
		return (i - 1) / 2
	}
}

type Stack struct {
	data int32
	next *Stack
}

func (s Stack) Add(n int32) Stack {
	if &s == nil {
		return Stack{n, nil}
	}
	return Stack{n, &s}
}

func (s Stack) Pop() (Stack, int32) {
	return *s.next, s.data
}

// 1 x -Push the element x onto the stack
// 2   -Delete the top element
// 3   -Print the MAX element in the stack
func getMax(queries []string) []int32 {
	maxElems := []int32{}
	var stack Stack
	var heap BinaryHeap
	for _, query := range queries {
		args := strings.Split(query, " ")
		switch args[0] {
		case "1":
			v64, err := strconv.ParseInt(args[1], 10, 32)
			v := int32(v64)
			if err != nil {
				fmt.Println("panic!")
			}

			// put the new value at the top of the stack
			stack = stack.Add(v)
			heap.Insert(v)
		case "2":
			var v int32
			stack, v = stack.Pop()
			heap.Delete(v)
		case "3":
			maxElems = append(maxElems, heap.Peek())
		}
	}
	return maxElems
}

func (s *Stack) String() string {
	if s == nil {
		return ""
	}
	sb := strings.Builder{}
	tmp := s
	for tmp.next != nil {
		v := strconv.FormatInt(int64(tmp.data), 10)
		sb.WriteString(v)
		sb.WriteString(" -> ")
		tmp = tmp.next
	}
	v := strconv.FormatInt(int64(tmp.data), 10)
	sb.WriteString(v)
	return sb.String()
}
