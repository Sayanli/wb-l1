package main

import (
	"fmt"
	"time"
)

func Sleep(d time.Duration) {
	<-time.After(d)
}

func main() {
	Sleep(4 * time.Second)
	fmt.Println("stop")
}
