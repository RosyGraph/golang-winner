package main

import (
	"bufio"
	"fmt"
	"os"

	"github.com/RosyGraph/music/music"
)

func main() {
	reader := bufio.NewReader(os.Stdin)

	fmt.Println("Enter the notes in a C major triad separated by spaces")

	text, _ := reader.ReadString('\n')

	triad := music.TriadFromString(text)

	if triad.Equals(music.TriadFromString("C E G")) {
		fmt.Println("\nCorrect!")
	} else {
		fmt.Println("\nIncorrect.")
	}
}
