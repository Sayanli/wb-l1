package main

import (
	"fmt"
	"slices"
)

func Intersection(set1, set2 []int) []int {
	var intersection []int
	for _, v := range set1 {
		if slices.Contains(set2, v) {
			intersection = append(intersection, v)
		}
	}
	return intersection
}

func main() {
	set1 := []int{1, 2, 3, 4, 5}
	set2 := []int{3, 4, 5, 6, 7}
	fmt.Println(Intersection(set1, set2))
}
