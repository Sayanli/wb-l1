package main

import "fmt"

type Human struct {
	name string
}

type Action struct {
	adress string
	Human
}

func main() {
	var a Action = Action{
		adress: "New York",
		Human: Human{
			name: "John",
		},
	}

	fmt.Println(a.name, a.adress)
}
