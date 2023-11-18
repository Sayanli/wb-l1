package main

import (
	"fmt"
)

type s1 struct {
	Str string
}

type s2 struct {
	str []rune
}

func Adapter(str []rune) string {
	return string(str)
}

func main() {
	f1 := s1{Str: "hello"}
	f2 := s2{str: []rune("world")}
	fmt.Println(f1.Str + " " + Adapter(f2.str))
}
