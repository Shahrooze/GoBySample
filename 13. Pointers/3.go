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
		id:        1,
		name:      "Shahrooz",
		lastName:  "Jafari",
		companies: []string{"Asa", "Agah"},
	}
	fmt.Println(aPerson)

	var otherPerson *Person = &aPerson
	otherPerson.name = "Other"
	fmt.Println(aPerson, otherPerson)
}
