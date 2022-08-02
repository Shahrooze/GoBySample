package main

import "fmt"

func main() {
	name := "ali"
	sayMessage(&name, 12)
	fmt.Println(name)
}
func sayMessage(name *string, age int) {
	*name = "Shahrooz"
	fmt.Printf("hello %v %v \n", *name, age)
}
