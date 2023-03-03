package main

import (
	"fmt"
)

type year int

var currentYear year

func main() {
	currentYear = 2021
	fmt.Println(currentYear)      //2021
	fmt.Printf("%T", currentYear) //returns main.year -> indicates it's an alias
}
