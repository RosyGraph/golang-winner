package main

type Stack struct {
	data rune
	next *Stack
}

func (s Stack) Add(r rune) Stack {
	if &s == nil {
		return Stack{r, nil}
	}
	return Stack{r, &s}
}

func (s Stack) Pop() (rune, Stack) {
	return s.data, *s.next
}

var (
	opening = map[rune]rune{'(': ')', '[': ']', '{': '}'}
	closing = map[rune]rune{')': '(', ']': '[', '}': '{'}
)

func isBalanced(s string) string {
	var stack Stack
	for _, r := range s {
		if _, ok := opening[r]; ok {
			stack = stack.Add(r)
		} else if match, ok := closing[r]; ok {
			if stack.data == 0 {
				return "NO"
			}
			var v rune
			v, stack = stack.Pop()
			if v != match {
				return "NO"
			}
		}
	}
	if stack.data != 0 {
		return "NO"
	}
	return "YES"
}
