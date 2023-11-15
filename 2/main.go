package main

import (
	"fmt"
	"sync"
)

func main() {
	slice := []int{2, 4, 6, 8, 10}
	wgn := len(slice)
	ans := make([]int, 0, wgn)

	var wg sync.WaitGroup
	wg.Add(wgn)

	var mx sync.Mutex

	for _, v := range slice {
		go func(v int) {
			defer wg.Done()
			mx.Lock()
			ans = append(ans, v*v)
			mx.Unlock()
		}(v)
	}
	wg.Wait()

	fmt.Println(ans)
}
