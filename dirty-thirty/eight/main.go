package main

import (
	"bufio"
	"fmt"
	"os"
	"strconv"
)

func main() {
	scanner := bufio.NewScanner(os.Stdin)
	scanner.Scan()

	n, _ := strconv.Atoi(scanner.Text())
	var name string
	var number string

	phoneNums := make(map[string]string)
	for i := 0; i < n; i++ {
		scanner.Scan()
		text := scanner.Text()
		name = text[:len(text)-9]
		number = text[len(text)-8:]
		phoneNums[name] = number
	}

	for scanner.Scan() {
		query := scanner.Text()
		response, ok := phoneNums[query]
		if ok {
			fmt.Printf("%s=%s\n", query, response)
		} else {
			fmt.Println("Not found")
		}
	}
}
