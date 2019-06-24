package counter

import "sync"

type Counter struct {
	mu    sync.Mutex
	value int
}

<<<<<<< HEAD
func NewCounter() *Counter {
	return &Counter{}
}

=======
>>>>>>> fc98406d6c82f55043fc17c8e635e99fa3e50917
func (c *Counter) Inc() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.value++

}

func (c *Counter) Value() int {
	return c.value
}
