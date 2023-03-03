package main

import (
	"fmt"
)

func main() {
	//Displays every other number comprised between 0 and 100 in descending order
	for i := 100; i >= 0; i -= 2 {
		fmt.Println(i)
	}
}
