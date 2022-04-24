package main

import "testing"

func TestHello(t *testing.T) {
	assertCorrectMessage := func(t testing.TB, got, want string) {
		/*
		   t.Helper() is needed to tell the test suite that this method is a helper.
		   By doing this when it fails the line number reported will be in our function
		   call rather than inside our test helper.
		*/
		t.Helper()
		if got != want {
			t.Errorf("got %q, want %q", got, want)
		}
	}
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Filip", "English")
		want := "Hello, Filip"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' when an empy string is supplied", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in Spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"
		assertCorrectMessage(t, got, want)
	})

	t.Run("in French", func(t *testing.T) {
		got := Hello("Jacques", "French")
		want := "Bonjour, Jacques"
		assertCorrectMessage(t, got, want)
	})
}
