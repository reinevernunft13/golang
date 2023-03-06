package main

import (
	"fmt"
)
//map through all students and their grades
func main() {
	grades := map[string]int{
		"Pedro":  2,
		"Miriam": 8,
		"Montse": 6,//last element MUST contain a comma, too
	}

	fmt.Println(grades)
}
