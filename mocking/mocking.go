package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const (
	countDownStart = 3
	finalWord      = "Go"
	sleepDuration  = 1
)

////////////////////////
type Sleeper interface {
	Sleep()
}

/////////////////
type DefaultSleeper struct{}

func (d *DefaultSleeper) Sleep() {
	time.Sleep(time.Second * sleepDuration)
}

//////////////////
type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (c *ConfigurableSleeper) Sleep() {
	c.sleep(c.duration)
}

///////////////////

func CountDown(writer io.Writer, s Sleeper) {
	for i := countDownStart; i > 0; i-- {
		s.Sleep()
		fmt.Fprintln(writer, i)
	}

	s.Sleep()
	fmt.Fprint(writer, finalWord)
}

func main() {
	//CountDown(os.Stdout, &DefaultSleeper{})
	//sleeper := &ConfigurableSleeper{1 * time.Duration, time.Sleep}         //// FIXME:
	//CountDown(os.Stdout, &ConfigurableSleeper{1 * time.Duration, sleeper}) //// FIXME:
}
