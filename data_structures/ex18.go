package main

import (
	"fmt"
)

func main() {
	//9 len , 10 for cap
	x := make([]int, 9, 10)

	x[2] = 8
	x[6] = 3

	fmt.Println(x)
	//display x's length
	fmt.Println(len(x))
	//displays x's capacity
	fmt.Println(cap(x))

	x = append(x, 10)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

	//x[11] = 11 //This will throw an error
	x = append(x, 12)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x)) //duplicates memory space

	x = append(x, 4)
	fmt.Println(x)
	fmt.Println(len(x))
	fmt.Println(cap(x))

}
