package main

import (
	"fmt"
)

type student struct {
	name  string
	lName string
}
type enrolled struct {
	student
	subject string
}

// Method example
func (e enrolled) introduce() { //D’aquesta manera afegim aquest metode al struct matriculat
	fmt.Println("My name is ", e.name, e.lName, " and I am learning ", e.subject)
}
func main() {
	student1 := enrolled{
		student: student{
			name:  "Jon",
			lName: "Ruiz",
		},
		subject: "Golang",
	}
	fmt.Println(student1) //{{Jon Ruiz} Go}
	student1.introduce()  //My name is  Jon Ruiz  and I am learning  Golang
}
