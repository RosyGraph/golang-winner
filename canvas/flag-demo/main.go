package main

import (
	"flag"
	"fmt"
)

var f = flag.String("filter", "", "filter options: grayscale, invert, brighten")

func main() {
	flag.Parse()
	fmt.Println("f has value ", *f)
}
