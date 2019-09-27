package main

import "fmt"

func BinarySearch(arr []int, n int) int {
	for len(arr) > 0 {
		var mid int
		if len(arr)%2 == 0 {
			mid = len(arr)/2 - 1
		} else {
			mid = len(arr) / 2
		}
		switch {
		case arr[mid] == n:
			return mid
		case arr[mid] > n:
			arr = arr[mid:]
		case arr[mid] < n:
			arr = arr[:mid]
		}
	}
	return -1
}

func main() {
	fmt.Println("vim-go")
}
