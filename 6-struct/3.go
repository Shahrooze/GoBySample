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
	fmt.Println(aPerson)
	//بهتره از این روش استفاده نکنید.
	//دلیلش رو باید بتونید حدس بزنید
}
