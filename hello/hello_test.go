package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("says hello to people", func(t *testing.T) {
		got := Hello("Kai", ENGLISH)
		want := "Hello, Kai"
		assertCorrectMessage(t, got, want)
	})

	t.Run("empty string defaults to 'world'", func(t *testing.T) {
		got := Hello("", ENGLISH)
		want := "Hello, world"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says hola in spanish", func(t *testing.T) {
		got := Hello("Kai", SPANISH)
		want := "Hola, Kai"
		assertCorrectMessage(t, got, want)
	})

	t.Run("says bonjour in french", func(t *testing.T) {
		got := Hello("Kai", "french")
		want := "Bonjour, Kai"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
