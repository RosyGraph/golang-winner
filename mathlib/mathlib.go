package mathlib

import (
	"fmt"
	"math"
)

func Intex() string {
	var d = `\draw[dashed]`
	f := func(x int) int {
		return x * x
	}
	var plot string
	for x := 0; x < 3; x++ {
		plot += fmt.Sprintf("%s (%d,%d) -- (%d,%d) -- (%d,%d);\n",
			d, x, 0, x, f(x+1), x+1, f(x+1))
	}
	plot += fmt.Sprintf("%s (%d,%d) -- (%d,%d);\n", d, 3, 0, 3, f(3))

	return plot
}

const pi float64 = math.Pi

func heron(a, b, c float64) float64 {
	s := (a + b + c) / 2
	return math.Sqrt(s * (s - a) * (s - b) * (s - c))
}

func primes(x int) {
	f := true
	for i := 2; i < x; i++ {
		for j := 2; j < i; j++ {
			if i%j == 0 {
				f = false
				break
			}
		}
		if f {
			fmt.Println(i)
		}
		f = true
	}
}

func h(t float64) float64 {
	return -8.5*math.Cos((2*pi/3)/3*t-13*pi/15) - 35.5
}

func pytID(n, d int) (int, int) {
	return Simplify(d*d-n*n, d*d)
}

func Simplify(n, d int) (int, int) {
	m := lcd(n, d)
	if m == -1 {
		return n, d
	}
	return Simplify(n/m, d/m)
}

func multiples(x int) []int {
	var ans []int
	for i := 2; i < x+1; i++ {
		if x%i == 0 {
			ans = append(ans, i)
		}
	}
	return ans
}

func gcd(x, y int) int {
	if x == 0 {
		return y
	}
	return gcd(y%x, x)
}

func lcd(x, y int) int {
	xm := multiples(x)
	ym := multiples(y)
	for _, v := range xm {
		for _, w := range ym {
			if v == w {
				return v
			}
		}
	}
	return -1
}

func degrees(x float64) float64 {
	y := 180. / math.Pi
	return x * y
}

func rads(x float64) float64 {
	y := math.Pi / 180
	return x * y
}

func sin(x float64) float64 { return math.Sin(rads(x)) }

func cos(x float64) float64 { return math.Cos(rads(x)) }

func tan(x float64) float64 { return math.Tan(rads(x)) }

func asin(x float64) float64 { return degrees(math.Asin(x)) }

func acos(x float64) float64 { return degrees(math.Acos(x)) }

func atan(x float64) float64 { return degrees(math.Atan(x)) }

func average(x, y float64) float64 { return (x + y) / 2. }
