package main

import (
	"fmt"
	"strings"
)

func CheckUnique(s string) bool {
	ls := strings.ToLower(s)
	m := make(map[rune]int)
	flag := true
	sr := []rune(ls)
	for _, v := range sr {
		if _, ok := m[v]; !ok {
			m[v] = 1
		} else {

			m[v] += 1
			flag = false
		}
	}
	return flag
}

func main() {
	fmt.Println(CheckUnique("Hello"))
	fmt.Println(CheckUnique("World"))
}
