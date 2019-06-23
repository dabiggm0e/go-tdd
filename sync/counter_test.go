package counter

import (
	"sync"
	"testing"
)

func TestCounter(t *testing.T) {
	t.Run("Incrementing the counter 3 times leaves it at 3", func(t *testing.T) {
		counter := Counter{}
		counter.Inc()
		counter.Inc()
		counter.Inc()

	})

	t.Run("it runs safely concurrently", func(t *testing.T) {
		wantedCounter := 1000

		counter := Counter{}
		var wg sync.WaitGroup
		wg.Add(wantedCounter)

		for i := 0; i < wantedCounter; i++ {

			go func(w *sync.WaitGroup) {
				counter.Inc()
				wg.Done()
			}(&wg)
		}

		wg.Wait()

		assertCounter(t, counter, wantedCounter)
	})
}

func assertCounter(t *testing.T, counter Counter, want int) {
	t.Helper()
	if counter.Value() != want {
		t.Errorf("got %d want %d", counter.Value(), want)
	}
}
