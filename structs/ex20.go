package main

import (
	"fmt"
)

type person struct {
	name string
	age  int
}

func main() {
	p1 := person{
		name: "Gerard",
		age:  22,
	}

	p2 := person{
		name: "Ona",
		age:  18,
	}

	fmt.Println(p1)
	fmt.Println(p2)
	fmt.Println(p1.name)
	fmt.Println(p2.name)
}
