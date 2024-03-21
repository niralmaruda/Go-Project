package main

import "testing"

func TestSearch(t *testing.T) {
	dictionary := Dictionary{"test": "this is just a test"}

	t.Run("known word", func(t *testing.T) {
		got, _ := dictionary.Search("test")
		want := "this is just a test"

		assertStrings(t, got, want)
	})
	t.Run("unknown word", func(t *testing.T) {
		_, got := dictionary.Search("unknown")

		assertError(t, got, ErrorNotFound)
	})
}

func TestAdd(t *testing.T) {
	t.Run("new word", func(t *testing.T) {
		dictionary := Dictionary{}
		word := "test"
		defination := "this is a test"
		dictionary.Add(word, defination)

		assertDefination(t, dictionary, word, defination)
	})
	t.Run("existing word", func(t *testing.T) {
		word := "test"
		defination := "this is a test"
		dictionary := Dictionary{word: defination}
		err := dictionary.Add(word, "new test")

		assertError(t, err, ErrorWordExists)
		assertDefination(t, dictionary, word, defination)
	})
}

func assertDefination(t testing.TB, dictionary Dictionary, word, defination string) {
	t.Helper()

	got, err := dictionary.Search(word)
	if err != nil {
		t.Fatal("Should find added word:", err)
	}

	assertStrings(t, got, defination)
}

func assertStrings(t testing.TB, got, want string) {
	t.Helper()

	if got != want {
		t.Errorf("got %q want %q given %q", got, want, "test")
	}
}

func assertError(t testing.TB, got, want error) {
	t.Helper()

	if got != want {
		t.Errorf("got error %q want %q", got, want)
	}
}
