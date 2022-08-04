package main

import "fmt"

type Person struct {
	id        int
	name      string
	lastName  string
	companies []string
}

func (p Person) sayHello() {
	fmt.Println("Hello", p)
}
func main() {

	aPerson := Person{
		1,
		"Shahrooz",
		"Jafari",
		[]string{"Asa", "Agah"},
	}
	aPerson.sayHello()
}
