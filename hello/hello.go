package main

import "fmt"

const (
	englishHelloPrefix = "Hello, "
)

// Hello prints a greeting message with the name in the parameter
func Hello(name string) string {
	return englishHelloPrefix + name
}

func main() {

	fmt.Println(Hello("world"))

}
