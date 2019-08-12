package kata

func LongestConsec(strarr []string, k int) (longest string) {
	for mainIndex := 0; mainIndex < len(strarr)-k+1; mainIndex++ {
		subString := subStringof(strarr, mainIndex, k)
		if len(subString) > len(longest) {
			longest = subString
		}
	}
	return
}

func subStringof(strarr []string, start, k int) (subString string) {
	for subIndex := 0; subIndex < k; subIndex++ {
		subString += strarr[start+subIndex]
	}
	return
}
