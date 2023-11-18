package main

import (
	"fmt"
)

func swap2(a, b int) (int, int) {
	a ^= b
	b ^= a
	a ^= b
	return a, b
}
func main() {
	//самое банальное
	a, b := 1, 2
	a, b = b, a
	fmt.Println(a, b)

	c, d := 3, 4
	c, d = swap2(c, d)
	fmt.Println(c, d)
}
