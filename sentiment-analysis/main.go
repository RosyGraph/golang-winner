package main

import (
	"bufio"
	"fmt"
	"os"
	"regexp"
	"sort"
	"strings"
)

func main() {
	// Read a file path from command line.
	path := os.Args[1]

	// Store lowercase words stripped of punctuation and stop words as a string
	// slice.
	words := ReadAsStringArray(path)
	sentiment := CalculateSentiment(words)

	// Print sentiment analysis for the file.
	fmt.Printf("Overall sentiment for '%s' is %f\n", path, sentiment)

}

// Read a file at path and create an alphabetized copy in "tmp/path".
func SortFile(path string) {
	words := ReadAsStringArray(path)
	sort.Strings(words)

	f, err := os.Create("tmp/" + path)
	if err != nil {
		panic(err)
	}
	w := bufio.NewWriter(f)
	for _, word := range words {
		w.WriteString(word + "\n")
	}
}

// Convert the space-delimited words in a text file into a string slice
func ReadAsStringArray(filename string) []string {
	// Open file, handling errors and closing at function end
	f, err := os.Open(filename)
	if err != nil {
		panic(err)
	}
	defer f.Close()

	// Split words in file at spaces
	sc := bufio.NewScanner(f)
	sc.Split(bufio.ScanWords)

	var words []string
	for sc.Scan() {
		words = append(words, Clean(sc.Text()))
	}
	return words
}

// Strip punctuation and convert to lowercase.
func Clean(s string) string {
	re := regexp.MustCompile(`[^a-zA-Z]`)
	return re.ReplaceAllString(strings.ToLower(s), "")
}

func Contains(arr []string, key string) bool {
	if strings.Compare(key, "") == 0 {
		return false
	}
	return binarySearch(0, len(arr)-1, key, arr)
}

func binarySearch(min, max int, key string, arr []string) bool {
	// Set pivot between min and max, avoiding integer overflow.
	pivot := min + (max-min)/2

	// Check sanity.
	if min > max {
		return false
	}

	switch strings.Compare(key, arr[pivot]) {
	case -1: // If key is lower lexicographically, check the preceding elements.
		return binarySearch(min, pivot-1, key, arr)
	case 1: // If key is higher lexicographically, check the following elements.
		return binarySearch(pivot+1, max, key, arr)
	}

	// The key must equal the element at the pivot at this point.
	return true
}

// Counts the number of times the keys occur in arr.
// Precondition: arr is sorted lexicographically and contains no empty strings.
func CountOccurences(keys, arr []string) int {
	var count int

	for _, key := range keys {
		if Contains(arr, key) {
			count++
		}
	}

	return count
}

func CalculateSentiment(words []string) float64 {
	negative := ReadAsStringArray("dat/negative.txt")
	positive := ReadAsStringArray("dat/positive.txt")
	stop := ReadAsStringArray("dat/stop_words.txt")

	var pos float64
	var neg float64
	var total float64

	for _, word := range words {
		if Contains(stop, word) {
			continue
		} else if Contains(positive, word) {
			pos++
		} else if Contains(negative, word) {
			neg++
		}
		total++
	}

	return (pos * 100.0 / total) - (neg * 100.0 / total)
}
