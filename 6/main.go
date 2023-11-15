package main

import (
	"context"
	"fmt"
	"time"
)

// первый способ через канал для остановки
func worker1(quit <-chan bool) {
	for {
		select {
		default:
			fmt.Println("work")
			time.Sleep(1 * time.Second)
		case <-quit:
			return
		}
	}
}

// второй способ через контекст
func worker2(ctx context.Context) {
	for {
		select {
		default:
			fmt.Println("work")
			time.Sleep(1 * time.Second)
		case <-ctx.Done():
			return
		}
	}
}

func main() {
	//можно завести отдельный канал для остановки
	quit := make(chan bool)
	ctx := context.Background()
	go worker1(quit)
	go worker2(ctx)
	time.Sleep(5 * time.Second)
	quit <- true
	fmt.Println("stop with channel")
	time.Sleep(5 * time.Second)
	ctx.Done()
	fmt.Println("stop with context")
}
