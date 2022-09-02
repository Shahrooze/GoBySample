package main

import "fmt"

type Animal struct {
	name string
}
type Bird struct {
	Animal
	canFly bool
}

func main() {
	aBird := Bird{canFly: true, Animal: Animal{name: "Kalagh"}}
	fmt.Println(aBird)
}
