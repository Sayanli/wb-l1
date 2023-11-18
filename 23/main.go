package main

import (
	"fmt"
)

func remove(slice []int, s int) []int {
	return append(slice[:s], slice[s+1:]...)
}

func main() {
	slice := []int{1, 2, 3, 4, 5}
	ns := remove(slice, 2)

	fmt.Println(ns)
}
