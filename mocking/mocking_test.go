package mocking
/*
import (
	"bytes"
	"reflect"
	"testing"
	"time"
)

////////////////
type CountDownOperationSpy struct {
	Calls []string
}

func (c *CountDownOperationSpy) Sleep() {
	c.Calls = append(c.Calls, sleep)
}

func (c *CountDownOperationSpy) Write(b []byte) (n int, err error) {
	c.Calls = append(c.Calls, write)
	return
}

/////////////

type SpyTime struct {
	durationSlept time.Duration
}

func (s *SpyTime) Sleep(duration time.Duration) {
	s.durationSlept = duration
}

//////////////////
const (
	sleep = "sleep"
	write = "write"
)

func TestCountDown(t *testing.T) {

	t.Run("print 3 to Go", func(t *testing.T) {
		buffer := &bytes.Buffer{}
		spySleepPrinter := &CountDownOperationSpy{}

		CountDown(buffer, spySleepPrinter)
		want := `3
2
1
Go`
		if buffer.String() != want {
			t.Errorf("got %s want %s", buffer, want)
		}
	})

	t.Run("sleep before every print", func(t *testing.T) {
		spySleepPrinter := &CountDownOperationSpy{}
		CountDown(spySleepPrinter, spySleepPrinter)

		want := []string{
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
			sleep,
			write,
		}

		if !reflect.DeepEqual(spySleepPrinter.Calls, want) {
			t.Errorf("got %v want %v", spySleepPrinter.Calls, want)
		}

	})
}

func TestConfigurableSleeper(t *testing.T) {
	sleepTime := time.Second * 5
	spyTime := &SpyTime{}

	sleeper := &ConfigurableSleeper{sleepTime, spyTime.Sleep}
	sleeper.Sleep()

	if spyTime.durationSlept != sleepTime {
		t.Errorf("should have slept for %v, but slept fot %v", sleepTime, spyTime.durationSlept)
	}
}
*/
