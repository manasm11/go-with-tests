package main

import (
	"bytes"
	"reflect"
	"testing"
)

const (
  sleep = "sleep"
  write = "write"
)

type SpySleepPrinter struct {
	Calls []string
}

func (ss *SpySleepPrinter) Sleep() {
	ss.Calls = append(ss.Calls, sleep)
}

func (ss *SpySleepPrinter) Write([]byte) (int, error) {
  ss.Calls = append(ss.Calls, write)
  return 0, nil
}

func TestCountdown(t *testing.T) {
	t.Run("prints 3 to Go!", func(t *testing.T) {
		buffer := &bytes.Buffer{}

		Countdown(buffer, &SpySleepPrinter{})

		got := buffer.String()
		want := "3\n2\n1\nGo!"

		if got != want {
			t.Errorf("got %q want %q", got, want)
		}
	})

  t.Run("sleep before every print", func(t *testing.T) {
    spySleepPrinter := &SpySleepPrinter{}

    Countdown(spySleepPrinter, spySleepPrinter)
    
    want := []string{write, sleep, write, sleep, write, sleep, write}

    if !reflect.DeepEqual(want, spySleepPrinter.Calls) {
      t.Errorf("got calls %v want %v", spySleepPrinter.Calls, want)
    }
  })
}
