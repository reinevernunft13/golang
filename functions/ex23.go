package main

import "fmt"

func main() {
	salut, valid := saludar("Anna")
	fmt.Println(salut, "How are you?", valid)
}

func saludar(name string) (string, bool) {
	salutacio := fmt.Sprint("How are you? ", name)
	validacio := false
	if name == "Anna" {
		validacio = true
	}

	return salutacio, validacio

}
