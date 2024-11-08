package main

import "testing"

func TestHello(t *testing.T) {
  got := Hello("Manas")
  want := "Hello, Manas"

  if got != want {
    t.Errorf("got %q want %q", got, want)
  }
}
