package main

import "fmt"

type Person struct {
	id        int
	name      string
	lastName  string
	companies []string
}

func main() {

	aPerson := Person{
		1,
		"Shahrooz",
		"Jafari",
		[]string{"Asa", "Agah"},
	}
	otherPerson := &aPerson
	otherPerson.name = "Ali"
	fmt.Println(aPerson)
	fmt.Println(otherPerson)
}
