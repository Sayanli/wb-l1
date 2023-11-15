package main

import (
	"fmt"
	"time"
)

func sender(ch chan int) {
	for i := 0; i < 10; i++ {
		ch <- i
		time.Sleep(1 * time.Second)
	}
	close(ch)
}

func main() {
	beginTime := time.Now()

	ch := make(chan int)
	wortTime := 7

	go sender(ch)

	timer := time.NewTimer(time.Duration(wortTime) * time.Second)
	defer timer.Stop()

	for {
		select {
		case <-timer.C:
			fmt.Println("timeout")
			endTime := time.Now()
			elapsedTime := endTime.Sub(beginTime)
			fmt.Printf("Время работы: %s секунд\n", elapsedTime)
			return
		case v, ok := <-ch:
			if !ok {
				fmt.Println("канал закрыт")
				return
			}
			fmt.Println(v)
		}
	}
}
