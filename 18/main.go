package main

import (
	"fmt"
	"sync"
)

type counter struct {
	count int
	mu    sync.Mutex
}

func (c *counter) add(wg *sync.WaitGroup) {
	defer wg.Done()
	c.mu.Lock()
	c.count++
	c.mu.Unlock()
}

func main() {
	cnt := new(counter)
	var wg sync.WaitGroup
	for i := 0; i < 5; i++ {
		wg.Add(1)
		go cnt.add(&wg)
	}
	wg.Wait()
	fmt.Println(cnt.count)
}
