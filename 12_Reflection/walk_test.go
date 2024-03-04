package walk

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age     int
	Country string
}

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
		{
			"nested fields",
			Person{
				"Marcos",
				Profile{33, "Brasil"},
			},
			[]string{"Marcos", "Brasil"},
		},
		{
			"pointers to things",
			&Person{
				"Marcos",
				Profile{33, "Brasil"},
			},
			[]string{"Marcos", "Brasil"},
		},
		{
			"slices",
			[]Profile{
				{33, "Brasil"},
				{35, "Portugal"},
			},
			[]string{"Brasil", "Portugal"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "Brasil"},
				{35, "Portugal"},
			},
			[]string{"Brasil", "Portugal"},
		},
		{
			"maps",
			map[string]string{
				"Cow":   "Moo",
				"Sheep": "Baa",
			},
			[]string{"Moo", "Baa"},
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
