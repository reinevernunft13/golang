//Slice 
package main

import (
	"fmt"
)

func main() {
	cart := []string{"carrots", "peppers", "water", "cheese"}

	cart = append(cart, "olive oil")
	fmt.Println(cart)
}
