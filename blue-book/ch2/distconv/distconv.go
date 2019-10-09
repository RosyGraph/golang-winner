package distconv

func FToM(ft Feet) Meter {
	return Meter(ft / 3.28084)
}

func MToF(m Meter) Feet {
	return Feet(m * 3.28084)
}
