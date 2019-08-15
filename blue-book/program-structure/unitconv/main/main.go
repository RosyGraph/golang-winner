// Unitconversion converts each argument from command-line or stdin
// Celsius = Fahrenheit, feet = meters, pounds = kilograms, and their inverse.
package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
	"strings"

	"github.com/RosyGraph/blue-book/program-structure/unitconv"
)

func main() {
	var args []string
	if len(os.Args) < 2 {
		sc := bufio.NewScanner(os.Stdin)
		sc.Scan()
		args = strings.Split(sc.Text(), " ")
	} else {
		args = os.Args[1:]
	}

	for _, v := range args {
		arg, err := strconv.ParseFloat(v, 64)
		if err != nil {
			panic(err)
		}
		c := unitconv.Celsius(arg)
		f := unitconv.Fahrenheit(arg)
		ft := unitconv.Feet(arg)
		m := unitconv.Meters(arg)
		lbs := unitconv.Pounds(arg)
		k := unitconv.Kilograms(arg)
		// Temperature
		fmt.Printf("%s = %s, %s = %s\n", c, unitconv.CToF(c), f, unitconv.FToC(f))
		// Length
		fmt.Printf("%s = %s, %s = %s\n", ft, unitconv.FtToM(ft), m, unitconv.MToFt(m))
		// Mass
		fmt.Printf("%s = %s, %s = %s\n\n", lbs, unitconv.LbsToKg(lbs), k, unitconv.KgToLbs(k))
	}
}
