package main

import (
	"fmt"
	"strconv"
	"strings"
)

type Stack struct {
	data int32
	next *Stack
}

// 1 x -Push the element x onto the stack
// 2   -Delete the top element
// 3   -Print the MAX element in the stack
func getMax(queries []string) []int32 {
	maxElems := []int32{}
	var oTop *Stack
	var vTop *Stack
	for _, query := range queries {
		args := strings.Split(query, " ")
		switch args[0] {
		case "1":
			v64, err := strconv.ParseInt(args[1], 10, 32)
			v := int32(v64)
			if err != nil {
				panic(err)
			}

			// put the new value at the top of the stack
			if oTop == nil {
				oTop = &Stack{v, nil}
				vTop = &Stack{v, nil}
			} else {
				oTop = &Stack{v, oTop}
				if v > vTop.data {
					vTop = &Stack{v, vTop}
				} else {
					tmp := vTop
					for tmp.next != nil && tmp.next.data > v {
						tmp = tmp.next
					}
					tmp.next = &Stack{v, tmp.next}
				}
			}
		case "2":
			v := int32(oTop.data)
			oTop = oTop.next
			if vTop.data == v {
				vTop = vTop.next
			} else {
				tmp := vTop
				for tmp.next.data != v {
					tmp = tmp.next
				}
				tmp.next = tmp.next.next
			}
		case "3":
			maxElems = append(maxElems, vTop.data)
		}
	}
	return maxElems
}

func (stack *Stack) toString() string {
	if stack == nil {
		return ""
	}
	sb := strings.Builder{}
	tmp := stack
	for tmp.next != nil {
		sb.WriteString(fmt.Sprintf("%d -> ", tmp.data))
		tmp = tmp.next
	}
	sb.WriteString(fmt.Sprintf("%d", tmp.data))
	return sb.String()
}

func main() {
	s1 := Stack{1, nil}
	s2 := Stack{2, &s1}
	fmt.Println(s2.toString())
}
