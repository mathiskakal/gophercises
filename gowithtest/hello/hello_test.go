package main

import "testing"

func TestHello(t *testing.T) {
	// Testing for a name
	t.Run("saying hello to people", func(t *testing.T) {
		got := Hello("Chris", "")
		want := "Hello, Chris"

		assertCorrectMessage(t, got, want)
	})

	// Testing for an empty name
	t.Run("say 'Hello, World' when an empty string is supplied", func(t *testing.T) {
		got := Hello("", "")
		want := "Hello, World"

		assertCorrectMessage(t, got, want)
	})

	// Testing for Spanish
	t.Run("in spanish", func(t *testing.T) {
		got := Hello("Elodie", "Spanish")
		want := "Hola, Elodie"

		assertCorrectMessage(t, got, want)
	})
	// Testing for French
	t.Run("in french", func(t *testing.T) {
		got := Hello("Marc", "French")
		want := "Bonjour, Marc"

		assertCorrectMessage(t, got, want)
	})
	// Testing for German
	t.Run("in french", func(t *testing.T) {
		got := Hello("Adolf", "German")
		want := "Halo, Adolf"

		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage(t testing.TB, got, want string) {
	t.Helper()
	if got != want {
		t.Errorf("got %q want %q", got, want)
	}
}

/*

 */
