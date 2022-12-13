package main

import "testing"

func TestHello(t *testing.T) {
	got := Hello("Kai")
	want := "Hello, Kai"

	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}
