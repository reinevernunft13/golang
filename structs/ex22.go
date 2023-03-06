package main

import (
	"fmt"
)

// anonymous structs -> these are useful when we don't need to reuse code
func main() {
	item1 := struct {
		name         string
		units, price int
	}{
		name:  "rosé wine",
		units: 3,
		price: 25,
	}

	item2 := struct {
		name         string
		units, price int
	}{
		name:  "cheese",
		units: 1,
		price: 63,
	}

	fmt.Println(item1)
	fmt.Println(item2)
}
