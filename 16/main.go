package main

import (
	"fmt"
)

func quicksort(s []int) []int {
	// если длина меньше 2, то дальше сортировать подмассив нет смысла
	if len(s) < 2 {
		return s
	}

	left, right := 0, len(s)-1

	// Выбираем опорный элемент
	pivot := len(s) / 2
	s[pivot], s[right] = s[right], s[pivot]

	for i := range s {
		if s[i] < s[right] {
			s[left], s[i] = s[i], s[left]
			left++
		}
	}

	s[left], s[right] = s[right], s[left]

	quicksort(s[:left])
	quicksort(s[left+1:])

	return s
}

func main() {
	arr := []int{5, 4, 3, 2, 1}
	fmt.Println(quicksort(arr))
}
