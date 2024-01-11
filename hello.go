package main

import "fmt"

const SPANISH = "Spanish"
const ENGLISH_HELLO_PREFIX = "Hello, "
const SPANISH_HELLO_PREFIX = "Hola, "

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}

	if language == SPANISH {
		return SPANISH_HELLO_PREFIX + name
	}

	return ENGLISH_HELLO_PREFIX + name
}

func main() {
	fmt.Println(Hello("", ""))
}
