package main

import (
	"fmt"
	"io"
	"os"
	"time"
)

const finalWord = "Go!"
const countdownStart = 3

type Sleeper interface {
	Sleep()
}

type ConfigurableSleeper struct {
	duration time.Duration
	sleep    func(time.Duration)
}

func (s *ConfigurableSleeper) Sleep() {
	s.sleep(s.duration)
}

func Countdown(writer io.Writer, sleeper Sleeper) {
	for i := countdownStart; i > 0; i-- {
		sleeper.Sleep()
		fmt.Fprintln(writer, i)
	}
	sleeper.Sleep()
	fmt.Fprintln(writer, finalWord)
}

func main() {
	sleeper := &ConfigurableSleeper{3 * time.Second, time.Sleep}
	Countdown(os.Stdout, sleeper)
}
