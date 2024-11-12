package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const startCounter = 3

type Sleeper interface {
	Sleep()
}

type DefaultSleeper struct{}

func (s *DefaultSleeper) Sleep() {
	time.Sleep(1 * time.Second)
}

func Countdown(w io.Writer, s Sleeper) {
	for i := startCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
    s.Sleep()
	}
	fmt.Fprint(w, finalWord)
}

func main() {
	Countdown(os.Stdout, &DefaultSleeper{})
}
