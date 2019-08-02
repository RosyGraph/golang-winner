package main

import (
	"fmt"
	"os"
)

var months = []string{
	"January",
	"February",
	"March",
	"April",
	"May",
	"June",
	"July",
	"August",
	"September",
	"October",
	"November",
	"December",
}

func main() {
	month := months[int(os.Args[1])]
	day := os.Args[2]
	year := os.Args[3]
	fmt.Printf("%v %v, %v", month, day, year)
}
