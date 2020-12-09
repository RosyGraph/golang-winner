package main

import "fmt"

type Node struct {
	data  int
	left  *Node
	right *Node
}

func main() {
	node := Node{data: 0, left: nil, right: nil}
	temp := node
	for node.left != nil {
		temp := temp.left
		fmt.Println(temp.data)
	}
}
