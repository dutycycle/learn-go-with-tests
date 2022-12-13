package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Kai", "english")
		want := "Hello, Kai"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", "english")
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says hola in spanish", func(t *testing.T) {
		got := Hello("Kai", "spanish")
		want := "Hola, Kai"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
