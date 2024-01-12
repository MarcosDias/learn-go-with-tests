package iteration

import (
	"fmt"
	"testing"
)

func TestRepeat(t *testing.T) {
	repeat := Repeat("a", 3)
	expected := "aaa"

	if repeat != expected {
		t.Errorf("expected %q but got %q", expected, repeat)
	}
}

func ExampleRepeat() {
	repeat := Repeat("b", 3)
	fmt.Println(repeat)
	// Output: bbb
}

func BenchmarkRepeat(b *testing.B) {
	for i := 0; i < b.N; i++ {
		Repeat("a", 5)
	}
}
