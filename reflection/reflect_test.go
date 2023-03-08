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

type WalkTest struct {
	Name          string
	Input         interface{}
	ExpectedCalls []string
}

func TestWalk(t *testing.T) {

	var got []string

	WalkTests := []WalkTest{
		{Name: "struct with one string field",
			Input: struct {
				Name string
			}{"Chris"}, ExpectedCalls: []string{"Chris"},
		},
		{Name: "struct with two string fields",
			Input: struct {
				Name string
				City string
			}{"Chris", "London"}, ExpectedCalls: []string{"Chris", "London"},
		},
		{Name: "nested fields",
			Input: Person{
				"Chris",
				Profile{Age: 33, City: "London"},
			}, ExpectedCalls: []string{"Chris", "London"},
		},
		{Name: "nested fields with pointer",
			Input: &Person{
				"Chris",
				Profile{Age: 33, City: "London"},
			}, ExpectedCalls: []string{"Chris", "London"},
		},
		{Name: "nested fields with slices",
			Input: []Profile{
				{Age: 33, City: "London"},
				{Age: 34, City: "Reyjavik"},
			}, ExpectedCalls: []string{"London", "Reyjavik"},
		},
		{
			"arrays",
			[2]Profile{
				{33, "London"},
				{34, "Reykjavík"},
			},
			[]string{"London", "Reykjavík"},
		},
	}
	for _, tt := range WalkTests {
		walk(tt.Input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, tt.ExpectedCalls) {
			t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
		}
		got = nil
	}
}

func TestWalkMap(t *testing.T) {

	var got []string
	// watch this test as map extraction isn't gauranteed to produce the same order - might just have to assert in succession that result contains each of Bar & Boz
	WalkTests := []WalkTest{
		{
			"maps",
			map[string]string{
				"Foo": "Bar",
				"Baz": "Boz",
			},
			[]string{"Bar", "Boz"},
		},
	}
	for _, tt := range WalkTests {
		walk(tt.Input, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, tt.ExpectedCalls) {
			t.Errorf("got %v, want %v", got, tt.ExpectedCalls)
		}
		got = nil
	}

}

func TestWalkChannelMap(t *testing.T) {
	t.Run("with channels", func(t *testing.T) {
		aChannel := make(chan Profile)

		go func() {
			aChannel <- Profile{33, "Berlin"}
			aChannel <- Profile{34, "Katowice"}
			close(aChannel)
		}()

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aChannel, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
	t.Run("with function", func(t *testing.T) {
		aFunction := func() (Profile, Profile) {
			return Profile{33, "Berlin"}, Profile{34, "Katowice"}
		}

		var got []string
		want := []string{"Berlin", "Katowice"}

		walk(aFunction, func(input string) {
			got = append(got, input)
		})

		if !reflect.DeepEqual(got, want) {
			t.Errorf("got %v, want %v", got, want)
		}
	})
}
