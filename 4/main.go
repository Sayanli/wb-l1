package main

import (
	"flag"
	"fmt"
	"os"
	"os/signal"
	"sync"
	"syscall"
	"time"
)

func worker(data <-chan int, wg *sync.WaitGroup, id int) {
	defer wg.Done()

	for data := range data {
		fmt.Printf("Worker %d received %d\n", id, data)
		time.Sleep(1 * time.Second)
	}
}
func main() {
	numWorkers := flag.Int("w", 5, "number of workers")
	flag.Parse()

	data := make(chan int)
	var wg sync.WaitGroup

	for i := 0; i < *numWorkers; i++ {
		wg.Add(1)
		go worker(data, &wg, i)
	}

	c := make(chan os.Signal, 1)
	signal.Notify(c, os.Interrupt, syscall.SIGINT, syscall.SIGTERM)

	go func() {
		<-c
		//если получили сигнал Ctrl+C, закроем канал data
		close(data)
	}()

	for i := 0; i < 20; i++ {
		data <- i
	}

	close(data)
	wg.Wait()

	fmt.Println("main done")
}
