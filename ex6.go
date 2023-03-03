package main

import (
	"fmt"
)

//bit shifting

func main() {
	a := 4
	fmt.Printf("%d\t\t%b\n", a, a) //4    100
	// \t imprimirá una tabulació i %b mostrara en binari
	b := a << 1
	fmt.Printf("%d\t\t%b\n", b, b) //8    1000
}
