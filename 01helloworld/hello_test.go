package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people in English if no languag is provided", func(t *testing.T) {
		got := Hello("Manas", "")
		want := "Hello, Manas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in English if incorrect languag is provided", func(t *testing.T) {
		got := Hello("Manas", "abcd")
		want := "Hello, Manas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say 'Hello, World' if empty string is provided.", func(t *testing.T) {
		got := Hello("", "English")
		want := "Hello, World"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in English", func(t *testing.T) {
		got := Hello("Manas", "English")
		want := "Hello, Manas"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in Spanish", func(t *testing.T) {
		got := Hello("Akanksha", "Spanish")
		want := "Hola, Akanksha"
		assertCorrectMessage(t, got, want)
	})

	t.Run("say hello to people in French", func(t *testing.T) {
		got := Hello("Akanksha", "French")
		want := "Bonjour, Akanksha"
		assertCorrectMessage(t, got, want)
	})
}

func assertCorrectMessage[T comparable](t testing.TB, got, want T) {
	t.Helper()
	if got != want {
		t.Errorf("got '%v' want '%v'", got, want)
	}
}
