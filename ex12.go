package main

import (
	"fmt"
)

func main() {
	students := []string{"Anna", "Pep", "Iria", "Oscar"}
	i := 2
	for {
		switch {
		case i < len(students):
			fmt.Println(students[i]) // prints first element
			i++
		default:
			break
		}
	}
}
