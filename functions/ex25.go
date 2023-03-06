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

// Exemple de mètode
func (s student) introduce() { //D’aquesta manera afegim aquest mètode al struct matriculat
	fmt.Println("El meu nom és ", s.name, s.lName, " i estic aprenent molt.")
}
func (e enrolled) introduce() { //D’aquesta manera afegim aquest mètode al struct matriculat
	fmt.Println("El meu nom és ", e.name, e.lName, " i estic impartint ", e.subject)
}

type user interface {
	introduce()
}

func room(u user) {
	fmt.Println("Estic dins de l'aula", u)
}
func main() {
	student1 := enrolled{
		student: student{
			name:  "Jon",
			lName: "Ruiz",
		},
		subject: "Go",
	}

	pr1 := student{
		name:  "Joan",
		lName: "Riera",
	}

	fmt.Println(student1) //{{Jon Ruiz} Go}
	student1.introduce()  //El meu nom és Jon Ruiz i estic aprenent Go
	pr1.introduce()       //El meu nom és  Joan Riera  i estic aprenent molt.

	room(pr1)      //Estic dins de l'aula {Joan Riera}
	room(student1) //Estic dins de l'aula {{Jon Ruiz} Go}
}
