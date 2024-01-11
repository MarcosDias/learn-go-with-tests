package main

import "fmt"

const ENGLISH_HELLO_PREFIX = "Hello, "

func Hello(name string) string {
	if name == "" {
		name = "World"
	}
	return ENGLISH_HELLO_PREFIX + name
}

func main() {
	fmt.Println(Hello(""))
}
