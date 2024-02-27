package walk

import "testing"

func TestWalk(t *testing.T) {
	expected := "Chris"
	var got []string

	x := struct {
		Nam string
	}{expected}

	walk(x, func(input string) {
		got = append(got, input)
	})

	if got[0] != expected {
		t.Errorf("got %q, want %q", got[0], expected)
	}
}
