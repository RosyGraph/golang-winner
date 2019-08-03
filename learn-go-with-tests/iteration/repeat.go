package iteration

func repeat(s string, n int) string {
	var repeated string
	for i := 0; i < n; i++ {
		repeated += s
	}
	return repeated
}
