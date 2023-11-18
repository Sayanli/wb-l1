package main

import (
	"fmt"
	"sync"
)

func Conveyor(nums []int) {
	l := len(nums)

	ch1 := make(chan int, l)
	ch2 := make(chan int, l)

	var wg sync.WaitGroup
	wg.Add(l * 2)

	for _, v := range nums {
		ch1 <- v

		go func() {
			defer wg.Done()
			ch2 <- (2 * <-ch1)
		}()

		go func() {
			defer wg.Done()
			fmt.Println(<-ch2)
		}()
	}
	wg.Wait()
}

func main() {
	nums := []int{1, 2, 3, 4, 5}
	Conveyor(nums)
}
