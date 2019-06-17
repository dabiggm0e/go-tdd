package main

import (
	"fmt"
	"io"
	"os"
)

const (
	countDownStart = 3
	finalWord      = "Go"
)

func CountDown(writer io.Writer) {
	for i := countDownStart; i > 0; i-- {
		fmt.Fprintln(writer, i)
	}
	fmt.Fprint(writer, finalWord)

}

func main() {
	CountDown(os.Stdout)
}
