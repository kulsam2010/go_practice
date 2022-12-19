package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Sameer", "")
		want := "Hello, Sameer"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello to strangers", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("saying hello in spanish", func(t *testing.T) {
		got := Hello("Sameer", "Spanish")
		want := "Hola, Sameer"
		assertCorrectMessage(t, got, want)
	})

	t.Run("Saying Hello in french", func(t *testing.T) {
		got := Hello("Sameer", "French")
		want := "Bonjour, Sameer"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got string, want string) {
	t.Helper()
	if got != want {
		t.Errorf("Got %q, want %q", got, want)
	}
}
