package main

import (
	"bytes"
	"testing"
)

func TestGreet(t *testing.T) {
	buffer := bytes.Buffer{}
	Greet(&buffer, "Marcos")

	got := buffer.String()
	want := "Hello, Marcos"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
