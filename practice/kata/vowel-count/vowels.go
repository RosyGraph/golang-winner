package kata

const vowels = "aeiou"

func GetCount(s string) int {
	var numVowels int
	for _, r := range s {
		if isVowel(r) {
			numVowels++
		}
	}
	return numVowels
}

func isVowel(r rune) bool {
	for _, v := range vowels {
		if r == v {
			return true
		}
	}
	return false
}
