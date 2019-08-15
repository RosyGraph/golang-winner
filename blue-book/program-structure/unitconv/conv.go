package unitconv

//CToF converts a Celsius temperature to Fahrenheit.
func CToF(c Celsius) Fahrenheit { return Fahrenheit(c*9/5 + 32) }

//CToK converts a Celsius temperature to Kelvin.
func CToK(c Celsius) Kelvin { return Kelvin(c + 273.15) }

//FToC converts a Fahrenheit temperature to Celsius.
func FToC(f Fahrenheit) Celsius { return Celsius((f - 32) * 5 / 9) }

//FToK converts a Fahrenheit temperature to Kelvin.
func FToK(f Fahrenheit) Kelvin { return CToK(FToC(f)) }

//KToC converts a Kelvin temperature to Celsius.
func KToC(k Kelvin) Celsius { return Celsius(k - 273.15) }

//KToF converts a Kelvin temperature to Fahrenheit.
func KToF(k Kelvin) Fahrenheit { return CToF(KToC(k)) }

//FtToM converts feet to meters.
func FtToM(f Feet) Meters { return Meters(f * 0.3048) }

//MToFt feet to meters.
func MToFt(m Meters) Feet { return Feet(m * 3.28084) }

//LbsToKg converts pounds to kilograms.
func LbsToKg(p Pounds) Kilograms { return Kilograms(p / 2.205) }

//KgToLbs converts kilograms to pounds.
func KgToLbs(k Kilograms) Pounds { return Pounds(k * 2.205) }
