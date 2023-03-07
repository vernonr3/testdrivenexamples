package main

import (
	"testing"
)

/*
map keys must be  comparable: i.e.
- Boolean
- string
- floating point types
- integer types
- pointer types
- channel types
- interface types that aren't type parameters
- struct types if all their fields are comparable
- array types
-

*/

type DictTest struct {
	dictionary Dictionary
	key        string
	definition string
	got        string
	want       string
	err        error
}

func assertStrings(t testing.TB, got, want string, err error) {
	t.Helper()
	if err != nil {
		if err.Error() != want {
			t.Errorf("got %q ", err.Error())
		}
	} else if got != want {
		t.Errorf("got %q want %q given, %q", got, want, "test")
	}
}

func TestSearch(t *testing.T) {

	dictests := []DictTest{
		{dictionary: Dictionary{"test": "this is just a test"}, key: "test", got: "", want: "this is just a test"},
		{dictionary: Dictionary{"test": "this is just a test"}, key: "unknown", got: "", want: ErrNotFound.Error()},
	}
	for _, tt := range dictests {
		tt.got, tt.err = tt.dictionary.Search(tt.key)
		assertStrings(t, tt.got, tt.want, tt.err)
	}
}

func TestAdd(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	mdictionary := Dictionary{}
	mdictionary2 := Dictionary{word: definition}
	mdictionary2.Add("test", "this is just a test")

	dictests := []DictTest{
		{dictionary: mdictionary, key: "test", definition: definition, err: nil},
		{dictionary: mdictionary2, key: "test", want: "duplicate entry", err: ErrDuplicate},
	}
	for _, tt := range dictests {
		tt.err = tt.dictionary.Add(tt.key, tt.definition)
		assertStrings(t, tt.got, tt.want, tt.err)
	}
}

func assertDefinition(t testing.TB, dictionary Dictionary, word, definition string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("should find added word:", err)
	}
	assertStrings(t, got, definition, err)
}

func TestUpdate(t *testing.T) {
	word := "test"
	definition := "this is just a test"
	dictionary := Dictionary{word: definition}
	newDefinition := "new definition"

	dictionary.Update(word, newDefinition)

	assertDefinition(t, dictionary, word, newDefinition)
}
