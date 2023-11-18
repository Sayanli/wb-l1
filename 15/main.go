package main

import (
	"fmt"
	"strings"
)

func createHugeString(size int) string {
	// для создание новой строки лучше использовать string builder
	var b strings.Builder

	for i := 0; i < size; i++ {
		fmt.Fprint(&b, "bober ")
	}

	return b.String()
}

func newSomeFunc() string {
	// лучше создадим локальную переменную
	var justString string

	v := createHugeString(1 << 10)

	//переведем нашу строку в слайс рун, т.к. символы могут занимать больше 1 байта
	justString = string([]rune(v)[:100])

	return justString
}

func main() {
	fmt.Println(newSomeFunc())
}
