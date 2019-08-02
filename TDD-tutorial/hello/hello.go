package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	frenchHelloPrefix  = "Bonjour, "
	spanishHelloPrefix = "Hola, "
	spanish            = "Spanish"
	french             = "French"
)

func Hello(name, lang string) string {
	if name == "" {
		name = "world"
	}
	return greetingPrefix(lang) + name + "!"
}

func greetingPrefix(lang string) (prefix string) {
	switch lang {
	case french:
		prefix = frenchHelloPrefix
	case spanish:
		prefix = spanishHelloPrefix
	default:
		prefix = englishHelloPrefix
	}
	return
}

func main() {
	fmt.Println(Hello("world", ""))
}
