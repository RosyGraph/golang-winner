package main

import "fmt"

const englishHelloPrefix = "Hello, "
const spanishHelloPrefix = "Hola, "

func Hello(name, language string) string {
	if name == "" {
		name = "World"
	}
	if language == "Spanish" {
		return spanishHelloPrefix + name + "."
	}
	if language == "French" {
		return "Bonjour, " + name + "."
	}
	return englishHelloPrefix + name + "."
}

func main() {
	fmt.Println(Hello("Chris", ""))
}
