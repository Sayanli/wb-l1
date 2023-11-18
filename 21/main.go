package main

import (
	"fmt"
	"strings"
)

func reverseWords(s string) string {
	var b strings.Builder

	words := strings.Fields(s)

	for i, j := 0, len(words)-1; i < j; i, j = i+1, j-1 {
		words[i], words[j] = words[j], words[i]
	}
	for _, word := range words {
		b.WriteString(" " + word)
	}
	return strings.TrimSpace(b.String())
}

func main() {
	str := "hello world"
	fmt.Println(reverseWords(str))
}
