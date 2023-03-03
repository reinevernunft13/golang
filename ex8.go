// Create a switch statement with the following conditions:
//1. energy is 100% charged or has a booster -> prints "Fully charged"
//2. if energy level is lower than 100% -> adds 20 units and prints current energy level
//3. A default clause displaying an error message.

package main

import "fmt"

func main() {

	energy := 80
	booster := false

	//switch statement without an initial condition
	switch {
	case energy == 100 || booster == true:
		fmt.Println("Fully charged")
	case energy < 100:
		energy += 20
		fmt.Println("Your current level of ", energy, " after having charged 20 units")
	default:
		fmt.Println("Wrong enery level")
	}

}
