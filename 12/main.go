package main

import (
	"fmt"
)

func makeSet(s []string) map[string]struct{} {

	set := make(map[string]struct{})
	for _, v := range s {
		set[v] = struct{}{}
	}
	return set
}
func main() {
	set := []string{"cat", "cat", "dog", "cat", "tree"}
	fmt.Println(makeSet(set))
}
