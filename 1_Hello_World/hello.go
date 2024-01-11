package main

import (
	"fmt"
)

const (
	SPANISH    = "Spanish"
	PORTUGUESE = "Portuguese"

	ENGLISH_HELLO_PREFIX    = "Hello, "
	SPANISH_HELLO_PREFIX    = "Hola, "
	PORTUGUESE_HELLO_PREFIX = "Olá, "
)

func Hello(name string, language string) string {
	if name == "" {
		name = "World"
	}
	return greetingPrefix(language) + name
}

func greetingPrefix(language string) (prefix string) {
	switch language {
	case PORTUGUESE:
		prefix = PORTUGUESE_HELLO_PREFIX
	case SPANISH:
		prefix = SPANISH_HELLO_PREFIX
	default:
		prefix = ENGLISH_HELLO_PREFIX
	}

	return
}

func main() {
	fmt.Println(Hello("", ""))
}
