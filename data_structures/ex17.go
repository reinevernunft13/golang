package main

import (
	"fmt"
)
//Join both data structures to form one
func main() {
	workers := []string{"Josep", "Cristina", "Helena", "Robert"}

	newHires := []string{"Ivan", "Estel", "Aitana"}

	workers = append(workers, newHires...)

	fmt.Println(workers)
}
