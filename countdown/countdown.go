package main

import (
	"fmt"
	"io"
	"os"
)

const finalWord = "Go!"
const startCounter = 3

func Countdown(w io.Writer) {
	for i := startCounter; i > 0; i-- {
		fmt.Fprintln(w, i)
	}
	fmt.Fprint(w,finalWord)
}

func main() {
	Countdown(os.Stdout)
}
