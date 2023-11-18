package main

import "fmt"

func checkType(i interface{}) {
	switch v := i.(type) {
	case int:
		fmt.Printf("i is an int: %d\n", v)
	case bool:
		fmt.Printf("i is a bool: %t\n", v)
	case string:
		fmt.Printf("i is a string: %s\n", v)
	case float64:
		fmt.Printf("i is a float64: %v\n", v)
	case rune:
		fmt.Printf("i is a rune: %c\n", v)
	case chan interface{}:
		fmt.Println("i is a channel")
	case nil:
		fmt.Println("i is nil")
	default:
		fmt.Println("unknown type")
	}
}

func main() {

	checkType(42)
	checkType("hello")
	checkType(true)
	checkType(make(chan interface{}))
	checkType(3.14)
	checkType('a')
	checkType(nil)

}
