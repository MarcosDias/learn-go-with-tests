package main

import "testing"

func TestSum(t *testing.T) {
	t.Run("Collection of 5 numbers", func(t *testing.T) {
		numbers := []int{1, 2, 3, 4, 5}
		got := Sum(numbers)
		want := 15

		if got != want {
			t.Errorf("got %d want %d give %v", got, want, numbers)
		}
	})
}

// Coverage tool:
// https://go.dev/blog/cover
// go test -cover
