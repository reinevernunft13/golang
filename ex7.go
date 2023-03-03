package main

import "fmt"

func main() {
	// hardcoded input
	age := 1
	myDay := "Thursday"

	if (age >= 18) || ((myDay == "Thursday") && age >= 16) {
		fmt.Println("You may come in.")

	} else if age == 0 {

		fmt.Println("Invalid age. You must enter an age")

	} else {

		fmt.Println("You may not come in")

	}

}
