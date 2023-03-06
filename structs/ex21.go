package main

import (
	"fmt"
)

type person struct {
	name  string
	age int
}
//embedded struct
type engineer struct {
	person
	makePlane bool
}

func main() {
	eng1 := engineer{
		person: person{
			name:  "Ruben",
			age: 26,
		},
		makePlane: true,
	}

	fmt.Println(eng1)
	fmt.Println(eng1.name, eng1.age, eng1.makePlane)
}
