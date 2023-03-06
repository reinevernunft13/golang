//Displays array data through a for range

package main
import (
	"fmt"
)

func main() {
	cart := []string{"carrots", "peppers", "water", "cheese"}

	for i, v := range cart {
		fmt.Println(i, "->", v)

	}
}
