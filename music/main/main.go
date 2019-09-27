package main

import "fmt"

func main() {
	fmt.Println("Enter the notes in a C major triad separated by spaces")

	music.TriadFromString(fmt.Scanln())

}
