package llist

type node struct {
	data int
	next *node
}

func push(head *node, data int) node {
	if head == nil {
		return node{data, nil}
	}
	return node{data, head}
}
