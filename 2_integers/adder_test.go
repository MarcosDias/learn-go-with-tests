package integers

import (
	"fmt"
	"testing"
)

func TestAdd(t *testing.T) {
	sum := Add(2, 2)
	expected := 4

	if sum != expected {
		t.Errorf("expected '%d' but got '%d'", expected, sum)
	}
}

// run:
// go test -v
func ExampleAdd() {
	sum := Add(1, 6)
	fmt.Println(sum)
	// Output: 7
}
