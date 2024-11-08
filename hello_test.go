package main

import "testing"

func TestHello(t *testing.T) {
	t.Run("say hello to people", func(t *testing.T) {
		got := Hello("Manas")
		want := "Hello, Manas"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

  t.Run("say 'Hello, World' if empty string is provided.", func(t *testing.T) {
		got := Hello("")
		want := "Hello, World"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})
}
