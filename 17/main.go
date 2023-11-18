package main

import (
	"fmt"
)

func binarySearch(s []int, target int) int {
	l := 0
	r := len(s) - 1
	ans := -1
	for l <= r {
		m := (l + r) / 2
		if s[m] == target {
			ans = m
			break
		} else if s[m] < target {
			l = m + 1
		} else {
			r = m - 1
		}
	}
	return ans
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(binarySearch(arr, 3))
}
