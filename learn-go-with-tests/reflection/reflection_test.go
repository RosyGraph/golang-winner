package main

import (
	"reflect"
	"testing"
)

type Person struct {
	Name    string
	Profile Profile
}

type Profile struct {
	Age  int
	City string
}

func TestWalk(t *testing.T) {
	t.Run("with maps", func(t *testing.T) {
		aMap := map[string]string{
			"Foo": "Bar",
			"Baz": "Boz",
		}

		var got []string
		walk(aMap, func(input string) {
			got = append(got, input)
		})

		assertContains(t, got, "Bar")
		assertContains(t, got, "Boz")
	})

	cases := []struct {
		Name          string
		Input         interface{}
		ExpectedCalls []string
	}{
		{
			"Struct with one string field",
			struct {
				Name string
			}{"Chris"},
			[]string{"Chris"},
		},
		{
			"Struct with two string fields",
			struct {
				First string
				Last  string
			}{"Chris", "Stevenson"},
			[]string{"Chris", "Stevenson"},
		},
		{
			"Struct with one string and one int field",
			struct {
				Name string
				Age  int
			}{"Chris", 27},
			[]string{"Chris"},
		},
		{
			"Nested fields",
			Person{
				"Chris",
				Profile{27, "Salt Lake City"},
			},
			[]string{"Chris", "Salt Lake City"},
		},
		{
			"Pointers to things",
			&Person{
				"Chris",
				Profile{27, "Salt Lake City"},
			},
			[]string{"Chris", "Salt Lake City"},
		},
		{
			"Slices",
			[]Profile{
				{27, "Salt Lake City"},
				{28, "New York City"},
			},
			[]string{"Salt Lake City", "New York City"},
		},
		{
			"Arrays",
			[2]Profile{
				{27, "Salt Lake City"},
				{28, "New York City"},
			},
			[]string{"Salt Lake City", "New York City"},
		},
	}

	for _, test := range cases {
		t.Run(test.Name, func(t *testing.T) {
			var got []string
			walk(test.Input, func(input string) {
				got = append(got, input)
			})

			if !reflect.DeepEqual(got, test.ExpectedCalls) {
				t.Errorf("got %v want %v", got, test.ExpectedCalls)
			}
		})
	}
}

func assertContains(t *testing.T, haystack []string, needle string) {
	contains := false
	for _, x := range haystack {
		if x == needle {
			contains = true
			break
		}
	}
	if !contains {
		t.Errorf("want %+v to contain %q but it didn't", haystack, needle)
	}
}
