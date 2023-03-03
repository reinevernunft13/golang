package main

import (
	"fmt"
)

func main() {

	c := "*"
	for i := 0; i < 10; i++ {
		fmt.Println() // no ha de imprimir res al principi
		for j := 0; j < 10; j++ {
			fmt.Print(c) // no fa salt de línea
		}
	}
}
