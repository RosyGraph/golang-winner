// Echo3 mimicks Unix's echo command using iteration and Printf.
package main

import (
	"fmt"
	"os"
)

func main() {
	for i, v := range os.Args {
		fmt.Printf("arg %s index %d\n", v, i)
	}
}
