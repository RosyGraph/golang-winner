package main

import (
	"flag"
	"fmt"
)

type Celsius float64
type Fahrenheit float64

var f = flag.Bool("f", false, "convert temp to Fahrenheit")
var c = flag.Bool("c", false, "convert temp to Celsius")

const (
	AbsoluteZeroC Celsius = -273.15
	FreezingC     Celsius = 0
	BoilingC      Celsius = 100
)

func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}

func main() {
	fmt.Println("vim-go")
}
