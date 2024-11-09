package main

import "fmt"

const (
	spanish            = "Spanish"
	french             = "French"
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
)

func Hello(name string, lang string) string {
	if name == "" {
		name = "World"
	}
	prefix := englishHelloPrefix
	if lang == spanish {
		prefix = spanishHelloPrefix
	} else if lang == french {
		prefix = frenchHelloPrefix
	}
	return prefix + name
}

func main() {
	fmt.Println(Hello("world", ""))
}
