package main

import "fmt"

type Person struct {
	id        int
	name      string
	lastName  string
	companies []string
}

func main() {
	var aPerson *Person
	fmt.Println(aPerson)
	aPerson = new(Person)
	//(*aPerson).lastName = "test"
	aPerson.lastName = "test"
	fmt.Println(aPerson)
}
