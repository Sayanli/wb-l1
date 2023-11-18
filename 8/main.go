package main

import "fmt"

func setBit(num int64, i int, flag bool) int64 {
	if flag {
		return num | (1 << i)
	}
	return num &^ (1 << i)
}

func main() {
	num := int64(500)
	newn := setBit(num, 1, false)
	fmt.Println(newn)
}
