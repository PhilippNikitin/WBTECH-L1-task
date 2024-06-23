package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    sync.Mutex
}

func (c *counter) increment() {
	c.mu.Lock()
	defer c.mu.Unlock()
	c.count++
}

func (c *counter) value() int {
	c.mu.Lock()
	defer c.mu.Unlock()
	return c.count
}

func main() {
	const numGoroutines = 500

	var wg sync.WaitGroup

	newCounter := counter{}

	wg.Add(numGoroutines)
	for i := 0; i < numGoroutines; i++ {
		go func() {
			defer wg.Done()
			newCounter.increment()
		}()
	}

	wg.Wait()
	fmt.Printf("Final counter value: %d\n", newCounter.value())
}
