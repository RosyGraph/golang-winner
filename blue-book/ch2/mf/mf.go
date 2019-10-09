package main

import (
	"fmt"
	"os"
	"strconv"

	"github.com/RosyGraph/blue-book/ch2/distconv"
)

func main() {
	if len(os.Args) > 1 {
		cmdline()
	} else {
		fmt.Println("Enter a distance in feet or meters.")
		var d float64
		_, err := fmt.Scanf("%f", &d)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		m := distconv.Meter(d)
		f := distconv.Feet(d)
		fmt.Printf("%s = %s, %s = %s\n",
			m, distconv.MToF(m), f, distconv.FToM(f))
	}
}

func cmdline() {
	for _, arg := range os.Args[1:] {
		d, err := strconv.ParseFloat(arg, 64)
		if err != nil {
			fmt.Fprintf(os.Stderr, "mf: %v\n", err)
			os.Exit(1)
		}
		m := distconv.Meter(d)
		f := distconv.Feet(d)
		fmt.Printf("%s = %s, %s = %s\n",
			m, distconv.MToF(m), f, distconv.FToM(f))
	}
}
