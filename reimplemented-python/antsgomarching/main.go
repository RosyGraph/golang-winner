package main

import "fmt"

func refrain(number, rhyme string) {
	for i := 0; i < 2; i++ {
		fmt.Println("The ants go marching", number, "by", number, "hurrah! Hurrah!")
	}
	fmt.Println("The ants go marching", number, "by", number+",")
	fmt.Println("The little one stops to " + rhyme + ",")
	fmt.Println("And they all go marching down")
	fmt.Println("To the ground")
	fmt.Println("To get out")
	fmt.Println("From the rain")
	fmt.Println("Boom boom boom!")
}

func main() {
	numbers := []string{
		"one",
		"two",
		"three",
		"four",
		"five",
	}
	rhymes := []string{
		"suck his thumb",
		"tie his shoe",
		"climb a tree",
		"shut the door",
		"take a dive",
	}
	for i := 0; i < 5; i++ {
		refrain(numbers[i], rhymes[i])
		fmt.Println()
	}
}
