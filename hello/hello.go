// https://quii.gitbook.io/learn-go-with-tests/go-fundamentals/hello-world
package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
	spanishHelloPrefix = "Hola, "
	frenchHelloPrefix  = "Bonjour, "
	spanish            = "Spanish"
	french             = "French"
	englishWorld       = "world"
)

// Hello prints a greeting message with the name in the parameter
func Hello(name string, language string) string {
	if name == "" {
		name = englishWorld
	}

	prefix := greetingPrefix(language)
	return prefix + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
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

	fmt.Println(Hello(englishWorld, ""))

}
