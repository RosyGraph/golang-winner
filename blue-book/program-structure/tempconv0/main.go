// Package tempconv performs Clsius and Fahrenit temperature computations.
package tempconv

import "fmt"

type Celsius float64

func (c Celsius) String() string { return fmt.Sprintf("%g°C", c) }

type Fahrenheit float64

func (f Fahrenheit) String() string { return fmt.Sprintf("%g°F", f) }

const (
	AbosulteZeroC Celsius = -273.15
	FreezingC             = 0
	BoilingC              = 100
)

func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }
