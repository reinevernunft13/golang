package main

import (
	"fmt"
)

// Displays even numbers from 0 to 20
func main() {
	i := 0
	for i <= 20 {
		fmt.Println(i)
		i += 2 //more optimal in terms of memory than conditional: if i%2 == 0 {fmt.Println(i)}
	}
}
