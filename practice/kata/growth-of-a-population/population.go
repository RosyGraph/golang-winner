package kata

func NbYear(population int, percent float64, migrants int, goal int) (years int) {
	percent /= 100
	for population < goal {
		population += int(float64(population)*percent) + migrants
		years++
	}
	return
}
