package main

import (
	"fmt"
	"sync"
)

func main() {
	var rwm sync.RWMutex
	//так же можно использовать sync.Map
	mp := make(map[int]int, 0)
	var wg sync.WaitGroup
	// запись
	for i := 0; i < 10; i++ {
		wg.Add(1)
		go func(m *sync.RWMutex, key, value int) {
			defer wg.Done()
			m.Lock()
			mp[key] = value
			m.Unlock()
		}(&rwm, i, i*i)
	}

	wg.Wait()

	// чтение
	for key, value := range mp {
		wg.Add(1)
		go func(m *sync.RWMutex, key, value int) {
			defer wg.Done()
			m.Lock()
			fmt.Printf("Key: %d, Value: %d\n", key, value)
			m.Unlock()
		}(&rwm, key, value)
	}

	wg.Wait()
}
