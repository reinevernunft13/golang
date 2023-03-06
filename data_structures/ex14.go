// Modify the array by replacing the element peppers with cucumber
package main

import (
	"fmt"
)

func main() {
	cart := []string{"carrots", "peppers", "water", "cheese"}

	cart[1] = "cucumber"

	fmt.Println(cart)
}
