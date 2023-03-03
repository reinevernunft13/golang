// Heu de realitzar una estructura condicional Switch que contempli els seguents case basant-se en un ascensor d'un centre comercial:
// 1. Planta primera i tercera, mostrara el missatge "Articles Llar".
// 2. Planta segona, mostrara el missatge "Moda infantil".
// 3. Planta cuarta, mostrara el missatge "Joguines"
// 4. Un default, mostrant un error.
package main

import (
	"fmt"
)

func main() {
	floor := 10

	// switch statement WITH initial condition
	switch floor {

	case 1, 3:
		fmt.Println("Home items")
	case 2:
		fmt.Println("Children clothes")
	case 4:
		fmt.Println("Toys")
	default:
		fmt.Println("Wrong floor. Pick an option between 1 and 4")
	}
}
