package walk

import (
	"reflect"
	"testing"
)

func TestWalk(t *testing.T) {
	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"struct with one string field",
			struct {
				Name string
			}{"Marcos"},
			[]string{"Marcos"},
		},
		{
			"struct with two string field",
			struct {
				Name    string
				Country string
			}{"Marcos", "Brasil"},
			[]string{"Marcos", "Brasil"},
		},
		{
			"struct with non string field",
			struct {
				Name string
				Age  int
			}{"Marcos", 33},
			[]string{"Marcos"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v, want %v", got, test.ExpectedCalls)
			}
		})
	}
}
