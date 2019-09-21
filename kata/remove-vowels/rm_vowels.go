package main

var vowels = map[rune]bool{'A': true, 'E': true, 'I': true, 'O': true, 'U': true,
	'a': true, 'e': true, 'i': true, 'o': true, 'u': true}

func RemoveVowels(s string) string {
	var v []rune

	for _, r := range s {
		if !vowels[r] {
			v = append(v, r)
		}
	}
	return string(v)
}
