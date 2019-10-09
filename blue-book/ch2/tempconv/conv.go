package tempconv

// CToF converts a C temp to F temp
func CToF(c Celsius) Fahrenheit {
	return Fahrenheit(c*9/5 + 32)
}

// FToC converts a F temp to C temp
func FToC(f Fahrenheit) Celsius {
	return Celsius((f - 32) * 5 / 9)
}
