package main

import (
	"fmt"
)

var city string = "Barcelona"

func main() {
	myString := fmt.Sprint("The city of ", city, " is in Catalonia")
	fmt.Println(myString) //La ciutat de Barcelona està a Catalunya

}
